package main

import (
	"encoding/json"
	"fmt"
	"log"

	"./model"
	"./repository"

	"github.com/buaazp/fasthttprouter"
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

const domainInfo = "https://api.ssllabs.com/api/v3/analyze?host="

func getInfoDomain(ctx *fasthttp.RequestCtx) {
	var host model.Domain
	domain := ctx.UserValue("domain").(string)
	url := domainInfo + domain
	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("Error al conectarse al api")
		return
	}
	if statusCode == 200 {
		json.Unmarshal(body, &host)
		repository.SaveInfoSearch(host)
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	json.NewEncoder(ctx).Encode(host)
}

/**func collectData(domain string, host *model.Domain) *model.Domain {
	c := colly.NewCollector()
	var title = ""

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Attr()
		fmt.Println(title)
	})
	c.Visit(domain)
	host.Title = title
	return host
}**/

func main() {
	repository.CreateTables()
	defer repository.Connet().Close()
	router := fasthttprouter.New()
	router.GET("/domains/:domain", getInfoDomain)
	log.Fatal(fasthttp.ListenAndServe(":5555", CORS(router.Handler)))
}
