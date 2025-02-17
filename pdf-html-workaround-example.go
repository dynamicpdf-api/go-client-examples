package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/converting-html-pdf-endpoint/"
	apiKey = "DP--api-key--"
	outputPath = "./output/pdf-html-css-go-output.pdf"
}

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.ApiKey = apiKey

	tempHtmlBytes, err := os.ReadFile(basePath + "example.html")
	if err != nil {
		fmt.Println("Error reading CSS file:", err)
		return
	}

	tempCssBytes, err := os.ReadFile(basePath + "example.css")
	if err != nil {
		fmt.Println("Error reading CSS file:", err)
		return
	}

	tempHtml := string(tempHtmlBytes)
	tempCss := string(tempCssBytes)

	linkIndex := strings.Index(tempHtml, "<link")
	if linkIndex == -1 {
		fmt.Println("No <link> tag found in HTML")
		return
	}

	sb := tempHtml[:linkIndex] + "<style>" + tempCss + "</style>"

	tempHtml = tempHtml[linkIndex:]
	closeTagIndex := strings.Index(tempHtml, "/>")
	if closeTagIndex != -1 {
		sb += tempHtml[closeTagIndex+2:]
	}

	inputOne := input.NewHtmlInputWithString(sb)
	pdfExample.Inputs = append(pdfExample.Inputs, inputOne)

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
