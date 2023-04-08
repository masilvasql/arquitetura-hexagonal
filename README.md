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

# Iniciando com COBRA (comand line interface)
* No terminal rodar o seguinte comando: ```cobra-cli init```

* Para adicionar um novo comando cli no cobra, rodar o seguinte comando: ```cobra-cli add cli```

* Para verificar os dados do seu cli, digitar o comando : ```go run main.go cli -h```

* Exemplo de código de criação de produto: ```go run main.go cli -a=create -n="Product CLI" -p=50.0```
* Exemplo de get: ```go run main.go cli -a=get -i=52660bf4-312c-47d3-bbd3-905328d593d7```