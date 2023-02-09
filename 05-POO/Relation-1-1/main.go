package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}
type Student struct {
	user   User
	codigo string
}

//Relación de uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	//Relacion de uno a uno
	/*
		Fran := User{
			nombre: "Fran",
			email:  "Fran@gmail.com",
			activo: true,
		}
		Molina := User{
			nombre: "Molina",
			email:  "Molina@gmail.com",
			activo: true,
		}
		FranStudent := Student{
			user:   Fran,
			codigo: "001arsd",
		}
		fmt.Println(Molina)
		fmt.Println(FranStudent.user.nombre)
	*/

	//Relacion de uno a muchus
	video1 := Video{titulo: "01-Introducción"}
	video2 := Video{titulo: "02-Instalación"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
