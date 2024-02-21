package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/drpulpo/gorn/model"
	"github.com/jackc/pgx/v4"
)

// GetUserData gets user data for a given id.
func GetOrderData(id int) (model.Order, error) {
	connString := `postgres://pgndsusa:password@fullstack-postgres:5432/fullstack_api`
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("error connecting to postgres using pgx. Error: %v", err)
	}

	defer conn.Close(context.Background())

	var order model.Order

	err = conn.QueryRow(context.Background(), `select id, firstname, lastname, phone, city, postalcode, state, country, crayon00, crayon01, crayon02, crayon03, crayon04, crayon05, crayon06, crayon07, crayon08, crayon09, crayon10, crayon11, crayon12, crayon13, crayon14, crayon15, crayon16, crayon17, crayon18, crayon19 from public.order where id = $1;`, id).
		Scan(
			&order.Id,
			&order.Firstname,
			&order.Lastname,
			&order.Phone,
			&order.City,
			&order.PostalCode,
			&order.State,
			&order.Country,
			&order.Crayon00,
			&order.Crayon01,
			&order.Crayon02,
			&order.Crayon03,
			&order.Crayon04,
			&order.Crayon05,
			&order.Crayon06,
			&order.Crayon07,
			&order.Crayon08,
			&order.Crayon09,
			&order.Crayon10,
			&order.Crayon11,
			&order.Crayon12,
			&order.Crayon13,
			&order.Crayon14,
			&order.Crayon15,
			&order.Crayon16,
			&order.Crayon17,
			&order.Crayon18,
			&order.Crayon19,
		)

	if err != nil {
		log.Fatal(err.Error())
	}

	return order, err

}

// InsertRecord inserts a user into postgres db
func InsertOrder(order model.Order) {
	log.Printf("value is %+v", order)
	connString := `postgres://pgndsusa:password@fullstack-postgres:5432/fullstack_api`
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("error connecting to postgres using pgx. Error: %v", err)
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `insert into public.order(
		firstname,
		lastname,
		phone,
		city,
		postalcode,
		state,
		country,
		crayon00,
		crayon01,
		crayon02,
		crayon03,
		crayon04,
		crayon05,
		crayon06,
		crayon07,
		crayon08,
		crayon09,
		crayon10,
		crayon11,
		crayon12,
		crayon13,
		crayon14,
		crayon15,
		crayon16,
		crayon17,
		crayon18,
		crayon19) values( 
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14,
			$15,
			$16,
			$17,
			$18,
			$19,
			$20,
			$21,
			$22,
			$23,
			$24,
			$25,
			$26,
			$27);`,
		order.Firstname,
		order.Lastname,
		order.Phone,
		order.City,
		order.PostalCode,
		order.State,
		order.Country,
		order.Crayon00,
		order.Crayon01,
		order.Crayon02,
		order.Crayon03,
		order.Crayon04,
		order.Crayon05,
		order.Crayon06,
		order.Crayon07,
		order.Crayon08,
		order.Crayon09,
		order.Crayon10,
		order.Crayon11,
		order.Crayon12,
		order.Crayon13,
		order.Crayon14,
		order.Crayon15,
		order.Crayon16,
		order.Crayon17,
		order.Crayon18,
		order.Crayon19)

	if err != nil {
		log.Fatalf("error occurred while inserting data to postgres %v", err)
	}

}

// get one user from the DB by its userid
func GetAllOrders() ([]model.Order, error) {
	connString := `postgres://pgndsusa:password@fullstack-postgres:5432/fullstack_api`
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("error connecting to postgres using pgx. Error: %v", err)
	}

	defer conn.Close(context.Background())

	var orders []model.Order

	rows, err := conn.Query(context.Background(), `select * from public.order`)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var order model.Order

		// unmarshal the row object to user
		err = rows.Scan(
			&order.Id,
			&order.Firstname,
			&order.Lastname,
			&order.Phone,
			&order.City,
			&order.PostalCode,
			&order.State,
			&order.Country,
			&order.Crayon00,
			&order.Crayon01,
			&order.Crayon02,
			&order.Crayon03,
			&order.Crayon04,
			&order.Crayon05,
			&order.Crayon06,
			&order.Crayon07,
			&order.Crayon08,
			&order.Crayon09,
			&order.Crayon10,
			&order.Crayon11,
			&order.Crayon12,
			&order.Crayon13,
			&order.Crayon14,
			&order.Crayon15,
			&order.Crayon16,
			&order.Crayon17,
			&order.Crayon18,
			&order.Crayon19,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		orders = append(orders, order)

	}

	// return empty user on error
	return orders, err
}

// update user in the DB
// func updateUser(id int64, user models.User) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the update sql query
// 	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

// 	// execute the sql statement
// 	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// check how many rows affected
// 	rowsAffected, err := res.RowsAffected()

// 	if err != nil {
// 		log.Fatalf("Error while checking the affected rows. %v", err)
// 	}

// 	fmt.Printf("Total rows/record affected %v", rowsAffected)

// 	return rowsAffected
// }

// delete user in the DB
func DeleteOrder(id int) (int64, error) {

	connString := `postgres://pgndsusa:password@fullstack-postgres:5432/fullstack_api`
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("error connecting to postgres using pgx. Error: %v", err)
	}

	defer conn.Close(context.Background())

	// create the delete sql query
	sqlStatement := `DELETE FROM public.order WHERE id=$1`

	res, err := conn.Exec(context.Background(), sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, err
}
