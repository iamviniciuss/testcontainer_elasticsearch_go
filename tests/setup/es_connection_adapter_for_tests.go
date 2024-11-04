package tests_setup

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	// "github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	testcontainers "github.com/testcontainers/testcontainers-go"
	elasticsearchTestContainer "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

type IndexInfo struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	UUID         string `json:"uuid"`
	Primary      string `json:"pri"`
	Replica      string `json:"rep"`
	DocsCount    string `json:"docs.count"`
	DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PrimaryStore string `json:"pri.store.size"`
}

type ESConnectionTest struct{}

const (
	CACERT_NAME    = "123"
	CONTAINER_NAME = "TestContainersES"
)

func NewESConnectionTests(nameOfSuite string) func() {

	var elasticsearchContainer *elasticsearchTestContainer.ElasticsearchContainer

	ctx := context.Background()
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	genericContainerRequest := &testcontainers.GenericContainerRequest{
		Logger: log.Default(),
		ContainerRequest: testcontainers.ContainerRequest{
			Name: CONTAINER_NAME,
			LifecycleHooks: []testcontainers.ContainerLifecycleHooks{
				{
					PreStarts: []testcontainers.ContainerHook{
						func(ctx context.Context, container testcontainers.Container) error {
							fmt.Println("PreStarts testcontainer for elasticsearch")
							return nil
						},
					},
					PreTerminates: []testcontainers.ContainerHook{
						func(ctx context.Context, container testcontainers.Container) error {
							fmt.Println("PreTerminates testcontainer for elasticsearch")
							return nil
						},
					},
					PostTerminates: []testcontainers.ContainerHook{
						func(ctx context.Context, container testcontainers.Container) error {
							fmt.Println("PostTerminates testcontainer for elasticsearch")
							return nil
						},
					},
					PostStarts: []testcontainers.ContainerHook{
						func(ctx context.Context, container testcontainers.Container) error {
							fmt.Println("PostStarts testcontainer for elasticsearch")

							// fn := func() {
							// 	waitingForStop := time.Second * 60
							// 	time.Sleep(waitingForStop)
							// 	elasticsearchContainer.Stop(ctx, &waitingForStop)
							// 	fmt.Println("****** Stop Container")
							// }

							// go fn()

							return nil
						},
					},
				},
			},
			Env: map[string]string{
				"ES_JAVA_OPTS":   "-Xms1g -Xmx1g",
				"discovery.type": "single-node",
				"node.name":      CONTAINER_NAME,
				"cluster.name":   CONTAINER_NAME,
			},
			SkipReaper: true,
		},
		ProviderType: 0,
		Started:      false,
		Reuse:        true,
	}

	elasticsearchContainerVar, err := elasticsearchTestContainer.RunContainer(
		ctx,
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:8.2.0"),
		testcontainers.CustomizeRequest(*genericContainerRequest),
	)

	if err != nil {
		fmt.Println("NewESConnectionTests/RunContainer")
		panic(err)
	}

	elasticsearchContainer = elasticsearchContainerVar

	address := strings.Replace(elasticsearchContainer.Settings.Address, "http://", "https://", -1)

	file, err := elasticsearchContainer.CopyFileFromContainer(ctx, "/usr/share/elasticsearch/config/certs/http_ca.crt")
	if err != nil {
		panic(err)
	}

	cacert, err := ioutil.ReadAll(ioutil.NopCloser(file))
	if err != nil {
		panic(err)
	}

	cfg := elasticsearch.Config{
		Logger: nil,
		Addresses: []string{
			address,
		},
		Username: elasticsearchContainer.Settings.Username,
		Password: elasticsearchContainer.Settings.Password,
		CACert:   cacert,
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	resp, err := esClient.Info()
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	certSourceFile := createTempCertES(string(cacert))

	os.Setenv("ES_URL", cfg.Addresses[0])
	os.Setenv("ES_USERNAME", elasticsearchContainer.Settings.Username)
	os.Setenv("ES_PASSWORD", elasticsearchContainer.Settings.Password)
	os.Setenv("ES_CRT", string(certSourceFile))

	return func(cont *elasticsearchTestContainer.ElasticsearchContainer) func() {
		return func() {
			log.Println("AFTER ALL TESTS: TearDown")
			deleteAllIndices(esClient)
		}
	}(elasticsearchContainer)
}

func createTempCertES(certData string) string {

	filePath := os.TempDir() + CACERT_NAME

	fileData := []byte(certData)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(fileData)
	if err != nil {
		log.Fatalf("error writing to file: %v", err)
	}

	return filePath
}

func deleteAllIndices(es *elasticsearch.Client) {
	indices, err := listAllIndices(es)
	if err != nil {
		log.Fatalf("Error listing indices: %s", err)
	}

	for _, index := range indices {
		err := deleteIndex(es, index.Index)
		if err != nil {
			log.Printf("Error deleting index %s: %s", index, err)
			continue
		}
		fmt.Printf("Index %s deleted successfully.\n", index)
	}
}

func listAllIndices(es *elasticsearch.Client) ([]IndexInfo, error) {
	req := esapi.CatIndicesRequest{
		Format: "json",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, fmt.Errorf("error listing indices: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.Status())
	}

	var indices []IndexInfo
	if err := json.NewDecoder(res.Body).Decode(&indices); err != nil {
		return nil, fmt.Errorf("error decoding indices response: %s", err)
	}

	return indices, nil
}

func deleteIndex(es *elasticsearch.Client, indexName string) error {
	req := esapi.IndicesDeleteRequest{
		Index: []string{indexName},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error deleting index %s: %s", indexName, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch for index %s: %s", indexName, res.String())
	}

	return nil
}
