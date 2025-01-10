package database

import (
	// "database/sql"
	// "log"

	// _ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBConnections struct {
	// MySQL   *sql.DB
	MongoDB *mongo.Client
}

func InitDB() (DBConnections, error) {
	// db_mysql, err := InitMySqlDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db_mongo := InitMongoDB()

	connections := DBConnections{
		// MySQL:   db_mysql,
		MongoDB: db_mongo,
	}

	return connections, nil
}
