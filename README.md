
```
% docker-compose exec mysql bash
# mysql -u root -ppassword
> create database if not exists m1_db;
> use m1_db;
> create table m1_db.user (id int, name varchar(10));
```

```
% docker-compose exec mysql bash
> mysql -uroot -ppassword m1_db < 001.sql
```

- add column
  - https://stackoverflow.com/questions/24571611/mysql-alter-table-if-column-not-exists
  - https://took.jp/post-660/
  - https://qiita.com/jas/items/0ffe72a96c3f8b6cf6dc