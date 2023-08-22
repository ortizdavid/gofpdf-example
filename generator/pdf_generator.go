package generator

import (
	"fmt"
	"log"

	"github.com/go-pdf/fpdf"
)

type PdfGenerator struct {
	Font				string

	LogoImage      		string
	LogoX				float64
	LogoY				float64
	LogoWidth			float64
	LogoHeight			float64
	LogoFlow			bool
	LogoLink        	int
	LogoTp				string
	LogoLinkStr			string

	HeaderStyleStr  	string
	HeaderFontSize		float64
	HeaderWidth			float64
	HeaderHeight		float64
	HeaderLn  			float64

	TitleStyleStr		string
	TitleFontSize		float64
	TitleWidth			float64
	TitleHeight			float64
	TitleLn  			float64

	BodyCellWidth  		float64
	BodyCellHeight 		float64
	BodyCellStyleStr 	string
	BodyCellFontSize 	float64
	BodyCellLn 			float64

	TableColStyleStr 	string
	TableColFontSize 	float64
	TableColWith		float64
	TableColHeight		float64
	TableLn  			float64

	TableRowStyleStr  	string
	TableRowFontSize	float64
	TableRowWith		float64
	TableRowHeight		float64

	FooterStyleStr  	string
	FooterFontSize		float64
	FooterY				float64
	FooterWidth			float64
	FooterHeight		float64
}

func (gen *PdfGenerator) NewFPdf() *fpdf.Fpdf {
	return fpdf.New("P", "mm", "A4", "")
}

func (gen *PdfGenerator) GetPdfStyle() PdfGenerator {
	return PdfGenerator {
		Font:             "Arial",
		LogoImage:        "",
		LogoX:            90,
		LogoY:            5,
		LogoWidth:        30,
		LogoHeight:       20,
		LogoFlow:         false,
		LogoLink:         0,
		LogoTp:           "",
		LogoLinkStr:      "",
		HeaderStyleStr:   "B",
		HeaderFontSize:   14,
		HeaderWidth:      40,
		HeaderHeight:     50,
		HeaderLn:         10,
		TitleStyleStr:    "B",
		TitleFontSize:    11,
		TitleWidth:       40,
		TitleHeight:      50,
		TitleLn:          10,
		BodyCellWidth:    40,
		BodyCellHeight:   60,
		BodyCellStyleStr: "",
		BodyCellFontSize: 9,
		BodyCellLn:       1,
		TableColStyleStr: "B",
		TableColFontSize: 10,
		TableColWith:     40,
		TableColHeight:   50,
		TableLn:          1,
		TableRowStyleStr: "",
		TableRowFontSize: 9,
		TableRowWith:     40,
		TableRowHeight:   60,
		FooterStyleStr:   "I",
		FooterFontSize:   8,
		FooterY:          -15,
		FooterWidth:      0,
		FooterHeight:     10,
	}
}


//pdf.Image("static/images/logo.png", 90, 5, 30, 20, false, "", 0, "")
func (gen *PdfGenerator) Header(pdf *fpdf.Fpdf) {
	s := gen.GetPdfStyle()
	pdf.AddPage()
	pdf.SetFont(s.Font, s.HeaderStyleStr, s.HeaderFontSize)
	if s.LogoImage != "" {
		pdf.Image(s.LogoImage, s.LogoX, s.LogoY, s.LogoWidth, s.LogoHeight, 
		s.LogoFlow, s.LogoTp, s.LogoLink, s.LogoLinkStr)
	}
	pdf.Cell(s.HeaderWidth, s.HeaderHeight, "Report")
	pdf.Ln(s.HeaderLn)
}

func (gen *PdfGenerator) Title(pdf *fpdf.Fpdf, title string) {
	s := gen.GetPdfStyle()
	pdf.SetFont(s.Font, s.TitleStyleStr, s.TitleFontSize)
	pdf.Cell(s.TitleWidth, s.TitleHeight, title)
	pdf.Ln(s.TitleLn)
}

func (gen *PdfGenerator) Footer(pdf *fpdf.Fpdf) {
	s := gen.GetPdfStyle()
	pdf.SetFooterFunc(func() {
        pdf.SetY(s.FooterY)
        pdf.SetFont(s.Font, s.FooterStyleStr, s.FooterFontSize) //I, 8
        pdf.Cell(s.FooterWidth, s.FooterHeight, fmt.Sprintf("Pag. %d", pdf.PageNo()))//0,10
    })
}

func (gen *PdfGenerator) BodyCell(pdf *fpdf.Fpdf, row any) {
	s := gen.GetPdfStyle()
	pdf.SetFont(s.Font, s.BodyCellStyleStr, s.BodyCellFontSize)
	pdf.Cell(s.BodyCellWidth, s.BodyCellHeight, fmt.Sprintf("%v", fmt.Sprintf("%v", row)))
	pdf.Ln(s.BodyCellLn)
}

func (gen *PdfGenerator) TableColumns(pdf *fpdf.Fpdf, columns ...string) {
	s := gen.GetPdfStyle()
	pdf.SetFont(s.Font, s.TableColStyleStr, s.TableColFontSize)
	for _, column := range columns {
		pdf.Cell(s.TableColWith, s.TableColHeight, column)
	}
	pdf.Ln(s.TableLn)
}

func (gen *PdfGenerator) TableRow(pdf *fpdf.Fpdf, row any) {
	s := gen.GetPdfStyle()
	//decoded := Utf8Decode(fmt.Sprintf("%v", row))
	pdf.SetFont(s.Font, s.TableRowStyleStr, s.TableRowFontSize)
	pdf.Cell(s.TableRowWith, s.TableRowHeight, fmt.Sprintf("%v", fmt.Sprintf("%v", row)))
}


func (gen *PdfGenerator) Line(pdf *fpdf.Fpdf) {
	//s := gen.GetPdfStyle()
	pdf.Line(300, 35, 35, 35)
}

func (gen *PdfGenerator) Output(pdf *fpdf.Fpdf, fileName string) {
	//s := gen.GetPdfStyle()
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		log.Fatal(err)
	}
}



