package main

import (
	"fmt"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var baseUrl string

func init() {
	basePath = "./resources/get-pdf-info-pdf-info-endpoint/"
	apiKey = "Dp--api-key--"
	baseUrl = "https://api.dpdf.io"
}

func main() {

	resource := resource.NewPdfResourceWithResourcePath(basePath+"fw4.pdf", "fw4.pdf")
	txt := endpoint.NewPdfText(resource, 1, 2)
	txt.Endpoint.BaseUrl = baseUrl
	txt.Endpoint.ApiKey = apiKey

	resp := txt.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		fmt.Print(string(res.Content().Bytes()))
	}
}
