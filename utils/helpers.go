package utils

import (
	"io"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// "github.com/olekukonko/tablewriter"
// renderTable renders a table with the specified headers and rows to the provided writer.
// It uses the tablewriter package to format the output.
func RenderTable(w io.Writer, headers []string, rows [][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true

	// Convert []string headers to table.Row (type alias for []interface{})
	headerRow := make(table.Row, len(headers))
	for i, h := range headers {
		headerRow[i] = h
	}
	t.AppendHeader(headerRow)

	// Convert [][]string to []table.Row
	for _, row := range rows {
		rowData := make(table.Row, len(row))
		for i, cell := range row {
			rowData[i] = cell
		}
		t.AppendRow(rowData)
	}

	t.Render()
}
