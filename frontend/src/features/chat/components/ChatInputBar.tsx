import React, { useState } from "react";
import { PaperAirplaneIcon, PlusIcon } from '@heroicons/react/24/outline';
import { Input } from "@headlessui/react"

export type ChatInputBarProps = {
  value?: string;
  onChange?: (value: string) => void;
  onSend?: (value: string) => void;
  onAttach?: () => void;
  disabled?: boolean;
  className?: string;
};

export default function ChatInputBar({
  value,
  onChange,
  onSend,
  onAttach,
  disabled = false,
  className = "",
}: ChatInputBarProps) {
  const [internalValue, setInternalValue] = useState("");
  const inputValue = value !== undefined ? value : internalValue;

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (onChange) {
      onChange(e.target.value);
    } else {
      setInternalValue(e.target.value);
    }
  };

  const handleSend = () => {
    if (inputValue.trim() && onSend) {
      onSend(inputValue);
      if (!onChange) setInternalValue("");
    }
  };

  return (
      <>
        <div
          className={`flex items-center gap-2 rounded-full p-3 bg-gray-300/20 gradient-border backdrop-blur fixed left-1/2 transform -translate-x-1/2 bottom-3 z-30 w-[calc(100%-1.5rem)] max-w-2xl ${className}`}
        >
          <button
            type="button"
            onClick={onAttach}
            className="w-10 h-10 flex items-center justify-center transition"
            aria-label="ファイルを添付"
          >
            <PlusIcon className="w-5 h-5" />
          </button>
          <Input
            type="text"
            value={inputValue}
            onChange={handleInputChange}
            placeholder="メッセージを入力..."
            className="px-4 py-2 transition flex-1 min-h-[40px] bg-transparent outline-none focus:ring-0"
          />
          <button
            type="button"
            onClick={handleSend}
            disabled={disabled || !inputValue.trim()}
            className="w-10 h-10 flex items-center justify-center border-none outline-none transition"
            aria-label="送信"
          >
            <PaperAirplaneIcon className="w-5 h-5" />
          </button>
        </div>
        <div className="fixed left-0 right-0 bottom-0 h-24 z-20 pointer-events-none bg-gradient-to-t from-white/80 via-white/60 to-white/20 backdrop-blur-lg" />
      </>
    );
}