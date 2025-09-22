import { useChats } from '../../chat/useChat';
import ChatListItem from '../../../components/common/ChatListItem';
import { useNavigate } from 'react-router-dom';

export default function CoachChatList() {
  const { data: chats, isLoading, error } = useChats();
  const navigate = useNavigate();

  const coachChats = chats?.filter(chat => chat.opponent.role === 'coach');

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div style={{ color: 'red' }}>Error: {error instanceof Error ? error.message : String(error)}</div>;

  return (
    <div style={{ maxWidth: 700}}>
      <h2 className="text-xl font-bold mb-4">コーチとのチャット一覧</h2>
      {coachChats && coachChats.length > 0 ? (
        <ul>
          {coachChats.map(chat => (
            <ChatListItem key={chat.id} chat={chat} />
          ))}
        </ul>
      ) : (
        <div>コーチとのチャットはありません</div>
      )}
    </div>
  );
}
