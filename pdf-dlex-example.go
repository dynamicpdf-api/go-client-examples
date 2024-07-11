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

func init() {
	basePath = "./resources/creating-pdf-pdf-endpoint/"
	apiKey = "Dp--api-key--"
	outputPath = "./output/merge-pdfs-go-output.pdf"
}

func main() {
	processCloudPdf()
	processLocalPdf()
}

func processLocalPdf() {

	pr := endpoint.NewPdf()
	pr.Endpoint.ApiKey = apiKey

	theDlexResource := resource.NewDlexResourceWithPath(basePath+"SimpleReportWithCoverPage.dlex", "SimpleReportWithCoverPage.dlex")
	pr.AddDlexWithDlexResourceNLayoutDataPath(*theDlexResource, basePath+"SimpleReportWithCoverPage.json")

	additionResource := endpoint.NewDlexWithAdditionalResource(basePath+"Northwind Logo.gif", "Northwind Logo.gif")

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
		os.WriteFile(basePath+"pdf-dlex-pdf-local-output.pdf",
			res.Content().Bytes(), os.ModeType)
	}
}

func processCloudPdf(baseUrl string, apiKey string, basePath string) {

	pr := endpoint.NewPdf()
	pr.Endpoint.BaseUrl = baseUrl
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
