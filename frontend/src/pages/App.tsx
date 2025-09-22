import { Amplify } from 'aws-amplify';
import { Authenticator } from '@aws-amplify/ui-react';
import { BrowserRouter, Routes, Route, useLocation } from 'react-router-dom';
import { Container } from '@mui/material';
import { ChatList } from '../components/common/ChatList';
import { ChatDetail } from '../features/chat/ChatDetail';
import { ChatPage } from '../features/user-dashboard/ChatPage';
import  HumMenu  from '../components/common/HumMenu';
import '../App.css';
import '@aws-amplify/ui-react/styles.css';
import { useCurrentUser } from '../features/authentication/useParticipant';
import { UserInitialRegister } from '../features/authentication/UserInitialRegister';
import CoachChatList from '../features/user-dashboard/components/CoachChatList';
import AIChatList from '../features/user-dashboard/components/AIChatList';
import CoachList from '../features/user-dashboard/components/CoachList';

Amplify.configure({
  Auth: {
    Cognito: {
      userPoolId: process.env.REACT_APP_AWS_COGNITO_USER_POOLS_ID || '',
      userPoolClientId: process.env.REACT_APP_AWS_COGNITO_USER_POOLS_CLIENT_ID || '',
      // identityPoolId: process.env.REACT_APP_AWS_COGNITO_IDENTITY_POOL_ID || '',
    }
  }
});

function AuthedApp() {
  const { data: currentUser, isLoading } = useCurrentUser();
  // useLocationはBrowserRouterの内側でしか使えないので、AuthedAppをRoutesの外でラップする
  const location = useLocation();
  const isChatDetail = /^\/chats\/[^/]+$/.test(location.pathname);
  return (
    <>
      <HumMenu isChatDetail={isChatDetail} />
      <Container maxWidth="sm" sx={{ mt: 4 }}>
        {isLoading && <div>Loading...</div>}
        {!isLoading && !currentUser && <UserInitialRegister />}
        {!isLoading && currentUser && (
          <Routes>
            <Route path="/" element={<ChatPage />} />
            <Route path="/coaches" element={<CoachList />} />
            <Route path="/chats" element={<CoachChatList />} />
            <Route path="/aichats" element={<AIChatList />} />
            <Route path="/chats/:id" element={<ChatDetail />} />
            <Route path="/setting" element={<div>Setting（仮実装）</div>} />
          </Routes>
        )}
      </Container>
    </>
  );
}


export default function App() {
  return (
    <BrowserRouter>
      <Authenticator loginMechanisms={["email"]}>
        <AuthedApp />
      </Authenticator>
    </BrowserRouter>
  );
}
