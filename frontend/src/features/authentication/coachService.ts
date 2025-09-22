import { apiFetch } from "../../lib/apiClient";
import type { ParticipantResponse } from "../../types";

export async function fetchCoachesBySport(sport: string): Promise<ParticipantResponse[]> {
  return apiFetch<ParticipantResponse[]>(`/coaches?sport=${encodeURIComponent(sport)}`);
}
