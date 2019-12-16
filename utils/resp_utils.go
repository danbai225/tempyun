package utils

import "tempyun/entity"

func Ok() *entity.Rjson {
	var r = entity.Rjson{}
	r.Message = "success"
	r.State = 0
	return &r
}
func Err() *entity.Rjson {
	var r = entity.Rjson{}
	r.Message = "err"
	r.State = 0
	return &r
}
