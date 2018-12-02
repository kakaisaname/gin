package model

import (
	db "wine/database"
)
type Person struct {
	Id 			int `json:"id" form:"id"`
	FirstName 	string `json:"first_name" form:"firstname"`
	LastName 	string `json:"last_name" form:"lastname"`
}

func (p *Person) AddPerson() (id int64,err error) {
	exec, err := db.Db.Exec("INSERT INTO(firstname,lastname) VALUES (?,?)",p.FirstName,p.LastName)
	if err != nil {
		return
	}
	id,err = exec.LastInsertId()
	return
}
