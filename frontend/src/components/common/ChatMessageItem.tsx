import type { ParticipantResponse, ChatItem } from '../../types';
import { ChatMessageContent } from './ChatMessageContent';
import { MessageHeader } from './MessageHeader';

interface ChatMessageItemProps {
  item: ChatItem;
  currentUserId: string;
  participants: ParticipantResponse[];
}

export function ChatMessageItem({ item, currentUserId, participants }: ChatMessageItemProps) {
  const participant = participants.find(p => p.id === item.participantId);
  const isMe = item.participantId === currentUserId;
  const isAnswer = item.type === 'answer' || item.type === 'ai_answer';

  return (
    <div
      style={{
        flexDirection: isAnswer ? 'column' : (isMe && !isAnswer ? 'row-reverse' : 'row'),
        justifyContent: isAnswer ? 'center' : 'flex-start',
        alignItems: isAnswer ? 'center' : 'flex-start',
        marginBottom: 24,
        width: '100%',
      }}
    >
      <MessageHeader participant={participant} isMe={isMe} createdAt={item.createdAt} />
      <div
        style={{
          display: 'flex',
          justifyContent: isMe ? 'flex-end' : 'flex-start',
          flexGrow: 0,
          marginLeft: isMe ? 50 : 0,
          marginRight: isMe ? 24 : 0,
        }}
      >
        <ChatMessageContent content={item.content} isAnswer={isAnswer} />
      </div>
    </div>
  );
}

export default ChatMessageItem;
