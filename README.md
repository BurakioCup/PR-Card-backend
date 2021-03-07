# PR-Card-backend  

## プログラムを書く上で
基本的に下記のサイトを準拠  
https://golang.org/doc/effective_go.html

## Github
#### Branch命名規則
- master
  - プロダクトとしてリリースするためのブランチ. 基本触らない
- develop(default)  
  - 開発ブランチ． コードが安定し,リリース準備ができたら master へマージする. リリース前はこのブランチが最新バージョンとなる.
- feature
  - 機能の追加. develop から分岐し, develop にマージする.
  - feature-{任意で詳細}
- fix
  - 現在のプロダクトのバージョンに対する変更・修正用.
  - fix-{任意で詳細}
#### コミットメッセージ
- add:新機能
- fix:バグ修正
- wip:作業中（WIP：Work In Progress）
- clean:整理（削除も含む）
- エンドポイントごとにブランチをdevelopからきる 

## できれば 
- ブランチ名はケバブケース
- DBはスネークケース
- コードはキャメルケース