# FROM golang:1.24.5

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o main .

# CMD ["./main"]


# Etapa 1: build
# imagem base Go
FROM golang:1.24.5

WORKDIR /app

# copia go.mod e go.sum para instalar dependências
COPY go.mod go.sum ./
RUN go mod download

# monta o código via bind mount em docker-compose
COPY . .

# expõe a porta da API
EXPOSE 8082

# comando para desenvolvimento: recompila e roda automaticamente
CMD ["sh", "-c", "go run main.go"]
