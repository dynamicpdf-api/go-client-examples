package main

import (
	"fmt"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	resource := resource.NewPdfResourceWithResourcePath("C:/temp/dynamicpdf-api-samples/get-xmp-metadata-pdf-xmp-endpoint/fw4.pdf", "fw4.pdf")
	xmp := endpoint.NewPdfXmp(resource)
	xmp.Endpoint.BaseUrl = "https://api.dpdf.io/"
	xmp.Endpoint.ApiKey = "DP--api-key--"

	resp := xmp.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		fmt.Print(string(res.Content().Bytes()))
	}
}
