import React from "react";

export type IconLinkProps = {
  href: string;
  icon?: React.ReactNode;
  label: string;
  className?: string;
  ariaLabel?: string;
};

export default function IconLink({ href, icon, label, className = "", ariaLabel }: IconLinkProps) {
  return (
    <a
      href={href}
  className={`flex items-center gap-2 transition ${className}`}
      aria-label={ariaLabel || label}
    >
      {icon && <span className="w-5 h-5 flex items-center justify-center">{icon}</span>}
      <span className="font-medium">{label}</span>
    </a>
  );
}
