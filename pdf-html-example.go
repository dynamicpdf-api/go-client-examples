package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	baseUrl := "https://api.dynamicpdf.com/"
	apiKey := "DP--api-key--"
	basePath := "./resources/converting-html-pdf-endpoint/"
	outputPath := "./output/"

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.BaseUrl = baseUrl
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
		os.WriteFile(outputPath+"pdf-html-go-output.pdf",
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}

}
