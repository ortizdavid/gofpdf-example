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
	pdfGen.TableRow("Row1")
	pdfGen.TableRow("Row1")
	pdfGen.TableRow("Row1")
	pdfGen.Fpdf.Ln(5)
	pdfGen.TableRow("Row2")
	pdfGen.TableRow("Row2")
	pdfGen.TableRow("Row2")
	pdfGen.Fpdf.Ln(5)
	// Footer ----------------------
	pdfGen.Footer()
	pdfGen.Output("files/table.pdf")
}

func GenerateSimplePdf()  {
	pdfGen := generator.NewPdfGenerator()
	pdfGen.Header()
	pdfGen.Title("Simple PDF")
	pdfGen.BodyCell("Simple Text")
	pdfGen.Output("files/simple.pdf")
}
