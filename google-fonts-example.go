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
	outputPath = "./output/pdf-google-fonts-go-example-output.pdf"
}

func main() {

	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.ApiKey = apiKey

	pageInput := input.NewPage()
	pageNumber := element.NewPageNumberingElement("test", "TopLeft", 0, 0)
	pageNumber.SetFontSize(72)
	pageNumber.TextFont()
	pageNumber.SetFont(*font.Google("Borel"))

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	myTextElement := element.NewText("myTest", element.TopLeft, 100, 200)
	myTextElement.SetFont(*font.Google("Borel"))

	pageInput.Elements = append(pageInput.Elements, myTextElement)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)

	resp := pdfCl.Process()
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
