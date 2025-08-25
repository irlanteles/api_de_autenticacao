# API REST de Autenticação em Go

Uma API REST desenvolvida em Go para fins de aprendizado, principais funcionalidades cadastro, login e recuperação de usuários, utilizando JWT para autenticação e chaves RSA para assinatura de tokens.


## Tecnologias
- Go
- PostgreSQL
- Docker
- JWT


## Como rodar
1. Clone o repositório:
   ```bash
   git clone https://github.com/irlanteles/api_de_autenticacao.git
2. Certifique-se de ter o Docker e o Docker Compose instalados.
3. Construa as imagens do Docker:
   
    docker compose build

4. Execute os containers:
   
    docker compose up -d

5. A API estará rodando na porta configurada (por padrão 8082). Você pode acessar, por exemplo:

    http://localhost:8082/login

  


