package main

import (
	"golang.org/x/net/context"
    httptransport "github.com/go-kit/kit/transport/http"
    endpoint "github.com/go-kit/kit/endpoint"
    kitlog "github.com/go-kit/kit/log"
    "net/http"
    "os"
)

var logger kitlog.Logger

func main(){
	ctx := context.Background()

	logger = kitlog.NewLogfmtLogger(os.Stderr)

	var svc WeatherService
	
	svc = weatherService{}
	
	
	var weather endpoint.Endpoint
	weather = makeWeatherEndpoint(svc)

	
	weatherHandler := httptransport.NewServer(
		ctx,
		weather,
		decodeRequest,
		encodeResponse,
	)
	
		
	http.Handle("/v1/weather",weatherHandler)
	logger.Log("err",http.ListenAndServe(":12202",nil))		
}