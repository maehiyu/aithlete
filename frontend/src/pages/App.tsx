
import { Routes, Route } from 'react-router-dom';
import { ChatDetail } from '../features/chat/ChatDetail';
import { ChatPage } from '../features/user-dashboard/ChatPage';
import '../App.css';
import '@aws-amplify/ui-react/styles.css';
import { useCurrentUser } from '../features/participant/hooks/useParticipant';
import { UserInitialRegister } from '../features/auth/UserInitialRegister';
import CoachChatList from '../features/user-dashboard/components/CoachChatList';
import AIChatList from '../features/user-dashboard/components/AIChatList';
import CoachList from '../features/user-dashboard/components/CoachList';
import { ConfirmDialogProvider } from '../contexts/ConfirmDialogContext';
import { MainLayout } from '../components/layout/MainLayout';

export default function App() {
  const { data: currentUser, isLoading } = useCurrentUser();
  
  return (
    <>
      <ConfirmDialogProvider>
          {isLoading && <div>Loading...</div>}
          {!isLoading && !currentUser && <UserInitialRegister />}
          {!isLoading && currentUser && (
            <Routes>
              <Route element={<MainLayout />}>
                <Route path="/" element={<ChatPage />} />
                <Route path="/coaches" element={<CoachList />} />
                <Route path="/chats" element={<CoachChatList />} />
                <Route path="/aichats" element={<AIChatList />} />
                <Route path="/chats/:id" element={<ChatDetail />} />
                <Route path="/setting" element={<div>Setting（仮実装）</div>} />
              </Route>
            </Routes>
          )}
      </ConfirmDialogProvider>
    </>
  );
}
