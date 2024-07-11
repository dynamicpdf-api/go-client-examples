package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/add-bookmarks/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/add-bookmarks-pdfs-go-output.pdf"
}

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	prInput := input.NewPdfWithResource(pdfResource)

	pr.Inputs = append(pr.Inputs, prInput)

	pdfResource2 := resource.NewPdfResourceWithResourcePath(basePath+"DocumentB.pdf", "DocumentB.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource2)
	pr.Inputs = append(pr.Inputs, prInput2)

	pdfResource3 := resource.NewPdfResourceWithResourcePath(basePath+"DocumentC.pdf", "DocumentC.pdf")
	prInput3 := input.NewPdfWithResource(pdfResource3)
	pr.Inputs = append(pr.Inputs, prInput3)

	outline := pr.Outlines.Add("Three Bookmarks")
	outline.Expanded = true
	outline.SetColor(color.NewRgbColorDefault().Red().Color)

	out := outline.Children()

	out.AddWithInputValue("DocumentA", prInput.Input, 0, endpoint.FitPage)
	out.AddWithInputValue("DocumentB", prInput2.Input, 0, endpoint.FitPage)
	out.AddWithInputValue("DocumentC", prInput3.Input, 0, endpoint.FitPage)

	out.AddWithURL("DynamicPDF Cloud API", endpoint.NewUrlAction("https://cloud.dynamicpdf.com/"))

	outline.GetChildren()[0].SetColor(color.NewRgbColorDefault().Orange().Color)
	outline.GetChildren()[1].SetColor(color.NewRgbColorDefault().Green().Color)
	outline.GetChildren()[3].SetColor(color.NewRgbColorDefault().Blue().Color)

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
