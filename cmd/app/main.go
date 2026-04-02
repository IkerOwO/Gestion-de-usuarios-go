package main

import (
	"fmt"
	"os"
	"time"
	"GestionDeUsuarios/cmd/internal"
)

/*
	Para el User Input se usa scan, scanf y scanln

	scan: Recive el input con formato raw como valores separados con espacios. Las NewLines cuentan como espacios

	scanf: Recive el input y lo guarda basado en determinados formatos 

	scanln: Similar a scan pero para de escanear cuando se le da al enter
*/

func printMenu(){
	fmt.Println("1. Crear Usuario")
	fmt.Println("2. Eliminar Usuario")
	fmt.Println("3. Actualizar Usuario")
	fmt.Println("4. Buscar Usuario")
	fmt.Println("5. Listar Usuarios")
	fmt.Println("6. Salir")
}

func crearUsuarios(){
	// Declaramos el usuario
	var( 
		usuario internal.Usuario
		nombreUsuario string
		contraseña string
		fechaAlta string
	)

	// Pedimos el nombre de usuario
	fmt.Println("Introduce el nombre de usuario: ")
	fmt.Scan(&nombreUsuario)
	// Pedimos la contraseña
	fmt.Println("Introduce la contraseña: ")
	fmt.Scan(&contraseña)
	// Pedimos la fecha de alta
	fmt.Println("Introduce la fecha de alta: ")
	fmt.Scan(&fechaAlta)

	// Guardamos el usuario
	usuario.NombreUsuario = nombreUsuario
	usuario.Contraseña = contraseña
	usuario.FechaAlta = fechaAlta
	usuario.Baneado = false

	// Guardamos el usuario en el array
	internal.UsuariosArray = append(internal.UsuariosArray, usuario)
	
}


func main(){
	var(
		opc int

	)

	for{
		// Limpiar terminal
		fmt.Print("\033[H\033[2J")

		fmt.Println("Gestion de Usuarios")
		printMenu()
		fmt.Println("Elige una opcion: ")
		fmt.Scan(&opc)
	
		switch opc{
			case 1:
				crearUsuarios()
			case 2:

			case 3:

			case 4:

			case 5:

			case 6:
				fmt.Println("Saliendo del programa...")
				time.Sleep(5 * time.Second)
				// Salir del programa
				os.Exit(0)
		}

	}
}