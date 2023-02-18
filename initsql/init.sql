create table if not exists users (
  id int AUTO_INCREMENT,
  name varchar(10),
  PRIMARY KEY (id)
);
create table if not exists post (id int, content varchar(255));

insert into users (name) values ("user-1");
insert into users (name) values ("user-2");
insert into users (name) values ("user-3");
insert into users (name) values ("user-4");
insert into users (name) values ("user-5");

DROP PROCEDURE IF EXISTS alter_table_procedure;
DELIMITER //
CREATE PROCEDURE alter_table_procedure()
BEGIN
  /* SQLEXCEPTIONを無視するように設定 */
  DECLARE CONTINUE HANDLER FOR SQLEXCEPTION BEGIN END;
  /* 以下のALTER TABLEで『Duplicate column name』エラーが発生してもプロシージャは正常終了する */
  ALTER TABLE users ADD COLUMN age INT NULL;
  ALTER TABLE users ADD COLUMN address VARCHAR(50) NULL;
END //
DELIMITER ;
CALL alter_table_procedure();
DROP PROCEDURE alter_table_procedure;
/* 追加対象のカラムをSELECTすることで、もしカラム追加できていなかったらエラーとなりFlywayのhistoryはfailedとなる */
SELECT age, address FROM users;
