package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "testdb"
)

func sess() (Db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}
	log.Info().Msg("Successfully connected!")
	return db, err
}

func insert(name string, age int, salary float64, designation string) {
	db, err := sess()
	if err != nil {
		log.Error().Msg("Error while connecting to db : " + err.Error())
		return
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Error().Msg("Error while beginning transaction")
		return
	}

	sqlstmnt := "insert into test (name,age,salary,designation) values ($1,$2,$3,$4)"
	stmt, err := tx.Prepare(sqlstmnt)
	if err != nil {
		tx.Rollback()
		log.Error().Msg("Error while preparing query : " + err.Error())
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, age, salary, designation)
	if err != nil {
		tx.Rollback()
		log.Error().Msg("Error in executing query : " + err.Error())
		return
	}
	log.Info().Msg("Query succwssfully executed")
	tx.Commit()
	selectOp()
}

func delete(id int) {
	db, err := sess()
	if err != nil {
		log.Error().Msg("Error while connecting to db : " + err.Error())
		return
	}
	defer db.Close()
	query := "delete from test where id=$1"
	_, err = db.Exec(query, id)
	if err != nil {
		log.Error().Msg("Error in executing query : " + err.Error())
		return
	}
	log.Info().Msg("Query succwssfully executed")
	selectOp()
}

func update(id int, name string) {
	db, err := sess()
	if err != nil {
		log.Error().Msg("Error while connecting to db : " + err.Error())
		return
	}
	defer db.Close()
	query := "update test set name=$2 where id=$1"
	_, err = db.Exec(query, id, name)
	if err != nil {
		log.Error().Msg("Error in executing query : " + err.Error())
		return
	}
	log.Info().Msg("Query succwssfully executed")
	selectOp()
}

func selectOp() {
	type row struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Age         int     `json:"age"`
		Salary      float64 `json:"salary"`
		Designation string  `json:"designation"`
	}
	var rows []row
	db, err := sess()
	if err != nil {
		log.Error().Msg("Error while connecting to db : " + err.Error())
		return
	}
	defer db.Close()
	query := "select * from test"
	resultRows, err := db.Query(query)
	if err != nil {
		log.Error().Msg("Error in executing query : " + err.Error())
		return
	}
	defer resultRows.Close()
	for resultRows.Next() {
		var resultRow row
		err = resultRows.Scan(&resultRow.ID, &resultRow.Name, &resultRow.Age, &resultRow.Salary, &resultRow.Designation)
		if err != nil {
			log.Error().Msg("Error while scanning row : " + err.Error())
			return
		}
		rows = append(rows, resultRow)
	}
	bytes, _ := json.Marshal(rows)
	log.Info().Msg("Select query result : " + string(bytes))
	//log.Info().Msg("Query succwssfully executed")
}

func createTestTable() {
	db, err := sess()
	if err != nil {
		log.Error().Msg("Error while connecting to db : " + err.Error())
		return
	}
	defer db.Close()
	query := "create table test(id serial primary key,name text,age int,salary real,designation varchar(20))"
	_, err = db.Exec(query)
	if err != nil {
		log.Error().Msg("Error in executing query : " + err.Error())
		return
	}
	log.Info().Msg("Table test created")
}

func dropTestTable() {
	db, err := sess()
	if err != nil {
		log.Error().Msg("Error while connecting to db : " + err.Error())
		return
	}
	defer db.Close()
	query := "drop table test"
	_, err = db.Exec(query)
	if err != nil {
		log.Error().Msg("Error in executing query : " + err.Error())
		return
	}
	log.Info().Msg("Table test dropped")
}

func main() {
	log.Info().Msg("Creating test table")
	createTestTable()
outer:
	for {
		var (
			op          string
			id          int
			name        string
			age         int
			salary      float64
			designation string
		)
		log.Info().Msg("Enter op to be executed. Valid Op can be insert, delete or update. Each valid op is followed by select op. Enter stop for exit")
		fmt.Scan(&op)
		switch op {
		case "insert":
			log.Info().Msg("Enter name age salary designation : ")
			fmt.Scan(&name, &age, &salary, &designation)
			insert(name, age, salary, designation)
		case "delete":
			log.Info().Msg("Enter id to be deleted : ")
			fmt.Scan(&id)
			delete(id)
		case "update":
			log.Info().Msg("Enter id for update : ")
			fmt.Scan(&id)
			log.Info().Msg("Enter name to be updated : ")
			fmt.Scan(&name)
			update(id, name)
		case "stop":
			log.Info().Msg("Exiting as stop op was entered")
			dropTestTable()
			break outer
		default:
			log.Info().Msg("Exiting as unknown op was entered")
			dropTestTable()
			break outer
		}
	}
}
