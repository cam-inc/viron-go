# サンプル API (Go)

`docker-compose` を使用したサンプル API 環境を提供します。

## クイックスタート

### 事前準備

#### Docker のインストール

- Docker for Desktop をインストールしてください。[ダウンロードサイト](https://www.docker.com/products/docker-desktop)

#### 証明書ファイル
`viron.crt`、`viron.csr`、`viron.key` ファイルを取得し、`viron-go/example/cert` ディレクトリに配置してください。

#### .env ファイル
```
cd viron-go/example
cp .env.template .env
```
その後、プロジェクトのシークレット情報を記入してください。

### サンプル用データベース起動

```
cd viron-go/example

# MySQL を使用する場合
docker compose -f docker-compose-store.yaml up --build mysql

# Mongo を使用する場合
docker compose -f docker-compose-store.yaml up --build mongo
```

### サンプル用バックエンドアプリケーション起動
```
cd viron-go

# MySQL を使用する場合
task example-dev-mysql

# Mongo を使用する場合
task example-dev-mongo
```