package repositories

import (
	"sync"

	"github.com/R58235/knock-knock/pkg/datastore"
	"github.com/jackc/pgx/v5"
)

var lock = &sync.Mutex{}
type Repository struct{
	conn *pgx.Conn
}

var repo *Repository
func GetRepository() *Repository {

    if repo == nil {
        lock.Lock()
        defer lock.Unlock()

        if repo == nil {
			conn := datastore.NewConnection()
            repo = &Repository{conn}
			return repo
        } else {
			return repo
        }
    } else {
        return repo
    }
}