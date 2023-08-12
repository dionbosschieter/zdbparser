package zdbparser

import (
	"bufio"
	"strings"
)

func parseDisksFromZdbOutput(scanner *bufio.Scanner) (disks []Disk) {
	sectionBuffer := ""
	sectionType := ""
	previousIndentLevel := 0

	for scanner.Scan() {
		originalLine := scanner.Text()
		line := strings.TrimSpace(originalLine)
		indentLevel := strings.Index(originalLine, strings.TrimLeft(originalLine, " "))

		if previousIndentLevel != indentLevel {
			if sectionType == "disk" {
				disks = append(disks, parseDisk(sectionBuffer))
			}
			sectionType = ""
			sectionBuffer = ""
		}

		parts := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(parts[0])
		value := ""
		if len(parts) == 2 {
			value = strings.TrimSpace(parts[1])
			value = strings.Replace(value, "'", "", -1)
		}

		switch key {
		case "type":
			sectionType = value
		}
		previousIndentLevel = indentLevel
		sectionBuffer += line + "\n"
	}

	return
}
