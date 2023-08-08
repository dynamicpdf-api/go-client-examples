package main

import (
	"fmt"
	"os"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func dlexResourceExample(key string, baseUrl string) *endpoint.Pdf {

	pdfD := endpoint.NewPdf()
	pdfD.AddDlexWithCloudResourceNLayoutDataPath("samples/getting-started/getting-started.dlex", "C:/temp/dynamicpdf-api-samples/getting-started.json")
	pdfD.Endpoint.BaseUrl = baseUrl
	pdfD.Endpoint.ApiKey = key

	return pdfD
}

func pdfFromResourceExample(key string, baseUrl string) *endpoint.Pdf {
	pr := endpoint.NewPdf()
	mergeOption := input.NewMergeOptions()
	prInput := input.NewPdfWithCloudPath("example/DocumentB.pdf", mergeOption)
	pr.Inputs = append(pr.Inputs, prInput)
	pr.Endpoint.BaseUrl = baseUrl
	pr.Endpoint.ApiKey = key

	pdfResource := resource.NewPdfResourceWithResourcePath("c:/temp/dynamicpdf-api-samples/DocumentA.pdf", "DocumentA.pdf")
	prInput2 := input.NewPdfWithResource(pdfResource)
	pr.Inputs = append(pr.Inputs, prInput2)

	return pr
}

func securityExample(key string, baseUrl string) *endpoint.Pdf {
	secPdf := endpoint.NewPdf()
	secRes := resource.NewPdfResourceWithResourcePath("c:/temp/dynamicpdf-api-samples/DocumentA.pdf", "DocumentA.pdf")
	secInput := input.NewPdfWithResource(secRes)
	secPdf.Endpoint.BaseUrl = baseUrl
	secPdf.Endpoint.ApiKey = key

	security := endpoint.NewAes128Security("user", "owner")
	secPdf.Security = endpoint.Security(security.Security)
	secPdf.Inputs = append(secPdf.Inputs, secInput)
	return secPdf
}

func barcodeExample(key string, baseUrl string) *endpoint.Pdf {
	barcodePdf := endpoint.NewPdf()
	barcodePdf.Endpoint.BaseUrl = baseUrl
	barcodePdf.Endpoint.ApiKey = key

	barcode := element.NewAztecBarcodeElement("Hello World", element.BottomRight, 50.0, 50.0)
	barcode.SetSymbolSize(element.R105xC105)
	barcode.SetXDimension(2)
	barcode.XOffset = 0
	barcode.YOffset = 10
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcode)
	barcodePdf.Inputs = append(barcodePdf.Inputs, pdfInput)

	return barcodePdf
}

func topLevelMetadata(key string, baseUrl string) *endpoint.Pdf {
	tlm := endpoint.NewPdf()
	tlm.Endpoint.BaseUrl = baseUrl
	tlm.Endpoint.ApiKey = key

	pageInput := input.NewPage()
	tlm.Author = "John Doe"
	tlm.Creator = "John Creator"
	tlm.Keywords = "dynamicpdf api example pdf java instructions"
	tlm.Title = "Sample PDF"
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008

	tlm.Inputs = append(tlm.Inputs, pageInput)

	return tlm
}

func  googleFontsExample(key string, b baseUrl string) *endpoint.PDF {

	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.BaseUrl = baseUrl
	pdfCl.Endpoint.ApiKey = key

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

	return pdfC1
}

func fontsExample(key string, baseUrl string) *endpoint.Pdf {

	fe := endpoint.NewPdf()
	fe.Endpoint.BaseUrl = baseUrl
	fe.Endpoint.ApiKey = key

	pageInput := input.NewPage()
	pageNumber := element.NewPageNumberingElement("A", "TopLeft", 0, 0)
	pageNumber.SetFontSize(42)
	pageNumber.SetFont(font.HelveticaBoldOblique())

	pageColor := color.NewRgbColorDefault().Red()
	pageNumber.SetColor(pageColor.Color)

	pageNumberingElementTwo := element.NewPageNumberingElement("B", "TopRight", 0, 0)
	pageNumberingElementTwo.SetColor(color.NewRgbColorDefault().DarkGreen().Color)
	myFont := font.NewFontResource("c:/temp/fonts-example/cnr.otf", "cnr.otf")
	pageNumberingElementTwo.SetFont(*myFont)
	pageNumberingElementTwo.SetFontSize(32)

	pageInput.Elements = append(pageInput.Elements, pageNumber)
	pageInput.Elements = append(pageInput.Elements, pageNumberingElementTwo)
	fe.Inputs = append(fe.Inputs, pageInput)
	return fe
}

func imageExample(key string, baseUrl string) *endpoint.Pdf {
	prImage := endpoint.NewPdf()
	imageResource := resource.NewImageResourceWithResourcePath("C:/temp/dynamicpdf-api-samples/image-info/getting-started.png", "getting-started.png")
	prImage.AddImage(imageResource)
	prImage.AddImageCloudPath("samples/image-info/getting-started.png")
	prImage.Endpoint.ApiKey = key
	prImage.Endpoint.BaseUrl = baseUrl
	return prImage
}

func acroFormExample(key string, baseUrl string) *endpoint.Pdf {
	pdfAcro := endpoint.NewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath("C:/temp/dynamicpdf-api-samples/simple-form-fill.pdf", "simple-form.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	pdfAcro.Inputs = append(pdfAcro.Inputs, pdfInput)

	field1 := endpoint.NewFormFieldWithValue("nameField", "DynamicPdf")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field1)
	field2 := endpoint.NewFormFieldWithValue("descriptionField", "RealRTime Pdf's. Real FAST!")
	pdfAcro.FormFields = append(pdfAcro.FormFields, *field2)

	pdfAcro.Endpoint.ApiKey = key
	pdfAcro.Endpoint.BaseUrl = baseUrl

	return pdfAcro
}

/*
func outlineExample(key string, baseUrl string) *endpoint.Pdf {

	pdfOut := endpoint.NewPdf()

	pageInputOne := pdfOut.AddPage()
	textOne := element.NewText("Hello World 1", element.BottomCenter, 250, 0)
	pageInputOne.Elements = append(pageInputOne.Elements, textOne)

	pdfOut.Inputs = append(pdfOut.Inputs, pageInputOne)

	pageInputTwo := pdfOut.AddPage()
	textTwo := element.NewText("Hello World 2", element.BottomCenter, 250, 0)
	pageInputTwo.Elements = append(pageInputTwo.Elements, textTwo)

	pageInputThree := pdfOut.AddPage()
	textThree := element.NewText("Hello World 3", element.BottomCenter, 250, 0)
	pageInputThree.Elements = append(pageInputThree.Elements, textThree)

	outline := pdfOut.Outlines.Add("Root Outline")
	outline.Expanded = true
	out := outline.Children()
	out.AddWithInputValue("Page 1", pageInputOne, 1, endpoint.FitPage)
	out.AddWithInputValue("Page 2", pageInputTwo, 2, endpoint.FitPage)
	out.AddWithInputValue("Page 3", pageInputOne, 3, endpoint.FitPage)

	pdfOut.Endpoint.ApiKey = key
	pdfOut.Endpoint.BaseUrl = baseUrl

	return pdfOut
}
*/

func templateExample(key string, baseUrl string) *endpoint.Pdf {
	pdfTemp := endpoint.NewPdf()

	pdfTemp.Endpoint.BaseUrl = baseUrl
	pdfTemp.Endpoint.ApiKey = key

	pdfResource := resource.NewPdfResourceWithResourcePath("c:/temp/dynamicpdf-api-samples/DocumentA.pdf", "DocumentA.pdf")
	pdfInput1 := input.NewPdfWithResource(pdfResource)

	temp := element.NewTemplate()
	txt := element.NewText("Hello World", element.TopCenter, 20, 10)
	txt.SetColor(color.NewRgbColorDefault().Red().Color)
	txt.SetFontSize(32)

	temp.Elements = append(temp.Elements, txt)
	pdfInput1.SetTemplate(temp)

	pdfTemp.Inputs = append(pdfTemp.Inputs, pdfInput1)

	return pdfTemp

}

func main() {

	//pdfTemp := templateExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfTemp, "c:/temp/dynamicpdf-api-samples/pdfTempExample.pdf")

	//	pdfOut := outlineExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//	process(pdfOut,"c:/temp/dynamicpdf-api-samples/pdfOutlineExample.pdf")

	//	pdfAcro := acroFormExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//	process(pdfAcro,"c:/temp/dynamicpdf-api-samples/pdfAcroExample.pdf")

	//pdfD := dlexResourceExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfD,"c:/temp/dynamicpdf-api-samples/pdfdlexExample.pdf")

	pdfIm := imageExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	process(pdfIm, "c:/temp/dynamicpdf-api-samples/out/pdfImage.pdf")

	///pdfTlm := topLevelMetadata("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfTlm, "c:/temp/dynamicpdf-api-samples/tlm.pdf")

	//pdfFnt := fontsExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfFnt, "c:/temp/dynamicpdf-api-samples/fnt.pdf")

	//pdfBar := barcodeExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfBar, "c:/temp/dynamicpdf-api-samples/bar.pdf")

	//	pdfSec := securityExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfSec, "c:/temp/dynamicpdf-api-samples/sec.pdf")

	//pdfPr := pdfFromResourceExample("DP---API-KEY---", "https://api.dynamicpdf.com/")
	//process(pdfPr, "c:/temp/dynamicpdf-api-samples/pdfPr.pdf")
}

func process(thePdf *endpoint.Pdf, outputFilePath string) {

	resp := thePdf.Process()
	res := <-resp

	if res.IsSuccessful() == true {
		os.WriteFile(outputFilePath,
			res.Content().Bytes(), os.ModeType)
	} else {
		fmt.Print(res.ErrorJson())
	}

}
