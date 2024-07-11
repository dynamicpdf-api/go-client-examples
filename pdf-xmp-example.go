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
	basePath = "./resources/get-xmp-metadata-pdf-xmp-endpoint//"
	apiKey = "Dp--api-key--"
	baseUrl = "https://api.dpdf.io"
}

func main() {

	resource := resource.NewPdfResourceWithResourcePath(basePath+"fw4.pdf", "fw4.pdf")
	xmp := endpoint.NewPdfXmp(resource)
	xmp.Endpoint.BaseUrl = baseUrl
	xmp.Endpoint.ApiKey = apiKey

	resp := xmp.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		fmt.Print(string(res.Content().Bytes()))
	}
}
