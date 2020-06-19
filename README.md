# TechTrain Mission Game API
## Overview
TechTrainのMissionの1つのCA Tech Dojoサーバーサイド(Go)編です。  
</br>

備わっている機能としては、Userの作成、取得、更新ができます。Gachaでは、Characterを作成しそれらをランダムに取得でき、取得したCharacter一覧を表示することができます。
</br>

## Demo
1. Dockerを起動します。
```
$ docker-compose up
```
</br>

2. APIサーバーを起動します。
```
$ ./CA-Tech-Dojo-Go
```
</br>

3. BashでUsageを使用する
</br>

## Usage
### ユーザー情報作成API
ユーザー情報を作成します。  
ユーザーの名前情報をリクエストで受け取り、ユーザーIDをデータベースに登録します。  
レスポンスとして認証トークンを与えます。
```
curl -X POST http://localhost:8080/user/create -H "accept: application/json" -H "Content-Type: application/json" -d "{“name”: string}”
```
</br>

### ユーザー情報取得API
ユーザー情報を取得します。  
ユーザーの認証トークンを読み取り、データベースに照会します。  
```
curl -H "Authorization: Bearer Your_Token”  localhost:8080/user/get
```
</br>

### ユーザー情報更新API
ユーザー情報を更新します。  
名前の更新を行います。  
```
curl -X PUT "http://localhost:8080/user/update" -H "Authorization: Bearer  Yout_Token” -H "Content-Type: application/json" -d "{ “name”: string }”
```
</br>

### キャラクター作成API
キャラクターを作成します。  
キャラクターの名前情報をリクエストで受け取り、キャラクターIDをデータベースに登録します。
```
curl -X POST http://localhost:8080/character/create -H 'Content-Type: application/json' -d '{"name": string}’
```
</br>

### ガチャ実行API
ガチャを引いてキャラクターを取得する処理をします。  
獲得したキャラクターはユーザー所持キャラクターテーブルに保存されます。  
同じキャラクターでも複数所持できるようになっています。  
```
curl -X POST http://localhost:8080/user/gacha -H 'Content-Type: application/json' -d '{"times": int}’ -H "Authorization: Bearer Your_Token”
```
</br>

### キャラクター一覧取得API
ユーザーが保持しているキャラクター一覧を取得し表示します。
```
curl -X GET http://localhost:8080/user/list -H "Authorization: Bearer Your_Token”
```
</br>

## Technology
言語 : Golang  
フレームワーク : Echo, Gorm  
トークン : Json Web Token  
アーキテクチャ : ドメイン駆動設計