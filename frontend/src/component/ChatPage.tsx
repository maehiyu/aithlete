import { useCoachesBySport, useCreateAICoach } from '../hooks/ParticipantHooks';
import { useCurrentUser } from '../hooks/ParticipantHooks';
import { Avatar } from '@mui/material';
import { useCreateChat } from '../hooks/chatHooks';
import { useNavigate } from 'react-router-dom';
import { useQueryClient } from '@tanstack/react-query';
import { useSendHandler } from '../context/SendHandlerContext';
import { useEffect } from 'react';

export function ChatPage() {
    const { data: currentUser, isLoading: isUserLoading, error: userError } = useCurrentUser();
    const { data: coaches, isLoading: isCoachesLoading, error: coachesError } = useCoachesBySport(currentUser?.sports?.[0] || '');
    const navigate = useNavigate();
    const createAICoachMutation = useCreateAICoach();
    const createChatMutation = useCreateChat();
    const queryClient = useQueryClient();
    const { setHandler } = useSendHandler();

    const handleCreateChat = (coachId: string) => {
        if (!coachId) return;
        createChatMutation.mutate(
            [coachId],
            {
                onSuccess: (data) => {
                    queryClient.setQueryData(["chat", data.id], data);
                    navigate(`/chats/${data.id}`);
                },
            }
        );
    };


    const handleCreateAIChat = async (message: string) => {
        if (!currentUser) return;
        try {
            const aiCoach = await createAICoachMutation.mutateAsync(currentUser.sports);
            const chat = await createChatMutation.mutateAsync([aiCoach.id]);

            navigate(`/chats/${chat.id}`, { state: { initialMessage: message } });
        } catch (e) {

        }
    };
    
    useEffect(() => {
        setHandler(() =>(input: string) => handleCreateAIChat(input));
    }, [setHandler, currentUser]);

    if (isUserLoading) return <div>Loading...</div>;
    if (userError) return <div style={{ color: 'red' }}>Error: {userError instanceof Error ? userError.message : String(userError)}</div>;
    if (!currentUser) return <div>スポーツ情報がありません</div>;

    return (
        <div style={{ maxWidth: 700, margin: '0 auto'}}>
            <h2>コーチ一覧（{currentUser?.sports?.[0] || ''}）</h2>
            {isCoachesLoading && <div>Loading coaches...</div>}
            {coachesError && <div style={{ color: 'red' }}>Error: {coachesError instanceof Error ? coachesError.message : String(coachesError)}</div>}
            {coaches && coaches.length > 0 ? (
                <ul>
                    {coaches.map(coach => (
                        <li
                            key={coach.id}
                            style={{ display: 'flex', alignItems: 'center', marginBottom: 8, cursor: 'pointer', opacity: createChatMutation.status === 'pending' ? 0.5 : 1 }}
                            onClick={() => handleCreateChat(coach.id)}
                        >
                            <Avatar src={coach.iconUrl || undefined} alt={coach.name} style={{ marginRight: 8 }} />
                            <span style={{ fontWeight: 'bold', marginRight: 8 }}>{coach.name}</span>
                            <span>({coach.sports.join(', ')})</span>
                        </li>
                    ))}
                </ul>
            ) : (
                !isCoachesLoading && <div>該当するコーチがいません</div>
            )}
            {createChatMutation.isError && (
                <div style={{ color: 'red' }}>チャット作成に失敗しました: {createChatMutation.error instanceof Error ? createChatMutation.error.message : String(createChatMutation.error)}</div>
            )}
        </div>
    );
}
