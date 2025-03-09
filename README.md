## **実行方法**

```sh
go run cmd/withCancel/main.go # キャンセルを実装
go run cmd/withTimeout/main.go # タイムアウトを実装
docker compose up -d # データベースを起動
go run cmd/db/main.go # データベースのタイムアウトを実装
docker compose down # データベースを停止
```

## **学習ポイント**

1. **`context`** を利用して、ワーカーの実行中に外部からキャンセル信号を送る方法を学べる
2. **`sync.WaitGroup`** と **`channel`** を組み合わせることで、並行処理の終了管理やリソースの安全な解放が実現できる
3. **`context.WithTimeout`** を利用して、タイムアウトを設定し、タイムアウトした場合にキャンセルする方法を学べる
4. **`ExecContext`** でコンテキスト付きクエリを実行し、タイムアウト時のエラー処理を実装する方法を習得できる

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
