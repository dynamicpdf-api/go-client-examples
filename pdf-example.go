package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/font"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

var apiKey string
var outputPath string

func init() {
	apiKey = "Dp--api-key--"
	outputPath = "./output/pdf-page-example-go-output.pdf"
}

func main() {

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.ApiKey = apiKey

	pageInput := input.NewPage()
	pdfExample.Author = "John Doe"
	pdfExample.Title = "My Blank PDF Page"
	pageInput.SetPageHeight(612)
	pageInput.SetPageWidth(1008)

	pageNumber := element.NewPageNumberingElement("1", "TopRight", 0, 0)
	pageNumber.SetFontSize(24)
	pageNumber.SetFont(font.Courier())

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pdfExample.Inputs = append(pdfExample.Inputs, pageInput)

	resp := pdfExample.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error - error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error json: " + res.ErrorJson())
		}
	} else {
		os.Remove(outputPath)
		os.WriteFile(outputPath,
			res.Content().Bytes(), os.ModeType)
	}

}
