package main

import (
	"fmt"
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
	fmt.Println("4. Banear Usuario")
	fmt.Println("5. Comprobar Baneo de Usuario")
	fmt.Println("6. Listar Usuarios")
	fmt.Println("7. Salir")
}

func main(){
	var(
		opc int
	)

	for {
		// Limpiar terminal
		fmt.Print("\033[H\033[2J")

		fmt.Println("Gestion de Usuarios")
		printMenu()
		fmt.Print("Elige una opcion: ")
		
		// Comprobamos si hay algún error al leer el input (e.g. meten una letra en lugar de un número)
		_, err := fmt.Scan(&opc)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, introduce un número.")
			// Limpiamos el buffer para evitar un bucle infinito
			var discard string
			fmt.Scanln(&discard)
			time.Sleep(2 * time.Second)
			continue
		}
	
		switch opc {
			case 1:
				var(
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

				internal.CrearUsuario(nombreUsuario, contraseña, fechaAlta)
			case 2:
				var(
					nombreUsuario string
				)
				fmt.Println("Introduce el nombre de usuario: ")
				fmt.Scan(&nombreUsuario)
				if internal.EliminarUsuario(nombreUsuario, internal.UsuariosArray){
					fmt.Println("Usuario eliminado correctamente")
				}else{
					fmt.Println("Usuario no encontrado")
				}
				time.Sleep(2 * time.Second)
			case 3:
				var(
					nombreUsuario string
					contraseña string
					fechaAlta string
				)
				fmt.Println("Introduce el nombre de usuario a actualizar: ")
				fmt.Scan(&nombreUsuario)
				fmt.Println("Introduce la nueva contraseña: ")
				fmt.Scan(&contraseña)
				fmt.Println("Introduce la nueva fecha de alta: ")
				fmt.Scan(&fechaAlta)
				if internal.ActualizarUsuario(nombreUsuario, contraseña, fechaAlta, internal.UsuariosArray){
					fmt.Println("Usuario actualizado correctamente")
				}else{
					fmt.Println("Usuario no encontrado")
				}
				time.Sleep(2 * time.Second)
			case 4:

			case 5:

			case 6:

			case 7:
				fmt.Println("Saliendo del programa...")
				time.Sleep(2 * time.Second)
				return
			default:
				fmt.Println("Opción no válida. Inténtalo de nuevo.")
				time.Sleep(2 * time.Second)
		}
	}
}