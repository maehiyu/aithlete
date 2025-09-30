import { useParams, useLocation } from 'react-router-dom';
import { useChat, useSendMessage } from './useChat';
import { useCurrentUser } from '../participant/hooks/useParticipant';
import { useEffect, useRef, useState } from 'react';
import { ChatMessageItem } from '../../components/common/ChatMessageItem';
import { useChatEvents } from './useChatEvents';
import ChatInputBar from './components/ChatInputBar';
import { ChatLayout, LoadingPage, ErrorPage, usePageState } from '../../components/layout/PageLayout';
import { ChatActionMenu, ChatActionMenuItemConfig } from './components/ChatActionMenu';
import { UsersIcon, CalendarDaysIcon, TrashIcon } from '@heroicons/react/24/outline';
import { AppointmentFormModal } from '../../features/appointment/components/AppointmentFormModal';
import { AppointmentListModal } from '../../features/appointment/components/AppointmentListModal';

export function ChatDetail() {
    const { id: chatId } = useParams<{ id: string }>();
    const location = useLocation();
    const { data, isLoading, error } = useChat(chatId ?? "");
    const { data: currentUser } = useCurrentUser();
    const timeline = data?.timeline || [];
    const latestQuestion = timeline.slice().reverse().find(item => item.type === 'question');
    const latestQuestionId = latestQuestion?.id;
    const sendMessage = useSendMessage(chatId ?? "");
    useChatEvents(chatId ?? "");
    const bottomRef = useRef<HTMLDivElement | null>(null);
    const isFirstScroll = useRef(true);
    const [message, setMessage] = useState("");
    const sentInitialMessage = useRef(false);

    // モーダルの開閉状態
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [isListModalOpen, setIsListModalOpen] = useState(false);

    // ページ状態管理
    const pageState = usePageState(data, isLoading, error);

    // --- チャット操作ハンドラ ---
    const handleDeleteChat = () => {
        // TODO: チャット削除処理を実装
        alert(`Chat ${chatId} will be deleted.`);
    };

    const handleShowParticipants = () => {
        // TODO: 参加者表示処理を実装
        alert('Showing participants...');
    };

    const handleCreateSchedule = () => {
        setIsCreateModalOpen(true);
    };

    const handleShowSchedule = () => {
        setIsListModalOpen(true);
    };

    // ユーザーの役割に応じてメニュー項目を動的に生成
    const menuItems: ChatActionMenuItemConfig[] = [
        {
            label: "参加者を表示",
            icon: UsersIcon,
            onClick: handleShowParticipants,
        },
    ];

    if (currentUser?.role === 'coach') {
        menuItems.push({
            label: "予約を作成",
            icon: CalendarDaysIcon,
            onClick: handleCreateSchedule,
        });
    } else {
        menuItems.push({
            label: "予約を確認",
            icon: CalendarDaysIcon,
            onClick: handleShowSchedule,
        });
    }

    menuItems.push({
        label: "チャットを削除",
        icon: TrashIcon,
        onClick: handleDeleteChat,
        isDestructive: true,
    });

    useEffect(() => {
        if (sentInitialMessage.current) return;
        const initial = location.state?.initialMessage;
        if (initial && currentUser) {
            sendMessage.mutate({
                content: initial,
                participantId: currentUser.id,
                questionId: latestQuestionId,
                type: "question",
            });
            sentInitialMessage.current = true;
            if (location.state?.initialMessage) {
                window.history.replaceState(
                    { ...window.history.state, usr: { ...location.state, initialMessage: undefined } },
                    ''
                );
            }
        }
    }, [currentUser]);

    useEffect(() => {
        if (bottomRef.current) {
            const rect = bottomRef.current.getBoundingClientRect();
            if (rect.bottom <= window.innerHeight && rect.top >= 0) {
                return;
            }
            if (isFirstScroll.current) {
                bottomRef.current.scrollIntoView({ behavior: "auto" });
                isFirstScroll.current = false;
            } else {
                bottomRef.current.scrollIntoView({ behavior: "smooth" });
            }
        }
    }, [data?.timeline]);

    // 状態に応じたレンダリング
    if (pageState.type === 'loading') {
        return <LoadingPage message="チャットを読み込み中..." />;
    }

    if (pageState.type === 'error') {
        return <ErrorPage error={pageState.error} />;
    }

    if (pageState.type === 'empty' || !data) {
        return (
            <div className="min-h-screen bg-gray-50 flex items-center justify-center">
                <div className="text-center">
                    <div className="text-gray-500 text-sm">チャットが見つかりません</div>
                </div>
            </div>
        );
    }

    const handleSend = (msg: string) => {
        if (!currentUser) return;
        const type = currentUser.role === "coach" ? "answer" : "question";
        sendMessage.mutate({
            content: msg,
            participantId: currentUser.id,
            questionId: latestQuestionId,
            type,
        });
        setMessage("");
    };

    return (
        <ChatLayout
            headerContent={
                <div className="fixed top-4 right-4 z-30">
                    <ChatActionMenu items={menuItems} />
                </div>
            }
            inputBar={
                <ChatInputBar
                    value={message}
                    onChange={setMessage}
                    onSend={handleSend}
                />
            }
        >
            <div className="space-y-4">
                {timeline.map(item => (
                    <ChatMessageItem
                        key={item.id}
                        item={item}
                        currentUserId={currentUser?.id || ''}
                        participants={data.participants}
                    />
                ))}
                <div ref={bottomRef} />
            </div>
            {/* モーダルコンポーネントを配置 */}
            <AppointmentFormModal
                isOpen={isCreateModalOpen}
                onClose={() => setIsCreateModalOpen(false)}
                chatId={chatId ?? ''}
                initialParticipantIds={data?.participants.map(p => p.id) || []}
            />
            <AppointmentListModal
                isOpen={isListModalOpen}
                onClose={() => setIsListModalOpen(false)}
                chatId={chatId ?? ''}
            />
        </ChatLayout>
    );
}