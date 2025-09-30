import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCoachesBySport, useCurrentUser } from '../../participant/hooks/useParticipant';
import Avatar from '@mui/material/Avatar';
import { PageLayout, LoadingPage, ErrorPage, usePageState } from '../../../components/layout/PageLayout';
import { createChat } from '../../chat/chatService';
import type { ChatCreateRequest } from '../../../types';

export default function CoachList() {
  const { data: currentUser } = useCurrentUser();

  const { data: coaches, isLoading, error } = useCoachesBySport(currentUser?.sports?.[0] || '');

  const navigate = useNavigate();
  const [isCreatingChat, setIsCreatingChat] = useState(false);
  const [chatError, setChatError] = useState<string | null>(null);

  const pageState = usePageState(coaches, isLoading, error);

  const handleCoachClick = async (coachId: string) => {
    if (!currentUser?.id) {
      setChatError('現在のユーザー情報が取得できません。');
      return;
    }
    if (isCreatingChat) return; 

    setIsCreatingChat(true);
    setChatError(null);

    try {
      const chatCreateRequest: ChatCreateRequest = {
        participantIds: [currentUser.id, coachId],
      };
      const newChatId = await createChat(chatCreateRequest);
      navigate(`/chats/${newChatId}`);
    } catch (err) {
      console.error('チャット作成エラー:', err);
      setChatError('チャットの作成に失敗しました。');
    } finally {
      setIsCreatingChat(false);
    }
  };

  if (pageState.type === 'loading') {
    return <LoadingPage message="コーチ情報を読み込み中..." />;
  }

  if (pageState.type === 'error') {
    return <ErrorPage error={pageState.error} />;
  }

  if (pageState.type === 'empty' || !coaches || coaches.length === 0) {
    return (
      <PageLayout title="コーチ一覧" maxWidth="4xl">
        <div className="text-center py-12">
          <div className="text-gray-500 text-sm">コーチが見つかりません</div>
        </div>
      </PageLayout>
    );
  }

  return (
    <PageLayout title="コーチ一覧" maxWidth="4xl">
      <div className="w-full max-w-none sm:max-w-4xl sm:mx-auto">
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3 sm:gap-4">
          {chatError && <div className="text-red-500 text-center mb-4">{chatError}</div>}
          {coaches.map(coach => (
            <div
              key={coach.id}
              className="bg-white shadow rounded-lg p-4 flex flex-col items-center cursor-pointer hover:bg-gray-50 transition-colors"
              onClick={() => handleCoachClick(coach.id)}
              style={{ opacity: isCreatingChat ? 0.6 : 1 }}
            >
              <Avatar src={coach.iconUrl || undefined} alt={coach.name} sx={{ width: 64, height: 64, mb: 2 }} />
              <span className="font-bold text-sm sm:text-base mb-1 text-center">{coach.name}</span>
              <span className="text-gray-500 text-xs sm:text-sm text-center">{coach.sports.join(', ')}</span>
            </div>
          ))}
        </div>
      </div>
    </PageLayout>
  );
}
