# hackathon-back-app

## 無限環境構築編
- [golang](https://go.dev/)
- [docker](https://www.docker.com/)
- [heroku cli](https://devcenter.heroku.com/ja/articles/heroku-cli#install-the-heroku-cli)
- [postgres](https://www.postgresql.org/)

## golang & postgres
    // golang & postgres の立ち上げ
    // golang はホットリロードではないため変更するごとにcontainerにいってリロードボタンを押す必要がある
    $ docker-compose up -d

## postgres(local環境)
- postgres local環境はdockerにて管理
- 初回リロード時~/init/init.sqlをもとにsqlインジェクションを行う
- localで共有したいdbデータがあるときは~/init/init.sqlに追記
    ### 初期実行方法
        // git repository取得
        $ git clone git@github.com:anbyjap/hackathon-back-app.git

        // dockerの立ち上げ
        $ docker-compose up -d
    ### 実行の仕方(2回目以降)
        // dockerの立ち上げ
        $ docker-compose up -d

        // terminalとpostgresを繋げる
        $ docker-compose exec postgres bash

        // postgresにlogin test dbに移動
        $ psql -U admin test
    ### init.sql等を更新した時にすること
        // dockerを落とす
        $ docker-compose down -v

        // キャッシュを無視してbuild
        $ docker-compose build --no-cache
        // 以下 実行の仕方(2回目以降)　と同様に
        $ docker-compose up -d

## postgres(prod環境)
- herokuの環境設定がめんどくさい
- 変更をコードベースで追えないため誰が変更したか分かりにくい
- credential系を共有したくない
- 以上三つの理由から基本的にBackendチームしか触らないと思うので後でBackendチームみんなで環境構築