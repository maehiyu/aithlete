import { Fragment, useState } from "react";
import { ChevronLeftIcon } from '@heroicons/react/24/outline';
import Avatar from "../../../components/common/Avatar";
import { useUpdateUser } from "../hooks/useParticipant";
import SportsSelector from "./SportsSelector";

type User = {
	name?: string;
	email?: string;
	iconUrl?: string | null;
	id?: string;
	sports?: string[];
};

type ProfileEditProps = {
	user?: User;
	onBack: () => void;
	onClose: () => void;
};

export default function ProfileEdit({ user, onBack, onClose }: ProfileEditProps) {
	const updateUser = useUpdateUser();
	const [formData, setFormData] = useState({
		name: user?.name || "",
		email: user?.email || "",
		sports: user?.sports || [],
	});
	const [isLoading, setIsLoading] = useState(false);

	const handleSave = async () => {
		if (!user?.id) return;

		setIsLoading(true);
		try {
			await updateUser.mutateAsync({
				participantId: user.id,
				data: {
					name: formData.name,
					sports: formData.sports,
					// email: formData.email, // emailの更新が可能な場合
				}
			});
			onClose(); // 保存後にシートを閉じる
		} catch (error) {
			console.error("プロフィール更新エラー:", error);
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<div className="h-full flex flex-col pt-0">
			{/* ヘッダー */}
			<div className="flex items-center justify-between px-6 pb-4 flex-shrink-0">
				<div className="flex items-center space-x-3">
					<button
						onClick={onBack}
						className="p-2 hover:bg-gray-100 rounded-full transition-colors"
						aria-label="設定に戻る"
					>
						<ChevronLeftIcon className="w-5 h-5 text-gray-600" />
					</button>
					<h2 className="text-lg font-semibold text-gray-900">プロフィール編集</h2>
				</div>
				<button
					onClick={handleSave}
					disabled={isLoading}
					className="px-4 py-2 bg-gray-900 text-white rounded-lg hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
				>
					{isLoading ? "保存中..." : "保存"}
				</button>
			</div>

			{/* プロフィール画像 */}
			<div className="px-6 py-4 border-b border-gray-100">
				<div className="flex items-center space-x-4">
					<Avatar 
						src={user?.iconUrl || undefined} 
						alt={user?.name}
						size="xl"
						fallback={user?.name?.charAt(0)}
					/>
					<div>
						<button className="text-sm text-gray-600 hover:text-gray-800 transition-colors">
							画像を変更
						</button>
					</div>
				</div>
			</div>

			{/* フォーム */}
			<div className="flex-1 px-6 py-4 space-y-4 overflow-y-auto">
				<div>
					<label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-2">
						名前
					</label>
					<input
						type="text"
						id="name"
						value={formData.name}
						onChange={(e) => setFormData(prev => ({ ...prev, name: e.target.value }))}
						className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-500 focus:border-transparent"
						placeholder="名前を入力"
					/>
				</div>

				<div>
					<label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
						メールアドレス
					</label>
					<input
						type="email"
						id="email"
						value={formData.email}
						onChange={(e) => setFormData(prev => ({ ...prev, email: e.target.value }))}
						className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-500 focus:border-transparent bg-gray-50"
						placeholder="メールアドレス"
						disabled
					/>
					<p className="text-xs text-gray-500 mt-1">
						メールアドレスの変更はできません
					</p>
				</div>

				<div>
					<label htmlFor="sports" className="block text-sm font-medium text-gray-700 mb-2">
						スポーツ
					</label>
					<SportsSelector
						selectedSports={formData.sports}
						onChange={(sports: string[]) => setFormData(prev => ({ ...prev, sports }))}
					/>
				</div>
			</div>
		</div>
	);
}
