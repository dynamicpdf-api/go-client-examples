package main

import (
	"fmt"
	"os"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
)

func main() {

	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	pdfCl.Endpoint.ApiKey = "DP---API-KEY---"


	pdfCl.Author = "Jane Doe"
	pdfCl.Title = "Sample PDF"
	pdfCl.Subject = "topLevel document metadata"
	pdfCl.Creator = "John Creator"
	pdfCl.Keywords = "dynamicpdf api example pdf java instructions"

	pageInput := input.NewPage()
	lineElement := element.NewLine(element.BottomCenter, 150, 200)
	pageInput.Elements = append(pageInput.Elements, lineElement)
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)
	fmt.Print("Endpoint Name: " + pdfCl.EndpointName())
	fmt.Print("Endpoint Url : " + pdfCl.BaseUrl())
	fmt.Print("Endpoint Key : " + pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile("C:/temp/dynamicpdf-api-samples/out/pdf-page-example-output.pdf", 
		res.Content().Bytes(), os.ModeType)
	}

}
