package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Convert list to [][]string
func listToString(list todoList) [][]string {
	var result [][]string
	for _, v := range list {
		result = append(result, []string{
			strconv.Itoa(v.ID),
			v.TaskTitle,
			v.CreatedAt.Format("2006-01-02 15:04:05"),
			strconv.FormatBool(v.IsDone),
			strconv.FormatBool(v.IsDeleted),
		})
	}
	return result
}

// Print tasks
func printTask(isAll bool, list todoList) {
	data := listToString(list)
	table := tablewriter.NewWriter(os.Stdout)
	if isAll {
		table.SetHeader([]string{"ID", "Task", "Create Time", "isFinished"})
		for _, v := range data {
			if v[4] == "false" {
				table.Append(v[:4])
			}
		}
	} else {
		table.SetHeader([]string{"ID", "Task", "Create Time"})
		for _, v := range data {
			if v[3] == "false" && v[4] == "false" {
				table.Append(v[:3])
			}
		}
	}
	table.Render()
}
