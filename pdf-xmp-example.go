package main

import (
	"fmt"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	resource := resource.NewPdfResourceWithResourcePath("C:/temp/dynamicpdf-api-samples/fw4.pdf", "fw4.pdf")
	xmp := endpoint.NewPdfXmp(resource)
	xmp.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	xmp.Endpoint.ApiKey = "DP.xxx-api-key-xxx"

	resp := xmp.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		fmt.Print(string(res.Content().Bytes()))
	}
}
