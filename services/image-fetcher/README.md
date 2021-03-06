## OGP 画像 URL 取得サービス

このディレクトリには OGP 画像 の URL を取得するサービス (Fetcher) の Go 言語による実装が含まれています。

主要なディレクトリ:

- `image-fetcher/`: OGP 画像 URL 取得の実装が含まれる
- `grpc/`: `fetcher/` を gRPC で操作するためのインターフェース (gRPC サーバー) を実装する

その他のディレクトリ:

- `pb/`: gRPC サービス定義から自動生成されたコード (リポジトリルートの `/pb/go` からコピーされたもの)
- `config/`: サーバーの設定を読み込む
- `log/`: ログ出力のためのユーティリティ

## テスト

以下のコマンドを実行します.

``` shell
make test
```
