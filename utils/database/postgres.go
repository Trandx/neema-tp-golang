// database/postgres.go
package database

import (
	"context"
	"fmt"
	log "log"
	"os"
	"sync"

	"gocloud.dev/postgres"
	_ "gocloud.dev/postgres/awspostgres"
	"xorm.io/xorm"
	"xorm.io/xorm/core"
	xlog "xorm.io/xorm/log"
	"xorm.io/xorm/names"

	//"gorm.io/driver/postgres"
	. "neema.co.za/rest/utils/logger"
)

type Database struct {
	*xorm.Engine
}

var engine *Database
var once sync.Once

func GetDatabase() *Database {
	once.Do(func() {
		dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"))

		db, err := postgres.Open(context.Background(), dbURL) //??
		if err != nil {
			return
		}
		//defer db.Close()

		coreDB := core.FromDB(db)
		xengine, err := xorm.NewEngineWithDB("postgres", dbURL, coreDB) //??

		xengine.DB().SetMaxOpenConns(1) // Définir le nombre maximal de connexions ouvertes
		xengine.NewSession()

		//session := xengine.NewSession()
		//defer session.Close()

		if err != nil {
			log.Fatalf("Error creating XORM engine: %v", err)
		}

		xengine.SetConnMaxLifetime(60000000000) // 60 seconds

		// Enable query logging
		xengine.SetLogger(xlog.NewSimpleLogger(os.Stdout))

		xengine.SetMapper(names.GonicMapper{})

		if err := xengine.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}

		Logger.Info("Connected to PostgreSQL database")
		engine = &Database{xengine}
	})
	return engine
}
