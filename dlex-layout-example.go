package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var outputPath string
var baseUrl string

func init() {
	basePath = "./resources/creating-pdf-dlex-layout/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/"
}

func main() {
	processCloud()
	processLocal()
}

func processCloud() {
	layoutDataResource := resource.NewLayoutDataResource(basePath+"creating-pdf-dlex-layout.json",
		"creating-pdf-dlex-layout.json")

	dlexEndpoint := endpoint.NewDlexEndpoint("samples/creating-pdf-dlex-layout-endpoint/creating-pdf-dlex-layout.dlex",
		layoutDataResource)

	dlexEndpoint.Endpoint.ApiKey = apiKey

	dlexEndpoint.Endpoint.BaseUrl = baseUrl
	resp := dlexEndpoint.Process()

	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath + "dlex-layout-remote-go-output.pdf")
		os.WriteFile(outputPath+"dlex-layout-remote-go-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}

func processLocal() {

	layoutDataResource := resource.NewLayoutDataResource(basePath+"creating-pdf-dlex-layout.json",
		"creating-pdf-dlex-layout.json")

	theDlexResource := resource.NewDlexResourceWithPath(basePath+"creating-pdf-dlex-layout.dlex",
		"creating-pdf-dlex-layout.dlex")

	theDlexEndpoint := endpoint.NewDlexEndpointWithResource(*theDlexResource, layoutDataResource)

	additionalResource := endpoint.NewDlexWithAdditionalResource(basePath+"creating-pdf-dlex-layout.png",
		"creating-pdf-dlex-layout.png")

	theDlexEndpoint.Resources = append(theDlexEndpoint.Resources, additionalResource)

	theDlexEndpoint.Endpoint.ApiKey = apiKey
	theDlexEndpoint.Endpoint.BaseUrl = baseUrl

	resp := theDlexEndpoint.Process()

	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath + "dlex-layout-local-go-output.pdf")
		os.WriteFile(outputPath+"dlex-layout-local-go-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}
