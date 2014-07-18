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
		// Level 4
		[]string{
			"=JXJ-   X^",
			"X   -J  XJ",
			"= X^  XJ-J",
			"-^=#=^-^-P",
			"    X   -J",
		},
		// Level 5
		[]string{
			"= - -P= =#",
			"-J=^XJX^=#",
			"XJX^X^- =#",
			"=#=#XJ  =J",
			"= -^=J  =J",
		},
		// Level 6
		[]string{
			"= =   XJ= ",
			"-J- XJ-^-^",
			"XJX =^XJ=J",
			"-^-^X^X = ",
			"=^=PX^=#X^",
		},
		// Level 7
		[]string{
			"XJ=   =J- ",
			"= -   -J-J",
			"=^-J  X =^",
			"=P-^XJXJ=^",
			"= - -^= - ",
		},
		// Level 8
		[]string{
			"=^X -J-   ",
			"-J- =#-JXJ",
			"-^=J-J=#- ",
			"  =^-J-^X ",
			"  = = -PX^",
		},
		// Level 9
		[]string{
			"X = -^XJX^",
			"-^X^X =J  ",
			"- =#X   XJ",
			"- =P= = - ",
			"=^X^-^-^-J",
		},
		// Level 10
		[]string{
			"=#  =J-JXP",
			"  X -^- -^",
			"XJ=#-   - ",
			"- X^=J  = ",
			"=JX^  - = ",
		},
		// Level 11
		[]string{
			"= =   -J- ",
			"X^=#=#  X ",
			"-^-^-J-J= ",
			"=#-^=^X - ",
			"=J-PXJXJ-J",
		},
		// Level 12
		[]string{
			"=^  =   X^",
			"= = =^-^XJ",
			"XJ- X -   ",
			"-PX -^=#-J",
			"=^-J=JX^= ",
		},
		// Level 13
		[]string{
			"X - = = X ",
			"X =J-^XJ= ",
			"-^=^X^-J=^",
			"- XJ=#-^=J",
			"-^-PXJ-J=#",
		},
		// Level 14
		[]string{
			"=JX^=#-^X^",
			"XJ=#-^= - ",
			"- -JXJ= X^",
			"  -PX^X   ",
			"=JX =^-J-^",
		},
		// Level 15
		[]string{
			"=JXJ-^X   ",
			"-^=J=^-JX^",
			"-P-J=#X^=#",
			"-^= =^X - ",
			"= XJ=#-J=^",
		},
		// Level 16
		[]string{
			"-^= =J  - ",
			"-^X - XJ-J",
			"=J=^X^=#= ",
			"=^=P-^=JX^",
			"-J  -^XJ= ",
		},
		// Level 17
		[]string{
			"X =J= =^= ",
			"-J-^=^X^X^",
			"-P=J=#- -J",
			"X X^- - -^",
			"  X -J  X ",
		},
		// Level 18
		[]string{
			"- X = XJ-^",
			"X -^=#-^=#",
			"X XJ-J-   ",
			"=J  =#XJ-J",
			"  -^=P-^=J",
		},
		// Level 19
		[]string{
			"XJX^-J=#=^",
			"- = =^=J=^",
			"-^  =#X XP",
			"-^=^=J  = ",
			"- -^=J-JXJ",
		},
		// Level 20
		[]string{
			"-J-J= =J  ",
			"- -J=#- X ",
			"-^= X -^=P",
			"= =^=#-J=J",
			"  -J- -^  ",
		},
		// Level 21
		[]string{
			"- -J=^  =#",
			"  XJ-^XJ-^",
			"X^  -^X^  ",
			"- = = X -^",
			"XJ-^X^=#XP",
		},
		// Level 22
		[]string{
			"-J  =JX =P",
			"=#XJ-JX =^",
			"  -^X =J=#",
			"= X^=^-^-^",
			"=#X^-J  =#",
		},
		// Level 23
		[]string{
			"-J=JXJ  - ",
			"X^X^- - = ",
			"-^X =#X^= ",
			"-JX^  -J-J",
			"- -^-JX =P",
		},
		// Level 24
		[]string{
			"  -JX - X^",
			"XJ=^X^XJ-P",
			"=^XJ- - =#",
			"- =JX^-J=#",
			"=J-^X - =#",
		},
		// Level 25
		[]string{
			"=^-J=#XP- ",
			"=^- -J- =#",
			"=^- - =J=J",
			"-J- - = =J",
			"= -^- =J-^",
		},
		// Level 26
		[]string{
			"  =^=J-PX^",
			"- - -^- - ",
			"X XJ- -   ",
			"-J  -^=#= ",
			"  - X -J=#",
		},
		// Level 27
		[]string{
			"X^XJXJ=J  ",
			"XP- =#-^=J",
			"=^-^- X -^",
			"- - X = XJ",
			"XJX^  XJ= ",
		},
		// Level 28
		[]string{
			"-P-^  XJ=^",
			"-^=#  =J= ",
			"-J-^=JXJ-J",
			"X -^- X^X ",
			"= - = =#-J",
		},
		// Level 29
		[]string{
			"-P=^X   -^",
			"X^X   - X ",
			"=^= -J-^= ",
			"-J-JX =#XJ",
			"=^X^X^- XJ",
		},
		// Level 30
		[]string{
			"X^- =J=#X ",
			"XJ=J=#X X^",
			"X^-J=#-^=J",
			"  -   - =P",
			"=J=#=#= - ",
		},
		// Level 31
		[]string{
			"- X^X -JX^",
			"-^XJ=#=J  ",
			"X - X^XJ-P",
			"-J  =#X^XJ",
			"  -   =^X ",
		},
		// Level 32
		[]string{
			"=J- - X   ",
			"= =^  -P=^",
			"=J-^=J  - ",
			"=^=J=#- = ",
			"-   =#=JXJ",
		},
		// Level 33
		[]string{
			"= -J-J= XJ",
			"-J=^=#X^- ",
			"- -J-^=J-P",
			"X =^=^X^XJ",
			"  -^-   -^",
		},
		// Level 34
		[]string{
			"=J=#-JXJ-P",
			"XJ-^=^-J=#",
			"  - X - X^",
			"X^-^-^-^-J",
			"= = = =   ",
		},
		// Level 35
		[]string{
			"= X -^=J= ",
			"-J-J- =^XP",
			"=#-^XJ-^=^",
			"X X^    =#",
			"X^X^-^X =J",
		},
		// Level 36
		[]string{
			"-PXJ- -^-^",
			"=J=^-J-J=^",
			"- X^-JX^X ",
			"-JX^-^= =#",
			"- X^=J  X ",
		},
		// Level 37
		[]string{
			"-JX =#  X^",
			"X^=J-J  -^",
			"  -^-J= = ",
			"= =J  -P=^",
			"=J  X =J= ",
		},
		// Level 38
		[]string{
			"X^XJ=^  XJ",
			"-^X =#-^XP",
			"=J- X^- -J",
			"X^=JXJ- =#",
			"XJ= =#X =#",
		},
		// Level 39
		[]string{
			"-JX   =^= ",
			"X^-J=#-^=J",
			"-J-^=^XJ- ",
			"- X XJ-PX ",
			"=#-^=#-J=#",
		},
		// Level 40
		[]string{
			"X -J= =P  ",
			"X   =#X -J",
			"- -J=^X -^",
			"XJ-   =#  ",
			"X -J-^- =J",
		},
		// Level 41
		[]string{
			"=JXJ- - X^",
			"-JX   X X ",
			"XJ=^  - =^",
			"X^=^= -J-P",
			"-   XJ=^=J",
		},
		// Level 42
		[]string{
			"=JX^-J-^=P",
			"X XJ-^  XJ",
			"X = X -^-^",
			"-J-^X     ",
			"X -^- =^X^",
		},
		// Level 43
		[]string{
			"=J-J- -^- ",
			"XJ=J-^  =J",
			"X^X = -^=^",
			"X^=^=^    ",
			"-J- = -P= ",
		},
		// Level 44
		[]string{
			"XJ  X^-^=J",
			"  XJXJXJ= ",
			"X^=^X -^-^",
			"-^X -J=J= ",
			"=P=^= = =^",
		},
		// Level 45
		[]string{
			"-^=^  X^-P",
			"XJX =   -J",
			"=^-^=#X^-^",
			"  =     XJ",
			"  XJXJX^=^",
		},
		// Level 46
		[]string{
			"-J  -^-JX ",
			"=^  -J=^-^",
			"XJ-^X^  XJ",
			"-^= -^X   ",
			"-P- = = X ",
		},
		// Level 47
		[]string{
			"X X =P=^= ",
			"XJ=^=#- =J",
			"=^-J=J=^XJ",
			"- =^=#X^-J",
			"-^=^=^=#XJ",
		},
		// Level 48
		[]string{
			"= = X -J=J",
			"=P-^=#X X^",
			"=J  -^X =J",
			"  =^-^=   ",
			"-^XJ= X^- ",
		},
		// Level 49
		[]string{
			"-J=^  X X ",
			"-   -PXJ-^",
			"  X -J-^X ",
			"-^-^= X^X ",
			"X^X^=J=J- ",
		},
		// Level 50
		[]string{
			"-^X =^  -J",
			"=^X - =#-^",
			"-J-^=J= - ",
			"X -^-^= -^",
			"=#-P=JXJ- ",
		},
		// Level 51
		[]string{
			"-P-^-JXJ-^",
			"XJ= =   =#",
			"-^=^-^X   ",
			"-J- -J  =J",
			"= - = XJ-^",
		},
		// Level 52
		[]string{
			"=#XJX^=JX ",
			"= =P-J=^- ",
			"= - =#=#X ",
			"- X^=J- -^",
			"-J= =#-JX ",
		},
		// Level 53
		[]string{
			"XJ- -   - ",
			"- - =^-J=#",
			"=^XJ-P-^-J",
			"-J-^-^X^X ",
			"X = - XJ=J",
		},
		// Level 54
		[]string{
			"XJ=JX^-J  ",
			"  -JX -^= ",
			"-J=^  - X^",
			"= =^= X - ",
			"X^- -PX^X ",
		},
		// Level 55
		[]string{
			"X - =J- - ",
			"- =^-P  -J",
			"X^-^X   =^",
			"X^X = X =^",
			"=^XJ  X XJ",
		},
		// Level 56
		[]string{
			"= =^= XJ=#",
			"-     X XP",
			"- =#=J= -^",
			"X XJ- -J=#",
			"=J-^X^XJ- ",
		},
		// Level 57
		[]string{
			"X XJ= XJ- ",
			"-^=#  X =#",
			"-J-J=#  X^",
			"X =^-P=^- ",
			"=J=J- - XJ",
		},
		// Level 58
		[]string{
			"X X =^=#-^",
			"-J-J-^-^= ",
			"XJ-^X^=J=#",
			"=J=J=^- =^",
			"XP- X - -J",
		},
		// Level 59
		[]string{
			"  X =J=^=J",
			"X^=#= =^  ",
			"-^XJ-J=#=^",
			"X^- -^- XP",
			"  X^=^X^-J",
		},
		// Level 60
		[]string{
			"=J= =^=#=J",
			"XJX = -^X ",
			"X^=J=J-^- ",
			"- X^=#-^=J",
			"- -J-P= = ",
		},
		// Level 61
		[]string{
			"=J  - - X ",
			"-J=#-J=^-J",
			"= =#= =^= ",
			"  XJX X^XP",
			"XJ=J=^=^-^",
		},
		// Level 62
		[]string{
			"=^X - =^-J",
			"-^-^-JXJ-J",
			"X X -^X^-P",
			"X = XJ-^X ",
			"X XJ=#X X^",
		},
		// Level 63
		[]string{
			"XJ  -JX -J",
			"  = = =^=^",
			"  XJ=#-^X ",
			"=P-^- X X^",
			"X^X XJ-JXJ",
		},
		// Level 64
		[]string{
			"  XJXJXJ  ",
			"XJ  XJ  XJ",
			"XJ  XJ  XJ",
			"XJ  XJXJXJ",
			"=PXJXJXJ  ",
		},
	}
)

func init() {
	for i, rawLevel := range rawLevels {
		Levels = append(Levels, MustParseLevel(i+1, rawLevel))
	}
}
