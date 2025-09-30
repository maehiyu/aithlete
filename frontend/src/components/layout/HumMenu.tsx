import { useState } from "react";
import { Bars2Icon } from '@heroicons/react/24/outline';
import NavButton from "../common/NavButton";
import MenuDialog from "./MenuDialog";

type User = {
	name?: string;
	email?: string;
	iconUrl?: string | null;
};

type HumMenuProps = {
	user?: User;
};

export default function HumMenu({ user }: HumMenuProps) {
	const [open, setOpen] = useState(false);

	return (
		<>
			<NavButton
				icon={<Bars2Icon className="w-full h-full" />}
				onClick={() => setOpen(true)}
				ariaLabel="メニューを開く"
				className="fixed top-4 left-4 z-30"
			/>
			<MenuDialog 
				open={open} 
				onClose={() => setOpen(false)}
				user={user}
			/>
		</>
	);
}
