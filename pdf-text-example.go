package main

import (
	"fmt"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	resource := resource.NewPdfResourceWithResourcePath("./resources/pdf-info/fw4.pdf", "fw4.pdf")
	txt := endpoint.NewPdfText(resource,1,2)
    txt.Endpoint.BaseUrl = "https://api.dpdf.io/"
    txt.Endpoint.ApiKey = "DP--api-key--"

    resp := txt.Process()
    res := <-resp
	
	if res.IsSuccessful() == true {
		fmt.Print(string(res.Content().Bytes()))
	}
}