import { Avatar } from '@mui/material';
import type { ParticipantResponse, ChatItem } from '../../types';
import { formatToMinute } from '../../utils/formatToMinute';

interface ChatMessageItemProps {
  item: ChatItem;
  currentUserId: string;
  participants: ParticipantResponse[];
}

export function ChatMessageItem({ item, currentUserId, participants }: ChatMessageItemProps) {
  const participant = participants.find(p => p.id === item.participantId);
  const isMe = item.participantId === currentUserId;
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: isMe ? 'row-reverse' : 'row',
        justifyContent: 'flex-start',
        alignItems: 'flex-start',
        marginBottom: 24,
        width: '100%', 
      }}
    >
      <Avatar src={participant?.iconUrl || undefined} alt={participant?.name || ''} />
      <div
        style={{
          marginLeft: isMe ? 0 : 8,
          marginRight: isMe ? 8 : 0,
          maxWidth: '80%',
          textAlign: isMe ? 'right' : 'left', 
        }}
      >
  <div style={{ color: '#888', fontSize: 12 }}>{formatToMinute(item.createdAt)}</div>
        <div style={{ fontWeight: 'bold', marginBottom: 4 }}>{participant?.name ?? (isMe ? 'You' : 'User')}</div>
        <div style={{
          background: isMe ? '#c8e6c9' : '#e3f2fd',
          borderRadius: 6,
          padding: 8,
          display: 'inline-block',
        }}>{item.content}</div>
      </div>
    </div>
  );
}

export default ChatMessageItem;
