package config

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug Debug

	Year                int    `env:"PLANNER_YEAR"`
	Month               int    `env:"PLANNER_MONTH"`
	SplitPrefix         string `env:"PLANNER_SPLIT_PREFIX"`
	WeekStart           time.Weekday
	Dotted              bool
	CalAfterSchedule    bool
	ClearTopRightCorner bool
	AMPMTime            bool
	AddLastHalfHour     bool
	// CombineWeekend puts Saturday and Sunday on one shared daily page
	// (the daily block's second template) and drops their notes/reflect pages.
	CombineWeekend bool

	Pages Pages

	Layout Layout
}

type Debug struct {
	ShowFrame bool
	ShowLinks bool
}

type Pages []Page
type Page struct {
	Name         string
	RenderBlocks RenderBlocks
}

type RenderBlocks []RenderBlock

func (r Pages) WeeklyEnabled() bool {
	for _, s := range r {
		for _, block := range s.RenderBlocks {
			if block.FuncName == "weekly" {
				return true
			}
		}
	}

	return false
}

func (r Pages) QuarterlyEnabled() bool {
	for _, s := range r {
		for _, block := range s.RenderBlocks {
			if block.FuncName == "quarterly" {
				return true
			}
		}
	}

	return false
}

type RenderBlock struct {
	FuncName string
	Tpls     []string
}

type Colors struct {
	Gray      string
	LightGray string
}

type Layout struct {
	Paper Paper

	Numbers Numbers
	Lengths Lengths
	Colors  Colors
}

type Numbers struct {
	ArrayStretch        float64
	QuarterlyLines      int
	WeeklyLines         int
	DailyTodos          int
	DailyNotes          int
	DailyPersonal       int
	DailyBottomHour     int
	DailyTopHour        int
	DailyDiaryGoals     int
	DailyDiaryGrateful  int
	DailyDiaryBest      int
	DailyDiaryLog       int
	TodoLinesInTodoPage int
	IndexMeetingNotes   int
	NotesIndexPages     int
	NotesOnPage         int
	DotHeightFull       int
	DotWidthFull        int
	DotWidthTwoThirds   float64
	ScheduleLineHeight  float64
}

type Paper struct {
	Width  string `env:"PLANNER_LAYOUT_PAPER_WIDTH"`
	Height string `env:"PLANNER_LAYOUT_PAPER_HEIGHT"`

	Margin Margin

	ReverseMargins bool
	MarginParWidth string
	MarginParSep   string
}

type Margin struct {
	Top    string `env:"PLANNER_LAYOUT_PAPER_MARGIN_TOP"`
	Bottom string `env:"PLANNER_LAYOUT_PAPER_MARGIN_BOTTOM"`
	Left   string `env:"PLANNER_LAYOUT_PAPER_MARGIN_LEFT"`
	Right  string `env:"PLANNER_LAYOUT_PAPER_MARGIN_RIGHT"`
}

func New(pathConfigs ...string) (Config, error) {
	var (
		bts []byte
		err error
		cfg Config
	)

	for _, filepath := range pathConfigs {
		if bts, err = ioutil.ReadFile(strings.ToLower(filepath)); err != nil {
			return cfg, fmt.Errorf("ioutil read file: %w", err)
		}

		if err = yaml.Unmarshal(bts, &cfg); err != nil {
			return cfg, fmt.Errorf("yaml unmarshal: %w", err)
		}
	}

	if err = env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("env parse: %w", err)
	}

	if cfg.Year == 0 {
		cfg.Year = time.Now().Year()
	}

	if cfg.Month < 0 || cfg.Month > 12 {
		return cfg, fmt.Errorf("month must be between 1 and 12 (or 0 for the whole year), got %d", cfg.Month)
	}

	if cfg.SplitPrefix == "" {
		cfg.SplitPrefix = "planner"
	}

	return cfg, nil
}

// SplitByMonth reports whether this build produces a single month of the year
// (PLANNER_MONTH=1..12) instead of the whole year.
func (c Config) SplitByMonth() bool {
	return c.Month != 0
}

// IncludesMonth reports whether pages of the given month belong in this build.
func (c Config) IncludesMonth(month time.Month) bool {
	return c.Month == 0 || time.Month(c.Month) == month
}

// MonthFile is the file name of the PDF holding the given month when the
// planner is split into one file per month. Cross-month links point to it,
// so it must match the names produced by the build script.
func (c Config) MonthFile(month time.Month) string {
	return fmt.Sprintf("%s-%d-%02d.pdf", c.SplitPrefix, c.Year, int(month))
}
