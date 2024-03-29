## APIサーバー起動
- コンテナ内起動時にサーバーも起動する
- `http://localhost:8000`

## Golang
- コンテナに入る

```
% docker-compose exec app sh
```

- main.goを実行

```
% docker-compose exec app go run main.go
```

- test実行

```
% go test ./utils/ -covermode=count -coverprofile=coverage.out -v
% go tool cover -html=coverage.out -o coverage.html
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
