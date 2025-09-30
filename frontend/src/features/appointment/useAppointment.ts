import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { createAppointment, listAppointments } from "./appointmentService";
import type { AppointmentCreateRequest, AppointmentResponse } from "../../types";

export function useCreateAppointment() {
  const queryClient = useQueryClient();
  return useMutation<{
    id: string;
  }, Error, AppointmentCreateRequest>({
    mutationFn: (newAppointment) => createAppointment(newAppointment),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["appointments"] });
    },
  });
}

export function useAppointments(params?: { userId?: string; coachId?: string; chatId?: string }) {
  return useQuery<AppointmentResponse[]>({
    queryKey: ["appointments", params],
    queryFn: () => listAppointments(params),
    enabled: !!params && (!!params.userId || !!params.coachId || !!params.chatId),
    staleTime: 1000 * 60 * 5, // 5 minutes
    refetchOnWindowFocus: false,
  });
}
