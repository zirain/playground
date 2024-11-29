package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("sample.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("sample.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	pageNumber := props.PageNumber{
		Pattern: "Page {current} of {total}",
		Place:   props.RightBottom,
		Family:  fontfamily.Courier,
		Style:   fontstyle.Bold,
		Size:    9,
		Color: &props.Color{
			Red: 255,
		},
	}

	cfg := config.NewBuilder().
		WithDebug(true).
		WithPageNumber(pageNumber).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	// footer
	footerRow := row.New().Add(
		text.NewCol(4, "left"),
		text.NewCol(4, "center"),
		text.NewCol(4, "right"),
	).WithStyle(&props.Cell{
		BorderColor: &props.BlackColor,
		BorderType:  border.Top,
	})
	_ = m.RegisterFooter(footerRow)

	for i := 0; i < 15; i++ {
		m.AddRows(text.NewRow(20, "dummy text"))
	}

	return m
}
