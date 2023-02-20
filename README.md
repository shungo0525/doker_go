## APIサーバー起動
- コンテナ内で下記のコマンドを実行
- `http://localhost:8000`

```
go run api/main.go 
```

## Golang
- コンテナに入る

```
% docker-compose exec app sh
```

- main.goを実行

```
% docker-compose exec app go run main.go
```


## Mysqlコンテナ
```
% docker-compose exec mysql bash
# mysql -u root -ppassword
```

```
% docker-compose exec mysql bash
> mysql -uroot -ppassword development < init.sql
```


## その他
```
INSERT INTO user(name)
SELECT
  CONCAT('NAME' , @rownum := @rownum + 1)
FROM
  user AS s1,
  user AS s2,
  user AS s3,
  user AS s4,
  (SELECT @rownum := 0) AS v;
```

## ref
- add column
  - https://stackoverflow.com/questions/24571611/mysql-alter-table-if-column-not-exists
  - https://took.jp/post-660/
  - https://qiita.com/jas/items/0ffe72a96c3f8b6cf6dc
- 1万件のデータ作成 (直積)
  - https://qiita.com/cobot00/items/8d59e0734314a88d74c7
