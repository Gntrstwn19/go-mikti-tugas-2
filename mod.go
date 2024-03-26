package main

import (
	"fmt"
)

func main() {
	// Data Uji
	// Data 1
	beratMark1 := 78.0  // kg
	tinggiMark1 := 1.69 // m
	beratJohn1 := 92.0  // kg
	tinggiJohn1 := 1.95 // m

	// Hitung BMI
	bmiMark1 := beratMark1 / (tinggiMark1 * tinggiMark1)
	bmiJohn1 := beratJohn1 / (tinggiJohn1 * tinggiJohn1)

	// Variabel Boolean 'markHigherBMI'
	markHigherBMI1 := bmiMark1 > bmiJohn1

	// Cetak hasil
	fmt.Printf("BMI Mark (Data 1): %.2f\n", bmiMark1)
	fmt.Printf("BMI John (Data 1): %.2f\n", bmiJohn1)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John? ", markHigherBMI1)

	// Data 2
	beratMark2 := 95.0  // kg
	tinggiMark2 := 1.88 // m
	beratJohn2 := 85.0  // kg
	tinggiJohn2 := 1.76 // m

	// Hitung BMI
	bmiMark2 := beratMark2 / (tinggiMark2 * tinggiMark2)
	bmiJohn2 := beratJohn2 / (tinggiJohn2 * tinggiJohn2)

	// Variabel Boolean 'markHigherBMI'
	markHigherBMI2 := bmiMark2 > bmiJohn2

	// Cetak hasil
	fmt.Printf("\nBMI Mark (Data 2): %.2f\n", bmiMark2)
	fmt.Printf("BMI John (Data 2): %.2f\n", bmiJohn2)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John? ", markHigherBMI2)
}
