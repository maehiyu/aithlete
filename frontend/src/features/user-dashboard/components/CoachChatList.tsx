import { useChats } from '../../chat/useChat';
import ChatListItem from '../../../components/common/ChatListItem';
import { PageLayout, LoadingPage, ErrorPage, EmptyState, usePageState } from '../../../components/layout/PageLayout';

export default function CoachChatList() {
  const { data: chats, isLoading, error } = useChats();
  const pageState = usePageState(chats, isLoading, error);
  const coachChats = chats?.filter(chat => chat.opponent.role === 'coach') || [];

  if (pageState.type === 'loading') {
    return <LoadingPage message="コーチのチャットを読み込み中..." />;
  }

  if (pageState.type === 'error') {
    return <ErrorPage error={pageState.error} />;
  }

  return (
    <PageLayout title="コーチとのチャット一覧" maxWidth="2xl">
      {/* データはあるがコーチのチャットがない場合 */}
      {pageState.type === 'success' && coachChats.length === 0 ? (
        <EmptyState 
          message="コーチとのチャットはまだありません"
          hint="コーチを探してチャットを始めましょう"
        />
      ) : (
        <div className="space-y-3">
          {coachChats.map(chat => (
            <ChatListItem key={chat.id} chat={chat} />
          ))}
        </div>
      )}
    </PageLayout>
  );
}
