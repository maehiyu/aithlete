import React from 'react';
import ReactMarkdown from 'react-markdown';

interface ChatMessageContentProps {
  content: string;
  isAnswer: boolean;
}

export function ChatMessageContent({ content, isAnswer }: ChatMessageContentProps) {
  // Base styles for the message bubble
  const baseBubbleStyle: React.CSSProperties = {
    borderRadius: 6,
    padding: 8,
    display: 'inline-block',
    wordBreak: 'break-word', // Ensure long words break
    whiteSpace: 'pre-wrap', // Preserve whitespace and line breaks
  };

  // Styles specific to answers (large display, no enclosing area)
  const answerStyle: React.CSSProperties = {
    background: 'none', // Remove background
    padding: 8, // Remove padding
    borderRadius: 0, // Remove border radius
    display: 'block', // Take full width
    width: '100%',
    fontSize: '1.1em', // Slightly larger font
    fontWeight: 'normal',
  };

  // Styles for regular messages (question or ai_answer)
  const regularMessageStyle: React.CSSProperties = {
    background: '#eeeeee', // Default background for others (will be overridden by isMe in parent)
  };

  const finalStyle = {
    ...baseBubbleStyle,
    ...(isAnswer ? answerStyle : regularMessageStyle),
  };

  return (
    <div style={finalStyle}>
      <ReactMarkdown>
        {content}
      </ReactMarkdown>
    </div>
  );
}

export default ChatMessageContent;
