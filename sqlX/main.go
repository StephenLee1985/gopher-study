package main

import (
	//	"database/sql"
	"fmt"

	log "github.com/cihub/seelog"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/driver/mysql"
	"github.com/mattes/migrate/migrate"
)

type MysqlConfig struct {
	userName string
	passwd   string
	host     string
	port     int16
	dbName   string
}

var mysqlConfig = MysqlConfig{
	userName: "root",
	passwd:   "111111",
	host:     "127.0.0.1",
	port:     3306,
	dbName:   "fang",
}

var schema = `
CREATE TABLE person (
	first_name text,
	last_name text,
	email text
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func InitDB() (*sqlx.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		mysqlConfig.userName,
		mysqlConfig.passwd,
		mysqlConfig.host,
		mysqlConfig.port,
		mysqlConfig.dbName)
	db, err := sqlx.Open("mysql", uri)
	if err != nil {
		fmt.Errorf("cat not connection mysql error: %v, uri:%s", err, uri)
		return db, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Errorf("can not ping mysql error: %s", err)
		return db, err
	}
	//	db.SetMaxIdleConns(int(GetConfig().Mc.MaxIdleConns))
	//	db.SetMaxOpenConns(int(GetConfig().Mc.MaxOpenConns))
	return db, err
}

func upgradeDB() {
	uri := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		mysqlConfig.userName,
		mysqlConfig.passwd,
		mysqlConfig.host,
		mysqlConfig.port,
		mysqlConfig.dbName)
	log.Info("upgrade db mysql drive: ", uri)
	errors, ok := migrate.UpSync(uri, "./sql")
	if errors != nil && len(errors) > 0 {
		for _, err := range errors {
			log.Error("db err", err)
		}
		log.Error("can't upgrade db", errors)
		log.Flush()
		panic(-1)
	}
	if !ok {
		log.Error("can't upgrade db")
		log.Flush()
		panic(-1)
	}
	log.Info("DB upgraded")
	log.Flush()
}

func main() {
	db, _ := InitDB()
	upgradeDB()
	//db.MustExec(schema)

	//tx := db.MustBegin()
	//	db.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	//	db.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	//tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	//tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	//tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	//	tx.Commit()

	// Query the database, storing results in a []Person (wrapped in []interface{})

	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// Person{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

	// You can also get a single result, a la QueryRow
	jason = Person{}
	db.Get(&jason, "SELECT * FROM person WHERE first_name=?", "Jason")
	fmt.Printf("%#v\n", jason)

	var strSlice = []string{"444", "555"}
	fmt.Println(strings.Join(strSlice, ","))
	/*	db.NamedExec("INSERT INTO application (uid,cid, name,instances,status) VALUES (:uid,:cid,:name,:instances,:status)",
		map[string]interface{}{
			"uid":       "1",
			"cid":       "444",
			"name":      "app1",
			"instances": 10,
			"status":    1,
		})*/
	sqlStr := fmt.Sprintf("update application set uid=10 where cid in (%s)", strings.Join(strSlice, ","))
	fmt.Println("sql : ", sqlStr)
	db.NamedExec(sqlStr,
		map[string]interface{}{},
	)
}
