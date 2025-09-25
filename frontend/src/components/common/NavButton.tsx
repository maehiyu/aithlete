import React from "react";
import { Button } from "@headlessui/react";

type NavButtonProps = {
	label?: string;
	icon?: React.ReactNode;
	onClick?: () => void;
	className?: string;
	ariaLabel?: string;
	type?: "button" | "submit" | "reset";
	disabled?: boolean;
};

export default function NavButton({
	label,
	icon,
	onClick,
	className = "",
	ariaLabel,
	type = "button",
	disabled = false,
}: NavButtonProps) {
	return (
		<Button
			type={type}
			onClick={onClick}
			className={`inline-flex items-center justify-center rounded-full bg-gray-300/20 hover:bg-white/80 gradient-border focus:outline-none transition min-h-[40px] min-w-[40px] w-12 h-12 backdrop-blur ${className}`}
			aria-label={ariaLabel || label}
			disabled={disabled}
		>
			{icon && (
				<span className={`flex items-center justify-center w-6 h-6${label ? ' mr-2' : ''}`}>
					{icon}
				</span>
			)}
			{label}
		</Button>
	);
}
