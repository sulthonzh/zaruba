package output

// Decoration is decorator for terminal
type Decoration struct {
	Normal      string
	Bold        string
	Faint       string
	Italic      string
	Underline   string
	BlinkSlow   string
	BlinkRapid  string
	Inverse     string
	Conceal     string
	CrossedOut  string
	Black       string
	Red         string
	Green       string
	Yellow      string
	Blue        string
	Magenta     string
	Cyan        string
	White       string
	BgBlack     string
	BgRed       string
	BgGreen     string
	BgYellow    string
	BgBlue      string
	BgMagenta   string
	BgCyan      string
	BgWhite     string
	NoStyle     string
	NoUnderline string
	NoInverse   string
	NoColor     string
	colorIndex  int
	colorList   []string
	iconIndex   int
	iconList    []string
	Skull       string
	Success     string
	Error       string
	Start       string
	Kill        string
	Inspect     string
	Run         string
	Empty       string
	Icon        func(string) string
}

func NewPlainDecoration() (d *Decoration) {
	return &Decoration{
		Normal:      "",
		Bold:        "",
		Faint:       "",
		Italic:      "",
		Underline:   "",
		BlinkSlow:   "",
		BlinkRapid:  "",
		Inverse:     "",
		Conceal:     "",
		CrossedOut:  "",
		Black:       "",
		Red:         "",
		Green:       "",
		Yellow:      "",
		Blue:        "",
		Magenta:     "",
		Cyan:        "",
		White:       "",
		BgBlack:     "",
		BgRed:       "",
		BgGreen:     "",
		BgYellow:    "",
		BgBlue:      "",
		BgMagenta:   "",
		BgCyan:      "",
		BgWhite:     "",
		NoStyle:     "",
		NoUnderline: "",
		NoInverse:   "",
		NoColor:     "",
		Skull:       "",
		Success:     "",
		Error:       "",
		Start:       "",
		Kill:        "",
		Inspect:     "",
		Run:         "",
		Empty:       "",
		colorIndex:  0,
		iconIndex:   0,
		iconList:    []string{""},
		colorList:   []string{""},
		Icon:        func(icon string) string { return "" },
	}
}

// NewDecoration Create New Decoration
func NewDecoration() (d *Decoration) {
	// source: https://gist.github.com/mxmerz/92e97cd27857a9ba787b
	d = &Decoration{
		Normal:      "\x1b[0m",
		Bold:        "\x1b[1m",
		Faint:       "\x1b[2m",
		Italic:      "\x1b[3m",
		Underline:   "\x1b[4m",
		BlinkSlow:   "\x1b[5m",
		BlinkRapid:  "\x1b[6m",
		Inverse:     "\x1b[7m",
		Conceal:     "\x1b[8m",
		CrossedOut:  "\x1b[9m",
		Black:       "\x1b[30m",
		Red:         "\x1b[31m",
		Green:       "\x1b[32m",
		Yellow:      "\x1b[33m",
		Blue:        "\x1b[34m",
		Magenta:     "\x1b[35m",
		Cyan:        "\x1b[36m",
		White:       "\x1b[37m",
		BgBlack:     "\x1b[40m",
		BgRed:       "\x1b[41m",
		BgGreen:     "\x1b[42m",
		BgYellow:    "\x1b[43m",
		BgBlue:      "\x1b[44m",
		BgMagenta:   "\x1b[45m",
		BgCyan:      "\x1b[46m",
		BgWhite:     "\x1b[47m",
		NoStyle:     "\x1b[0m",
		NoUnderline: "\x1b[24m",
		NoInverse:   "\x1b[27m",
		NoColor:     "\x1b[39m",
		Skull:       "💀",
		Success:     "🎉",
		Error:       "🔥",
		Start:       "🏁",
		Kill:        "🔪",
		Inspect:     "🔎",
		Run:         "🚀",
		Empty:       "  ",
		colorIndex:  0,
		iconIndex:   0,
		iconList: []string{
			"🍏", "🍎", "🍌", "🍉", "🍇", "🍐", "🍊", "🍋", "🍓", "🍈", "🍒", "🍑", "🍍", "🥝", "🍅", "🍆", "🥑",
		},
		Icon: func(icon string) string { return icon },
	}
	d.colorList = []string{
		d.Green,
		d.Yellow,
		d.Blue,
		d.Magenta,
		d.Cyan,
		d.Bold + d.Green,
		d.Bold + d.Yellow,
		d.Bold + d.Blue,
		d.Bold + d.Magenta,
		d.Bold + d.Cyan,
	}
	return d
}

// GenerateColor new color
func (d *Decoration) GenerateColor() string {
	if d.colorIndex >= len(d.colorList) {
		d.colorIndex = 0
	}
	color := d.colorList[d.colorIndex]
	d.colorIndex++
	return color
}

// GenerateIcon new icon
func (d *Decoration) GenerateIcon() string {
	if d.iconIndex >= len(d.iconList) {
		d.iconIndex = 0
	}
	icon := d.iconList[d.iconIndex]
	d.iconIndex++
	return icon
}
