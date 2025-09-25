import { useCoachesBySport, useCurrentUser } from '../../participant/hooks/useParticipant';
import Avatar from '@mui/material/Avatar';
import { PageLayout, LoadingPage, ErrorPage, usePageState } from '../../../components/layout/PageLayout';

export default function CoachList() {
  const { data: currentUser } = useCurrentUser();
  console.log('sport:', currentUser);
  const { data: coaches, isLoading, error } = useCoachesBySport(currentUser?.sports?.[0] || '');

  // ページ状態管理
  const pageState = usePageState(coaches, isLoading, error);

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
      {/* スマホでは画面いっぱい、大画面では中央寄せ */}
      <div className="w-full max-w-none sm:max-w-4xl sm:mx-auto">
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3 sm:gap-4">
          {coaches.map(coach => (
            <div key={coach.id} className="bg-white shadow rounded-lg p-4 flex flex-col items-center">
              <Avatar src={coach.iconUrl || undefined} alt={coach.name} sx={{ width: 64, height: 64, mb: 2 }} />
              <span className="font-bold text-sm sm:text-base mb-1 text-center">{coach.name}</span>
              <span className="text-gray-500 text-xs sm:text-sm text-center">{coach.sports.join(', ')}</span>
              {/* 追加情報やボタンなどもここに */}
            </div>
          ))}
        </div>
      </div>
    </PageLayout>
  );
}
