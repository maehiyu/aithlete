import type { ParticipantCreateRequest, ParticipantResponse, ParticipantUpdateRequest } from "../../types";
import { apiFetch } from "../../lib/apiClient";


export async function getCurrentUser(): Promise<ParticipantResponse> {
  return apiFetch<ParticipantResponse>("/participants/me");
}

export async function fetchParticipant(participantId: string): Promise<ParticipantResponse> {
  return apiFetch<ParticipantResponse>(`/participants/${participantId}`);
}

export async function createUser(data: ParticipantCreateRequest): Promise<string> {
  const res = await apiFetch<ParticipantResponse>("/participants", {
    method: "POST",
    body: data,
  });
  return res.id;
}

export async function updateUser(participantId: string, data: ParticipantUpdateRequest): Promise<void> {
  await apiFetch<ParticipantResponse>(`/participants/${participantId}`, {
    method: "PUT",
    body: data,
  });
}

export async function createAICoach(sports: string[]): Promise<string> {
  const req: ParticipantCreateRequest = {
    name: 'AIコーチ',
    role: 'ai_coach',
    sports,
    email: '',
  };
  const res = await apiFetch<ParticipantResponse>("/participants", {
    method: "POST",
    body: req,
  });
  return res.id;
}
