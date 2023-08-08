package main

import (
	"fmt"
	"os"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
)

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pdfExample.Endpoint.ApiKey = "DP---API-KEY---"

	inputOne := input.NewHtmlInputWithString("<html>An example HTML fragment.</html>")
	inputTwo := input.NewHtmlInputWithString("<html><p>HTML with basepath.</p><img src='./images/logo.png'></html>")
	inputTwo.SetBasePath("https://www.dynamicpdf.com")

	pdfExample.Inputs = append(pdfExample.Inputs, inputOne)
	pdfExample.Inputs = append(pdfExample.Inputs, inputTwo)

	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("c:/temp/html-to-pdf/html-output-go.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}
}
