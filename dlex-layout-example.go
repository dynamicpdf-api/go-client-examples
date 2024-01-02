package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func main() {

	processCloud()
	processLocal()

}

func processCloud() {
	layoutDataResource := resource.NewLayoutDataResource("./resources/dlex-layout/SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")

	layoutData := endpoint.NewDlexEndpoint("samples/dlex-layout/SimpleReportWithCoverPage.dlex", layoutDataResource)
	layoutData.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	layoutData.Endpoint.ApiKey = "DP--api-key--"
	resp := layoutData.Process()

	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("./output/dlex-layout-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}

func processLocal() {
	layoutDataResource := resource.NewLayoutDataResource("./resources/dlex-layout/SimpleReportWithCoverPage.json", "SimpleReportWithCoverPage.json")

	theDlexResource := resource.NewDlexResourceWithPath("./resources/dlex-layout/SimpleReportWithCoverPage.dlex", "SimpleReportWithCoverPage.dlex")


	layoutData := endpoint.NewDlexEndpointWithResource(*theDlexResource, layoutDataResource)

	additionalResource := endpoint.NewDlexWithAdditionalResource("./resources/dlex-layout/NorthwindLogo.gif", "NorthwindLogo.gif")

	layoutData.Resources = append(layoutData.Resources, additionalResource);

	layoutData.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
	layoutData.Endpoint.ApiKey = "DP--api-key--"
	resp := layoutData.Process()

	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile("./output/dlex-layout-local-output.pdf", res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Println("errorId: " + res.ErrorId().String())
		fmt.Println("errorMsg: " + res.ErrorMessage())
		fmt.Println("Failed with error json: " + res.ErrorJson())
	}
}
