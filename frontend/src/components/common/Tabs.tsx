import { NavLink, Outlet, useLocation } from "react-router-dom";
import { useState, useEffect } from "react";
import { useSendHandler } from '../../context/SendHandlerContext';
import { Box, IconButton, TextField } from "@mui/material";
import SettingIcon from '@mui/icons-material/Settings';
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone';
import HistoryIcon from '@mui/icons-material/History';
import ChatIcon from '@mui/icons-material/ChatBubbleOutline';
import MenuIcon from '@mui/icons-material/Menu';
import SendIcon from '@mui/icons-material/Send';

export default function Tabs() {
    const [isInputMode, setIsInputMode] = useState(false);
    const [input, setInput] = useState("");
    const [isComposing, setIsComposing] = useState(false);
    const location = useLocation();
    const { handler } = useSendHandler();

    useEffect(() => {
        if (/^\/chat(\/|$)/.test(location.pathname) || /^\/chats\/[\w-]+$/.test(location.pathname)) {
            setIsInputMode(true);
        } else {
            setIsInputMode(false);
        }
    }, [location.pathname]);

    return (
        <>
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
                    {isInputMode ? (
                        <>
                            <Box sx={{ p: 1, border: '1px solid #eee', borderRadius: 50 }}>
                                <NavLink to="/chats">
                                    {({ isActive }) => (
                                        <IconButton color={isActive ? "primary" : "default"}>
                                            <MenuIcon color={isActive ? "primary" : "inherit"} />
                                        </IconButton>
                                    )}
                                </NavLink>
                            </Box>
                            <Box sx={{ p: 1, border: '1px solid #eee', borderRadius: 50, ml: 2, flex: 1, display: 'flex', alignItems: 'center' }}>
                                <TextField
                                    placeholder={""}
                                    value={input}
                                    onChange={e => setInput(e.target.value)}
                                    variant="standard"
                                    multiline
                                    minRows={1}
                                    maxRows={4}
                                    InputProps={{
                                        disableUnderline: true,
                                        sx: {
                                            p: 0,
                                            alignItems: 'center',
                                        },
                                        inputProps: {
                                            sx: {
                                                display: 'flex',
                                                alignItems: 'center',
                                                py: 0,
                                                resize: 'none',
                                            }
                                        }
                                    }}
                                    sx={{ width: '100%', flex: 1, minWidth: 0, transition: 'width 0.3s', border: 'none', textAlign: 'center' }}
                                    onCompositionStart={() => setIsComposing(true)}
                                    onCompositionEnd={() => setIsComposing(false)}
                                    onKeyDown={e => {
                                        if (e.key === 'Enter' && !e.shiftKey && !isComposing) {
                                            e.preventDefault();
                                            if (typeof handler === 'function' && input.trim()) {
                                                handler(input);
                                                setInput("");
                                            }
                                        }
                                    }}
                                />
                                <IconButton onClick={() => {
                                    if (typeof handler === 'function') {
                                        handler(input);
                                        setInput("");
                                    }
                                }} disabled={!input.trim()}>
                                    <SendIcon />
                                </IconButton>
                            </Box>
                        </>
                    ) : (
                        <Box
                            sx={{
                                display: 'flex',
                                alignItems: 'center',
                                flex: 1,
                                minWidth: 0,
                                justifyContent: 'space-between',
                            }}
                        >
                            <Box
                                sx={{
                                    border: '1px solid #eee',
                                    borderRadius: 50,
                                    p: 1,
                                    width: '100%',
                                    display: 'flex',
                                    justifyContent: 'space-between',
                                    alignItems: 'center',
                                    maxWidth: 320,
                                }}
                            >
                                <NavLink to="/settings">
                                    {({ isActive }) => (
                                        <IconButton color={isActive ? "primary" : "default"}>
                                            <SettingIcon color={isActive ? "primary" : "inherit"} />
                                        </IconButton>
                                    )}
                                </NavLink>
                                <NavLink to="/notifications">
                                    {({ isActive }) => (
                                        <IconButton color={isActive ? "primary" : "default"}>
                                            <NotificationsNoneIcon color={isActive ? "primary" : "inherit"} />
                                        </IconButton>
                                    )}
                                </NavLink>
                                <NavLink to="/chats">
                                    {({ isActive }) => (
                                        <IconButton color={isActive ? "primary" : "default"}>
                                            <HistoryIcon color={isActive ? "primary" : "inherit"} />
                                        </IconButton>
                                    )}
                                </NavLink>
                            </Box>
                            <Box sx={{ ml: 2, border: '1px solid #eee', borderRadius: 50, p: 1 }}>
                                <NavLink to="/chat">
                                    {({ isActive }) => (
                                        <IconButton color={isActive ? "primary" : "default"}>
                                            <ChatIcon color={isActive ? "primary" : "inherit"} />
                                        </IconButton>
                                    )}
                                </NavLink>
                            </Box>
                        </Box>
                    )}
                </Box>
            </Box>
            <Outlet />
        </>
    );
}
