package main

import "fmt"

//Creo la Struct, que seria la "clase":

type Character struct {
	//Atributos:
	name string
	age  int
}

//Metodos:

//Mostrar datos:
func (c *Character) printData() {
	fmt.Printf("El nombre es: %s y la edad es: %d \n", c.name, c.age)
}

//Cambiar nombre:
func (c *Character) changeName(data string) {
	c.name = data
	fmt.Printf("El nombre se cambio por: %s \n", c.name)
}

//Herencia:

type Client struct {
	Character
	payd float32
}

//Cambiar ver info de Cliente:
func (c *Client) printInfo() {
	fmt.Printf("El nombre es: %s, su edad es: %d y tiene: %v \n", c.name, c.age, c.payd)
}

func main() {
	oneCharacter := Character{"Francisco", 30}

	// fmt.Println(oneCharacter)
	oneCharacter.printData()

	secondCharacter := Character{name: "Fran", age: 20}

	// fmt.Println(secondCharacter)
	secondCharacter.changeName("Moli")
	secondCharacter.printData()

	firstClient := Client{payd: 155.75}
	firstClient.name = "Elon"
	firstClient.age = 50
	firstClient.printInfo()

}
