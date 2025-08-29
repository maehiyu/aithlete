import { useQuery } from "@tanstack/react-query";
import { fetchCoachesBySport } from "../services/coachService";
import type { ParticipantResponse } from "../type/type";

export function useCoachesBySport(sport: string) {
  return useQuery<ParticipantResponse[]>({
    queryKey: ["coaches", sport],
    queryFn: () => fetchCoachesBySport(sport),
    enabled: !!sport,
  });
}
