package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


type Store struct {
}


func (s *Store) Open()(*sql.DB, error) {
	config:=NewConfig()
	d,err :=sql.Open("postgres",config.DatabaseURL)
if err!=nil{
	return nil,err
}

if err:=d.Ping(); err!=nil{
	return nil,err
}
	fmt.Println("Data base ping done!")
return d,nil
}



