import React from "react";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import Typography from "@mui/material/Typography";
import Avatar from "@mui/material/Avatar";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import { Link } from "react-router-dom";
import type { ChatSummaryResponse } from "../../types";

export type ChatListItemProps = {
  chat: ChatSummaryResponse;
};

export default function ChatListItem({ chat }: ChatListItemProps) {
  return (
    <ListItem alignItems="flex-start" component={Link} to={`/chats/${chat.id}`}> 
      <ListItemAvatar>
        <Avatar src={chat.opponent?.iconUrl} alt={chat.opponent?.name} />
      </ListItemAvatar>
      <ListItemText
        primary={<Typography variant="h6">{chat.opponent?.name ?? "(No Title)"}</Typography>}
        secondary={
          <Typography
            variant="body2"
            color="text.secondary"
            noWrap
            sx={{
              overflow: 'hidden',
              textOverflow: 'ellipsis',
              whiteSpace: 'nowrap',
              display: 'block',
            }}
          >
            {chat.latestQa ?? "No recent activity"}
          </Typography>
        }
      />
    </ListItem>
  );
}
