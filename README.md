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
curl -H 'X-API-KEY:testkey' localhost:8080/hello
```

## メモ
- ミドルウェアを使ってAPIキーによる認証を実施
