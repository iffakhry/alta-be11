package score

func KonversiNilai(nilai int) string {
	var result string
	if nilai >= 80 && nilai <= 100 {
		result = "Nilai A"
	} else if nilai >= 65 && nilai <= 79 {
		result = "Nilai B+"
	} else if nilai >= 50 && nilai <= 64 {
		result = "Nilai B"
	} else if nilai >= 35 && nilai <= 49 {
		result = "Nilai C"
	} else if nilai >= 0 && nilai <= 34 {
		result = "Nilai D"
	} else {
		result = "Nilai harus antara 0 - 100"
	}
	return result
}
