package entities

type Guru struct {
	ID     int
	Nama   string
	Nip    string
	Status string
}

// select menggunakan join
type GuruResponse struct {
	ID    int
	Nama  string
	Nip   string
	Kelas string
}
