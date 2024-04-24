package exporter

import (
	"fmt"

	"github.com/earl86/sniffer-agent/model"
)

type cliExporter struct {
}

func NewCliExporter() *cliExporter {
	return &cliExporter{}
}

func (c *cliExporter) Export(qp model.QueryPiece) (err error) {
	fmt.Println(*qp.String())
	return
}
