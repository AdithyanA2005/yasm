package utils

import (
	"io"
	"os"
	"strings"

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

// ExtractAfterSubstring returns the portion of the string s that appears after the first occurrence
// of the substring substr. If substr is not found in str, it returns an empty string and false.
// Otherwise, it returns the substring after substr and true.
func ExtractAfterSubstring(str string, substr string) (string, bool) {
	index := strings.Index(str, substr)
	if index == -1 {
		return "", false
	}

	// Add length of substr to get the portion after it
	return str[index+len(substr):], true
}
