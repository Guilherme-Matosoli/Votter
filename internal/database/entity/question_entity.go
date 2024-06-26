package entity

type question struct {
	Id          string `json:"id"`
	Poll_id     string `json:"poll_id"`
	Description string `json:"description"`
}
