package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	layoutDataResource := resource.NewLayoutDataResource("c:/temp/dynamicpdf-api-samples/error-handling-example/SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")

	layoutData := endpoint.NewDlexEndpoint("samples/error-handling-example/SimpleReportWithCoverPage.dlex", layoutDataResource)
	layoutData.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	layoutData.Endpoint.ApiKey = "<API-KEY>"
	resp := layoutData.Process()
	res := <-resp

	fmt.Println("HTTP Response: " + strconv.Itoa(res.StatusCode()))

	if res.IsSuccessful() == true {
		os.WriteFile("c:/temp/dynamicpdf-api-samples/dlex-layout-example/output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}
