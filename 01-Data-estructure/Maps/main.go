package main

import "fmt"

func main() {
	//Los mapas son elementos desordenados:

	//Se define un map con clave del tipo entero y los valores en string:
	allDays := make(map[int]string)

	fmt.Println(allDays) //map[]

	allDays[1] = "Domingo"
	allDays[2] = "Lunes"
	allDays[3] = "Martes"
	allDays[4] = "Miercoles"
	allDays[5] = "Jueves"
	allDays[6] = "Viernes"
	allDays[7] = "Sabado"

	fmt.Println(allDays) //map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:Sabado]

	//Para borrar un elemento del mapa se usa "delete":

	delete(allDays, 1)
	fmt.Println(allDays) //map[2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:Sabado]

	//Mapas un poco más complejos:

	students := make(map[string][]int)

	students["Francisco"] = []int{1, 22, 75}
	students["Molina"] = []int{92, 10}

	fmt.Println(students) //map[Francisco:[1 22 75] Molina:[92 10]

	//Para ver un dato espefico de ese map:

	fmt.Println(students["Francisco"][1]) //22

}
