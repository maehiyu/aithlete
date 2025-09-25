import { useCreateAICoach, useCurrentUser } from '../authentication/useParticipant';
import { useCreateChat } from '../chat/useChat';
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import ChatInputBar from '../chat/components/ChatInputBar';

export function ChatPage() {
    const { data: currentUser, isLoading: isUserLoading, error: userError } = useCurrentUser();
    const [message, setMessage] = useState("");
    const navigate = useNavigate();
    const createAICoachMutation = useCreateAICoach();
    const createChatMutation = useCreateChat();

    const handleCreateAIChat = async (message: string) => {
        if (!currentUser) return;
        try {
            const aiCoachId = await createAICoachMutation.mutateAsync(currentUser.sports);
            const chatId = await createChatMutation.mutateAsync([aiCoachId]);
            setMessage("");

            navigate(`/chats/${chatId}`, { state: { initialMessage: message } });
        } catch (e) {

        }
    };
    

    if (isUserLoading) return <div>Loading...</div>;
    if (userError) return <div style={{ color: 'red' }}>Error: {userError instanceof Error ? userError.message : String(userError)}</div>;
    if (!currentUser) return <div>スポーツ情報がありません</div>;

    return (
    <div style={{ maxWidth: 700, margin: '0 auto', position: 'relative', minHeight: '80vh' }}>

            <ChatInputBar
                value={message}
                onChange={setMessage}
                onSend={handleCreateAIChat}
            />
        </div>
    );
}
