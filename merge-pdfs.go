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
	basePath = "./resources/merge-pdfs-pdf-endpoint/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/merge-pdfs-go-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	prInput := input.NewPdfWithResource(pdfResource)
	prInput.StartPage = 1
	prInput.PageCount = 1

	pr.Inputs = append(pr.Inputs, prInput)

	pdfResource2 := resource.NewPdfResourceWithResourcePath(basePath+"DocumentB.pdf", "DocumentB.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource2)
	pr.Inputs = append(pr.Inputs, prInput2)

	mergeOption := input.NewMergeOptions()
	prInput3 := input.NewPdfWithCloudPath("samples/merge-pdfs-pdf-endpoint/DocumentC.pdf", mergeOption)
	pr.Inputs = append(pr.Inputs, prInput3)

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
		os.WriteFile(outputPath, res.Content().Bytes(), os.ModeType)
	}

}
