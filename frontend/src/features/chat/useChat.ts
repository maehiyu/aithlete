import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  fetchChats,
  fetchChat,
  createChat,
  updateChat,
  deleteChat,
  sendQuestion,
  sendAnswer
} from "./chatService";
import type { ChatDetailResponse, ChatSummaryResponse, ChatUpdateRequest, ChatCreateRequest } from "../../types";


export function useChats() {
  return useQuery<ChatSummaryResponse[]>({
    queryKey: ["chats"],
    queryFn: fetchChats,
    staleTime: 1000 * 60 * 5, // 5分
    refetchOnWindowFocus: false,
  });
}

export function useChat(chatId: string) {
  return useQuery<ChatDetailResponse>({
    queryKey: ["chat", chatId],
    queryFn: () => fetchChat(chatId),
    enabled: !!chatId,
    staleTime: 1000 * 60 * 5, // 5分
    refetchOnWindowFocus: false,
  });
}

export function useCreateChat() {
  const queryClient = useQueryClient();
  return useMutation<
    ChatDetailResponse,
    Error,
    string[]
  >({
    mutationFn: (opponentIds) => createChat({ participantIds: opponentIds }),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
      queryClient.setQueryData(["chat", data.id], data);
    },
  });
}

export function useCreateChatWithQuestion() {
  const queryClient = useQueryClient();
  return useMutation<
    ChatDetailResponse,
    Error,
    { opponentIds: string[]; questions: string[] }
  >({
    mutationFn: ({ opponentIds, questions}) =>
      createChat({ participantIds: opponentIds, questions: questions }),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
      queryClient.setQueryData(["chat", data.id], data);
    },
  });
}

export function useUpdateChat() {
  const queryClient = useQueryClient();
  return useMutation<
    ChatDetailResponse, // mutationFnの戻り値
    Error,              // エラー型
    { chatId: string; data: ChatUpdateRequest } // mutationFnの引数型
  >({
    mutationFn: ({ chatId, data }) => updateChat(chatId, data),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ["chat", variables.chatId] });
      queryClient.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useDeleteChat() {
  const queryClient = useQueryClient();
  return useMutation<
    void,    // mutationFnの戻り値
    Error,   // エラー型
    string   // mutationFnの引数型（chatId）
  >({
    mutationFn: (chatId) => deleteChat(chatId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useSendAnswer(chatId: string) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: { content: string; participantId: string; questionId: string }) => sendAnswer(chatId, data),
    onSuccess: () => {
      // WebSocketイベントでキャッシュが更新されるため、ここでは何もしない
    },
  });
}

export function useSendQuestion(chatId: string) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: { content: string; participantId: string }) => sendQuestion(chatId, data),
    onSuccess: () => {
      // WebSocketイベントでキャッシュが更新されるため、ここでは何もしない
    },
  });
}

export function useSendMessage(chatId: string, role: string, questionId?: string) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: { content: string; participantId: string, questionId?: string }) => {
      if (role === "coach") {
        if (!questionId) throw new Error("questionId is required for answer");
        return sendAnswer(chatId, { content: data.content, participantId: data.participantId, questionId });
      } else {
        return sendQuestion(chatId, { content: data.content, participantId: data.participantId });
      }
    },
    onSuccess: (result, variables) => {
      if (role === "coach") {
        // 自分が送った回答の場合、WebSocketイベントでキャッシュが更新されるため、ここでは何もしない
      } else {
        // 自分が送った質問の場合、WebSocketイベントでキャッシュが更新されるため、ここでは何もしない
      }
    },
  });
}