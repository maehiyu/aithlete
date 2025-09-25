import { useEffect } from 'react';
import { useChats } from '../../chat/useChat';
import ChatListItem from '../../../components/common/ChatListItem';
import { useNavigate } from 'react-router-dom';
import { PageLayout, LoadingPage, ErrorPage, EmptyState, usePageState } from '../../../components/layout/PageLayout';

export default function AIChatList() {
  const { data: chats, isLoading, error } = useChats();
  const pageState = usePageState(chats, isLoading, error);
  const aiChats = chats?.filter(chat => chat.opponent.role === 'ai_coach') || [];

  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  if (pageState.type === 'loading') {
    return <LoadingPage message="AIコーチのチャットを読み込み中..." />;
  }

  if (pageState.type === 'error') {
    return <ErrorPage error={pageState.error} />;
  }

  return (
    <PageLayout title="AIコーチとのチャット一覧" maxWidth="2xl">
      {pageState.type === 'success' && aiChats.length === 0 ? (
        <EmptyState 
          message="AIコーチとのチャットはまだありません"
          hint="新しいスポーツを選択してAIコーチとチャットを始めましょう"
        />
      ) : (
        <div className="space-y-3">
          {aiChats.map(chat => (
            <ChatListItem key={chat.id} chat={chat} />
          ))}
        </div>
      )}
    </PageLayout>
  );
}
