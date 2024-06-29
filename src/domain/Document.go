package domain

type Document struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	CreateByEmployeeID string `json:"create_by_employee_id"`
}
