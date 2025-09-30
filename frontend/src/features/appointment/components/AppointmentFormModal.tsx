import { Dialog, Transition } from '@headlessui/react';
import { Fragment, useState } from 'react';
import { useCreateAppointment } from '../useAppointment';
import { AppointmentCreateRequest } from '../../../types';
import { useCurrentUser } from '../../participant/hooks/useParticipant';
import { useNavigate } from 'react-router-dom';

type AppointmentFormModalProps = {
  isOpen: boolean;
  onClose: () => void;
  chatId: string; // どのチャットに関連付けるか
  initialParticipantIds: string[]; // 初期参加者ID (例: 自分とコーチ)
};

export function AppointmentFormModal({ isOpen, onClose, chatId, initialParticipantIds }: AppointmentFormModalProps) {
  const { data: currentUser } = useCurrentUser();
  const navigate = useNavigate();
  const { mutate: createAppointmentMutation, isPending, error } = useCreateAppointment();

  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [scheduledAt, setScheduledAt] = useState(''); // ISO 8601形式の文字列
  const [duration, setDuration] = useState(60); // デフォルト60分

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!currentUser?.id) {
      alert('ユーザー情報が取得できません。');
      return;
    }

    const newAppointment: AppointmentCreateRequest = {
      chatId: chatId,
      title: title,
      description: description,
      scheduledAt: new Date(scheduledAt).toISOString(), // ISO 8601形式に変換
      duration: duration,
      participantIds: initialParticipantIds,
    };

    createAppointmentMutation(newAppointment, {
      onSuccess: (data) => {
        alert('予約が作成されました！');
        onClose();
        // 予約詳細ページへ遷移するなら
        // navigate(`/appointments/${data.id}`);
      },
      onError: (err) => {
        alert(`予約作成に失敗しました: ${err.message}`);
      },
    });
  };

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog as="div" className="relative z-10" onClose={onClose}>
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-black/25" />
        </Transition.Child>

        <div className="fixed inset-0 overflow-y-auto">
          <div className="flex min-h-full items-center justify-center p-4 text-center">
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 scale-95"
              enterTo="opacity-100 scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 scale-100"
              leaveTo="opacity-0 scale-95"
            >
              <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
                <Dialog.Title
                  as="h3"
                  className="text-lg font-medium leading-6 text-gray-900"
                >
                  新しい予約を作成
                </Dialog.Title>
                <div className="mt-2">
                  <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                      <label htmlFor="title" className="block text-sm font-medium text-gray-700">タイトル</label>
                      <input
                        type="text"
                        id="title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                        className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                      />
                    </div>
                    <div>
                      <label htmlFor="description" className="block text-sm font-medium text-gray-700">説明</label>
                      <textarea
                        id="description"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        rows={3}
                        className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                      />
                    </div>
                    <div>
                      <label htmlFor="scheduledAt" className="block text-sm font-medium text-gray-700">日時</label>
                      <input
                        type="datetime-local"
                        id="scheduledAt"
                        value={scheduledAt}
                        onChange={(e) => setScheduledAt(e.target.value)}
                        required
                        className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                      />
                    </div>
                    <div>
                      <label htmlFor="duration" className="block text-sm font-medium text-gray-700">時間 (分)</label>
                      <input
                        type="number"
                        id="duration"
                        value={duration}
                        onChange={(e) => setDuration(parseInt(e.target.value))}
                        required
                        min="1"
                        className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" 
                      />
                    </div>

                    <div className="mt-4">
                      <button
                        type="submit"
                        disabled={isPending}
                        className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2"
                      >
                        {isPending ? '作成中...' : '作成'}
                      </button>
                      <button
                        type="button"
                        className="ml-2 inline-flex justify-center rounded-md border border-transparent bg-gray-100 px-4 py-2 text-sm font-medium text-gray-900 hover:bg-gray-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-gray-500 focus-visible:ring-offset-2"
                        onClick={onClose}
                      >
                        キャンセル
                      </button>
                    </div>
                    {error && <p className="text-red-500 text-sm mt-2">エラー: {error.message}</p>}
                  </form>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition>
  );
}
