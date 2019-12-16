package entity

type File struct {
	Isdir bool   `json:"isdir"`
	Mode  string `json:"mode"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	Read  bool   `json:"read"`
	Size  int64  `json:"size"`
	Time  int64  `json:"time"`
	Type  string `json:"type"`
	Write bool   `json:"write"`
}
