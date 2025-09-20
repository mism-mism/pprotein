# Run Linters

Go と フロントエンドのリンターを実行してコード品質をチェックします。

## 実行内容
- golangci-lint による Go コードの検査
- ESLint による TypeScript/Vue コードの検査
- Prettier による フォーマットチェック

## コマンド
```bash
# Go linter
golangci-lint run

# フロントエンド linter
cd view && npm run lint

# フロントエンド自動修正
cd view && npm run fix
```

## 修正が必要な場合
1. 自動修正可能な問題は `npm run fix` で修正
2. Go の問題は手動修正またはコメント追加
3. 型エラーは TypeScript 定義を確認
4. Vue コンポーネントのベストプラクティスに従う

Cursor Rules に従ってコードスタイルを統一してください。
