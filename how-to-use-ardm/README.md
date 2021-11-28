# ブログ用サンプルRedis環境
Dockerを用いてローカル用のRedisを立ち上げるためのコード

## 事前準備
以下ツールの利用を前提とする
- Docker
- Docker Compose

## 使い方
環境立ち上げ
```
$ cd how-to-use-ardm
$ docker-compose up -d
```

確認
```
$ docker ps
```

停止
```
$ docker-compose down
```