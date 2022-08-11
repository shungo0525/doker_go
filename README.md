
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
