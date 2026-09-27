package models

type Seed struct {
	SeedSignature string   `json:"_seed_signature"`
	Directories   []string `json:"directories"`
	Files         []string `json:"files"`
}
