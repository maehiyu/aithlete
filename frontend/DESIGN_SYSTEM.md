# Frontend Design System & Layout Guidelines

## レイアウトパターン規約

### 1. Page Layout Structure (ページレイアウト構造)

すべてのページコンポーネントは以下の階層構造に従う：

```tsx
// ✅ 推奨パターン
export default function PageComponent() {
  return (
    <div className="min-h-screen bg-gray-50 px-4 sm:px-6 lg:px-8">  {/* Page wrapper */}
      <div className="mx-auto max-w-2xl pt-8 pb-12">                {/* Content container */}
        <h1 className="text-xl font-bold mb-6 text-gray-900">       {/* Page title */}
          ページタイトル
        </h1>
        {/* コンテンツ */}
      </div>
    </div>
  );
}
```

#### 各層の責務:
- **Page wrapper**: 全画面レイアウト、背景、基本レスポンシブ
- **Content container**: コンテンツ幅制限、中央配置、縦方向余白
- **Page title**: 統一されたタイトルスタイル

### 2. Responsive Spacing (レスポンシブスペーシング)

#### 水平方向パディング:
```scss
px-4      // モバイル (16px)
sm:px-6   // タブレット (24px) - 640px以上  
lg:px-8   // デスクトップ (32px) - 1024px以上
```

#### 縦方向マージン:
```scss
pt-8      // ページ上部 (32px)
pb-12     // ページ下部 (48px)
mb-6      // タイトル下 (24px)
```

### 3. Content Width Standards (コンテンツ幅規格)

| 用途 | クラス | 幅 | 使用場面 |
|------|--------|-----|----------|
| 標準コンテンツ | `max-w-2xl` | 672px | 記事、フォーム、リスト |
| 狭いコンテンツ | `max-w-lg` | 512px | ログイン、モーダル |
| 広いコンテンツ | `max-w-4xl` | 896px | ダッシュボード、テーブル |
| 全幅 | `max-w-full` | 100% | 画像、ビデオ |

### 4. Component Spacing (コンポーネント間スペーシング)

```scss
space-y-2    // 密接な関連要素 (8px)
space-y-3    // リスト項目 (12px)  
space-y-4    // セクション内要素 (16px)
space-y-6    // セクション間 (24px)
space-y-8    // 大きなセクション間 (32px)
```

### 5. Color Palette (カラーパレット)

#### 背景色:
```scss
bg-white     // メインコンテンツ
bg-gray-50   // ページ背景
bg-gray-100  // カードの境界
```

#### テキスト色:
```scss
text-gray-900  // 主要テキスト
text-gray-700  // 補助テキスト  
text-gray-500  // ラベル、プレースホルダー
text-gray-400  // 非活性、ヒント
```

### 6. Loading & Error States (ローディング・エラー状態)

```tsx
// ✅ 統一されたローディング状態
if (isLoading) {
  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center">
      <div className="text-gray-500">読み込み中...</div>
    </div>
  );
}

// ✅ 統一されたエラー状態
if (error) {
  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center">
      <div className="text-center">
        <div className="text-red-600 text-sm font-medium">エラーが発生しました</div>
        <div className="text-gray-500 text-xs mt-1">
          {error instanceof Error ? error.message : String(error)}
        </div>
      </div>
    </div>
  );
}
```

### 7. Empty States (エンプティ状態)

```tsx
// ✅ 統一されたエンプティ状態
<div className="text-center py-12">
  <div className="text-gray-500 text-sm">
    データがありません
  </div>
  <div className="text-gray-400 text-xs mt-1">
    操作方法やヒント
  </div>
</div>
```

### 8. Anti-Patterns (避けるべきパターン)

```tsx
// ❌ 避けるべき
<Container maxWidth="sm">              // MUIコンポーネント依存
  <div style={{ padding: '20px' }}>   // インラインスタイル
    <Typography variant="h1">         // フレームワーク固有コンポーネント

// ❌ 避けるべき任意の数値
<div style={{ 
  maxWidth: '700px',    // 任意の数値
  padding: '2rem 0.5rem 3rem 1rem'  // 不統一なスペーシング
}}>

// ❌ レスポンシブ対応なし  
<div className="px-6 max-w-xl">      // ブレークポイント未考慮
```

### 9. Implementation Checklist (実装チェックリスト)

新しいページコンポーネント作成時:

- [ ] `min-h-screen`でフルハイトレイアウト
- [ ] レスポンシブパディング (`px-4 sm:px-6 lg:px-8`)
- [ ] 適切なコンテンツ幅制限 (`max-w-*`)
- [ ] 中央配置 (`mx-auto`)
- [ ] 統一された縦スペーシング (`pt-8 pb-12`)
- [ ] 一貫したタイトルスタイル
- [ ] ローディング・エラー・エンプティ状態の実装
- [ ] MUI等のフレームワーク依存コンポーネント不使用

### 10. Migration Guide (移行ガイド)

既存コンポーネントのモダン化:

1. **MUI Container削除**:
   ```tsx
   // Before
   <Container maxWidth="sm" sx={{ mt: 4 }}>
   
   // After  
   <div className="mx-auto max-w-2xl pt-8">
   ```

2. **インラインスタイル削除**:
   ```tsx
   // Before
   <div style={{ padding: '2rem', maxWidth: '600px' }}>
   
   // After
   <div className="p-8 max-w-xl">
   ```

3. **任意数値を標準スケールに**:
   ```tsx
   // Before  
   <div className="max-w-[650px] pt-[30px]">
   
   // After
   <div className="max-w-2xl pt-8">
   ```

この設計指針により、コードベース全体で一貫したUX・保守性・パフォーマンスを実現する。
