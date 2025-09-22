import { useState, Fragment } from "react";

import { Dialog, DialogPanel, Transition, TransitionChild } from "@headlessui/react";
import NavButton from "./NavButton";
import { Bars2Icon, Bars3BottomLeftIcon, ChatBubbleLeftRightIcon, ChevronLeftIcon, UsersIcon } from '@heroicons/react/24/outline';
import IconLink from "./IconLink";

type HumMenuProps = {
	isChatDetail?: boolean;
};

export default function HumMenu({ isChatDetail = false }: HumMenuProps) {
	const [open, setOpen] = useState(false);

		return (
			<>

				{isChatDetail ? (
					<NavButton
						icon={<ChevronLeftIcon className="w-full h-full" />}
						onClick={() => {
							window.history.back();
						}}
						ariaLabel="前の画面に戻る"
						className="fixed top-4 left-4 z-30"
					/>
				) : (
					<NavButton
						icon={<Bars2Icon className="w-full h-full" />}
						onClick={() => setOpen(true)}
						ariaLabel="メニューを開く"
						className="fixed top-4 left-4 z-30"
					/>
				)}
				{/* 通常メニューはチャット詳細以外で表示 */}
				{!isChatDetail && (
					<Transition show={open} as={Fragment}>
						<Dialog as="div" className="fixed inset-0 z-40" onClose={setOpen}>
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
										<button
											className="mb-4 text-gray-500"
											onClick={() => setOpen(false)}
											aria-label="メニューを閉じる"
										>
											閉じる
										</button>
										<nav>
											<ul className="space-y-4">
												<li><IconLink href="/coaches" label="コーチ一覧" icon={<UsersIcon />} /></li>
												<li><IconLink href="/chats" label="チャット履歴" icon={<Bars3BottomLeftIcon />} /></li>
												<li><IconLink href="/aichats" label="AIチャット履歴" icon={<ChatBubbleLeftRightIcon />} /></li>
											</ul>
										</nav>
									</DialogPanel>
								</TransitionChild>
							</div>
						</Dialog>
					</Transition>
				)}
			</>
		);
}
