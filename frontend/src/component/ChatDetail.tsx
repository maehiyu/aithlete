import { useParams } from 'react-router-dom';
import { useChat } from '../hooks/useChat';
import { Avatar } from '@mui/material';

export function ChatDetail() {
    const { id } = useParams<{ id: string }>();
    const { data, isLoading, error } = useChat(id ?? "");

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div style={{ color: 'red' }}>Error: {error instanceof Error ? error.message : String(error)}</div>;
    if (!data) return <div>Not found</div>;

    return (
        <div style={{ maxWidth: 600, margin: '0 auto'}}>
            <div>
                    {/* 質問（コーチ側） */}
                    {data.questions.map((q) => {
                        const participant = data.participants.find(p => p.id === q.participantId);
                        const isCoach = participant?.role === 'coach';
                        return (
                            <div
                                key={q.id}
                                style={{
                                    marginBottom: 24,
                                    textAlign: isCoach ? 'left' : 'right',
                                    display: 'flex',
                                    flexDirection: isCoach ? 'row' : 'row-reverse',
                                    alignItems: 'flex-start',
                                }}
                            >
                                <Avatar src={participant?.iconUrl || undefined} alt={participant?.name || ''} />
                                <div style={{ marginLeft: isCoach ? 8 : 0, marginRight: isCoach ? 0 : 8, maxWidth: '80%' }}>
                                    <div style={{ color: '#888', fontSize: 12 }}>{q.createdAt}</div>
                                    <div style={{ fontWeight: 'bold', marginBottom: 4 }}>{participant?.name ?? (isCoach ? 'Coach' : 'User')}</div>
                                    <div style={{ background: isCoach ? '#e3f2fd' : '#c8e6c9', borderRadius: 6, padding: 8, display: 'inline-block' }}>{q.content}</div>
                                </div>
                            </div>
                        );
                    })}
                    {/* 回答（ユーザー側） */}
                    {data.answers.map((a) => {
                        const participant = data.participants.find(p => p.id === a.participantId);
                        const isCoach = participant?.role === 'coach';
                        return (
                            <div
                                key={a.id}
                                style={{
                                    marginBottom: 24,
                                    textAlign: isCoach ? 'left' : 'right',
                                    display: 'flex',
                                    flexDirection: isCoach ? 'row' : 'row-reverse',
                                    alignItems: 'flex-start',
                                }}
                            >
                                <Avatar src={participant?.iconUrl || undefined} alt={participant?.name || ''} />
                                <div style={{ marginLeft: isCoach ? 8 : 0, marginRight: isCoach ? 0 : 8, maxWidth: '80%' }}>
                                    <div style={{ color: '#888', fontSize: 12 }}>{a.createdAt}</div>
                                    <div style={{ fontWeight: 'bold', marginBottom: 4 }}>{participant?.name ?? (isCoach ? 'Coach' : 'User')}</div>
                                    <div style={{ background: isCoach ? '#e3f2fd' : '#c8e6c9', borderRadius: 6, padding: 8, display: 'inline-block' }}>{a.content}</div>
                                </div>
                            </div>
                        );
                    })}
            </div>
        </div>
    );
}