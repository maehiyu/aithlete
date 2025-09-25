import type { ChatCreateRequest, ChatDetailResponse, ChatItemRequest, ChatSummaryResponse, ChatUpdateRequest } from "../../types";
import { apiFetch } from "../../lib/apiClient";
import toCamelCase from "../../utils/toCamelCase";
import toSnakeCase from "../../utils/toSnakeCase";
import { v4 as uuidv4 } from 'uuid';


export async function fetchChats(): Promise<ChatSummaryResponse[]> {
	const res = await apiFetch<any[]>("/chats");
    return toCamelCase<ChatSummaryResponse[]>(res);
}


export async function fetchChat(chatId: string): Promise<ChatDetailResponse> {
	const res = await apiFetch<any>(`/chats/${chatId}`);
	return toCamelCase<ChatDetailResponse>(res);
}



export async function createChat(data: ChatCreateRequest): Promise<string> {
	const res = await apiFetch<{ id: string }>("/chats", {
		method: "POST",
		body: toSnakeCase(data),
	});
	return res.id;
}

export async function updateChat(chatId: string, data: ChatUpdateRequest): Promise<void> {
	await apiFetch<any>(`/chats/${chatId}`, {
		method: "PUT",
		body: toSnakeCase(data),
	});
}

export async function deleteChat(chatId: string): Promise<void> {
	await apiFetch(`/api/chats/${chatId}`, {
		method: "DELETE"
	});
}

export async function sendMessage(chatId: string, data: ChatItemRequest): Promise<void> {
	console.log('sendMessage chatItemRequest:', data);
	const tempId = uuidv4();
	await apiFetch<any>(`/chats/${chatId}/messages`, {
		method: "POST",
		body: toSnakeCase({ ...data, tempId }),
	});
}
