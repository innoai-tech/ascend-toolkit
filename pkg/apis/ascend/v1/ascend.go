package v1

type Ascend struct {
	Version      string   `json:"version"`
	Type         string   `json:"type"`
	ProductTypes []string `json:"product_types,omitzero"`
	Chips        []*Chip  `json:"chips"`
}
