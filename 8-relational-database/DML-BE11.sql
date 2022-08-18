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
VALUES ("DM-0002",7, 3, "12A", "09:00", "12:00");
INSERT INTO data_mengajar (id, guru_id, mata_pelajaran_id, kelas, jam_mulai, jam_selesai)
VALUES ("DM-0003",7, 4, "12B", "07:00", "09:00");
INSERT INTO data_mengajar (id, guru_id, mata_pelajaran_id, kelas, jam_mulai, jam_selesai)
VALUES ("DM-0004",8, 5, "11A", "07:00", "09:00");
INSERT INTO data_mengajar (id, guru_id, mata_pelajaran_id, kelas, jam_mulai, jam_selesai)
VALUES ("DM-0005",8, 5, "10A", "13:00", "15:00");

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

-- inner join 
-- menampilkan list jadwal seluruh guru yang sudah punya jadwal mengajar
SELECT g.id, g.nama, g.nip, dm.kelas, dm.jam_mulai, dm.jam_selesai from guru g
INNER JOIN data_mengajar dm ON g.id = dm.guru_id;

SELECT g.id, g.nama, g.nip, dm.kelas, dm.jam_mulai, dm.jam_selesai, mp.nama_pelajaran from guru g
INNER JOIN data_mengajar dm ON g.id = dm.guru_id
INNER JOIN mata_pelajaran mp ON dm.mata_pelajaran_id = mp.id
where dm.kelas like "%A"
;

SELECT g.id, g.nama, g.nip, mp.nama_pelajaran from guru g
INNER JOIN data_mengajar dm ON g.id = dm.guru_id
INNER JOIN mata_pelajaran mp ON dm.mata_pelajaran_id = mp.id
;

-- menampilkan list seluruh guru, baik yang punya jadwal maupun yang belum
-- left join
SELECT g.id, g.nama, dm.kelas, dm.jam_mulai from guru g
LEFT JOIN data_mengajar dm ON g.id = dm.guru_id;

SELECT g.id, g.nama, g.nip, dm.kelas, dm.jam_mulai, dm.jam_selesai from data_mengajar dm
LEFT JOIN guru g ON g.id = dm.guru_id;

-- right join
SELECT g.id, g.nama, g.nip, dm.kelas, dm.jam_mulai, dm.jam_selesai from guru g
RIGHT JOIN data_mengajar dm ON g.id = dm.guru_id
WHERE dm.jam_mulai = "07:00";

SELECT g.id, g.nama, g.nip, dm.kelas, dm.jam_mulai, dm.jam_selesai from data_mengajar dm
RIGHT JOIN guru g ON g.id = dm.guru_id;

-- AGGREGATION
SELECT * FROM mata_pelajaran;

SELECT min(kkm) FROM mata_pelajaran where id between 1 and 4;
SELECT MAX(kkm) FROM mata_pelajaran;

SELECT SUM(kkm) as jumlah_kkm FROM mata_pelajaran;

-- SUM with GROUP BY
select dm.guru_id, guru.nama, SUM(mp.kkm) as jumlah_kkm FROM data_mengajar dm
inner join mata_pelajaran mp on dm.mata_pelajaran_id = mp.id
inner join guru ON dm.guru_id = guru.id
group by dm.guru_id;

SELECT guru.id, guru.nama, SUM(mp.kkm) as jumlah_kkm FROM guru
INNER JOIN data_mengajar dm ON guru.id = dm.guru_id
INNER JOIN mata_pelajaran mp ON dm.mata_pelajaran_id = mp.id
WHERE guru.id BETWEEN 5 AND 10
GROUP BY guru.id;


SELECT AVG(kkm) as avg_kkm FROM mata_pelajaran;

SELECT COUNT(id) as count_kkm FROM mata_pelajaran;
SELECT COUNT(id) as count_kkm, count(kkm) FROM mata_pelajaran where kkm >= 60;


SELECT g.id, g.nama, g.nip, dm.kelas, dm.jam_mulai, dm.jam_selesai, mp.nama_pelajaran, mp.kkm from guru g
INNER JOIN data_mengajar dm ON g.id = dm.guru_id
INNER JOIN mata_pelajaran mp ON dm.mata_pelajaran_id = mp.id
;

-- HAVING
select dm.guru_id, guru.nama FROM data_mengajar dm
inner join mata_pelajaran mp on dm.mata_pelajaran_id = mp.id
inner join guru ON dm.guru_id = guru.id
group by dm.guru_id
having sum(mp.kkm) >= 100;

select guru.id, guru.nama FROM guru
inner join data_mengajar dm ON guru.id = dm.guru_id
group by guru.id
having count(mata_pelajaran_id) >= 2;







