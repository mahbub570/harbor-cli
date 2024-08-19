package summary

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goharbor/harbor-cli/pkg/views/base/tablelist"
)

var colors = map[string]string{
	"red":    "\033[31m",
	"purple": "\033[35m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"green":  "\033[32m",
	"reset":  "\033[0m",
}

type SecuritySummary struct {
	CriticalCnt   int64 `json:"critical_cnt"`
	FixableCnt    int64 `json:"fixable_cnt"`
	HighCnt       int64 `json:"high_cnt"`
	LowCnt        int64 `json:"low_cnt"`
	MediumCnt     int64 `json:"medium_cnt"`
	ScannedCnt    int64 `json:"scanned_cnt"`
	TotalArtifact int64 `json:"total_artifact"`
	TotalVuls     int64 `json:"total_vuls"`
}

func colorize(count int64, color string) string {
	return fmt.Sprintf("%s%d%s", colors[color], count, colors["reset"])
}

func DisplaySecuritySummary(summaryData *SecuritySummary) {
	columns := []table.Column{
		{Title: "Type", Width: 30},
		{Title: "Quantity", Width: 15},
	}

	rows := []table.Row{
		{"Critical Vulnerabilities", colorize(summaryData.CriticalCnt, "red")},
		{"High Vulnerabilities", colorize(summaryData.HighCnt, "purple")},
		{"Medium Vulnerabilities", colorize(summaryData.MediumCnt, "yellow")},
		{"Low Vulnerabilities", colorize(summaryData.LowCnt, "blue")},
		{"Fixable Vulnerabilities", colorize(summaryData.FixableCnt, "green")},
		{"Scanned Artifacts", fmt.Sprintf("%d", summaryData.ScannedCnt)},
		{"Total Artifacts", fmt.Sprintf("%d", summaryData.TotalArtifact)},
		{"Total Vulnerabilities", fmt.Sprintf("%d", summaryData.TotalVuls)},
	}

	m := tablelist.NewModel(columns, rows, len(rows))

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}