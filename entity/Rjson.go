package entity

type Rjson struct {
	State   int         `json:"state"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
