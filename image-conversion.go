package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/position"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var baseUrl string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/image-conversion/"
	apiKey = "DP--api-key--"
	outputPath = "./output/image-conversion-go-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	imageResource1 := resource.NewImageResourceWithResourcePath(basePath+"testimage.tif", "testimage.tif")

	imageInput1 := input.NewImagewithImageResource(imageResource1)
	imageInput1.Align = position.Center

	imageInput1.ExpandToFit = false
	imageInput1.SetPageHeight(1008)
	imageInput1.SetPageWidth(612)
	imageInput1.VAlign = position.Top
	imageInput1.Align = position.Center

	pr.Inputs = append(pr.Inputs, imageInput1)

	imageResource2 := resource.NewImageResourceWithResourcePath(basePath+"dynamicpdfLogo.png", "dynamicpdfLogo.png")
	imageInput2 := pr.AddImage(imageResource2)
	imageInput2.ExpandToFit = false
	imageInput2.SetPageWidth(1008)
	imageInput2.SetPageHeight(612)
	imageInput2.VAlign = position.Middle
	imageInput2.Align = position.Center
	imageInput2.ScaleX = .5
	imageInput2.ScaleY = .5

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
