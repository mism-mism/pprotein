# Start Development Environment

開発環境を起動します。

## 実行内容
- pprotein メインサーバーの起動（ポート 9000）
- pprotein-agent の起動（ポート 19000）
- フロントエンド開発サーバーの起動（Vite）

## コマンド
```bash
# バックエンド（別ターミナルで実行）
make run          # メインサーバー
make run-agent    # エージェント

# フロントエンド（別ターミナルで実行）
cd view && npm run dev
```

## アクセス先
- Web UI: http://localhost:9000
- Agent: http://localhost:19000
- Frontend Dev: http://localhost:5173

## 開発時のワークフロー
1. バックエンドとフロントエンドを並行起動
2. ブラウザで Web UI にアクセス
3. コード変更時は自動リロード
4. Agent でプロファイリングデータを収集
