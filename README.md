
```
% docker-compose exec mysql bash
# mysql -u root -ppassword
> create database if not exists m1_db;
> use m1_db;
>
create table m1_db.user (
  id int AUTO_INCREMENT,
  name varchar(10),
  PRIMARY KEY (id)
);

insert into user (name)
values ('1'),('2'),('3'),('4'),('5'),('6'),('7'),('8'),('9'),('10');

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

```
% docker-compose exec mysql bash
> mysql -uroot -ppassword m1_db < 001.sql
```

- add column
  - https://stackoverflow.com/questions/24571611/mysql-alter-table-if-column-not-exists
  - https://took.jp/post-660/
  - https://qiita.com/jas/items/0ffe72a96c3f8b6cf6dc
- 1万件のデータ作成 (直積)
  - https://qiita.com/cobot00/items/8d59e0734314a88d74c7

## Golang
- コンテナに入る

```
% docker-compose exec app sh
```

- main.goを実行

```
% docker-compose exec app go run main.go
```
