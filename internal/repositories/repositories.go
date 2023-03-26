package repositories

import (
	"database/sql"
	"sync"

	"github.com/R58235/knock-knock/pkg/datastore"
)

var lock = &sync.Mutex{}
type Repository struct{
	db *sql.DB
}

var repo *Repository
func GetRepository() *Repository {

    if repo == nil {
        lock.Lock()
        defer lock.Unlock()

        if repo == nil {
			db := datastore.NewConnection()
            repo = &Repository{db}
			return repo
        } else {
			return repo
        }
    } else {
        return repo
    }
}