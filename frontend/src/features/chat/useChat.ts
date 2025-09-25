import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  fetchChats,
  fetchChat,
  createChat,
  updateChat,
  deleteChat,
  sendMessage
} from "./chatService";
import type { ChatDetailResponse, ChatItemRequest, ChatSummaryResponse, ChatUpdateRequest} from "../../types";
import { send } from "process";


export function useChats() {
  return useQuery<ChatSummaryResponse[]>({
    queryKey: ["chats"],
    queryFn: fetchChats,
    staleTime: 1000 * 60 * 5, 
    refetchOnWindowFocus: false,
  });
}

export function useChat(chatId: string) {
  return useQuery<ChatDetailResponse>({
    queryKey: ["chat", chatId],
    queryFn: () => fetchChat(chatId),
    enabled: !!chatId,
    staleTime: 1000 * 60 * 5, 
    refetchOnWindowFocus: false,
  });
}

export function useCreateChat() {
  const queryClient = useQueryClient();
  return useMutation<string, Error, string[]>({
    mutationFn: (opponentIds) => createChat({ participantIds: opponentIds }),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useCreateChatWithQuestion() {
  const queryClient = useQueryClient();
  return useMutation<string,Error, { opponentIds: string[]; questions: string[] }>({
    mutationFn: ({ opponentIds, questions}) =>
      createChat({ participantIds: opponentIds, questions: questions }),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useUpdateChat() {
  const queryClient = useQueryClient();
  return useMutation<void, Error, { chatId: string; data: ChatUpdateRequest }>({
    mutationFn: ({ chatId, data }) => updateChat(chatId, data),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ["chat", variables.chatId] });
      queryClient.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useDeleteChat() {
  const queryClient = useQueryClient();
  return useMutation<void, Error, string>({
    mutationFn: (chatId) => deleteChat(chatId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["chats"] });
    },
  });
}

export function useSendMessage(chatId: string) {
  const queryClient = useQueryClient();
  return useMutation<void, Error, ChatItemRequest>({
    mutationFn: (data) => sendMessage(chatId, data),
    onSuccess: (result, variables) => {

    },
  });
}