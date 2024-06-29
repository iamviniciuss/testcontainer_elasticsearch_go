package elasticsearch

type Document struct {
	Source interface{} `json:"_source"`
}
