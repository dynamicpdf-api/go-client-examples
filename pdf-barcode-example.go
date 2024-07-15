package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var apiKey string
var outputPath string
var basePath string

func init() {
	apiKey = "DP--api-key--"
	outputPath = "./output/"
	basePath = "./resources/merge-pdfs-pdf-endpoint/"
}

func main() {
	barcodeNew()
	barcodeExisting()

}

func barcodeNew() {
	barcodePdf := endpoint.NewPdf()
	barcodePdf.Endpoint.ApiKey = apiKey

	pageInput := input.NewPage()
	pageInput.SetPageHeight(621)
	pageInput.SetPageWidth(1008)

	barcode := element.NewCode11Barcode("12345678910", element.TopCenter, 200, 50, 50)
	barcode.SetColor(color.NewRgbColorDefault().Red().Color)

	pageInput.Elements = append(pageInput.Elements, barcode)
	barcodePdf.Inputs = append(barcodePdf.Inputs, pageInput)

	fmt.Print(barcodePdf.GetInstructionsJson(true))

	resp := barcodePdf.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath + "barcode-new-example-go-output.pdf")
		os.WriteFile(outputPath + "barcode-new-example-go-output.pdf",
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("Error Message:" + res.ErrorMessage())
		fmt.Print(res.ErrorJson())
	}
}

func barcodeExisting() {

	barcodePdf := endpoint.NewPdf()
	barcodePdf.Endpoint.ApiKey = apiKey

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath+"DocumentA.pdf", "DocumentA.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	barcodePdf.Inputs = append(barcodePdf.Inputs, pdfInput)

	pdfTemplate := element.NewTemplate()
	pdfTemplate.Id = "Temp1"

	barcode := element.NewAztecBarcodeElement("Hello World", element.TopCenter, 0, 500)
	barcode.SetColor(color.NewRgbColorDefault().Red().Color)

	pdfTemplate.Elements = append(pdfTemplate.Elements, barcode)

	pdfInput.SetTemplate(pdfTemplate)

	resp := barcodePdf.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath + "barcode-existing-example-go-output.pdf")
		os.WriteFile(outputPath + "barcode-existing-example-go-output.pdf",
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("Error Message:" + res.ErrorMessage())
		fmt.Print(res.ErrorJson())
	}
}
