package internal

// Estructura de un usuario
type Usuario struct {
	NombreUsuario string
	Contraseña    string
	FechaAlta     string
	Baneado       bool
}

// Array de usuarios
var UsuariosArray []Usuario

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