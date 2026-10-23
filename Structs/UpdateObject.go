package structs

type UpdateObject struct {
	Change string `json:"change"`
	Value  any    `json:"value"`
}
