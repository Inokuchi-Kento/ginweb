package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Group_ID int    `json:"group_id"`
}
