create table if not exists user (id int, name varchar(10));
create table if not exists post (id int, content varchar(255));

DROP PROCEDURE IF EXISTS alter_table_procedure;
DELIMITER //
CREATE PROCEDURE alter_table_procedure()
BEGIN
  /* SQLEXCEPTIONを無視するように設定 */
  DECLARE CONTINUE HANDLER FOR SQLEXCEPTION BEGIN END;
  /* 以下のALTER TABLEで『Duplicate column name』エラーが発生してもプロシージャは正常終了する */
  ALTER TABLE user ADD COLUMN age INT NULL;
--   ALTER TABLE user ADD COLUMN address VARCHAR(50) NULL;
END //
DELIMITER ;
CALL alter_table_procedure();
DROP PROCEDURE alter_table_procedure;
/* 追加対象のカラムをSELECTすることで、もしカラム追加できていなかったらエラーとなり
   Flywayのhistoryはfailedとなる */
SELECT age, address FROM user;