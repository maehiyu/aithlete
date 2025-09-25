import { useCreateAICoach, useCurrentUser } from '../participant/hooks/useParticipant';
import { useCreateChat } from '../chat/useChat';
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import ChatInputBar from '../chat/components/ChatInputBar';
import { PageLayout, LoadingPage, ErrorPage, usePageState } from '../../components/layout/PageLayout';

export function ChatPage() {
    const { data: currentUser, isLoading: isUserLoading, error: userError } = useCurrentUser();
    const [message, setMessage] = useState("");
    const navigate = useNavigate();
    const createAICoachMutation = useCreateAICoach();
    const createChatMutation = useCreateChat();

    // ページ状態管理
    const pageState = usePageState(currentUser, isUserLoading, userError);

    const handleCreateAIChat = async (message: string) => {
        if (!currentUser) return;
        try {
            const aiCoachId = await createAICoachMutation.mutateAsync(currentUser.sports);
            const chatId = await createChatMutation.mutateAsync([aiCoachId]);
            setMessage("");

            navigate(`/chats/${chatId}`, { state: { initialMessage: message } });
        } catch (e) {
            console.error("AIチャット作成エラー:", e);
        }
    };
    
    // 状態に応じたレンダリング
    if (pageState.type === 'loading') {
        return <LoadingPage message="ユーザー情報を読み込み中..." />;
    }

    if (pageState.type === 'error') {
        return <ErrorPage error={pageState.error} />;
    }

    if (pageState.type === 'empty' || !currentUser) {
        return (
            <PageLayout title="チャットを始める" maxWidth="2xl">
                <div className="text-center py-12">
                    <div className="text-gray-500 text-sm">
                        スポーツ情報がありません
                    </div>
                    <div className="text-gray-400 text-xs mt-1">
                        プロフィール設定でスポーツを選択してください
                    </div>
                </div>
            </PageLayout>
        );
    }

    return (
        <>
            <PageLayout title="AIコーチとチャットを始める" maxWidth="2xl">
                <div className="space-y-6 pb-24">
                    <div className="text-center py-8">
                        <div className="text-gray-600 mb-4">
                            選択したスポーツ: <span className="font-medium">{currentUser.sports?.join(', ')}</span>
                        </div>
                        <div className="text-sm text-gray-500">
                            下のメッセージボックスに質問や相談内容を入力してください
                        </div>
                    </div>
                </div>
            </PageLayout>
            <ChatInputBar
                value={message}
                onChange={setMessage}
                onSend={handleCreateAIChat}
            />
        </>
    );
}
