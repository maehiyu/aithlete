import { ChevronLeftIcon } from '@heroicons/react/24/outline';
import NavButton from "./NavButton";

export default function BackButton() {
	return (
		<NavButton
			icon={<ChevronLeftIcon className="w-full h-full" />}
			onClick={() => {
				window.history.back();
			}}
			ariaLabel="前の画面に戻る"
			className="fixed top-4 left-4 z-30"
		/>
	);
}
