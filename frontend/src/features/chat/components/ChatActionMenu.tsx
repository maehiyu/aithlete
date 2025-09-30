import { Menu, Transition } from "@headlessui/react";
import { EllipsisVerticalIcon, UsersIcon, CalendarDaysIcon, TrashIcon } from "@heroicons/react/24/outline";
import { Fragment } from "react";

type Props = {
    onDelete: () => void;
    onShowParticipants: () => void;
    onShowSchedule: () => void;
}

export function ChatActionMenu({ onDelete, onShowParticipants, onShowSchedule }: Props) {
    return (
        <Menu as="div" className="relative inline-block text-left">
            <div>
                <Menu.Button className="inline-flex items-center justify-center rounded-full bg-gray-300/20 hover:bg-white/80 gradient-border focus:outline-none transition min-h-[40px] min-w-[40px] w-12 h-12 backdrop-blur">
                    <EllipsisVerticalIcon className="w-6 h-6" aria-hidden="true" />
                </Menu.Button>
            </div>
            <Transition
                as={Fragment}
                enter="transition ease-out duration-100"
                enterFrom="transform opacity-0 scale-95"
                enterTo="transform opacity-100 scale-100"
                leave="transition ease-in duration-75"
                leaveFrom="transform opacity-100 scale-100"
                leaveTo="transform opacity-0 scale-95"
            >
                <Menu.Items className="absolute right-0 mt-2 w-56 origin-top-right divide-y divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black/5 focus:outline-none">
                    <div className="px-1 py-1 ">
                        <Menu.Item>
                            {({ active }) => (
                                <button
                                    onClick={onShowParticipants}
                                    className={`${active ? 'bg-gray-100' : ''} group flex w-full items-center rounded-md px-2 py-2 text-sm text-gray-900`}
                                >
                                    <UsersIcon className="mr-2 h-5 w-5" aria-hidden="true" />
                                    参加者を表示
                                </button>
                            )}
                        </Menu.Item>
                        <Menu.Item>
                            {({ active }) => (
                                <button
                                    onClick={onShowSchedule}
                                    className={`${active ? 'bg-gray-100' : ''} group flex w-full items-center rounded-md px-2 py-2 text-sm text-gray-900`}
                                >
                                    <CalendarDaysIcon className="mr-2 h-5 w-5" aria-hidden="true" />
                                    予約を確認
                                </button>
                            )}
                        </Menu.Item>
                    </div>
                    <div className="px-1 py-1">
                        <Menu.Item>
                            {({ active }) => (
                                <button
                                    onClick={onDelete}
                                    className={`${active ? 'bg-red-500 text-white' : 'text-gray-900'} group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                                >
                                    <TrashIcon className="mr-2 h-5 w-5" aria-hidden="true" />
                                    チャットを削除
                                </button>
                            )}
                        </Menu.Item>
                    </div>
                </Menu.Items>
            </Transition>
        </Menu>
    );
}