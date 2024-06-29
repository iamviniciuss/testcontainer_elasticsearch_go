package elasticsearch

import (
	"io/ioutil"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

type ESConnection struct {
	client *elasticsearch.Client
}

func NewESConnection() *ESConnection {
	esConnection := new(ESConnection)
	cert, err := ioutil.ReadFile(os.Getenv("ES_CRT"))
	if err != nil {
		panic(err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ES_URL"),
		},
		CACert:   cert,
		Username: os.Getenv("ES_USERNAME"),
		Password: os.Getenv("ES_PASSWORD"),
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	esConnection.client = es

	return esConnection
}

func (esc *ESConnection) Info() {
	res, err := esc.client.Info()

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func (esc *ESConnection) Client() *elasticsearch.Client {
	return esc.client
}
