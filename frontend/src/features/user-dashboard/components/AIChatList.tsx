import { useEffect } from 'react';
import { useChats } from '../../chat/useChat';
import { Avatar } from '@mui/material';
import ChatListItem from '../../../components/common/ChatListItem';
import { useNavigate } from 'react-router-dom';

export default function AIChatList() {
  const { data: chats, isLoading, error } = useChats();
  const navigate = useNavigate();

  // コーチとのチャットのみ抽出（例: chat.type === 'coach' など、実際の型に合わせて調整）
  const aiChats = chats?.filter(chat => chat.opponent.role === 'ai_coach');

  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div style={{ color: 'red' }}>Error: {error instanceof Error ? error.message : String(error)}</div>;

  return (
    <div style={{ maxWidth: 700, margin: '0 auto', padding: '2rem 0' }}>
      <h2 className="text-xl font-bold mb-4">AIコーチとのチャット一覧</h2>
      {aiChats && aiChats.length > 0 ? (
        <ul>
          {aiChats.map(chat => (
            <ChatListItem key={chat.id} chat={chat} />
          ))}
        </ul>
      ) : (
        <div>コーチとのチャットはありません</div>
      )}
    </div>
  );
}
