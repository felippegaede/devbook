package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Verificar senha compara uma senha e um hash  e retorna se elas tem o mesmo valor
func VerificarSenha(senhaString, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaString), []byte(senhaHash))
}
