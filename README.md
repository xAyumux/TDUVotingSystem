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

|             |      |     |
| ----------- | ---- | --- |
| id          | INT  | PK  |
| title       | TEXT |     |
| description | TEXT |     |

PK(id)

### Votes

|        |      |             |
| ------ | ---- | ----------- |
| userid | TEXT | PK          |
| id     | INT  | FK(pous.id) |
| result | INT  |             |

PK(userid)
FK(id)

## API 設計

- GET /get/poll

  今日の投票結果を取得
  title:array
  各選択肢と票数を格納

- GET /poll?id=xxxx

  QueryString を取得
  id から表示する選択肢を取得

- GET /list

  array[]:title,ID

- POST vote

  UserID,id,result（選択した番号）
