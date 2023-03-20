package main

import (
	"fmt"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)
      
    func main() {
        
      resource := resource.NewPdfResourceWithResourcePath("C:/temp/dynamicpdf-api-samples/pdf-info/fw4.pdf", "fw4.pdf")
      text := endpoint.NewPdfInfoResource(resource)
      
      text.Endpoint.BaseUrl = "https://api.dynamicpdf.com/"
      text.Endpoint.ApiKey = "DP.DVs/HNreImRwrMAr4T5C8rLvcg0dCPdvpQ5187Fy1+25SbHeCHpudOMR"

      resp := text.Process()
      res := <-resp
	
		  if res.IsSuccessful() == true {
						fmt.Print(string(res.Content().Bytes()))
		  }
    }