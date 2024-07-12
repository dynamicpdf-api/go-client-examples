package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

var apiKey string
var outputPath string

func init() {
	apiKey = "DP--api-key--"
	outputPath = "./output/barcode-example-go-output.pdf"
}

func main() {

	barcodePdf := endpoint.NewPdf()
	barcodePdf.Endpoint.ApiKey = apiKey

	pageInput := input.NewPage()
	pageInput.SetPageHeight(621)
	pageInput.SetPageWidth(1008)

	barcode := element.NewCode11Barcode("12345678910", element.TopCenter, 200, 50, 50)
	//barcode.Placement = element.TopCenter
	barcode.SetColor(color.NewRgbColorDefault().Red().Color)

	pageInput.Elements = append(pageInput.Elements, barcode)
	barcodePdf.Inputs = append(barcodePdf.Inputs, pageInput)

	fmt.Print(barcodePdf.GetInstructionsJson(true))

	resp := barcodePdf.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.Remove(outputPath)
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("Error Message:" + res.ErrorMessage())
		fmt.Print(res.ErrorJson())
	}
}
