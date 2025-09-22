export function createChatWebSocket(onMessage: (data: any) => void, token: string): WebSocket {
  const protocol = window.location.protocol === "https:" ? "wss" : "ws";
  const host = window.location.hostname;
  const url = `${protocol}://${host}:9000/ws?token=${encodeURIComponent(token)}`;
  const ws = new WebSocket(url);
  ws.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data);
      onMessage(data);
    } catch (e) {
      // ignore
    }
  };
  return ws;
}