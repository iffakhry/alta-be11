package main

import "fmt"

type Student struct {
	Nama         string
	JenisKelamin string
	NomorInduk   string
	Jurusan      string
	Telp         []string
	Alamat       Address
}

type Address struct {
	Jalan string
	Nomor int
	Kota  string
}

func main() {
	var student1 Student
	student1.Nama = "Adi"
	student1.NomorInduk = "NIS012345"
	student1.Telp = append(student1.Telp, "0812345")
	student1.Telp = append(student1.Telp, "0823456789")
	student1.Telp = append(student1.Telp, "0823")
	fmt.Println("student1:", student1)
	fmt.Println("student1:", student1.Telp[1])
	for _, value := range student1.Telp {
		fmt.Println(value)
	}

	//cara1
	var alamat1 Address
	alamat1.Jalan = "Jl Perjuangan"
	alamat1.Nomor = 100
	alamat1.Kota = "Jakarta"
	student1.Alamat = alamat1

	//cara2
	student1.Alamat = Address{
		Jalan: "jl lurus",
		Kota:  "Surabaya",
		Nomor: 123,
	}
	fmt.Println("student1:", student1)
	fmt.Println("nama jalan student1:", student1.Alamat.Jalan)

	var students []Student
	students = append(students, Student{
		Nama: "Jono",
	})
	students = append(students, Student{
		Nama: "Toni",
	})
	fmt.Println(students[0].Nama)

	var mapstudents map[string]Student
	mapstudents["tono"] = Student{
		Nama: "tono toni",
	}

}
