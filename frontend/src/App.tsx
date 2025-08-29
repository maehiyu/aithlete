import { Amplify } from 'aws-amplify';
import { Authenticator } from '@aws-amplify/ui-react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Container } from '@mui/material';
import TabBar from './component/TabBar';
import { ChatList } from './component/ChatList';
import { ChatDetail } from './component/ChatDetail';
import { ChatPage } from './component/ChatPage';
import './App.css';
import '@aws-amplify/ui-react/styles.css';
import { useCurrentUser } from './hooks/useParticipant';
import { UserInitialRegister } from './component/UserInitialRegister';

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

  return (
    <Container maxWidth="sm" sx={{ mt: 4 }}>
      {isLoading && <div>Loading...</div>}
      {!isLoading && !currentUser && <UserInitialRegister />}
      {!isLoading && currentUser && (
        <BrowserRouter>
          <TabBar />
          <Routes>
            <Route path="/" element={<ChatList />} />
            <Route path="/chat" element={<ChatPage />} />
            <Route path="/chats" element={<ChatList />} />
            <Route path="/chats/:id" element={<ChatDetail />} />
            <Route path="/setting" element={<div>Setting（仮実装）</div>} />
          </Routes>
        </BrowserRouter>
      )}
    </Container>
  );
}


export default function App() {
  return (
    <Authenticator loginMechanisms={["email"]}>
      <AuthedApp />
    </Authenticator>
  );
}
