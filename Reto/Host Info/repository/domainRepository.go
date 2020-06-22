package repository

import (
	"database/sql"
	"log"
	"sync"

	"../model"
	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var instance *sql.DB

//Connet - Here de database is connecting
func Connet() *sql.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		db, err := sql.Open("postgres",
			"postgresql://ash@localhost:26257/domain_servise?ssl=true&sslmode=require&sslrootcert=C:/Users/Usuario/Downloads/cockroach-v20.1.2.windows-6.2-amd64/certs/ca.crt&sslkey=C:/Users/Usuario/Downloads/cockroach-v20.1.2.windows-6.2-amd64/certs/client.ash.key&sslcert=C:/Users/Usuario/Downloads/cockroach-v20.1.2.windows-6.2-amd64/certs/client.ash.crt")
		if err != nil {
			log.Fatal("error conecting to the database", err)
		}
		instance = db
	}
	return instance
}

//SaveInfoSearch - Save information about the host
func SaveInfoSearch(domai model.Domain) {
	if _, err := Connet().Exec(
		"INSERT INTO domains (host, servers_change, is_down, ssl_grade, previous_ssl_grade, logo, title) VALUES ($1, $2, $3, $4, $5, $6, $7)", domai.Host, domai.ServersChange, domai.IsDown, domai.SslGrade, domai.PreviousSslGrade, domai.Logo, domai.Title); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(domai.Servers); i++ {
		if _, err := Connet().Exec(
			"INSERT INTO servers (address, ssl_grade, country, owner, domain_host) VALUES ($1, $2, $3, $4, $5)", domai.Servers[i].Address, domai.Servers[i].SslGrade, domai.Servers[i].Country, domai.Servers[i].Owner, domai.Host); err != nil {
			log.Fatal(err)
		}
	}
}

//CreateTables - Create necessary tables
func CreateTables() {
	if _, err := Connet().Exec(
		"DROP TABLE IF EXISTS servers"); err != nil {
		log.Fatal(err)
	}
	if _, err := Connet().Exec(
		"DROP TABLE IF EXISTS domains"); err != nil {
		log.Fatal(err)
	}
	if _, err := Connet().Exec(
		"CREATE TABLE IF NOT EXISTS domains (host VARCHAR PRIMARY KEY, servers_change BOOL, is_down BOOL, ssl_grade VARCHAR, previous_ssl_grade VARCHAR, logo VARCHAR, title VARCHAR)"); err != nil {
		log.Fatal(err)
	}
	if _, err := Connet().Exec(
		"CREATE TABLE IF NOT EXISTS servers (address VARCHAR, ssl_grade VARCHAR, country VARCHAR, owner VARCHAR, domain_host VARCHAR)"); err != nil {
		log.Fatal(err)
	}
	if _, err := Connet().Exec(
		"ALTER TABLE servers ADD CONSTRAINT domains_fk FOREIGN KEY (domain_host) REFERENCES domains (host) ON DELETE CASCADE"); err != nil {
		log.Fatal(err)
	}
}
