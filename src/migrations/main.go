package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"path"
	"regexp"
)

func main() {
	var (
		dir      = flag.String("dir", "", "Path to the directory that contains migration files.")
		username = flag.String("U", "", "Name of the user with rights to connect to the database.")
		password = flag.String("P", "", "Password for connecting to the database.")
		database = flag.String("db", "", "Name of the database to connect to.")
		host     = flag.String("H", "localhost", "Host of the database server.")
		port     = flag.String("p", "5432", "Number of the port the database server accepts connections.")
		// dbUrl = flag.String(
		// 	"db_url",
		// 	"dude_expenses:dudeExpen$es123@localhost:5432/dude_expenses_development",
		// 	"The Postgresql database URL. Should be in 'user:password@host:port/database' format.",
		// )
		required = []string{"U", "P", "db", "dir"}
	)
	flag.Parse()

	passed := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { passed[f.Name] = true })
	for _, req := range required {
		if !passed[req] {
			log.Fatalf("Missing required argument -%s. Run with -h to see required arguments.\n", req)
		}
	}

	dbUrl := "postgres://" + *username + ":" + *password + "@" + *host + ":" + *port + "/" + *database + "?sslmode=require"
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}
	defer db.Close()

	if err = enableSchemaMigrations(db); err != nil {
		log.Fatal(err)
	}
	version, err := getSchemaMigrationsVersion(db)
	if err != nil {
		log.Fatal(err)
	}
	err = runMigrations(*dir, db, version)
	if err != nil {
		log.Fatal(err)
	}
}

func enableSchemaMigrations(db *sql.DB) error {
	var res string
	queryStr := `
	SELECT EXISTS (
  	SELECT 1
   	FROM pg_tables
   	WHERE schemaname = 'public'
   	AND tablename = 'schema_migrations'
  )
  `
	if err := db.QueryRow(queryStr).Scan(&res); err != nil {
		return err
	}
	if !(res == "true") {
		return createSchemaMigrations(db)
	}
	return nil
}

func createSchemaMigrations(db *sql.DB) error {
	queryStr := "CREATE TABLE schema_migrations(version VARCHAR(15) PRIMARY KEY)"
	rows, err := db.Query(queryStr)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func getSchemaMigrationsVersion(db *sql.DB) (string, error) {
	var res string
	queryStr := "SELECT version FROM schema_migrations ORDER BY version DESC LIMIT(1)"
	if err := db.QueryRow(queryStr).Scan(&res); err != nil && err != sql.ErrNoRows {
		return res, err
	}
	return res, nil
}

type Migration struct {
	Name    string
	Version string
	Sql     string
}

func runMigrations(dir string, db *sql.DB, currVersion string) error {
	migrations := make([]Migration, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	var version string
	var name string
	var content []byte
	for _, file := range files {
		version, name, err = getVersionAndName(file.Name())
		content, err = ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return err
		}
		migrations = append(
			migrations,
			Migration{Version: version, Sql: string(content), Name: name},
		)
	}

	const queryStr = "INSERT INTO schema_migrations VALUES($1)"
	for _, m := range migrations {
		if currVersion < m.Version {
			fmt.Println("== " + m.Name + ": migrating ===========================")
			fmt.Println(m.Sql)

			rows, err := db.Query(m.Sql)
			if err != nil {
				return err
			}
			defer rows.Close()
			fmt.Println("== " + m.Name + ": migrated ============================")
			fmt.Println("")
			db.QueryRow(queryStr, m.Version)
		}
	}
	return nil
}

func getVersionAndName(filename string) (string, string, error) {
	var version string
	var name string
	r, err := regexp.Compile("^(\\d+)_(.+)\\.sql")
	if err != nil {
		return version, name, err
	}
	matches := r.FindStringSubmatch(filename)
	if len(matches) > 2 {
		version = matches[1]
		name = matches[2]
	}
	return version, name, nil
}
