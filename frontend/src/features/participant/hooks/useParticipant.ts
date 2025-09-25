import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  fetchParticipant,
  createUser,
  updateUser,
  getCurrentUser,
  createAICoach
} from "../services/participantService";
import type { ParticipantResponse, ParticipantCreateRequest, ParticipantUpdateRequest } from "../../../types";
import { fetchCoachesBySport } from "../services/coachService";

export function useCurrentUser() {
  return useQuery<ParticipantResponse>({
    queryKey: ['participant', 'me'],
    queryFn: getCurrentUser,
    staleTime: 5 * 60 * 1000, // 5分キャッシュ
    refetchOnWindowFocus: false,
  })
}

export function useParticipant(participantId: string) {
  return useQuery<ParticipantResponse>({
    queryKey: ["participant", participantId],
    queryFn: () => fetchParticipant(participantId),
    enabled: !!participantId,
    staleTime: 5 * 60 * 1000, // 5分キャッシュ
    refetchOnWindowFocus: false,
  });
}

export function useCoachesBySport(sport: string) {
  return useQuery<ParticipantResponse[]>({
    queryKey: ["coaches", sport],
    queryFn: () => fetchCoachesBySport(sport),
    enabled: !!sport,
  });
}


export function useCreateUser() {
  const queryClient = useQueryClient();
  return useMutation<string, Error, ParticipantCreateRequest>({
    mutationFn: (participantData: ParticipantCreateRequest) => {
      return createUser(participantData);
    },
  });
}

export function useUpdateUser() {
  const queryClient = useQueryClient();
  return useMutation<void, Error, { participantId: string; data: ParticipantUpdateRequest }>({
    mutationFn: ({ participantId, data }) => updateUser(participantId, data),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ["participant", variables.participantId] });
      queryClient.invalidateQueries({ queryKey: ["participants"] });
    },
  });
}

export function useCreateAICoach() {
  return useMutation<string, Error, string[]>({
    mutationFn: (sports: string[]) => createAICoach(sports),
  });
}