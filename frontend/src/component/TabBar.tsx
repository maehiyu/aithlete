import React from "react";
import { BottomNavigation, BottomNavigationAction, Box, TextField, IconButton, Menu } from "@mui/material";
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';
import HistoryIcon from '@mui/icons-material/History';
import SettingsIcon from '@mui/icons-material/Settings';
import MenuIcon from '@mui/icons-material/Menu';
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone';
import { useLocation, useNavigate } from "react-router-dom";

const TabBar: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();

  let tabValue = 0;
  if (location.pathname === "/chat" || location.pathname.startsWith("/chat/")) tabValue = 0;
  else if (location.pathname === "/chats" || location.pathname.startsWith("/chats/")) tabValue = 1;
  else if (location.pathname === "/setting" || location.pathname.startsWith("/setting")) tabValue = 2;

  const isChatPage =
    location.pathname === "/chat" ||
    location.pathname.startsWith("/chat/") ||
    /^\/chats\/[\w-]+$/.test(location.pathname);

  return (
    <Box sx={{ position: 'fixed', bottom: 0, left: 0, right: 0, zIndex: 1000, px: 2, pb: 1, pointerEvents: 'none' }}>
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          maxWidth: 600,
          mx: 'auto',
          pointerEvents: 'auto',
          width: '100%',
        }}
      >
        {isChatPage ? (
          <>
            <Box
              sx={{
                p: 1,
                border: '1px solid #eee',
                borderRadius: 50,
              }}
            >
              <IconButton onClick={() => navigate('/chats')}>
                <MenuIcon />
              </IconButton>
            </Box>
            <Box
              sx={{
                p: 1,
                border: '1px solid #eee',
                borderRadius: 50,
                ml: 2,
                flex: 1,
                display: 'flex',
                alignItems: 'center',
              }}
            >
              <TextField
                placeholder="質問を入力..."
                variant="standard"
                InputProps={{
                  disableUnderline: true,
                  inputProps: {
                    sx: {
                      height: '40px',
                      display: 'flex',
                      alignItems: 'center',
                      py: 0,
                    }
                  }
                }}
                sx={{ width: '100%', flex: 1, minWidth: 0, transition: 'width 0.3s', border: 'none', height: '40px', textAlign: 'center' }}
              />
            </Box>
          </>
        ) : (
          <Box
            sx={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'space-between',
              width: '100%',
              maxWidth: 600,
              mx: 'auto',
              pointerEvents: 'auto',
            }}
          >
            <Box
              sx={{
                display: 'flex',
                alignItems: 'center',
                border: '1px solid #eee',
                borderRadius: 50,
                p: 1,
                flex: 1,
                minWidth: 0,
                justifyContent: 'space-between',
              }}
            >
              <IconButton onClick={() => navigate('/setting')} color={tabValue === 2 ? 'primary' : 'default'}>
                <SettingsIcon />
              </IconButton>
              <IconButton onClick={() => navigate('/notifications')} color={tabValue === 3 ? 'primary' : 'default'}>
                <NotificationsNoneIcon />
              </IconButton>
              <IconButton onClick={() => navigate('/chats')} color={tabValue === 1 ? 'primary' : 'default'}>
                <HistoryIcon />
              </IconButton>
            </Box>
            <Box
              sx={{
                display: 'flex',
                alignItems: 'center',
                ml: 2,
                border: '1px solid #eee',
                borderRadius: 50,
                p: 1,
              }}
            >
              <IconButton onClick={() => navigate('/chat')} color={tabValue === 0 ? 'primary' : 'default'}
                sx={{

                }}
                >
                <ChatBubbleOutlineIcon />
              </IconButton>
            </Box>
          </Box>
        )}
      </Box>
    </Box>
  );
};

export default TabBar;
