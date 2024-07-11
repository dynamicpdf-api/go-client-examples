package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/outlines/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/outlines-go-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	pageInput := pr.AddPage()
	textElement := element.NewText("Hello World 1", element.TopCenter, 0, 0)
	pageInput.Elements = append(pageInput.Elements, textElement)

	pageInput2 := pr.AddPage()
	textElement2 := element.NewText("Hello World 2", element.TopCenter, 0, 0)
	pageInput2.Elements = append(pageInput2.Elements, textElement2)

	pageInput3 := pr.AddPage()
	textElement3 := element.NewText("Hello World 3", element.TopCenter, 0, 0)
	pageInput3.Elements = append(pageInput3.Elements, textElement3)

	outline := pr.Outlines.Add("Root Outline")
	outline.Children().AddWithInputValue("Page 1", pageInput.Input, 0, endpoint.FitPage)
	outline.Children().AddWithInputValue("Page 2", pageInput2.Input, 0, endpoint.FitPage)
	outline.Children().AddWithInputValue("Page 3", pageInput3.Input, 0, endpoint.FitPage)

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"PdfOutlineInput.pdf", "PdfOutlineInput.pdf")
	mergeOptions := input.NewMergeOptions()

	outlines := false
	mergeOptions.Outlines = &outlines

	pdfInput := pr.AddPdf(pdfResource, mergeOptions)
	pdfInput.Input.SetId("pdfoutlineinput")
	outline.Children().AddPdfOutlines(*pdfInput)

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
