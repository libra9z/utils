package main

import (
	"golang.org/x/net/context"
    "github.com/go-kit/kit/endpoint"
    "encoding/json"
    "net/http"
)

type WeatherService interface{
	GetWeather(params ...string)(string ,error)
}

type weatherService struct {}

type weatherRequest struct {
	Areaid	string
	Date	string
}

type weatherResponse struct {
    Result   	string 	`json:"result"`
    Err 		string 	`json:"err,omitempty"` // errors don't define JSON marshaling	
}


func (svc weatherService)GetWeather(params ...string)(string ,error){
	return "",nil
}


func makeWeatherEndpoint(svc WeatherService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(weatherRequest)
        v, err := svc.GetWeather(req.Areaid)
        if err != nil {
            return weatherResponse{v, err.Error()}, nil
        }
        return weatherResponse{v, "success"}, nil
    }
}

func decodeRequest(r *http.Request) (interface{}, error) {
    var request weatherRequest
    
    return request, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
