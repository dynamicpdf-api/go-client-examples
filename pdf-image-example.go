package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/dynamicpdf-api/go-client/v2/imaging"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var baseUrl string
var outputPath string

func init() {
	basePath = "./resources/pdf-image/"
	apiKey = "DP--api-key--"
	baseUrl = "https://api.dpdf.io"
	outputPath = "./output/pdf-image-go-output"
}

func main() {

	resource := resource.NewPdfResourceWithResourcePath(basePath+"pdfnumberedinput.pdf", "pdfnumberedinput.pdf")
	pdfImage := imaging.NewPdfImage(resource)
	pdfImage.ImageFormat = imaging.NewPngImageFormat().ImageFormat
	pdfImage.Endpoint.BaseUrl = baseUrl
	pdfImage.Endpoint.ApiKey = apiKey

	resp := pdfImage.Process()
	res := <-resp

	if res.IsSuccessful() == true {

		for i, image := range res.Images {
			img, err := base64.StdEncoding.DecodeString(image.Data)
			if err != nil {
				fmt.Print(err)
				return
			}

			filePath := outputPath + strconv.Itoa(i) + ".png"
			os.Remove(filePath)
			os.WriteFile(filePath, img, os.ModeType)
		}

	} else {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	}
}
