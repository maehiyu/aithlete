import { Outlet, useLocation, useNavigate } from 'react-router-dom';
import HumMenu from './HumMenu';
import BackButton from '../common/BackButton';
import { useCurrentUser } from '../../features/participant/hooks/useParticipant';
import { ChatBubbleLeftRightIcon } from '@heroicons/react/24/outline';
import NavButton from '../common/NavButton';

const AIChatButton = () => {
  const navigate = useNavigate();

  return (
    <NavButton
      icon={<ChatBubbleLeftRightIcon className="w-full h-full text-white" />}
      onClick={() => navigate('/')}
      ariaLabel="チャットページを開く"
      className="fixed bottom-5 right-5 z-30 !bg-gray-800 hover:!bg-gray-600"
    />
  );
};


export function MainLayout() {
  const { data: currentUser } = useCurrentUser();
  const location = useLocation();

  const isChatDetail = /^\/chats\/[^/]+$/.test(location.pathname);

  const showAIChatButton = location.pathname !== '/' && !isChatDetail;

  return (
    <div className="app-layout">
      {isChatDetail ? <BackButton /> : <HumMenu user={currentUser} />}
      
      <main className="main-content">
        <Outlet />
      </main>

      {showAIChatButton && <AIChatButton />}
    </div>
  );
}
