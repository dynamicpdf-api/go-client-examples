package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	pdfAcro := endpoint.NewPdf()
	pdfAcro.Endpoint.BaseUrl = "https://api.dpdf.io"
	pdfAcro.Endpoint.ApiKey = "DP--api-key--"
	outputPath := "./output/"
	basePath := "./resources/fill-acro-form-pdf-endpoint/"

	pdfResource := resource.NewPdfResourceWithResourcePath(basePath + "fw9AcroForm_18.pdf", "fw9AcroForm_18.pdf")

	pdfInput := input.NewPdfWithResource(pdfResource)
	pdfAcro.Inputs = append(pdfAcro.Inputs, pdfInput)

	field1 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_1[0]", "Any Company, Inc.")
	field1.Remove = true
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field1)
	
	field2 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_2[0]", "Any Company")
	field2.Flatten = true
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field2)

	field3 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].FederalClassification[0].c1_1[0]", "1")
	field3.Flatten = false
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field3)

	field4 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].Address[0].f1_7[0]", "123 Main Street")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field4)
	
	field5 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].Address[0].f1_8[0]", "Washington, DC  22222")
	field5.Remove = true

	pdfAcro.FormFields = append(pdfAcro.FormFields, *field5)
	
	resp := pdfAcro.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(outputPath + "form-flatten-delete-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}
}
