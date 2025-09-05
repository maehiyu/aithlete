import { useEffect, useRef } from "react";
import { getIdToken } from "../services/authToken";
import { useQueryClient } from "@tanstack/react-query";
import { createChatWebSocket } from "../services/websocketService";
import toCamelCase from "../utils/toCamelCase";

// ChatEvent型（必要に応じて拡張）
type ChatEvent = {
  id: string;
  chat_id: string;
  type: string;
  from: string;
  to: string[];
  timestamp: number;
  payload: any;
};

export function useChatEvents(chatId: string) {
  const queryClient = useQueryClient();
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!chatId) return;
    let isMounted = true;
    let ws: WebSocket | null = null;
    (async () => {
      const token = await getIdToken();
      if (!token || !isMounted) return;
      ws = createChatWebSocket((data: ChatEvent) => {
        if (data.chat_id !== chatId) return;
        queryClient.setQueryData(["chat", chatId], (old: any) => {
          if (!old) return old;
          if (data.type === "answer") {
            const payload = toCamelCase(data.payload);
            return {
              ...old,
              answers: [...(old.answers || []), payload],
              streamMessages: undefined,
            };
          }
          if (data.type === "question") {
            const payload = toCamelCase(data.payload);
            return {
              ...old,
              questions: [...(old.questions || []), payload],
              streamMessages: undefined,
            };
          }
          if (data.type === "stream") {
            const payload = toCamelCase(data.payload);
            return {
              ...old,
              streamMessages: payload,
            };
          }
          return old;
        });
      }, token);
      wsRef.current = ws;
    })();
    return () => {
      isMounted = false;
      if (ws) ws.close();
    };
  }, [chatId]);
}