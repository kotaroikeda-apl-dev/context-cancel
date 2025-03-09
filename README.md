## **実行方法**

```sh
go run cmd/withCancel/main.go # キャンセルを実装
```

## **学習ポイント**

1. **`context`** を利用して、ワーカーの実行中に外部からキャンセル信号を送る方法を学べる
2. **`sync.WaitGroup`** と **`channel`** を組み合わせることで、並行処理の終了管理やリソースの安全な解放が実現できる

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)