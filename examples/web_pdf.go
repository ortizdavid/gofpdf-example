package examples

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ortizdavid/gofpdf-example/generator"
)


func pdfHandler(w http.ResponseWriter, r *http.Request) {
	pdfGen := generator.NewPdfGenerator()
	fileName := "files/web1.pdf"
	productList := getProductList()

	pdfGen.Header()
	pdfGen.Title("Product List")
	pdfGen.BodyCell(fmt.Sprintf("Num. of records: %d", len(productList)))
	pdfGen.Fpdf.Ln(10)
	// Columns --------------------
	pdfGen.TableColumns("Id", "Name", "Price")
	// Rows ---------------------
	for _, product := range productList {
		pdfGen.TableRow(product.id)
		pdfGen.TableRow(product.name)
		pdfGen.TableRow(product.price)
		pdfGen.Fpdf.Ln(5)
	}
	pdfGen.Footer()
	pdfGen.Output(fileName)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename="+fileName)
	pdfBytes, _ := os.ReadFile(fileName)
	w.Write(pdfBytes)
}

func GenerateWebPdf() {
	http.HandleFunc("/pdf", pdfHandler)
	http.ListenAndServe(":8080", nil)
}

type product struct {
	id int
	name string
	price float32
}

func getProductList() []product {
	return []product{
		{id: 1, name: "Book", price: 500.99},
		{id: 2, name: "Pen", price: 10.99},
		{id: 3, name: "Dictionary", price: 200.99},
		{id: 4, name: "Pencil", price: 8.99},
		{id: 5, name: "Phone", price: 1000.99},
		{id: 6, name: "Computer", price: 3100.99},
		{id: 7, name: "Printer", price: 4100.99},
	}
}

