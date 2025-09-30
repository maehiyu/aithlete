import { Fragment, useState } from "react";
import { Dialog, DialogPanel, Transition, TransitionChild, Disclosure, DisclosureButton, DisclosurePanel } from "@headlessui/react";
import { Bars3BottomLeftIcon, ChatBubbleLeftRightIcon, UsersIcon, XMarkIcon, ChevronRightIcon } from '@heroicons/react/24/outline';
import Avatar from "../common/Avatar";
import SettingsSheet from "../../features/participant/components/SettingsSheet";
import { useChats } from '../../features/chat/useChat';

type User = {
	name?: string;
	email?: string;
	iconUrl?: string | null;
};

type MenuDialogProps = {
	open: boolean;
	onClose: () => void;
	user?: User;
};

export default function MenuDialog({ open, onClose, user }: MenuDialogProps) {
	const [settingsOpen, setSettingsOpen] = useState(false);
	const { data: chats } = useChats();
	
	// AIチャットのみをフィルタリング（最新3件）
	const aiChats = chats?.filter(chat => chat.opponent.role === 'ai_coach').slice(0, 3) || [];
	// 通常のチャットのみをフィルタリング（最新3件）  
	const humanChats = chats?.filter(chat => chat.opponent.role !== 'ai_coach').slice(0, 3) || [];

	return (
		<Transition show={open} as={Fragment}>
			<Dialog as="div" className="fixed inset-0 z-40" onClose={onClose}>
				<div className="absolute inset-0 bg-black bg-opacity-30" />
				<div className="fixed inset-y-0 left-0 flex max-w-full">
					<TransitionChild
						as={Fragment}
						enter="transform transition ease-in-out duration-300"
						enterFrom="-translate-x-full"
						enterTo="translate-x-0"
						leave="transform transition ease-in-out duration-300"
						leaveFrom="translate-x-0"
						leaveTo="-translate-x-full"
					>
						<DialogPanel className="w-[90vw] max-w-xs sm:w-64 bg-white h-full shadow-xl p-6">
							<div className="flex items-center justify-between mb-6">
								<button
									onClick={() => setSettingsOpen(true)}
									className="flex items-center space-x-3 p-2 transition-colors flex-1 mr-2"
								>
									<Avatar 
										src={user?.iconUrl || undefined} 
										alt={user?.name}
										size="md"
										fallback={user?.name?.charAt(0)}
									/>
									<div className="text-left">
										<p className="font-medium text-gray-900">{user?.name || "ユーザー"}</p>
										<p className="text-sm text-gray-500">{user?.email}</p>
									</div>
								</button>
								<button
									onClick={onClose}
									className="p-2 hover:bg-gray-100 rounded-full transition-colors"
									aria-label="メニューを閉じる"
								>
									<XMarkIcon className="w-5 h-5 text-gray-400" />
								</button>
							</div>

							<nav>
								<div className="space-y-2">
									<Disclosure>
										{({ open }) => (
											<div>
												<DisclosureButton className="flex items-center justify-between w-full p-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
													<div className="flex items-center space-x-3">
														<Bars3BottomLeftIcon className="w-5 h-5" />
														<span>チャット履歴</span>
													</div>
													<ChevronRightIcon 
														className={`w-4 h-4 transition-transform ${open ? 'rotate-90' : ''}`} 
													/>
												</DisclosureButton>
												<DisclosurePanel className="pl-8 pb-2">
													<div className="space-y-1">
														<div className="space-y-1 max-h-32 overflow-y-auto">
															{humanChats.length > 0 ? (
																humanChats.map(chat => (
																	<a 
																		key={chat.id}
																		href={`/chats/${chat.id}`}
																		className="flex items-center space-x-2 p-2 text-xs text-gray-600 hover:bg-gray-50 rounded cursor-pointer"
																	>
																		<Avatar 
																			src={chat.opponent.iconUrl || undefined} 
																			alt={chat.opponent.name}
																			size="sm"
																			fallback={chat.opponent.name?.charAt(0)}
																		/>
																		<div className="flex-1 min-w-0">
																			<div className="font-medium truncate">{chat.opponent.name}とのチャット</div>
																			<div className="text-gray-400">{new Date(chat.lastActiveAt).toLocaleDateString('ja-JP')}</div>
																		</div>
																	</a>
																))
															) : (
																<div className="p-2 text-xs text-gray-500">
																	チャット履歴はありません
																</div>
															)}
														</div>
														<a href="/chats" className="block p-2 text-sm text-blue-600 hover:text-blue-800 hover:bg-blue-50 rounded text-right">
															すべて見る
														</a>
													</div>
												</DisclosurePanel>
											</div>
										)}
									</Disclosure>

									<Disclosure>
										{({ open }) => (
											<div>
												<DisclosureButton className="flex items-center justify-between w-full p-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
													<div className="flex items-center space-x-3">
														<ChatBubbleLeftRightIcon className="w-5 h-5" />
														<span>AIチャット履歴</span>
													</div>
													<ChevronRightIcon 
														className={`w-4 h-4 transition-transform ${open ? 'rotate-90' : ''}`} 
													/>
												</DisclosureButton>
												<DisclosurePanel className="pl-8 pb-2">
													<div className="space-y-1">
														<div className="space-y-1 max-h-32 overflow-y-auto">
															{aiChats.length > 0 ? (
																aiChats.map(chat => (
																	<a 
																		key={chat.id}
																		href={`/chats/${chat.id}`}
																		className="block p-2 text-xs text-gray-600 hover:bg-gray-50 rounded cursor-pointer"
																	>
																		<div className="font-medium truncate">{chat.latestQa || 'AIコーチとのチャット'}</div>
																		<div className="text-gray-400">{new Date(chat.lastActiveAt).toLocaleDateString('ja-JP')}</div>
																	</a>
																))
															) : (
																<div className="p-2 text-xs text-gray-500">
																	AIチャット履歴はありません
																</div>
															)}
														</div>
														<a href="/aichats" className="block p-2 text-sm text-blue-600 hover:text-blue-800 hover:bg-blue-50 rounded text-right">
															すべて見る
														</a>
													</div>
												</DisclosurePanel>
											</div>
										)}
									</Disclosure>

									<a 
										href="/coaches" 
										className="flex items-center space-x-3 w-full p-3 text-left text-gray-700 hover:bg-gray-50 rounded-lg transition-colors"
									>
										<UsersIcon className="w-5 h-5" />
										<span>コーチ一覧</span>
									</a>
								</div>
							</nav>
						</DialogPanel>
					</TransitionChild>
				</div>
			</Dialog>

			{/* Settings Sheet */}
			<SettingsSheet 
				open={settingsOpen}
				onClose={() => setSettingsOpen(false)}
				user={user}
			/>
		</Transition>
	);
}
