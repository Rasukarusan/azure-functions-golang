# Azure functions Golang with Docker

## Required

`func`コマンドを使えるようにするため下記でインストール

```bash
npm i -g azure-functions-core-tools@4 --unsafe-perm true
```

確認

```bash
$ func --version
4.0.3971
```

## Run

- Mac で起動する場合

```bash
go build main.go && func start
```

確認

```bash
$ curl http://localhost:7071/api/SampleHttpTrigger
{"Status":200,"Rssult":"ok"}
```

- docker で起動する場合

M1Mac だと起動に失敗します。

```bash
docker-compose up --build
```

確認

```bash
curl http://localhost:8080/api/SampleHttpTrigger
```

## Deploy

下記コマンドで Docker イメージを push 後、「Azure Portal > 関数アプリ」 より対象のアプリを選択し再起動することで Docker イメージを pull して終了です。

```bash
DOCKER_IMAGE_TAG='your_azr_acr.azurecr.io/tools/demo'

az acr login --name your_azr_acr

docker build -t tools-docker -f Dockerfile . \
&& docker tag tools-docker:latest ${DOCKER_IMAGE_TAG} \
&& docker push ${DOCKER_IMAGE_TAG}
```
