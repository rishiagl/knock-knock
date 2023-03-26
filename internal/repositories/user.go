// Repositories package is used to store all the funtions that directly interact with the database.
package repositories

// TODO: Context to be added.
import (
	"fmt"
	"log"

	"github.com/R58235/knock-knock/internal/models"
)

func(r *Repository) GetUserById(id []byte) (models.User, error){
	var u models.User

	err := r.db.QueryRow("select id, name, passwd from knock_knock.user where id = $1", id).Scan(&u.Id, &u.Usrname, &u.Passwd)
	if err != nil {
		log.Println(err)
		return u, fmt.Errorf("User not found: %v", err)
	}
	return u, nil
}

func(r *Repository) GetUserByUsrname(usrname string) (models.User, error){
	var u models.User

	err := r.db.QueryRow("select id, name, passwd from knock_knock.user where usrname = $1", usrname).Scan(&u.Id, &u.Usrname, &u.Passwd)
	if err != nil {
		log.Println(err)
		return u, fmt.Errorf("User not found: %v", err)
	}
	return u, nil
}

func (r *Repository) AddUser(u models.User) (int, error){
	result, err := r.db.Exec("INSERT INTO knock_knock.user VALUES (?, ?, ?)", u.Id, u.Usrname, u.Passwd)
    if err != nil {
        return 0, fmt.Errorf("AddUser Error: %v", err)
    }

    // Get the new user's generated ID for the client.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddUser Error: %v", err)
    }
    // Return the new user's ID.
    return id, nil
}

func (r *Repository) UpdateUser(u models.User) (error){
	_, err := r.db.Exec("UPDATE knock_knock.user SET usrname = $1, passwd = $2 WHERE id = $3", u.Usrname, u.Passwd, u.Id)
    if err != nil {
        return fmt.Errorf("UpdateUser Error: %v", err)
    }
    return nil
}

func (r *Repository) DeleteUserWithId(id []byte) (error){
	_, err := r.db.Exec("DELETE FROM knock_knock.user WHERE id = $1", id)
    if err != nil {
        return fmt.Errorf("DeleteUSer Error: %v", err)
    }
    return nil
}