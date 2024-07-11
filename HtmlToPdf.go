package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	apiKey = "Dp--api-key--"
	outputPath = "./output/html-output-go.pdf"
}

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.ApiKey = apiKey

	inputOne := input.NewHtmlInputWithString("<html>An example HTML fragment.</html>")
	inputTwo := input.NewHtmlInputWithString("<html><p>HTML with basepath.</p><img src='./images/logo.png'></html>")
	inputTwo.SetBasePath("https://www.dynamicpdf.com")

	pdfExample.Inputs = append(pdfExample.Inputs, inputOne)
	pdfExample.Inputs = append(pdfExample.Inputs, inputTwo)

	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath)
		os.WriteFile(outputPath, res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}
}
