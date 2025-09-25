import React from "react";

type AvatarProps = {
  src?: string;
  alt?: string;
  size?: "sm" | "md" | "lg" | "xl";
  fallback?: string;
  className?: string;
};

const sizeClasses = {
  sm: "w-8 h-8 text-sm",
  md: "w-10 h-10 text-base",
  lg: "w-12 h-12 text-lg", 
  xl: "w-16 h-16 text-xl"
};

export default function Avatar({ 
  src, 
  alt, 
  size = "md", 
  fallback, 
  className = "" 
}: AvatarProps) {
  const [imageError, setImageError] = React.useState(false);
  
  const sizeClass = sizeClasses[size];
  const baseClasses = `${sizeClass} rounded-full bg-gray-200 flex items-center justify-center overflow-hidden ${className}`;

  // 画像が利用可能で、エラーがない場合
  if (src && !imageError) {
    return (
      <div className={baseClasses}>
        <img 
          src={src} 
          alt={alt || "Avatar"} 
          className="w-full h-full object-cover"
          onError={() => setImageError(true)}
        />
      </div>
    );
  }

  // フォールバック（イニシャルや名前の最初の文字）
  const displayFallback = fallback || alt?.charAt(0)?.toUpperCase() || "?";

  return (
    <div className={`${baseClasses} bg-gray-300 text-gray-600 font-medium`}>
      {displayFallback}
    </div>
  );
}
