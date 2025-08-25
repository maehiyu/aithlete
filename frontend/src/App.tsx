import React from "react";
import { BrowserRouter, Routes, Route, useLocation, useNavigate } from "react-router-dom";
import { Container } from "@mui/material";
import TabBar from "./component/TabBar";
import "./App.css";
import { ChatList } from "./component/ChatList";
import { ChatDetail } from "./component/ChatDetail";


function App() {
  return (
    <BrowserRouter>
      <TabBar />
      <Container maxWidth="sm" sx={{ mt: 4 }}>
        <Routes>
          <Route path="/" element={<ChatList />} />
          <Route path="/chat" element={<div>Chat（仮実装）</div>} />
          <Route path="/chats" element={<ChatList />} />
          <Route path="/chats/:id" element={<ChatDetail />} />
          <Route path="/setting" element={<div>Setting（仮実装）</div>} />
        </Routes>
      </Container>
    </BrowserRouter>
  );
}

export default App;
