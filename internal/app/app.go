package app

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"market/internal/adapters/brand"
	"market/internal/adapters/category"
	"market/internal/adapters/product"
	"market/internal/adapters/worker"
	"market/internal/config"
	brand2 "market/internal/database/brand"
	category2 "market/internal/database/category"
	product2 "market/internal/database/product"
	worker2 "market/internal/database/worker"
	"market/pkg/database/postgresql"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Run(pathConfig string) {
	conf := config.GetConfig(pathConfig)
	db, err := postgresql.NewConnection(5, conf)
	if err != nil {
		fmt.Println(err)
	}
	serverStart(routerCreate(db), conf)
}

func routerCreate(db *pgxpool.Pool) *httprouter.Router {
	router := httprouter.New()

	brSt := brand2.NewBrand(db)
	handlerBrand := brand.NewHeandler(brSt)

	catSt := category2.NewCategory(db)
	handlerCat := category.NewHeandler(catSt)

	worSt := worker2.NewWorker(db)
	handlerWorker := worker.NewHeandler(worSt)

	proSt := product2.NewProduct(db)
	handlerProduct := product.NewHeandler(proSt)

	handlerWorker.Register(router)
	handlerProduct.Register(router)
	handlerBrand.Register(router)
	handlerCat.Register(router)

	return router
}

func serverStart(router *httprouter.Router, conf *config.Config) {
	host := fmt.Sprintf("%s:%s", conf.SrvHost, conf.SrvPort)
	fmt.Printf("Start server http://%s\n", host)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	if err := server.Serve(listener); err != nil {
		log.Fatal("Can't serve")
	}
}
