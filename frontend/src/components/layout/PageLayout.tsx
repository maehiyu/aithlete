import React from 'react';

/**
 * Layout Design System Types
 * デザインシステムで定義された標準パターンの型定義
 */

// ページレイアウトの基本構造を強制する型
export interface PageLayoutProps {
  children: React.ReactNode;
  title?: string;
  maxWidth?: 'lg' | '2xl' | '4xl' | 'full';
  className?: string;
}

export const PageLayout: React.FC<PageLayoutProps> = ({ 
  children, 
  title, 
  maxWidth = '2xl',
  className = '' 
}) => {
  const maxWidthClass = {
    'lg': 'max-w-lg',
    '2xl': 'max-w-2xl', 
    '4xl': 'max-w-4xl',
    'full': 'max-w-full'
  }[maxWidth];

  return (
    <div className="min-h-screen bg-gray-50 px-4 sm:px-6 lg:px-8">
      <div className={`mx-auto ${maxWidthClass} pt-8 pb-12 ${className}`}>
        {title && (
          <h1 className="text-xl font-bold mb-6 text-gray-900">
            {title}
          </h1>
        )}
        {children}
      </div>
    </div>
  );
};

export const LoadingPage: React.FC<{ message?: string }> = ({ 
  message = '読み込み中...' 
}) => (
  <div className="min-h-screen bg-gray-50 flex items-center justify-center">
    <div className="text-gray-500">{message}</div>
  </div>
);

export const ErrorPage: React.FC<{ error: Error | string }> = ({ error }) => (
  <div className="min-h-screen bg-gray-50 flex items-center justify-center">
    <div className="text-center">
      <div className="text-red-600 text-sm font-medium">エラーが発生しました</div>
      <div className="text-gray-500 text-xs mt-1">
        {error instanceof Error ? error.message : String(error)}
      </div>
    </div>
  </div>
);

export const EmptyState: React.FC<{
  message: string;
  hint?: string;
}> = ({ message, hint }) => (
  <div className="text-center py-12">
    <div className="text-gray-500 text-sm">{message}</div>
    {hint && (
      <div className="text-gray-400 text-xs mt-1">{hint}</div>
    )}
  </div>
);

export function usePageState<T>(
  data: T | undefined,
  isLoading: boolean,
  error: Error | null
) {
  if (isLoading) return { type: 'loading' as const };
  if (error) return { type: 'error' as const, error };

  if (data !== undefined) return { type: 'success' as const, data };
  return { type: 'empty' as const };
}

// チャット専用レイアウト（固定入力欄付き）
export const ChatLayout: React.FC<{
  children: React.ReactNode;
  inputBar: React.ReactNode;
  className?: string;
}> = ({ children, inputBar, className = '' }) => (
  <div className="min-h-screen bg-white flex flex-col">
    {/* チャット領域（スクロール可能） */}
    <div className={`flex-1 overflow-y-auto px-4 sm:px-6 ${className}`}>
      <div className="mx-auto max-w-2xl py-4">
        {children}
      </div>
    </div>
    {/* 固定入力欄 */}
    <div className="flex-shrink-0 border-t border-gray-200 bg-white">
      <div className="mx-auto max-w-2xl px-4 sm:px-6 py-4">
        {inputBar}
      </div>
    </div>
  </div>
);
