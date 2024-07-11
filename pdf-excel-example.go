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
	basePath = "./resources/users-guide/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/excel-to-pdf-output-go.pdf"
}

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.ApiKey = apiKey
	excelResource := resource.NewExcelResourceWithResourcePath(basePath+"sample-data.xlsx", "sample-data.xlsx")
	pdfInput := input.NewExcelInputWithResource(excelResource)
	pdfExample.Inputs = append(pdfExample.Inputs, pdfInput)
	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath)
		os.WriteFile(outputPath, res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}

}
