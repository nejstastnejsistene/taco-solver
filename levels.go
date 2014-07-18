package main

var (
	Levels []Level

	rawLevels = [][]string{
		// Level 1
		[]string{
			"=#=#=#",
			"XPX XJ",
			"= =^= ",
		},
		// Level 2
		[]string{
			"-P- = = ",
			"      = ",
			"-J- -J- ",
		},
		// Level 3
		[]string{
			"-P-^  =J",
			"-^    -^",
			"-#      ",
			"XJX X -J",
		},
	}
)

func init() {
	for _, rawLevel := range rawLevels {
		Levels = append(Levels, MustParseLevel(rawLevel))
	}
}
