import type { ChatCreateRequest, ChatDetailResponse, ChatSummaryResponse, ChatUpdateRequest } from "../type/type";
import { apiFetch } from "./apiClient";
import toCamelCase from "../utils/toCamelCase";
import toSnakeCase from "../utils/toSnakeCase";


export async function fetchChats(): Promise<ChatSummaryResponse[]> {
	const res = await apiFetch<any[]>("/chats");
    return toCamelCase<ChatSummaryResponse[]>(res);
}


export async function fetchChat(chatId: string): Promise<ChatDetailResponse> {
	const res = await apiFetch<any>(`/chats/${chatId}`);
	return toCamelCase<ChatDetailResponse>(res);
}



export async function createChat(data: ChatCreateRequest): Promise<ChatDetailResponse> {
	const res = await apiFetch<any>("/chats", {
		method: "POST",
		body: toSnakeCase(data),
	});
	return toCamelCase<ChatDetailResponse>(res);
}

export async function updateChat(chatId: string, data: ChatUpdateRequest): Promise<ChatDetailResponse> {
	const res = await apiFetch<any>(`/chats/${chatId}`, {
		method: "PUT",
		body: toSnakeCase(data),
	});
	return toCamelCase<ChatDetailResponse>(res);
}

export async function deleteChat(chatId: string): Promise<void> {
	await apiFetch(`/api/chats/${chatId}`, {
		method: "DELETE"
	});
}

export async function sendQuestion(chatId: string, data: { content: string; participantId: string }) {
	const res = await apiFetch<any>(`/chats/${chatId}/questions`, {
		method: "POST",
		body: toSnakeCase(data),
	});
	return toCamelCase(res);
}


export async function sendAnswer(chatId: string, data: { content: string; participantId: string; questionId: string }) {
	const res = await apiFetch<any>(`/chats/${chatId}/answers`, {
		method: "POST",
		body: toSnakeCase(data),
	});
	return toCamelCase(res);
}
