## 目的
対人およびChatGPTとのチャットが可能なモバイルチャットアプリ

## 概要
### 機能一覧
https://qiita.com/q_hallelujah/private/24d5720ff7d07ba7996f

## 開発環境構築
```
~lambada関数実行~

$ go mod tidy

*local環境でlambda関数単体実行
$ make invoke-(lambda関数名) EVENT_FILE=(eventファイルパス)

*local環境でURL払出し実行
$ make run
```
```
~DynamoDB admin起動~

$ docker-compose build

$ docker compose up
```
### 動作確認
- ブラウザで、`http://localhost:8001`にアクセス
### テスト
```
$ make test
```

## トラブルシューティング
```
$ make clean

$ make build
```