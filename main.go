package main

import (
	"fmt"
)

func main() {
	// Data Uji
	// Data 1
	lumbaLumba := []float64{96, 108, 89}
	koala := []float64{88, 91, 110}

	// Hitung skor rata-rata
	avgLumbaLumba := calculateAverage(lumbaLumba)
	avgKoala := calculateAverage(koala)

	// Cetak skor rata-rata
	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", avgLumbaLumba)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", avgKoala)

	// Bandingkan skor rata-rata untuk menentukan pemenang
	if avgLumbaLumba > avgKoala && avgLumbaLumba >= 100 {
		fmt.Println("Tim Lumba-lumba memenangkan kompetisi!")
	} else if avgKoala > avgLumbaLumba && avgKoala >= 100 {
		fmt.Println("Tim Koala memenangkan kompetisi!")
	} else if avgLumbaLumba == avgKoala && avgLumbaLumba >= 100 && avgKoala >= 100 {
		fmt.Println("Hasil seri!")
	} else {
		fmt.Println("Tidak ada pemenang. Skor rata-rata kurang dari 100.")
	}

	// Data Bonus 1
	lumbaLumbaBonus1 := []float64{97, 112, 101}
	koalaBonus1 := []float64{109, 95, 123}

	// Hitung skor rata-rata
	avgLumbaLumbaBonus1 := calculateAverage(lumbaLumbaBonus1)
	avgKoalaBonus1 := calculateAverage(koalaBonus1)

	// Cetak skor rata-rata
	fmt.Printf("\nSkor rata-rata Tim Lumba-lumba (Bonus 1): %.2f\n", avgLumbaLumbaBonus1)
	fmt.Printf("Skor rata-rata Tim Koala (Bonus 1): %.2f\n", avgKoalaBonus1)

	// Bandingkan skor rata-rata untuk menentukan pemenang
	if avgLumbaLumbaBonus1 > avgKoalaBonus1 && avgLumbaLumbaBonus1 >= 100 {
		fmt.Println("Tim Lumba-lumba memenangkan kompetisi (Bonus 1)!")
	} else if avgKoalaBonus1 > avgLumbaLumbaBonus1 && avgKoalaBonus1 >= 100 {
		fmt.Println("Tim Koala memenangkan kompetisi (Bonus 1)!")
	} else if avgLumbaLumbaBonus1 == avgKoalaBonus1 && avgLumbaLumbaBonus1 >= 100 && avgKoalaBonus1 >= 100 {
		fmt.Println("Hasil seri! (Bonus 1)")
	} else {
		fmt.Println("Tidak ada pemenang. Skor rata-rata kurang dari 100. (Bonus 1)")
	}

	// Data Bonus 2
	lumbaLumbaBonus2 := []float64{97, 112, 101}
	koalaBonus2 := []float64{109, 95, 106}

	// Hitung skor rata-rata
	avgLumbaLumbaBonus2 := calculateAverage(lumbaLumbaBonus2)
	avgKoalaBonus2 := calculateAverage(koalaBonus2)

	// Cetak skor rata-rata
	fmt.Printf("\nSkor rata-rata Tim Lumba-lumba (Bonus 2): %.2f\n", avgLumbaLumbaBonus2)
	fmt.Printf("Skor rata-rata Tim Koala (Bonus 2): %.2f\n", avgKoalaBonus2)

	// Bandingkan skor rata-rata untuk menentukan pemenang
	if avgLumbaLumbaBonus2 > avgKoalaBonus2 && avgLumbaLumbaBonus2 >= 100 {
		fmt.Println("Tim Lumba-lumba memenangkan kompetisi (Bonus 2)!")
	} else if avgKoalaBonus2 > avgLumbaLumbaBonus2 && avgKoalaBonus2 >= 100 {
		fmt.Println("Tim Koala memenangkan kompetisi (Bonus 2)!")
	} else if avgLumbaLumbaBonus2 == avgKoalaBonus2 && avgLumbaLumbaBonus2 >= 100 && avgKoalaBonus2 >= 100 {
		fmt.Println("Hasil seri! (Bonus 2)")
	} else {
		fmt.Println("Tidak ada pemenang. Skor rata-rata kurang dari 100. (Bonus 2)")
	}

	fmt.Println("")

	Tugas2()
}

// Fungsi untuk menghitung skor rata-rata
func calculateAverage(scores []float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}
