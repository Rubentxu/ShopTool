package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	prod "shopTool/products"

	"github.com/go-kit/kit/log"
)

func main() {
	var (
		httpAddr  = flag.String("http.addr", ":8081", "HTTP listen address")
		mongoAddr = flag.String("mongo.addr", "localhost", "mongoDb address")
	)
	flag.Parse()
	fmt.Println("ListenAndServe url: ", *httpAddr)
	fmt.Println("mongoDb url: ", *mongoAddr)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	/*
		var s prod.ProductService
		{
			s, _ = prod.NewProductService()
			s = prod.LoggingMiddleware(logger)(s)
		}

		var h http.Handler
		{
			h = prod.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
		}

		errs := make(chan error)
		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errs <- fmt.Errorf("%s", <-c)
		}() */

	/* 	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}() */
	h, _ := prod.NewHandler(*mongoAddr)
	logger.Log(http.ListenAndServe(*httpAddr, h))

}
