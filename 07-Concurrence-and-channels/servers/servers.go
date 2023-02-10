package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(name string, channel chan string) {
	_, error := http.Get(name)

	if error != nil {
		// fmt.Println(name, "no esta funciondo :c")
		channel <- name + " no esta funciondo :c"
	} else {
		// fmt.Println(name, "esta activo.")
		channel <- name + " esta activo :D"
	}

}

func main() {
	start := time.Now()
	channel := make(chan string)
	servidores := []string{
		"https://www.google.com.ar/",
		"https://www.mercadolibre.com.ar/",
		"https://www.pedidosya.com.ar/",
	}

	for _, server := range servidores {
		go checkServer(server, channel)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-channel)
	}

	ending := time.Since(start)

	fmt.Println("Se ejecuto en: ", ending)
}
