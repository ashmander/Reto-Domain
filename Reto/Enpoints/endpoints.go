package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"./model"

	"github.com/buaazp/fasthttprouter"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

var (
	corsAllowOrigin  = "*"
	corsAllowMethods = "GET,POST"
)

//CORS - Define origins an methods
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		next(ctx)
	}
}

func getEndpoints(ctx *fasthttp.RequestCtx) {
	db, err := sql.Open("postgres",
		"postgresql://ash@localhost:26257/domain_servise?ssl=true&sslmode=require&sslrootcert=C:/Users/Usuario/Downloads/cockroach-v20.1.2.windows-6.2-amd64/certs/ca.crt&sslkey=C:/Users/Usuario/Downloads/cockroach-v20.1.2.windows-6.2-amd64/certs/client.ash.key&sslcert=C:/Users/Usuario/Downloads/cockroach-v20.1.2.windows-6.2-amd64/certs/client.ash.crt")
	if err != nil {
		log.Fatal("error conecting to the database", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM domains")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var endpoints []model.Domain
	for rows.Next() {
		var hostConsulted model.Domain
		var host, ssl_grade, previous_ssl_grade, logo, title string
		var servers_change, is_down bool
		if err := rows.Scan(&host, &servers_change, &is_down, &ssl_grade, &previous_ssl_grade, &logo, &title); err != nil {
			log.Fatal(err)
		}
		hostConsulted.Host = host
		hostConsulted.SslGrade = ssl_grade
		hostConsulted.PreviousSslGrade = previous_ssl_grade
		hostConsulted.Logo = logo
		hostConsulted.Title = title
		hostConsulted.ServersChange = servers_change
		hostConsulted.IsDown = is_down
		endpoints = append(endpoints, hostConsulted)
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	json.NewEncoder(ctx).Encode(endpoints)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/endpoints", getEndpoints)

	log.Fatal(fasthttp.ListenAndServe(":7777", CORS(router.Handler)))
}
