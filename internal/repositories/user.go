package repositories

import (
	"fmt"
	"log"

	"github.com/R58235/knock-knock/internal/models"
)

func(r *Repository) GetUser(id int16) models.User{
	rows, err := r.db.Query("select id, name, passwd from knock_knock.user where id = $1", id)
	if err != nil {
		fmt.Println("Error in rows")
		log.Fatal(err)
	}
	defer rows.Close()
	var u models.User
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name, &u.Passwd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.Name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return u
}