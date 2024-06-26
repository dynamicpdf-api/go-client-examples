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

	pdfExample := endpoint.NewPdf()
	pdfExample.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pdfExample.Endpoint.ApiKey = "DP--api-key--"

	pageInput := input.NewPage()
	pdfExample.Author = "John Doe"
	pdfExample.Title = "My Blank PDF Page"
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008

	pageNumber := element.NewPageNumberingElement("1", "TopRight", 0, 0 )
	pageNumber.SetFontSize(24)
	pageNumber.SetFont(font.Courier())
	
	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pdfExample.Inputs = append(pdfExample.Inputs, pageInput)

	fmt.Print("here")

    resp := pdfExample.Process()
    res := <-resp
	

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error - error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error json: " + res.ErrorJson())
		}
	} else{
		os.WriteFile("C:/temp/dynamicpdf-api-samples/pdf-page-example-output.pdf", 
		res.Content().Bytes(), os.ModeType)
	}

}
