import { useCoachesBySport, useCurrentUser } from '../../authentication/useParticipant';
import Avatar from '@mui/material/Avatar';

export default function CoachList() {
  const { data: currentUser } = useCurrentUser();
  const { data: coaches, isLoading, error } = useCoachesBySport(currentUser?.sports?.[0] || '');

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div style={{ color: 'red' }}>Error: {error instanceof Error ? error.message : String(error)}</div>;
  if (!coaches || coaches.length === 0) return <div>コーチが見つかりません</div>;

  return (
    <div style={{ maxWidth: 900, margin: '0 auto', padding: '2rem 0' }}>
      <h2 className="text-xl font-bold mb-4">コーチ一覧</h2>
      <div className="grid grid-cols-3 sm:grid-cols-3 md:grid-cols-3">
        {coaches.map(coach => (
          <div key={coach.id} className="bg-white rounded-xl shadow p-6 flex flex-col items-center">
            <Avatar src={coach.iconUrl || undefined} alt={coach.name} sx={{ width: 64, height: 64, mb: 2 }} />
            <span className="font-bold text-lg mb-1">{coach.name}</span>
            <span className="text-gray-500 mb-2">{coach.sports.join(', ')}</span>
            {/* 追加情報やボタンなどもここに */}
          </div>
        ))}
      </div>
    </div>
  );
}
