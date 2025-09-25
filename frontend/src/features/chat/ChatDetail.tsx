import { useParams, useLocation } from 'react-router-dom';
import { useChat, useSendMessage } from './useChat';
import { useCurrentUser } from '../authentication/useParticipant';
import { useEffect, useRef } from 'react';
import { useState } from 'react';
import { ChatMessageItem } from '../../components/common/ChatMessageItem';
import { useChatEvents } from './useChatEvents';
import ChatInputBar from './components/ChatInputBar';

export function ChatDetail() {
    const { id } = useParams<{ id: string }>();
    const location = useLocation();
    const { data, isLoading, error } = useChat(id ?? "");
    const { data: currentUser } = useCurrentUser();
    const timeline = data?.timeline || [];
    const latestQuestion = timeline.slice().reverse().find(item => item.type === 'question');
    const latestQuestionId = latestQuestion?.id;
    const sendMessage = useSendMessage(id ?? "");
    useChatEvents(id ?? "");
    const bottomRef = useRef<HTMLDivElement | null>(null);
    const isFirstScroll = useRef(true);
    const [message, setMessage] = useState("");
    const sentInitialMessage = useRef(false);

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

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div style={{ color: 'red' }}>Error: {error instanceof Error ? error.message : String(error)}</div>;
    if (!data) return <div>Not found</div>;

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
        <div style={{ maxWidth: 600, margin: '0 auto', paddingBottom: '64px' }}>
            <div>
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
            <ChatInputBar
                value={message}
                onChange={setMessage}
                onSend={handleSend}
            />
        </div>
    );
}