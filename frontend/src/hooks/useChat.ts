import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  fetchChats,
  fetchChat,
  createChat,
  updateChat,
  deleteChat,
  sendAnswer
} from "../services/chatService";
import type { ChatDetailResponse, ChatSummaryResponse, ChatUpdateRequest, ChatCreateRequest } from "../type/type";


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
    string[] // opponentIdsのみを引数に
  >({
    mutationFn: (opponentIds) => createChat({ opponentIds }),
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
    onSuccess: (answer) => {
      queryClient.setQueryData(["chat", chatId], (old: any) => {
        if (!old) return old;
        return {
          ...old,
          answers: [...(old.answers || []), answer],
        };
      });
    },
  });
}