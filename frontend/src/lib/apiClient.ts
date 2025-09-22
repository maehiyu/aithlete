

import { getIdToken } from './authToken';

const baseUrl = process.env.REACT_APP_API_BASE_URL;

type ApiFetchOptions = {
    method?: "GET" | "POST" | "PUT" | "DELETE";
    body?: any;
};

export async function apiFetch<T>(endpoint: string, options: ApiFetchOptions = {}): Promise<T> {
    const { method = "GET", body } = options;
    const token = await getIdToken();
    const defaultHeaders: Record<string, string> = {};
    if (token) {
        defaultHeaders["Authorization"] = `Bearer ${token}`;
    }
    if (body !== undefined) {
        defaultHeaders["Content-Type"] = "application/json";
    }
    const fetchOptions: RequestInit = {
        method,
        headers: defaultHeaders,
    };
    if (body !== undefined) {
        fetchOptions.body = JSON.stringify(body);
    }
    const response = await fetch(`${baseUrl}${endpoint}`, fetchOptions);
    if (!response.ok) {
        throw new Error("Failed to fetch data");
    }
    return response.json();
}
