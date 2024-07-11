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
	basePath = "./resources/delete-pages/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/delete-pages-go-output.pdf.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"pdfnumberedinput.pdf", "pagenumberedinput1.pdf")
	prInput := pr.AddPdf(pdfResource, input.NewMergeOptions())
	prInput.StartPage = 1
	prInput.PageCount = 3

	pr.Inputs = append(pr.Inputs, prInput)

	pdfResource2 := resource.NewPdfResourceWithResourcePath(basePath+"pdfnumberedinput.pdf", "pagenumberedinput2.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource2)
	prInput2.StartPage = 6
	prInput2.PageCount = 2

	pr.Inputs = append(pr.Inputs, prInput2)

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.Remove(outputPath)
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	}

}
