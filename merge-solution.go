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
	basePath = "./resources/"
	apiKey = "DP--api-key--"
	outputPath = "./output/merge-solution-pdfs-go-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"merge-pdfs-pdf-endpoint/DocumentA.pdf", "DocumentA.pdf")
	prInput := input.NewPdfWithResource(pdfResource)
	pr.Inputs = append(pr.Inputs, prInput)

	pdfResource2 := resource.NewPdfResourceWithResourcePath(basePath+"merge-pdfs-pdf-endpoint/DocumentB.pdf", "DocumentB.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource2)
	pr.Inputs = append(pr.Inputs, prInput2)

	wordResource := resource.NewWordResourceWithResourcePath(basePath+"users-guide/Doc1.docx", "Doc1.docx")
	pdfInput := input.NewWordInputWithResource(wordResource)
	pr.Inputs = append(pr.Inputs, pdfInput)

	imageResource1 := resource.NewImageResourceWithResourcePath(basePath+"image-conversion/testimage.tif", "testimage.tif")

	imageInput1 := input.NewImagewithImageResource(imageResource1)
	pr.Inputs = append(pr.Inputs, imageInput1)

	layoutDataResource := resource.NewLayoutDataResource(basePath+"/creating-pdf-dlex-layout/creating-pdf-dlex-layout.json",
		"creating-pdf-dlex-layout.json")

	pr.AddDlexWithCloudResourceNLayoutData("samples/creating-pdf-dlex-layout-endpoint/creating-pdf-dlex-layout.dlex",
		layoutDataResource)

	htmlResource := resource.NewHtmlResource(basePath+"users-guide/products.html", "products.html")
	htmlInput := input.NewHtmlInputWithResource(htmlResource)
	pr.Inputs = append(pr.Inputs, htmlInput)

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
