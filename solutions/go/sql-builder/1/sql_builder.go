package sqlbuilder

import (
	"fmt"
	"strings"
)

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	builder := strings.Builder{}

	builder.WriteString("SELECT ")
	if len(columns) > 0 {
		builder.WriteString(strings.Join(columns, ", "))
	} else {
		builder.WriteString("*")
	}

	builder.WriteString(" FROM " + table + ";")

	return builder.String()
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	var output []string
	for _, row := range rows {
		output = append(output, buildInsert(table, row))
	}
	return output
}

func buildInsert(table string, row []ColumnValue) string {
	columns := make([]string, len(row))
	values := make([]string, len(row))
	for i, cv := range row {
		columns[i] = cv.Column
		values[i] = "'" + cv.Value + "'"
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		table, strings.Join(columns, ", "), strings.Join(values, ", "))
}
