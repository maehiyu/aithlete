import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  fetchChats,
  fetchChat,
  createChat,
  updateChat,
  deleteChat
} from "../services/chatService";
import type { ChatDetailResponse, ChatSummaryResponse, ChatUpdateRequest, ChatCreateRequest } from "../type/type";

export function useChats() {
  return useQuery<ChatSummaryResponse[]>({
    queryKey: ["chats"],
    queryFn: fetchChats,
  });
}

export function useChat(chatId: string) {
  return useQuery<ChatDetailResponse>({
    queryKey: ["chat", chatId],
    queryFn: () => fetchChat(chatId),
    enabled: !!chatId,
  });
}

export function useCreateChat() {
  const queryClient = useQueryClient();
  return useMutation<
    ChatDetailResponse, // mutationFnの戻り値
    Error,              // エラー型
    ChatCreateRequest   // mutationFnの引数型
  >({
    mutationFn: createChat,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
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
