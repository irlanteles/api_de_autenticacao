package main

import (
	"fmt"
    "log"
    "net/http"
     "os"
   

    "github.com/sema/apiAuth/db"
    "github.com/sema/apiAuth/internal/response"
    
)



func main() {
	err := db.Connect()
    if err != nil {
        log.Fatalf("Erro ao conectar no banco: %v", err)
    }
    fmt.Println("Conexão com o banco bem-sucedida!")

    http.HandleFunc("/login", response.LoginResponse)
    http.HandleFunc("/cadastro", response.CadastroResponse)
    http.HandleFunc("/recuperarsenha", response.RecuperarSenhaResponse)
    
    port := os.Getenv("API_PORT")
    addr := fmt.Sprintf(":%s", port)
    log.Println("Servidor rodando na porta "+ addr)
    log.Fatal(http.ListenAndServe(addr, nil))
    

   
}
