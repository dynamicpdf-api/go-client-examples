package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	processCloud()
	processLocal()

}
func processCloud() {
	layoutDataResource := resource.NewLayoutDataResource("./resources/creating-pdf-dlex-layout/creating-pdf-dlex-layout.json", 
		"creating-pdf-dlex-layout.json")

	layoutData := endpoint.NewDlexEndpoint("samples/creating-pdf-dlex-layout-endpoint/creating-pdf-dlex-layout.dlex", 
		layoutDataResource)
	layoutData.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	layoutData.Endpoint.ApiKey = "DP--api-key--"
	resp := layoutData.Process()

	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("./output/dlex-layout-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}

func processLocal() {
	layoutDataResource := resource.NewLayoutDataResource("./resources/creating-pdf-dlex-layout/creating-pdf-dlex-layout.json", 
		"creating-pdf-dlex-layout.json")

	theDlexResource := resource.NewDlexResourceWithPath("./resources/creating-pdf-dlex-layout/creating-pdf-dlex-layout.dlex", 
		"creating-pdf-dlex-layout.dlex")

	theDlexEndpoint := endpoint.NewDlexEndpointWithResource(*theDlexResource, layoutDataResource)

	additionalResource := endpoint.NewDlexWithAdditionalResource("./resources/creating-pdf-dlex-layout/creating-pdf-dlex-layout.png", 
		"creating-pdf-dlex-layout.png")

	theDlexEndpoint.Resources = append(theDlexEndpoint.Resources, additionalResource)

	theDlexEndpoint.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	theDlexEndpoint.Endpoint.ApiKey = "DP--api-key--"
	resp := theDlexEndpoint.Process()

	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("./output/dlex-layout-local-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}
