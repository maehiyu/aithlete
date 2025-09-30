import { Menu, Transition } from "@headlessui/react";
import { EllipsisVerticalIcon} from "@heroicons/react/24/outline";
import { Fragment } from "react";
import React from "react";

// メニュー項目の型定義
export type ChatActionMenuItemConfig = {
    label: string;
    icon: React.ElementType; // Heroiconsのコンポーネントを受け取るため
    onClick: () => void;
    isDestructive?: boolean; // 削除系アクション用
};

// 各メニュー項目をレンダリングするヘルパーコンポーネント
function ChatActionMenuItem({ onClick, label, icon: Icon, isDestructive = false }: ChatActionMenuItemConfig) {
    return (
        <div className="px-1 py-1">
            <Menu.Item>
                {({ active }) => (
                    <button
                        onClick={onClick}
                        className={`${active ? (isDestructive ? 'bg-red-500 text-white' : 'bg-gray-100') : (isDestructive ? 'text-red-600' : 'text-gray-900')} group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                    >
                        <Icon className="mr-2 h-5 w-5" aria-hidden="true" />
                        {label}
                    </button>
                )}
            </Menu.Item>
        </div>
    );
}

type Props = {
    items: ChatActionMenuItemConfig[];
}

export function ChatActionMenu({ items }: Props) {
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
                        {items.map((item, index) => (
                            <ChatActionMenuItem key={index} {...item} />
                        ))}
                    </div>
                </Menu.Items>
            </Transition>
        </Menu>
    );
}