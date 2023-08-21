package domain

type Group struct {
	ID   int    `json:"id" ent:"id"`
	Name string `json:"name" ent:"name"`
}
