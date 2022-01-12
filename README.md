# TDUVotingSystem

学生証を使った投票システム

Docker で環境構築
API - Unity
API - Vue.js

## 使用技術

- Unity
- Vue.js
- Golang
- MySQL

## テーブル設計

### Polls

|             |      |                   |
| ----------- | ---- | ----------------- |
| id          | INT  | PK,AUTO_INCREMENT |
| title       | TEXT |                   |
| description | TEXT |                   |

PK(id)

### Votes

|        |      |              |
| ------ | ---- | ------------ |
| userid | TEXT | PK           |
| id     | INT  | FK(polls.id) |
| result | INT  |              |

PK(userid)
FK(id)

## API 設計

- GET /get/poll

  ID が一番大きいものを返す（最新）
  return JSON {
  Title
  ID
  Votes（result の件数）
  Description
  }

2 つのテーブルからデータを持ってくる
今日の投票結果を取得
title:array
各選択肢と票数を格納
最新の ID の投票トピックの各選択肢に何票入っているかのデータ
votes テーブル

- GET /poll?id=xxxx

  id から表示する投票トピックを指定する
  return JSON {
  Title
  ID
  Votes（result の件数）
  Description
  }

2 つのテーブルからデータを持ってくる
QueryString を取得
polls テーブル

- GET /list

return JSON{
Title
ID
}
Polls のテーブルの内容を全件持ってくる
array[]:title,ID
投票トピックの名前と ID のリストが返される
polls テーブル

- POST /polls

  polls テーブル
  AUTO_INCREMENT

- POST /votes

  UserID,id,result（選択した番号）
  投票のデータを追加する
  ※可能なら
  一つの投票トピックに同じ人が複数の票を入れようとしていたら拒否する
  votes テーブル
