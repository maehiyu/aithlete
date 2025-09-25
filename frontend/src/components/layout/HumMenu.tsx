import { useState } from "react";
import { Bars2Icon, ChevronLeftIcon } from '@heroicons/react/24/outline';
import NavButton from "../common/NavButton";
import MenuDialog from "./MenuDialog";

type User = {
	name?: string;
	email?: string;
	iconUrl?: string | null;
};

type HumMenuProps = {
	isChatDetail?: boolean;
	user?: User;
};

export default function HumMenu({ isChatDetail = false, user }: HumMenuProps) {
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
			{!isChatDetail && (
				<MenuDialog 
					open={open} 
					onClose={() => setOpen(false)}
					user={user}
				/>
			)}
		</>
	);
}
