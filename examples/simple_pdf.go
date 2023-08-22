package examples

import (
	"github.com/ortizdavid/gofpdf-example/generator"
)

func GenerateReport() {

	var pdfGen generator.PdfGenerator
	pdf := pdfGen.NewFPdf()

	pdfGen.Header(pdf)
	pdfGen.Title(pdf, "Report Title XXXX")

	// Columns
	pdfGen.TableColumns(pdf, "Column 1", "Column 2", "Column 3")

	// Rows------------------
	pdfGen.TableRow(pdf, "Val1")
	pdfGen.TableRow(pdf, "Val1")
	pdfGen.TableRow(pdf, "Val1")
	pdf.Ln(5)
	pdfGen.TableRow(pdf, "Val2")
	pdfGen.TableRow(pdf, "Val2")
	pdfGen.TableRow(pdf, "Val2")
	pdf.Ln(5)
	//----------------------

	pdfGen.Footer(pdf)
	pdfGen.Output(pdf, "files/report1.pdf")
}

func GenerateSimplePdf()  {
	var pdfGen generator.PdfGenerator
	pdf := pdfGen.NewFPdf()
	pdfGen.Header(pdf)
	pdfGen.Title(pdf, "Simple PDF")
	pdfGen.BodyCell(pdf, "Simple Text")
	pdfGen.Output(pdf, "files/simple1.pdf")
}