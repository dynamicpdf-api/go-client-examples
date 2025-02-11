package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

var basePath string
var apiKey string
var outputPath string
var basePathLocal string

func init() {
	basePath = "./resources/creating-pdf-pdf-endpoint/"
	basePathLocal = "./resources/creating-pdf-dlex-layout/"
	apiKey = "DP.Ugguk0HRu/CxWc+I7xeOH0NSqYZZwYdn0jYWB65rgfn2T6ImN1zUhLpS"
	outputPath = "./output/merge-pdfs-go-output.pdf"
}

func main() {
	processCloudPdf()
	processLocalPdf()
	processLocalPdfTwo()
}

func processLocalPdf() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	theDlexResource := resource.NewDlexResourceWithPath(basePath+"SimpleReportWithCoverPage.dlex", "SimpleReportWithCoverPage.dlex")
	pr.AddDlexWithDlexResourceNLayoutDataPath(*theDlexResource, basePath+"SimpleReportWithCoverPage.json")

	pr.AddAdditionalResource(basePath+"Northwind Logo.gif", "Northwind Logo.gif")

	// The dlex has an additional image in it, I'm using local dlex, need to add to the pdf endpoint

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"pdf-dlex-pdf-local-two-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}
}

func processLocalPdfTwo() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	theDlexResource := resource.NewDlexResourceWithPath(basePathLocal+"ExampleTemplate.dlex", "ExampleTemplate.dlex")
	pr.AddDlexWithDlexResourceNLayoutDataPath(*theDlexResource, basePathLocal+"ExampleTemplate.json")

	pr.AddAdditionalResource(basePathLocal+"signature-one.png", "signature-one.png")
	pr.AddAdditionalResource(basePathLocal+"template_example.pdf", "template_example.pdf")

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"pdf-dlex-pdf-local-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}
}

func processCloudPdf() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	layoutDataResource := resource.NewLayoutDataResource(basePath+"SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")

	pr.AddDlexWithCloudResourceNLayoutData("samples/creating-pdf-pdf-endpoint/SimpleReportWithCoverPage.dlex", layoutDataResource)

	resp := pr.Process()
	res := <-resp

	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			fmt.Print("Failed with error: " + res.ClientError().Error())
		} else {
			fmt.Print("Failed with error: " + res.ErrorJson())
		}
	} else {
		os.WriteFile(basePath+"pdf-dlex-pdf-remote-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}

}
