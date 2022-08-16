-- membuat database
CREATE DATABASE db_be11;

-- menghapus db
DROP DATABASE db_be11;

-- memilih db yang akan digunakan
USE db_be11;

show tables;

CREATE TABLE guru(
id INT PRIMARY KEY NOT NULL auto_increment,
nama varchar(50) not null,
nip varchar(100)
);

create table mata_pelajaran(
id int not null auto_increment,
nama_pelajaran varchar(200),
kkm int,
primary key(id)
);

create table data_mengajar(
id varchar(30) primary key not null,
guru_id int,
mata_pelajaran_id int,
kelas varchar(20),
jam_mulai varchar(15),
jam_selesai varchar(15),
CONSTRAINT fk_datamengajar_guru FOREIGN KEY (guru_id) REFERENCES guru(id),
CONSTRAINT fk_datamengajar_matpel FOREIGN KEY (mata_pelajaran_id) REFERENCES mata_pelajaran(id)
);

-- menambah column
alter table dummy
add column created_at datetime; 

-- menambah column dengan default value
alter table dummy
add column updated_at datetime default current_timestamp; 

-- menambah foreign key
alter table dummy 
add column mata_pelajaran_id int;

alter table dummy
ADD CONSTRAINT fk_dummy_matpel FOREIGN KEY (mata_pelajaran_id) references mata_pelajaran(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- menghapus table
DROP TABLE table_name;

-- menghapus foreign key
ALTER TABLE dummy
DROP FOREIGN KEY fk_dummy_matpel;


show columns from dummy;

-- Error Code: 1452. Cannot add or update a child row: a foreign key constraint fails (`db_be11`.`#sql-85_74`, CONSTRAINT `fk_dummy_matpel` FOREIGN KEY (`mata_pelajaran_id`) REFERENCES `mata_pelajaran` (`id`))





