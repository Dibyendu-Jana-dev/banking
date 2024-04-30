package domain

import "go.mongodb.org/mongo-driver/mongo"

//func (conn SQLConnDetails) FindAll() ([]Customer, error) {
//	PgSchema := "public"
//	//if Pool == nil{
//	//	Pool = GetPool()
//	//}
//	//
//	//conn = SQLConnDetails{
//	//	PgSchema: PgSchema,
//	//	Pool: Pool,
//	//}
//	sqlQuery := `SELECT id, "name", date_of_birth, city, zipcode, status FROM ` + PgSchema + `.customers`
//	rows, err := conn.Pool.Query(sqlQuery)
//	if err != nil {
//		logger.Error("error while customer query table" + err.Error())
//		return nil, err
//	}
//	customers := make([]Customer, 0)
//	for rows.Next() {
//		var customer Customer
//		err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.Zipcode, &customer.Status)
//		if err != nil {
//			logger.Error("error while customer scanning table" + err.Error())
//			return nil, err
//		}
//		customers = append(customers, customer)
//	}
//	return customers, nil
//}
//
//func (conn SQLConnDetails) ById(customerId int) (*Customer, *errs.AppError) {
//	PgSchema := "public"
//
//	sqlQuery := `SELECT id, "name", date_of_birth, city, zipcode, status FROM ` + PgSchema + `.customers where id = $1`
//	rows := conn.Pool.QueryRow(sqlQuery, customerId)
//	var customer Customer
//	err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.Zipcode, &customer.Status)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, errs.NewNotFoundError("Customer Not Found")
//		} else {
//			logger.Error("error while customer scanning table" + err.Error())
//			return nil, errs.NewUnExpectedError("unexpected Database error")
//		}
//	}
//	return &customer, nil
//}




/*package main

import (
"log"
"os"
"gopkg.in/gomail.v2"
)

func main() {
	// Retrieve email credentials from environment variables
	email := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")

	if email == "" || password == "" {
		log.Fatal("Email credentials not found in environment variables")
		return
	}

	// Create a new message
	m := gomail.NewMessage()

	// Set the sender address
	m.SetHeader("From", email)

	// Set the recipient's email address
	m.SetHeader("To", "deev0924cgmail.com")

	// Set the subject of the email
	m.SetHeader("Subject", "Hello, Gomail!")

	// Set the body of the email
	m.SetBody("text/plain", "This is the email body.")

	// Create an SMTP client and send the email
	d := gomail.NewDialer("smtp.example.com", 587, email, password)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}

	log.Println("Email sent successfully!")
}*/

type UserRepositoryDb struct {
	client     *mongo.Client
	database   string
	collection map[string]string
}

func NewUserRepositoryDb(dbClient *mongo.Client, database string, collection map[string]string) UserRepositoryDb {
	return UserRepositoryDb{
		client:     dbClient,
		database:   database,
		collection: collection,
	}
}