# carshare-backend

# 説明
このアプリは知り合い同士で車をシェアするアプリのAPIです。
[APIのplayground](https://my-go-api-onion0904-2d2c780f.koyeb.app/)

## 内容
クリーンアーキテクチャーを採用して物のシェア管理アプリのAPIを作成しました。軽くユニットテストをしてみたり(信頼性)、インターフェース経由でバックエンドの処理の機能を触るようにしたり(更新性)などを意識して開発しました。また、sqlcやgqlgenなどの新しい技術を取り入れつつ、初めてデプロイまでも実装するという経験もできました。

## 目的
- CleanArchitectureを学ぶ
- UnitTestを学ぶ
- GraphQLとsqlcを学ぶ
- CI/CDとdeployを学ぶ
- 友達と車をシェアしたい(最初の目的)

## 主なドメイン
- サインイン(auth)
    - メールアドレス
    - パスワード
- プロフィール設定(user)
    - 名前
- メンバーを招待(group)
    - QRコードで招待
    - 招待リンクで招待
- 日程決め(event)
    - カレンダー形式のイメージ
    - 予約作成でメモを追加できる
    - 通学で使うかどうかを追加できる
    - 一週間前(月曜日)から次の週(月曜日から日曜日)の登録可能
    - ルール
        - アイテムルール
            - デフォルトで重要0個、普通7個を一週間に登録できるようになってる
            - これはアイテムルールの変更APIで変更可能
        - かぶった場合
            - 重要な用事と重要な用事(早い者勝ち)
            - 重要な用事と普通の用事(重要な用事を優先)
            - 普通の用事と普通の用事(早い者勝ち)


## 使用技術
- API形式
    - GraphQL
- ORM
    - sqlc
- DB
    - supabase(postgresql)
- ログイン認証
    - JWT
- deploy
    - koyeb
    - docker
- test
    - testing
    - gomock

## ディレクトリの説明

```
├──app # アプリケーションコード
|   ├── cmd # アプリケーションのスタート地点 (main.go等)
|   ├── config # 各種設定値
|   ├── docs # 色んな使い方のドキュメント
|   ├── domain # ビジネスロジックの中核 、各種ドメインオブジェクト
|   ├── infrastructure # データベースや外部APIへの実装詳細
|   ├── presentation # ユーザーへの表⽰や⼊⼒ (HTTP等)
|   |└── server # HTTPサーバーの設定やルーティング
|   └── usecase # ユースケース
└── pkg # ドメインロジックをもたない汎⽤的な処理
```

## 参考

- [Goで学ぶGraphQLサーバーサイド入門](https://zenn.dev/hsaki/books/golang-graphql)
- [Go ⾔語で構築するクリーンアーキテクチャ設計](https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn?productVariantID=itEzQN5gKZX8gXMmLTEXAB)
- [supabase](https://qiita.com/FrohleinYoshie/items/4acf666572e54232589a)
- [koyeb](https://www.koyeb.com/docs/deploy/go)