# Variável para o diretório de testes
TEST_DIR := ./src/...
INTEGRATION_TEST_DIR := ./tests/integration_tests/...

# Alvo para rodar todos os testes
test:
	go test -v $(TEST_DIR) $(INTEGRATION_TEST_DIR)

# Alvo para rodar todos os testes e gerar o relatório de cobertura
coverage:
	go test -coverprofile=output_tests/coverage.out -coverpkg=$(TEST_DIR) $(INTEGRATION_TEST_DIR)
	go tool cover -html=output_tests/coverage.out -o output_tests/coverage.html
	open coverage.html

# Alvo padrão
all: test coverage
