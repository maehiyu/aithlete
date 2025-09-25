import { useEffect, useRef } from "react";
import { getIdToken } from "../../lib/authToken";
import { useQueryClient } from "@tanstack/react-query";
import { createChatWebSocket } from "../../lib/websocketService";
import toCamelCase from "../../utils/toCamelCase";

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
            const payload = toCamelCase(data.payload);
            const tempId = payload.tempId;
            let newTimeline = old.timeline || [];
            if (tempId) {
              let replaced = false;
              newTimeline = newTimeline.map((item: { tempId?: string }) => {
                if (item.tempId === tempId) {
                  replaced = true;
                  return payload;
                }
                return item;
              });
              if (!replaced) {
                newTimeline = [...newTimeline, payload];
              }
            } else {
              newTimeline = [...newTimeline, payload];
            }
            return {
              ...old,
              timeline: [...newTimeline],
            };
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
