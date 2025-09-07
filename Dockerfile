# Usa Go oficial
FROM golang:1.25-alpine

# Instalar bash y netcat para wait-for-it
RUN apk add --no-cache bash curl

# Carpeta de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum primero (para cachear dependencias)
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Copiar wait-for-it.sh
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Construir la app
RUN go build -o main ./cmd/server

# Puerto expuesto
EXPOSE 8080

# Comando por defecto: espera a Postgres y luego corre Go
CMD ["./wait-for-it.sh", "postgres:5432", "--", "./main"]
