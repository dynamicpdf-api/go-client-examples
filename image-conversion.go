package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var baseUrl string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/image-conversion/"
	apiKey = "DP--api-key--"
	outputPath = "./output/image-conversion-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	imageResource1 := resource.NewImageResourceWithResourcePath(basePath+"testimage.tif", "testimage.tif")

	imageInput1 := input.NewImagewithImageResource(imageResource1)
	pr.Inputs = append(pr.Inputs, imageInput1)

	imageResource2 := resource.NewImageResourceWithResourcePath(basePath+"dynamicpdfLogo.png", "dynamicpdfLogo.png")
	imageInput2 := pr.AddImage(imageResource2)
	imageInput2.ScaleX = 25
	imageInput2.ScaleY = 25

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	}

}
