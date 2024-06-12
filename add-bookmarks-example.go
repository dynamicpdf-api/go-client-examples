package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pr.Endpoint.ApiKey = "DP--api-key--"
	basePath := "c:/temp/dynamicpdf-api-samples/"

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
		os.WriteFile(basePath+"add-bookmarks-pdf-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}
}
