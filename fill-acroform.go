package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

var basePath string
var apiKey string
var outputPath string

func init() {
	basePath = "./resources/add-bookmarks/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/acroform-pdfs-go-output.pdf"
}

func main() {

	pdfAcro := endpoint.NewPdf()
	pdfAcro.Endpoint.ApiKey = apiKey

	pdfInput := input.NewPdfWithCloudPath("samples/fill-acro-form-pdf-endpoint/fw9AcroForm_18.pdf", input.NewMergeOptions())
	pdfAcro.Inputs = append(pdfAcro.Inputs, pdfInput)

	field1 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_1[0]", "Any Company, Inc.")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field1)
	field2 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_2[0]", "Any Company")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field2)
	field3 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].FederalClassification[0].c1_1[0]", "1")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field3)
	field4 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].Address[0].f1_7[0]", "123 Main Street")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field4)
	field5 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].Address[0].f1_8[0]", "Washington, DC  22222")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field5)
	field6 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_9[0]", "Any Requester")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field6)
	field7 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_10[0]", "17288825617")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field7)
	field8 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].EmployerID[0].f1_14[0]", "52")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field8)
	field9 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].EmployerID[0].f1_15[0]", "1234567")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field9)

	resp := pdfAcro.Process()
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
