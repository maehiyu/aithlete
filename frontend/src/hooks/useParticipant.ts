import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  fetchParticipant,
  createParticipant,
  updateParticipant
} from "../services/participantService";
import type { ParticipantResponse, ParticipantCreateRequest, ParticipantUpdateRequest } from "../type/type";

export function useParticipant(participantId: string) {
  return useQuery<ParticipantResponse>({
    queryKey: ["participant", participantId],
    queryFn: () => fetchParticipant(participantId),
    enabled: !!participantId,
  });
}

export function useCreateParticipant() {
  const queryClient = useQueryClient();
  return useMutation<ParticipantResponse, Error, ParticipantCreateRequest>({
    mutationFn: createParticipant,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["participants"] });
    },
  });
}

export function useUpdateParticipant() {
  const queryClient = useQueryClient();
  return useMutation<ParticipantResponse, Error, { participantId: string; data: ParticipantUpdateRequest }>({
    mutationFn: ({ participantId, data }) => updateParticipant(participantId, data),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ["participant", variables.participantId] });
      queryClient.invalidateQueries({ queryKey: ["participants"] });
    },
  });
}
