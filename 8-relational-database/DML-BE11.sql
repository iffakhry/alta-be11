-- insert data
INSERT INTO guru (nama, nip)
VALUES ("Budi", "NIP0001");
INSERT INTO guru (nama, nip)
VALUES ("Adi", "NIP0002");
INSERT INTO guru (nama, nip)
VALUES ("Cindy", "NIP0004");

insert into mata_pelajaran (nama_pelajaran, kkm)
VALUES ("Programming", 75);
insert into mata_pelajaran (nama_pelajaran, kkm)
VALUES ("Statistika", 75);
insert into mata_pelajaran (nama_pelajaran, kkm)
VALUES ("Menggambar", 75);

INSERT INTO data_mengajar (id, guru_id, mata_pelajaran_id, kelas, jam_mulai, jam_selesai)
VALUES ("DM-0001",4, 1, "10A", "07:00", "09:00");
INSERT INTO data_mengajar (id, guru_id, mata_pelajaran_id, kelas, jam_mulai, jam_selesai)
VALUES ("DM-0004",6, 3, "12A", "09:00", "12:00");

-- membaca data
SELECT * FROM guru;
select * from mata_pelajaran;
select * from data_mengajar;

SELECT guru_id, mata_pelajaran_id, kelas FROM data_mengajar;

-- WHERE clause
SELECT * FROM data_mengajar WHERE id = "DM-0001";

SELECT * FROM data_mengajar WHERE kelas = "11B";

SELECT * FROM data_mengajar WHERE guru_id = 1;

SELECT kelas FROM data_mengajar WHERE jam_mulai IS NOT NULL;

SELECT kelas FROM data_mengajar WHERE jam_mulai IS NULL;

-- update data
UPDATE guru SET nama = "Adi" where nip = "NIP0002";

UPDATE guru SET nama = "Aldi", nip = "nip0001a" where id = 1;

-- delete data
DELETE FROM guru WHERE id = 3;

DELETE FROM guru WHERE id = 6;
DELETE FROM data_mengajar WHERE id = "DM-0004";
-- Error Code: 1451. Cannot delete or update a parent row: a foreign key constraint fails (`db_be11`.`data_mengajar`, CONSTRAINT `fk_datamengajar_guru` FOREIGN KEY (`guru_id`) REFERENCES `guru` (`id`))

-- Like
SELECT * FROM guru WHERE nama LIKE "%i%";

-- between
SELECT * FROM guru WHERE id between 4 and 8;

-- AND  OR
SELECT * FROM guru WHERE id between 4 and 7 AND nama LIKE "%y";
SELECT * FROM guru WHERE id between 4 and 7 OR nama LIKE "%y";

-- ORDER BY ASC/DESC
SELECT * FROM guru WHERE id between 4 and 7 ORDER BY nip ASC;
SELECT * FROM guru WHERE id between 4 and 7 ORDER BY nip DESC;

-- limit
SELECT * FROM guru ORDER BY nip DESC LIMIT 2;

-- Alias
SELECT g.id, g.nama, g.nip from guru g;
SELECT guru.id, guru.nama, guru.nip from guru;
SELECT g.id, g.nama AS nama_guru, g.nip from guru g;

-- join 
SELECT g.id, g.nama, g.nip, dm.kelas from guru g
INNER JOIN data_mengajar dm ON g.id = dm.guru_id;