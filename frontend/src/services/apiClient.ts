
const baseUrl = process.env.REACT_APP_API_BASE_URL;

type ApiFetchOptions = {
    method?: "GET" | "POST" | "PUT" | "DELETE";
    body?: any;
};

export async function apiFetch<T>(endpoint: string, options: ApiFetchOptions = {}): Promise<T> {
    const { method = "GET", body } = options;
    // デフォルトヘッダー
    const defaultHeaders: Record<string, string> = {
        "Authorization": "Bearer user1", // ダミー
    };
    // Content-Typeはbodyがある場合のみ追加
    if (body !== undefined) {
        defaultHeaders["Content-Type"] = "application/json";
    }
    // ユーザー指定のheadersがあればマージ
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
