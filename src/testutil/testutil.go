package testutil

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"
)

type TestState struct {
	Initialized                   bool
	Clear_elasticsearch_container func()
}

var (
	once     sync.Once
	instance *TestState
)

const syncFilePath = "/tmp/test_state_initialized2"

func GetTestState() *TestState {
	once.Do(func() {
		if _, err := os.Stat(syncFilePath); os.IsNotExist(err) {
			instance = &TestState{
				Initialized: true,
			}
			file, err := os.Create(syncFilePath)
			if err != nil {
				fmt.Println("Error creating sync file:", err)
			}
			file.Close()

			instance.Clear_elasticsearch_container = tests.NewESConnectionTests("infra")

			fmt.Println("Shared state initialized")
		} else {
			// Carregar ou inicializar estado existente
			instance = &TestState{
				Initialized: true,
				// Carregar outros campos conforme necessário
			}
			fmt.Println("Shared state already initialized", " - GoID:", GoID(), " ES:", instance.Clear_elasticsearch_container)
		}
	})
	return instance
}

// func TestMain(m *testing.M) {
// 	state := GetTestState()

// 	code := m.Run()

// 	state.Clear_elasticsearch_container()

// 	CleanupTestState()

// 	os.Exit(code)
// }

// func CleanupTestState() {
// 	// Remove o arquivo de sincronização
// 	if err := os.Remove(syncFilePath); err != nil {
// 		fmt.Println("Error removing sync file:", err)
// 	}
// 	fmt.Println("Teardown shared state")
// }

// func GoID() int {
// 	var buf [64]byte
// 	n := runtime.Stack(buf[:], false)
// 	idField := bytes.Fields(buf[:n])[1]
// 	id, err := strconv.Atoi(string(idField))
// 	if err != nil {
// 		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
// 	}
// 	return id
// }

// package testutil

// import (
// 	"fmt"
// 	"os"
// 	"sync"
// 	"testing"

// 	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/elasticsearch"
// 	"github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"
// )

// type TestState struct {
// 	// Adicione aqui os campos que você precisa compartilhar
// 	Initialized bool
// 	// Outros campos...
// 	*elasticsearch.ESConnection
// }

// var (
// 	once     sync.Once
// 	instance *TestState
// )

// func GetTestState() *TestState {
// 	once.Do(func() {
// 		instance = &TestState{
// 			Initialized: true,
// 			// Inicialize outros campos conforme necessário
// 		}
// 		fmt.Println("Shared state initialized")
// 	})
// 	return instance
// }

func TestMain(m *testing.M) {
	// Inicializa o estado compartilhado
	_ = GetTestState()

	// Coloque aqui suas configurações compartilhadas, como inicialização de banco de dados, variáveis de ambiente, etc.

	fmt.Println("@@@@@@ TestMain testutil")
	clear_elasticsearch_container := tests.NewESConnectionTests("infra")

	clear_elasticsearch_container()

	// Rodar os testes
	code := m.Run()

	// Teardown, se necessário
	fmt.Println("Teardown shared state")

	// Saída com o código de status dos testes
	// Normalmente, não chamamos os.Exit aqui em código compartilhado,
	// porque isso encerraria o processo de teste imediatamente.
	// Os frameworks de teste fazem isso depois que TestMain retorna.
	os.Exit(code)
}
