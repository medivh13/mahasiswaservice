package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sqlx.DB
	// ConnMongo *mongo.Database
}

func Initialize(host, username, password, database, port string) (Database, error) {
	db := Database{}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

// func InitializeMongo(host, user, password, database, port string) (*mongo.Database, *mongo.Client, error) {

// 	db := Database{}

// 	dsn := fmt.Sprintf("mongodb://%s:%s", host, port)

// 	clientOptions := options.Client().ApplyURI(dsn).SetAuth(options.Credential{
// 		Username:   user,
// 		Password:   password,
// 		AuthSource: database,
// 	})
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, nil, err
// 	}

// 	err = client.Connect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, nil, err
// 	}

// 	db.ConnMongo = client.Database(database)
// 	return db.ConnMongo, client, nil
// }
