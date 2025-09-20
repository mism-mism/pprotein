# Deploy to Production

本番環境へのデプロイ手順を実行します。

## 実行内容
- Docker イメージのビルド
- タグ付けとプッシュ
- 本番環境での起動確認

## コマンド
```bash
# Docker ビルドとプッシュ
docker-compose build
docker tag pprotein:latest ghcr.io/mism-mism/pprotein:latest
docker push ghcr.io/mism-mism/pprotein:latest

# 本番環境起動
docker-compose up -d
```

## デプロイ前チェックリスト
1. [ ] 全テストが通過していることを確認
2. [ ] リンターエラーがないことを確認
3. [ ] ビルドが成功することを確認
4. [ ] 環境変数が正しく設定されていることを確認
5. [ ] データベースマイグレーションが必要かチェック

## ロールバック手順
問題が発生した場合は以前のイメージに戻します：
```bash
docker tag ghcr.io/mism-mism/pprotein:previous ghcr.io/mism-mism/pprotein:latest
docker-compose restart
```
