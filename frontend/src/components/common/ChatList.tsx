import React from "react";
import { useChats } from "../../features/chat/useChat";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Avatar from "@mui/material/Avatar";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import { Link } from "react-router-dom";

export function ChatList() {
  const { data, isLoading, error } = useChats();

  if (isLoading) return <Typography>Loading...</Typography>;
  if (error) return <Typography color="error">Error: {error instanceof Error ? error.message : String(error)}</Typography>;
  if (!data || data.length === 0) return <Typography>No chats found.</Typography>;

  return (
    <List>
      {data.map((chat, idx) => (
        <React.Fragment key={chat.id}>
          <ListItem
            alignItems="flex-start"
            component={Link}
            to={`/chats/${chat.id}`}
          >
            <ListItemAvatar>
              <Avatar src={chat.opponent.iconUrl} alt={chat.opponent.name} />
            </ListItemAvatar>
            <ListItemText
              primary={<Typography variant="h6">{chat.opponent.name ?? "(No Title)"}</Typography>}
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
          {idx < data.length - 1 && <Divider component="li" />}
        </React.Fragment>
      ))}
    </List>
  );
}
