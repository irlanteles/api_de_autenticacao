package models 

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CadastroRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RecuperarSenhaRequest struct{
	Email    string `json:"email"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`	
}

type UserPermission struct {
	Sistema   string `json:"sistema"`
	Permissao string `json:"permissao"`
}
