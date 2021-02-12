package structs

type Status struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}
