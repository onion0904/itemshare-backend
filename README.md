# carshare-backend

# 説明
このアプリは知り合い同士で車をシェアするアプリのAPIです。
[APIのplayground](https://my-go-api-onion0904-2d2c780f.koyeb.app/)

## 内容
クリーンアーキテクチャーを採用して車シェア管理アプリのAPIを作成しました。テストを行い信頼性を高めたり、インターフェース経由でバックエンドの処理の機能を触るようにしたりして更新性を高めたりなど、設計や実装をするところを意識して開発しました。また、sqlcやgqlgenなどの新しい技術を取り入れつつ、初めてデプロイまでも実装するという経験もできました。 

## 目的
- CleanArchitectureを学ぶ
- Testを学ぶ
- GraphQLとsqlcを学ぶ
- CI/CDとdeployを学ぶ
- 友達と車をシェアしたい

## ドメイン
- サインイン(auth)
    -メールアドレス
    - パスワード
    - パスワードの確認
- プロフィール設定(user)
    - 名前
    - アイコン(選ぶ形式にする)
- メンバーを招待(group)
    - QRコードで招待
    - 招待リンクで招待
- 日程決め(event)
    - カレンダー形式のイメージ
    - 予約作成でメモを追加できる
    - 通学で使うかどうかを追加できる
    - 一週間前(月曜日)から次の週(月曜日から日曜日)の登録可能
    - ルール
        - 重要な用事を2個+普通の用事を2個以上>=4
        - 四日まで登録可能
        - かぶった場合
        - 重要な用事と重要な用事(早い者勝ち)
        - 重要な用事と普通の用事(重要な用事を優先)
        - 普通の用事と普通の用事(早い者勝ち)
- メンバーリストの表示


## 使用技術
- API形式
    - GraphQL
- orm
    - sqlc
- DB
    - supabase(postgresql)
- ログイン認証
    - JWT
- SQL
    - MySQL
- deploy
    - koyeb
    - docker
- mock_test
    - gomock

## ディレクトリの説明

```
├──app # アプリケーションコード
|   ├── cmd # アプリケーションのスタート地点 (main.go等)
|   ├── config # 各種設定値
|   ├── docs # APIドキュメント
|   ├── domain # ビジネスロジックの中核 、各種ドメインオブジェクト
|   ├── infrastructure # データベースや外部APIへの実装詳細
|   ├── presentation # ユーザーへの表⽰や⼊⼒ (HTTP等)
|   ├── server # HTTPサーバーの設定やルーティング
|   └── usecase # ユースケース
└── pkg # ドメインロジックをもたない汎⽤的な処理
```

## 参考

- [Goで学ぶGraphQLサーバーサイド入門](https://zenn.dev/hsaki/books/golang-graphql)
- [Go ⾔語で構築するクリーンアーキテクチャ設計](https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn?productVariantID=itEzQN5gKZX8gXMmLTEXAB)
- [supabase](https://qiita.com/FrohleinYoshie/items/4acf666572e54232589a)
- [koyeb](https://www.koyeb.com/docs/deploy/go)