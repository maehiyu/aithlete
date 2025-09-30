import { apiFetch } from "../../lib/apiClient";
import { AppointmentResponse, AppointmentCreateRequest } from "../../types";
import toCamelCase from "../../utils/toCamelCase";
import toSnakeCase from "../../utils/toSnakeCase";

export async function fetchAppointmentsByUserId(userId: string): Promise<AppointmentResponse[]> {
	const response = await apiFetch<any>(`/appointments?userId=${userId}`);
    return toCamelCase<AppointmentResponse[]>(response);
}

export async function createAppointment(request: AppointmentCreateRequest): Promise<{ id: string }> {
    const response = await apiFetch<{ id: string }>(`/appointments`, {
        method: 'POST',
        body: toSnakeCase(request),
    });
    return toCamelCase(response);
}

export async function listAppointments(params?: { userId?: string; coachId?: string; chatId?: string }): Promise<AppointmentResponse[]> {
    const queryParams = new URLSearchParams();
    if (params?.userId) queryParams.append('user_id', params.userId);
    if (params?.coachId) queryParams.append('coach_id', params.coachId);
    if (params?.chatId) queryParams.append('chat_id', params.chatId);

    const queryString = queryParams.toString();
    const url = `/appointments${queryString ? `?${queryString}` : ''}`;

    const response = await apiFetch<any[]>(url);
    return toCamelCase<AppointmentResponse[]>(response);
}
