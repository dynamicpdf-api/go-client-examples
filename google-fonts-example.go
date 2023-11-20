package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/input"
)

func main() {

	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pdfCl.Endpoint.ApiKey = "DP.xxx-api-key-xxx"

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
		os.WriteFile("C:/temp/fonts-example/pdf-google-fonts-go-example-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
