package util

var locations = []string{
	"Aurskog",
	"Aurskog-Høland",
	"Aursmoen",
	"Bjørkelangen",
	"Blaker",
	"Fet",
	"Fetsund",
	"Fosser",
	"Gan",
	"Grasåsen",
	"Landet Rundt",
	"Lierfoss",
	"Lillestrøm",
	"Løken",
	"Lørenfallet",
	"Momoen",
	"Mork",
	"Rømskog",
	"Setskog",
	"Sverige",
	"Søndre Mangen",
	"Sørumsand",
	"Ukraina",
}

func IsValidLocation(a string) bool {
	for _, b := range locations {
		if b == a {
			return true
		}
	}
	return false
}

var aHLocations = []string{
	"Aurskog",
	"Aurskog-Høland",
	"Aursmoen",
	"Bjørkelangen",
	"Fosser",
	"Grasåsen",
	"Lierfoss",
	"Løken",
	"Momoen",
	"Rømskog",
	"Setskog",
	"Søndre Mangen",
}

func AmountInAH(locations map[string]int) int {
	count := 0
	for _, location := range aHLocations {
		count += locations[location]
	}
	return count
}

var lsLocations = []string{
	"Blaker",
	"Fetsund",
	"Gan",
	"Lillestrøm",
	"Lørenfallet",
	"Mork",
	"Sørumsand",
}

func AmountInLS(locations map[string]int) int {
	count := 0
	for _, location := range lsLocations {
		count += locations[location]
	}
	return count
}
