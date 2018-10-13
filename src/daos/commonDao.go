package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// http://tdm-gcc.tdragon.net/download need to download 64bit exe to windows for handling gcc not found issue.
	_ "github.com/mattn/go-sqlite3"
	"entities"
)

/*
  Before running program, create database first. 

  $ sqlite3 test.db
  sqlite> sqlite3    <---- login to database
  sqlite> .database  <---- show database file place
  $ sqlite3 testDB.db .dump > testDB.sql    <---- export
  $ sqlite3 testDB.db < testDB.sql          <---- import

  sqlite> .table     <---- show tables
  sqlite> select * from xxx;
*/
func Test() (error){
	// db object maintians connection pool
	var db *sqlx.DB

	// exactly the same as the built-in
	db, err := sqlx.Open("sqlite3", "./test.db")
	
	// from a pre-existing sql.DB; note the required driverName
	//db = sqlx.NewDb(sql.Open("sqlite3", ":memory:"), "sqlite3")
	
	// force a connection and test that it worked
	err = db.Ping()

	if(err != nil) {
		fmt.Println(err)
		return err
	}

	/*
	createTableSql := `CREATE TABLE place (
		country text,
		city text NULL,
		telcode integer);`
	_, err = db.Exec(createTableSql)
	
	if(err != nil) {
		fmt.Println(err)
		return err
	}
	*/

	countryCitySql := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
	res := db.MustExec(countryCitySql, "China", "Hong Kong", 852)
	fmt.Println(res)

    placeEntity := []entities.PlaceEntity{}
    db.Select(&placeEntity, "SELECT * FROM place ORDER BY country ASC")
	fmt.Println(placeEntity)
	return err
}


