package vulnerabilities

import (
	"fmt"
	"strings"
)

var colors = map[string]string{
	"Critical": "\033[31m",
	"High":     "\033[35m",
	"Medium":   "\033[33m",
	"Low":      "\033[34m",
	"reset":    "\033[0m",
}

type Vulnerability struct {
	CVEID        string  `json:"cve_id"`
	CVSSv3Score  float64 `json:"cvss_v3_score"`
	Severity     string  `json:"severity"`
	Package      string  `json:"package"`
	Version      string  `json:"version"`
	FixedVersion string  `json:"fixed_version"`
	Description  string  `json:"desc"`
}

func colorize(text string, color string) string {
	return fmt.Sprintf("%s%s%s", colors[color], text, colors["reset"])
}

func wrapText(text string, width int) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}
	wrapped := words[0]
	spaceLeft := width - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n    " + word
			spaceLeft = width - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}

func DisplayVulnerabilities(vulnerabilities []Vulnerability) {
	for _, vuln := range vulnerabilities {
		fmt.Printf("CVE ID: %s\n", colorize(vuln.CVEID, vuln.Severity))
		fmt.Printf("Severity: %s\n", colorize(vuln.Severity, vuln.Severity))
		fmt.Printf("CVSS Score: %.1f\n", vuln.CVSSv3Score)
		fmt.Printf("Package: %s\n", vuln.Package)
		fmt.Printf("Current Version: %s\n", vuln.Version)
		fmt.Printf("Fixed Version: %s\n", vuln.FixedVersion)
		fmt.Printf("Description: %s\n", wrapText(vuln.Description, 80))
		fmt.Println(strings.Repeat("-", 80))
	}
}