package units

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed guance.json
var guanceUnitsBytes []byte

var supportedUnits map[string][]string

func init() {
	supportedUnits = make(map[string][]string)
	guanceUnits, err := parseGuanceUnits()
	if err != nil {
		log.Fatal("parse guance units: %w", err)
	}
	for _, category := range guanceUnits {
		for _, unit := range category.Formats {
			if unit.Ref == "" {
				continue
			}
			supportedUnits[unit.Ref] = []string{category.Name, unit.ID}
		}
	}
}

type Category struct {
	Name    string `json:"name"`
	Formats []Unit `json:"formats"`
}

type Unit struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Ref  string `json:"ref"`
}

func parseGuanceUnits() ([]Category, error) {
	var guanceUnits []Category
	if err := json.Unmarshal(guanceUnitsBytes, &guanceUnits); err != nil {
		return nil, fmt.Errorf("unmarshal guance units: %w", err)
	}
	return guanceUnits, nil
}

// convertUnit converts unit from Grafana Dashboard to Guance Cloud
func convertUnit(id string) []string {
	if units, ok := supportedUnits[id]; ok {
		return units
	}
	return nil
}
