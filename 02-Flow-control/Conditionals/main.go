package main

import "fmt"

/** App para restaurante
* Descuentos de 10% hastas 100 de consumo
* Descuentos de 20% hastas 200 de consumo
* Descuentos de 30% mayor 200 de consumo
* igv 19%
 */
func main() {

	var consumo, descuento float64
	var datosDescuento string

	//Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {

		if consumo == 0 {
			fmt.Println("AVISO: El consumo debe ser mayor que 0.")

		}

		if consumo <= 100 { //Descuento del 10%:
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 { //Descuento del 20%:
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 { //Descuento del 30%:
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}

	} else {
		fmt.Println("El consumo debe ser mayor que 0.")
	}

	resultadoDelDescuento := consumo - descuento
	igv := resultadoDelDescuento * 0.19
	total := resultadoDelDescuento + igv
	fmt.Println("****** Facturador ******")
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento que aplica: ", datosDescuento)
	fmt.Println("Descuento: ", descuento)
	fmt.Println("Monto con descuento: ", resultadoDelDescuento)
	fmt.Println("IGV: ", igv)
	fmt.Println("El total a pagar es: ", total)
	fmt.Println("****** Facturador ******")

}
