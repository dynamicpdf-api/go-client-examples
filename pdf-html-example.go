package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/converting-html-pdf-endpoint/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/pdf-html-go-output.pdf"
}

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.ApiKey = apiKey

	inputOne := input.NewHtmlInputWithString("<html>An example HTML fragment.</html>")

	inputTwo := input.NewHtmlInputWithString("<html><p>HTML with basePath.</p><img src='./images/logo.png'></html>")
	inputTwo.SetBasePath("https://www.dynamicpdf.com")

	htmlResource := resource.NewHtmlResource(basePath+"products.html", "products.html")

	inputThree := input.NewHtmlInputWithResource(htmlResource)

	pdfExample.Inputs = append(pdfExample.Inputs, inputOne)
	pdfExample.Inputs = append(pdfExample.Inputs, inputTwo)
	pdfExample.Inputs = append(pdfExample.Inputs, inputThree)

	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath)
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}

}
