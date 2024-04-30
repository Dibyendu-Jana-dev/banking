package domain

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/url"
)

//	"fmt"
//	"log"
//	"time"
//
//	_ "github.com/lib/pq"
//)
//
//type SQLConnDetails struct {
//	Pool     *sql.DB
//	PgSchema string
//}
//
//var Pool *sql.DB
//
//func InitDbPool() SQLConnDetails{
//	var port, user, password, dbname string
//	host := "localhost"
//	port = "5432"
//	user = "postgres"
//	password = "dev0924cc"
//	dbname = "postgres"
//
//	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//	config, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		log.Print(err)
//		log.Print("Error in config.")
//	}
//	config.SetConnMaxLifetime(time.Minute * 3)
//	config.SetMaxOpenConns(10)
//	config.SetMaxIdleConns(10)
//
//	err = config.Ping()
//	if err != nil {
//		log.Print(err)
//		log.Print("Could not connect to Postgres.")
//	} else {
//		log.Println("Postgres connected!")
//	}
//
//	Pool = config
//	return SQLConnDetails{Pool: Pool, PgSchema: "public"}
//}

//func GetPool() *sql.DB {
//	if Pool == nil {
//		InitDbPool()
//	}
//
//	return Pool
//}

const DriverName = "mongodb"

type Config struct {
	Host                   string
	Port                   string
	Username               string
	Password               string
	MaxPool                string
	Database               string
	UserCollection         map[string]string
	AdminCollection        map[string]string
	ProviderCollection     map[string]string
	ConnectorsCollection   map[string]string
	CwpCollection          map[string]string
	EventLogsCollection    map[string]string
	FileListingCollection  map[string]string
	CloudListingCollection map[string]string
	EnforcerCollection     map[string]string
	NotificationCollection map[string]string
	PolicyCollection       map[string]string
	IntelligenceCollection map[string]string
	CdrCollection          map[string]string
}

func Init(dbConfig *Config) (*mongo.Client, error) {
	dataSource := fmt.Sprintf("%s://%s:%s@%s:%s/?maxPoolSize=%s&w=majority",
		DriverName, dbConfig.Username, url.QueryEscape(dbConfig.Password), dbConfig.Host, dbConfig.Port, dbConfig.MaxPool)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dataSource))
	if err != nil {
		return nil, err
	}

	// verifies connection is db is working
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}
