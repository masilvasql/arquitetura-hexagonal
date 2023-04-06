# Testes com go

* Na pasta do projeto, execute o comando `go test ./...` para executar os testes.

# Criando mock para testes de unidade
* mockgen -destination=application/mocks/application.go -source=application/product.go application