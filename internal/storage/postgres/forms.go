package postgres

type Form struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Schema     string `json:"schema"`
	CreativeAt string `json:"creative_at"`
}
