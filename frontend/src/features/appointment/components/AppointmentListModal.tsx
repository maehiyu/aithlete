import { Dialog, Transition } from '@headlessui/react';
import { Fragment } from 'react';
import { useAppointments } from '../useAppointment';
import { useCurrentUser } from '../../participant/hooks/useParticipant';

type AppointmentListModalProps = {
  isOpen: boolean;
  onClose: () => void;
  chatId?: string; // 特定のチャットの予約を表示する場合
  userId?: string; // 特定のユーザーの予約を表示する場合
  coachId?: string; // 特定のコーチの予約を表示する場合
};

export function AppointmentListModal({ isOpen, onClose, chatId, userId, coachId }: AppointmentListModalProps) {
  const { data: appointments, isLoading, error } = useAppointments({
    userId: userId,
    coachId: coachId,
    chatId: chatId,
  });

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
              <Dialog.Panel className="w-full max-w-2xl transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
                <Dialog.Title
                  as="h3"
                  className="text-lg font-medium leading-6 text-gray-900"
                >
                  予約一覧
                </Dialog.Title>
                <div className="mt-2">
                  {isLoading && <p>予約を読み込み中...</p>}
                  {error && <p className="text-red-500">エラー: {error.message}</p>}
                  {appointments && appointments.length > 0 ? (
                    <ul className="divide-y divide-gray-200">
                      {appointments.map((appointment) => (
                        <li key={appointment.id} className="py-4">
                          <p className="text-sm font-medium text-gray-900">{appointment.title}</p>
                          <p className="text-xs text-gray-500">日時: {new Date(appointment.scheduledAt).toLocaleString()}</p>
                          <p className="text-xs text-gray-500">時間: {appointment.duration}分</p>
                          <p className="text-xs text-gray-500">ステータス: {appointment.status}</p>
                          <p className="text-xs text-gray-500">参加者:</p>
                          <ul className="list-disc list-inside text-xs text-gray-500 ml-4">
                            {appointment.participants.map(p => (
                              <li key={p.participantId}>{p.participantId} ({p.participationStatus})</li>
                            ))}
                          </ul>
                        </li>
                      ))}
                    </ul>
                  ) : (
                    !isLoading && <p>予約はありません。</p>
                  )}
                </div>

                <div className="mt-4">
                  <button
                    type="button"
                    className="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    onClick={onClose}
                  >
                    閉じる
                  </button>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition>
  );
}
