package dbase

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const (
	user     = "rylan"
	password = "password"
	net      = "tcp"
	addr     = "127.0.0.1:6603"
	dbname   = "omni"
)

func InitDatabase() (*sql.DB, error) {

	cfg := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    net,
		Addr:   addr,
		DBName: dbname,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

// func (db *Database) DoesTableExist(tn string) error {
// 	q := fmt.Sprintf("SELECT * FROM %s;", tn)
// 	_, err := db.db.Query(q)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (db *Database) CreateTable(tn string, m map[string]string) error {
// 	var t string
// 	for k, v := range m {
// 		t += fmt.Sprintf("%s %s,", k, v)
// 	}

// 	q := fmt.Sprintf("CREATE TABLE %s (%s)", tn, strings.TrimSuffix(t, ","))

// 	_, err := db.db.Query(q)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

// func (db *Database) Query(q string) (*sql.Rows, error) {
// 	rows, err := db.db.Query(q)

// 	if err != nil {
// 		fmt.Println(err)
// 		return rows, err
// 	}
// 	return rows, nil
// }

// func (db *Database) QueryWarg(q string, a []string) (*sql.Rows, error) {
// 	rows, err := db.db.Query(q, strings.Join(a, ","))
// 	fmt.Print(rows)
// 	return rows, err
// }
