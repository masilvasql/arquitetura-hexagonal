# Testes com go

* Na pasta do projeto, execute o comando `go test ./...` para executar os testes.

# Criando mock para testes de unidade
* mockgen -destination=application/mocks/application.go -source=application/product.go application

# Usando SQLIT3 (já está no dockerfile os comandos, menos a criação do db)
* apt-get update
* apt-get install sqlite3
* criar arquivo de banco de dados: touch db.sqlite3

# Acessando o banco de dados
* sqlite3 sqlite.db
* No console que abrir, digitar os comandos:
    * create table products(id string, name string, price float, status string);

* para saber se a tabela foi criada, digitar o comando: .tables