import { apiFetch } from "./apiClient";
import type { ParticipantResponse } from "../type/type";

export async function fetchCoachesBySport(sport: string): Promise<ParticipantResponse[]> {
  return apiFetch<ParticipantResponse[]>(`/coaches?sport=${encodeURIComponent(sport)}`);
}
