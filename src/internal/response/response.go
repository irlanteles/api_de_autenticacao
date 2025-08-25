package response

import (
	"encoding/json"
	"net/http"
	"github.com/sema/apiAuth/internal/models"
	"github.com/sema/apiAuth/internal/repository"
	"fmt"
	
	"github.com/sema/apiAuth/internal/service"
	
)


func LoginResponse(response http.ResponseWriter, request *http.Request){
	if request.Method !=http.MethodPost{
		http.Error(response , "Metodo não permitido", http.StatusMethodNotAllowed)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// json.NewEncoder(w).Encode(map[string]string{"message":"Metodo nao permitido"})
		
		return
	}

	var modelLogin models.LoginRequest
	if err := json.NewDecoder(request.Body).Decode(&modelLogin); err != nil{
		http.Error(response, "Erro ao processar requisicao", http.StatusBadRequest)
		return
	}
	
	user,err := repository.GetUserByEmail(request.Context(),modelLogin.Email)
	if err != nil{
		http.Error(response, "Erro ao buscar usuário", http.StatusUnauthorized)
		return
	}
	
	if !service.CheckPassword(modelLogin.Password, user.Password){
		http.Error(response, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}
	

	response.Header().Set("Content-Type","application/json")
	
	token , err:=service.GenerateJWT(modelLogin.Email)
	if err != nil{
		json.NewEncoder(response).Encode(map[string]string{"message":"Erro ao gerar token"})
		
		return
	}
	json.NewEncoder(response).Encode(map[string]string{"message":"Login realizado com sucesso","token":token})		

}

func CadastroResponse(response http.ResponseWriter, request *http.Request){
	if request.Method !=http.MethodPost{
		http.Error(response , "Metodo não permitido", http.StatusMethodNotAllowed)
		// json.NewEncoder(w).Encode(map[string]string{"message":"Metodo nao permitido"})
		return
	}

	var modelLogin models.LoginRequest
	if err := json.NewDecoder(request.Body).Decode(&modelLogin); err != nil{
		http.Error(response, "Erro ao processar requisicao", http.StatusBadRequest)
		return
	}

	if err := repository.CreateUser(request.Context(), modelLogin.Email, modelLogin.Password); err != nil{
		http.Error(response, "Erro ao criar usuário", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{"message":"Cadastro realizado com sucesso"})


}

func RecuperarSenhaResponse(response http.ResponseWriter, request *http.Request){

	if request.Method !=http.MethodPost{
		http.Error(response , "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	var modelRecuperarSenha models.RecuperarSenhaRequest
	if err := json.NewDecoder(request.Body).Decode(&modelRecuperarSenha); err != nil{
		http.Error(response, "Erro ao processar requisicao", http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type","application/json")
	json.NewEncoder(response).Encode(map[string]string{"message":"Instruções enviadas para o seu email"})
}
