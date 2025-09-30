import React from 'react';
import { Avatar } from '@mui/material';
import type { ParticipantResponse } from '../../types';
import { formatToMinute } from '../../utils/formatToMinute';

interface MessageHeaderProps {
  participant: ParticipantResponse | undefined;
  isMe: boolean;
  createdAt: string;
}

export function MessageHeader({ participant, isMe, createdAt }: MessageHeaderProps) {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: isMe ? 'row-reverse' : 'row', 
        justifyContent: 'flex-start', 
        alignItems: 'flex-start', 
        width: '100%',
        marginBottom: 0, 
      }}
    >
      <Avatar src={participant?.iconUrl || undefined} alt={participant?.name || ''} />
      <div
        style={{
          marginLeft: isMe ? 0 : 8,
          marginRight: isMe ? 8 : 0,
          textAlign: isMe ? 'right' : 'left', 
          flexGrow: 1,
        }}
      >
        <div style={{ color: '#888', fontSize: 12 }}>{formatToMinute(createdAt)}</div>
        <div style={{ fontWeight: 'bold', marginBottom: 4 }}>{participant?.name ?? (isMe ? 'You' : 'User')}</div>
      </div>
    </div>
  );
}

export default MessageHeader;
