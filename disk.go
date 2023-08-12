package zdbparser

import (
	"strconv"
	"strings"
)

type Disk struct {
	Type      string `yaml:"type"`
	ID        int    `yaml:"id"`
	GUID      string `yaml:"guid"`
	Path      string `yaml:"path"`
	Devid     string `yaml:"devid"`
	PhysPath  string `yaml:"phys_path"`
	WholeDisk int    `yaml:"whole_disk"`
	DTL       int    `yaml:"DTL"`
	CreateTXG int    `yaml:"create_txg"`
}

func parseDisk(child string) (disk Disk) {
	lines := strings.Split(child, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Replace(value, "'", "", -1)

		switch key {
		case "type":
			disk.Type = value
		case "id":
			disk.ID, _ = strconv.Atoi(value)
		case "guid":
			disk.GUID = value
		case "path":
			disk.Path = value
		case "devid":
			disk.Devid = value
		case "phys_path":
			disk.PhysPath = value
		case "whole_disk":
			disk.WholeDisk, _ = strconv.Atoi(value)
		case "DTL":
			disk.DTL, _ = strconv.Atoi(value)
		case "create_txg":
			disk.CreateTXG, _ = strconv.Atoi(value)
		}
	}

	return
}
