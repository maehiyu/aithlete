import { Fragment, useState } from "react";
import { Dialog, DialogPanel, Transition, TransitionChild } from "@headlessui/react";
import { 
	XMarkIcon, 
	UserIcon, 
	BellIcon, 
	ShieldCheckIcon,
	ArrowRightOnRectangleIcon,
	CogIcon
} from '@heroicons/react/24/outline';
import Avatar from "../../../components/common/Avatar";
import { useConfirm } from "../../../contexts/ConfirmDialogContext";
import { useAuthenticator } from "@aws-amplify/ui-react";
import { useUpdateUser } from "../hooks/useParticipant";
import ProfileEdit from "./ProfileEdit";

type User = {
	name?: string;
	email?: string;
	iconUrl?: string | null;
	id?: string;
};

type SettingsSheetProps = {
	open: boolean;
	onClose: () => void;
	user?: User;
};

export default function SettingsSheet({ open, onClose, user }: SettingsSheetProps) {
	const { confirm } = useConfirm();
	const { signOut } = useAuthenticator();
	const updateUser = useUpdateUser();
	const [currentView, setCurrentView] = useState<'settings' | 'profile'>('settings');

	const handleProfileSettings = () => {
		setCurrentView('profile');
	};

	const handleBackToSettings = () => {
		setCurrentView('settings');
	};

	const handleLogout = async () => {
		const confirmed = await confirm({
			title: "ログアウトの確認",
			message: "本当にログアウトしますか？",
			confirmText: "ログアウト",
			cancelText: "キャンセル",
			type: "warning"
		});

		if (confirmed) {
			try {
				await signOut();
				onClose(); // シートを閉じる
			} catch (error) {
				console.error("ログアウトエラー:", error);
			}
		}
	};

	const settingsItems = [
		{
			icon: <UserIcon className="w-5 h-5" />,
			label: "プロフィール設定",
			onClick: handleProfileSettings,
		},
		{
			icon: <BellIcon className="w-5 h-5" />,
			label: "通知設定",
			onClick: () => console.log("通知設定"),
		},
		{
			icon: <ShieldCheckIcon className="w-5 h-5" />,
			label: "プライバシー",
			onClick: () => console.log("プライバシー"),
		},
		{
			icon: <CogIcon className="w-5 h-5" />,
			label: "一般設定",
			onClick: () => console.log("一般設定"),
		},
	];

	return (
		<Transition show={open} as={Fragment}>
			<Dialog as="div" className="fixed inset-0 z-50" onClose={onClose}>
				<div className="absolute inset-0 bg-black bg-opacity-50" />
				<div className="fixed inset-x-0 bottom-0 flex justify-center">
					<TransitionChild
						as={Fragment}
						enter="transform transition ease-in-out duration-300"
						enterFrom="translate-y-full"
						enterTo="translate-y-0"
						leave="transform transition ease-in-out duration-300"
						leaveFrom="translate-y-0"
						leaveTo="translate-y-full"
					>
						<DialogPanel className="w-full max-w-md bg-white rounded-t-xl shadow-xl h-[80vh] flex flex-col relative overflow-hidden">
							{/* ハンドル */}
							<div className="flex justify-center pt-4 pb-2 flex-shrink-0">
								<div className="w-12 h-1 bg-gray-300 rounded-full" />
							</div>

							{/* 設定画面（常に表示） */}
							<div className="flex-1 overflow-hidden">
								<SettingsView 
									user={user}
									onClose={onClose}
									onProfileSettings={handleProfileSettings}
									onLogout={handleLogout}
								/>
							</div>

							{/* プロフィール編集画面（スライドイン） */}
							<Transition
								show={currentView === 'profile'}
								as={Fragment}
							>
								<TransitionChild
									as={Fragment}
									enter="transform transition ease-in-out duration-300"
									enterFrom="translate-x-full"
									enterTo="translate-x-0"
									leave="transform transition ease-in-out duration-300"
									leaveFrom="translate-x-0"
									leaveTo="translate-x-full"
								>
									<div className="absolute inset-0 bg-white rounded-t-xl">
										{/* ハンドル（プロフィール編集画面用） */}
										<div className="flex justify-center pt-4 pb-2">
											<div className="w-12 h-1 bg-gray-300 rounded-full" />
										</div>
										<div className="h-full">
											<ProfileEdit 
												user={user}
												onBack={handleBackToSettings}
												onClose={onClose}
											/>
										</div>
									</div>
								</TransitionChild>
							</Transition>

						</DialogPanel>
					</TransitionChild>
				</div>
			</Dialog>
		</Transition>
	);
}

// SettingsView コンポーネントを同じファイル内に定義
type SettingsViewProps = {
	user?: User;
	onClose: () => void;
	onProfileSettings: () => void;
	onLogout: () => void;
};

function SettingsView({ user, onClose, onProfileSettings, onLogout }: SettingsViewProps) {
	const settingsItems = [
		{
			icon: <UserIcon className="w-5 h-5" />,
			label: "プロフィール設定",
			onClick: onProfileSettings,
		},
		{
			icon: <BellIcon className="w-5 h-5" />,
			label: "通知設定",
			onClick: () => console.log("通知設定"),
		},
		{
			icon: <ShieldCheckIcon className="w-5 h-5" />,
			label: "プライバシー",
			onClick: () => console.log("プライバシー"),
		},
		{
			icon: <CogIcon className="w-5 h-5" />,
			label: "一般設定",
			onClick: () => console.log("一般設定"),
		},
	];

	return (
		<div className="h-full flex flex-col">
			{/* ヘッダー */}
			<div className="flex items-center justify-between px-6 pb-4 flex-shrink-0">
				<h2 className="text-lg font-semibold text-gray-900">設定</h2>
				<button
					onClick={onClose}
					className="p-2 hover:bg-gray-100 rounded-full transition-colors"
					aria-label="設定を閉じる"
				>
					<XMarkIcon className="w-5 h-5 text-gray-400" />
				</button>
			</div>

			{/* ユーザー情報 */}
			<div className="px-6 py-4 border-b border-gray-100 flex-shrink-0">
				<div className="flex items-center space-x-4">
					<Avatar 
						src={user?.iconUrl || undefined} 
						alt={user?.name}
						size="lg"
						fallback={user?.name?.charAt(0)}
					/>
					<div>
						<h3 className="font-medium text-gray-900">{user?.name || "ユーザー"}</h3>
						<p className="text-sm text-gray-500">{user?.email}</p>
					</div>
				</div>
			</div>

			{/* スクロール可能なコンテンツエリア */}
			<div className="flex-1 overflow-y-auto">
				{/* 設定項目 */}
				<div className="px-6 py-4">
					<div className="space-y-1">
						{settingsItems.map((item, index) => (
							<button
								key={index}
								onClick={item.onClick}
								className="w-full flex items-center justify-between px-3 py-3 hover:bg-gray-50 rounded-lg transition-colors text-left"
							>
								<div className="flex items-center space-x-3">
									<div className="text-gray-600">{item.icon}</div>
									<span className="text-gray-900">{item.label}</span>
								</div>
								<div className="text-gray-400">
									<svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
										<path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
									</svg>
								</div>
							</button>
						))}
					</div>
				</div>

				{/* ログアウト */}
				<div className="px-6 py-4 border-t border-gray-100">
					<button
						onClick={onLogout}
						className="w-full flex items-center justify-between px-3 py-3 hover:bg-red-50 rounded-lg transition-colors text-left"
					>
						<div className="flex items-center space-x-3">
							<div className="text-red-600">
								<ArrowRightOnRectangleIcon className="w-5 h-5" />
							</div>
							<span className="text-red-600">ログアウト</span>
						</div>
					</button>
				</div>

				{/* 安全エリア */}
				<div className="h-safe-area-inset-bottom" />
			</div>
		</div>
	);
}
