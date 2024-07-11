package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var baseUrl string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/image-conversion/"
	apiKey = "DP--api-key--"
	outputPath = "./output/image-conversion-multipage-tiff-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	imageResource := resource.NewImageResourceWithResourcePath(basePath+"MultiPageTiff.tif", "multipage.tif")
	pr.AddImage(imageResource)

	fmt.Print(pr.GetInstructionsJson(true))

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(outputPath, res.Content().Bytes(), os.ModeType)
	}

}
