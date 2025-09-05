import { useParams, useLocation } from 'react-router-dom';
import { useChat, useSendMessage } from '../hooks/chatHooks';
import { useCurrentUser } from '../hooks/ParticipantHooks';
import { useSendHandler } from '../context/SendHandlerContext';
import { useEffect, useRef } from 'react';
import { useChatTimeline } from '../hooks/utils/useChatTimeline';
import { ChatMessageItem } from './ChatMessageItem';
import { useChatEvents } from '../hooks/useChatEvents';

type ChatDetailProps = {
    initialMessage?: string;
};


export function ChatDetail(props: ChatDetailProps) {
    const { id } = useParams<{ id: string }>();
    const location = useLocation();
    const { data, isLoading, error } = useChat(id ?? "");
    const { data: currentUser } = useCurrentUser();
    const latestQuestionId = data?.questions[data?.questions.length - 1]?.id;
    const sendMessage = useSendMessage(id ?? "", currentUser?.role ?? "", latestQuestionId);
    const { setHandler } = useSendHandler();

    useChatEvents(id ?? "");

    const bottomRef = useRef<HTMLDivElement | null>(null);
    const isFirstScroll = useRef(true);

    useEffect(() => {
        const initial = props.initialMessage ?? location.state?.initialMessage;
        if (initial && currentUser) {
            sendMessage.mutate({
                content: initial,
                participantId: currentUser.id,
                questionId: latestQuestionId,
            });

            if (location.state?.initialMessage) {
                window.history.replaceState(
                    { ...window.history.state, usr: { ...location.state, initialMessage: undefined } },
                    ''
                );
            }
        }
    }, []);

    useEffect(() => {
        setHandler(() => (input: string) => {
            if (!input.trim() || !currentUser) return;
            const latestQuestionId = data?.questions[data?.questions.length - 1]?.id;
            sendMessage.mutate({ content: input, participantId: currentUser.id, questionId: latestQuestionId });
        });
    }, [currentUser, setHandler, data]);


    useEffect(() => {
        if (bottomRef.current) {
            if (isFirstScroll.current) {
                bottomRef.current.scrollIntoView({ behavior: "auto" });
                isFirstScroll.current = false;
            } else {
                bottomRef.current.scrollIntoView({ behavior: "smooth" });
            }
        }
    }, [data?.questions, data?.answers]);

    const timeline = useChatTimeline(data?.questions, [
        ...(data?.answers || []),
        ...(data?.streamMessages ? [data.streamMessages] : [])
    ]);

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div style={{ color: 'red' }}>Error: {error instanceof Error ? error.message : String(error)}</div>;
    if (!data) return <div>Not found</div>;


    return (
        <div style={{ maxWidth: 600, margin: '0 auto'}}>
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
        </div>
    );
}