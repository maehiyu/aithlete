import type { ParticipantCreateRequest, ParticipantResponse, ParticipantUpdateRequest } from "../type/type";
import { apiFetch } from "./apiClient";

export async function fetchParticipant(participantId: string): Promise<ParticipantResponse> {
  return apiFetch<ParticipantResponse>(`/api/participants/${participantId}`);
}

export async function createParticipant(data: ParticipantCreateRequest): Promise<ParticipantResponse> {
  return apiFetch<ParticipantResponse>("/api/participants", {
    method: "POST",
    body: data,
  });
}

export async function updateParticipant(participantId: string, data: ParticipantUpdateRequest): Promise<ParticipantResponse> {
  return apiFetch<ParticipantResponse>(`/api/participants/${participantId}`, {
    method: "PUT",
    body: data,
  });
}
