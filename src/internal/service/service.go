package service

import(
	"os"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"github.com/sema/apiAuth/db"
	"github.com/sema/apiAuth/internal/models"
	"golang.org/x/crypto/bcrypt"
	

)

func GenerateJWT(username string) (string, error) {
	// Lê chave privada do arquivo
	// privateKeyData, err := os.ReadFile("/apiAuth/src/keys/private.pem")
	privateKeyData, err := os.ReadFile("./keys/private.pem")
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return "", err
	}
	
	rows , err := db.DB.Query("SELECT sistema,permissao FROM usuario.vw_credencial WHERE email = $1", username)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	// 	type UserPermission struct {
	// 	Sistema   string `json:"sistema"`
	// 	Permissao string `json:"permissao"`
	// }

	var permission []models.UserPermission
	for rows.Next() {
		var userPermission models.UserPermission
		if err := rows.Scan(&userPermission.Sistema, &userPermission.Permissao); err != nil {
			return "", err
		}
		permission = append(permission, userPermission)
	}
	

	// Define payload (claims)
	claims := jwt.MapClaims{
		"username": username,
		"permission": permission,                    // Identificador do usuário
		"exp": time.Now().Add(time.Hour).Unix(), // Expira em 1 hora
		"iat": time.Now().Unix(),            // Criado agora
	}

	// Cria o token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Assina com a chave privada
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func CheckPassword(password, passwordHash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}