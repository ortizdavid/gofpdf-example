package examples

import (
	"github.com/ortizdavid/gofpdf-example/generator"
)

func GenerateTablePdf() {

	pdfGen := generator.NewPdfGenerator()

	// Header---------------------
	pdfGen.Header()
	pdfGen.Title("PDF with a Table")
	// Columns --------------------
	pdfGen.TableColumns("Column 1", "Column 2", "Column 3")
	// Rows -----------------------
	pdfGen.TableRow("Val1")
	pdfGen.TableRow("Val1")
	pdfGen.TableRow("Val1")
	pdfGen.Fpdf.Ln(5)
	pdfGen.TableRow("Val2")
	pdfGen.TableRow("Val2")
	pdfGen.TableRow("Val2")
	pdfGen.Fpdf.Ln(5)
	// Footer ----------------------
	pdfGen.Footer()
	pdfGen.Output("files/table1.pdf")
}

func GenerateSimplePdf()  {
	pdfGen := generator.NewPdfGenerator()
	pdfGen.Header()
	pdfGen.Title("Simple PDF")
	pdfGen.BodyCell("Simple Text")
	pdfGen.Output("files/simple1.pdf")
}
