package main

import (
	"fmt"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var outputPath string
var baseUrl string

func init() {
	basePath = "./resources/image-info/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/form-flatten-delete-go-output.pdf"
	baseUrl = "https://api.dpdf.io"
}

func RunOne() {

	resource := resource.NewImageResourceWithResourcePath(basePath+"/getting-started.png", "")
	imageInfo := endpoint.NewImageInfo(resource)

	imageInfo.Endpoint.BaseUrl = baseUrl
	imageInfo.Endpoint.ApiKey = apiKey

	resp := imageInfo.Process()
	res := <-resp

	if res.IsSuccessful() == true {

		fmt.Print(string(res.Content().Bytes()))
	}
}

func RunTwo() {

	resource := resource.NewImageResourceWithResourcePath(basePath+"/multipage.tiff", "")
	imageInfo := endpoint.NewImageInfo(resource)

	imageInfo.Endpoint.ApiKey = apiKey
	imageInfo.Endpoint.BaseUrl = baseUrl

	resp := imageInfo.Process()
	res := <-resp

	if res.IsSuccessful() == true {

		fmt.Print(string(res.Content().Bytes()))
	}
}

func main() {
	RunOne()
	RunTwo()
}
