package repo


import (
	"net/url"
	"github.com/jinzhu/gorm"
	// We need the postgres driver
  _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/lib/pq"
)

// Repository yup
type Repository struct {
	db *gorm.DB
}



// InitDatabase will be used to create a database client/repo
func InitDatabase(url *url.URL) (*Repository, error) {

	// Figure out the options
	db, err := gorm.Open(url.Scheme, url.String())
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(20)
	//defer db.Close()
	
	

	return &Repository{db}, nil
}