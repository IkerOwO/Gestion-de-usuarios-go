package internal

import (
	"fmt"
	"time"
)

// Estructura de un usuario
type Usuario struct {
	NombreUsuario string
	Contraseña    string
	FechaAlta     string
	Baneado       bool
}

// Array de usuarios
var UsuariosArray []Usuario

// Crear Usuarios
func CrearUsuario(nombreUsuario string, contraseña string, fechaAlta string){
	// Guardamos el usuario
	nuevoUsuario := Usuario{
		NombreUsuario: nombreUsuario,
		Contraseña:    contraseña,
		FechaAlta:     fechaAlta,
		Baneado:       false,
	}

	// Guardamos el usuario en el array
	UsuariosArray = append(UsuariosArray, nuevoUsuario)
	
	fmt.Print("Usuario creado")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
}

// Funcion eliminar usuarios
func EliminarUsuario(usuario string, usuarios []Usuario) bool {
	for i,u := range usuarios{
		if u.NombreUsuario == usuario{
			// Eliminamos el usuario
			UsuariosArray = append(UsuariosArray[:i], UsuariosArray[i+1:]...)
			fmt.Println("Usuario eliminado")
			time.Sleep(2 * time.Second)
			fmt.Print("\033[H\033[2J")
			return true
		}
	}
	return false
}

// Funcion actualizar usuarios
func ActualizarUsuario(usuario string, contraseña string, fechaAlta string, usuarios []Usuario) bool {
	for i, u := range usuarios {
		if u.NombreUsuario == usuario {
			usuarios[i].Contraseña = contraseña
			usuarios[i].FechaAlta = fechaAlta
			fmt.Println("Usuario Actualizado")
			time.Sleep(2 * time.Second)
			fmt.Print("\033[H\033[2J")
			return true
		}
	}
	return false
}

// Funcion banear usuarios
func BanUser(usuario string, usuarios []Usuario) bool {
	isBaneado := false
	for i, u := range usuarios {
		if u.NombreUsuario == usuario {
			usuarios[i].Baneado = true
			isBaneado = true
			break
		}
	}
	return isBaneado
}

// Funcion comprobar si un usuario esta baneado
func IsUsuarioBaneado(usuario string, usuarios []Usuario) bool {
	isBaneado := false
	for _, u := range usuarios {
		if u.NombreUsuario == usuario {
			isBaneado = u.Baneado
			break
		}
	}
	return isBaneado
}

// Funcion listar usuarios