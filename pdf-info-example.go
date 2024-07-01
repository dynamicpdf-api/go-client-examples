package main

import (
	"fmt"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {
	basePath := "./resources/get-pdf-info-pdf-info-endpoint/"
	resource := resource.NewPdfResourceWithResourcePath(basePath+"fw4.pdf", "fw4.pdf")
	text := endpoint.NewPdfInfoResource(resource)

	text.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	text.Endpoint.ApiKey = "DP--api-key--"

	resp := text.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		fmt.Print(string(res.Content().Bytes()))
	}
}
