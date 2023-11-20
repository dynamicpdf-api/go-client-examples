package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	layoutDataResource := resource.NewLayoutDataResource("c:/temp/dynamicpdf-api-samples/creating-pdf-dlex-layout.json", "creating-pdf-dlex-layout.json")

	layoutData := endpoint.NewDlexEndpoint("samples/creating-pdf-dlex-layout-endpoint/creating-pdf-dlex-layout.dlex", layoutDataResource)
	layoutData.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	layoutData.Endpoint.ApiKey = "DP.xxx-api-key-xxx"
	resp := layoutData.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("c:/temp/dynamicpdf-api-samples/dlex-layout-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}
