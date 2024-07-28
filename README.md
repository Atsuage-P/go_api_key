# APIキーによる認証
## 初回設定
- `.env.local` をコピーして `.env` を作成
```
cp .env.local .env
```

## 動作確認
```
# APIサーバ起動
go run main.go

# 別ターミナルでAPIコール
curl localhost:8080/hello

# POSTでリクエストボディの数字に1加えたものを返す
curl -X POST -H 'X-API-KEY:testkey' -H 'Content-Type:application/json' -d '{"num": 1}' localhost:8080/number

# DELETEでリクエストボディの数字から1引いたものを返す
curl -X DELETE -H 'X-API-KEY:testkey' -H 'Content-Type:application/json' -d '{"num": 1}' localhost:8080/number
```

## メモ
- ミドルウェアを使ってAPIキーによる認証を実施
- `echo` を使用
  - 標準のミドルウェアによる実装
  - ~~自作ミドルウェアによる実装~~
- `oapi-codegen` でOpenAPIのドキュメントからコードの生成
  - `oapi_cfg.yaml` の設定は [ここ](https://github.com/oapi-codegen/oapi-codegen/blob/main/configuration-schema.json) を参照する
```
oapi-codegen -config oapi_cfg.yaml openapi.yaml > ./oapi/oapi.gen.go
```
- `slog` によるログ出力
  - Go v1.21から導入された `slog` での構造化ログを確認する(サードパーティ製では `zap` が有名)
  - [参考にした実装](https://github.com/PumpkinSeed/slog-context/blob/main/examples/main.go)
  - [参考記事](https://blog.arthur1.dev/entry/2024/05/18/212731)
- `cotroller`パッケージのテスト実行時はcontrollerディレクトリ以下に.envをコピーする必要あり
- `panic` からの復帰と500エラーの応答は `echo` のミドルウェアが使えるが練習兼ねて自作の `defer recover` で返す
- `golangci-lint` による静的コードチェック
```
golangci-lint run
```
