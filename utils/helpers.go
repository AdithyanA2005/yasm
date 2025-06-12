package utils

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

// renderTable renders a table with the specified headers and rows to the provided writer.
// It uses the tablewriter package to format the output.
func RenderTable(w io.Writer, headers []string, rows [][]string) {
	table := tablewriter.NewWriter(w)
	table.Header(headers)
	for _, row := range rows {
		table.Append(row)
	}
	table.Render()
}
