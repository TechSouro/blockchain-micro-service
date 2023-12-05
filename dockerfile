FROM golang:1.21

# Instalação do PostgreSQL
RUN apt-get update && \
    apt-get install -y postgresql-client

WORKDIR /app
COPY . .

# Construção da aplicação Go
RUN cd cmd/app && go build -o server

# Configuração do PostgreSQL
ENV PGHOST=db
ENV PGPORT=5432
ENV PGUSER=admin
ENV PGPASSWORD=123@lunave
ENV PGDATABASE=blockservice

# Comando para iniciar a aplicação e o PostgreSQL
CMD ["./cmd/app/server"]
