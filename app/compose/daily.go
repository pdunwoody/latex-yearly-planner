package compose

import (
	"errors"

	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

var Daily = DailyStuff("", "")
var DailyReflect = DailyStuff("Reflect", "Reflect")
var DailyNotes = DailyStuff("More", "Notes")

func DailyStuff(prefix, leaf string) func(cfg config.Config, tpls []string) (page.Modules, error) {
	return func(cfg config.Config, tpls []string) (page.Modules, error) {
		year := cal.NewYear(cfg.WeekStart, cfg.Year)
		modules := make(page.Modules, 0, 366)
		combine := cfg.CombineWeekend

		if combine && len(leaf) == 0 && len(tpls) < 2 {
			return nil, errors.New("combineweekend needs a second (weekend) template in the daily block")
		}

		// With combined weekends only the weekday pages have notes/reflect pages.
		var skip func(cal.Day) bool
		if combine && len(leaf) > 0 {
			skip = cal.Day.IsWeekend
		}

		for _, quarter := range year.Quarters {
			for _, month := range quarter.Months {
				if !cfg.IncludesMonth(month.Month) {
					continue
				}

				days := monthDays(month)

				for i := 0; i < len(days); i++ {
					day := days[i]
					week := weekOf(month, day)

					if combine && day.IsWeekend() {
						if len(leaf) > 0 {
							continue
						}

						// Saturday and the Sunday after it share a page, unless a
						// month ends between them; then each gets the page alone.
						span := cal.Days{day}
						if i+1 < len(days) && day.Next().Time.Equal(days[i+1].Time) && days[i+1].IsWeekend() {
							span = append(span, days[i+1])
							i++
						}

						modules = append(modules, page.Module{
							Cfg: cfg,
							Tpl: tpls[1],
							Body: map[string]interface{}{
								"Year":         year,
								"Quarter":      quarter,
								"Month":        month,
								"Week":         week,
								"Day":          *day,
								"Days":         span,
								"Breadcrumb":   cal.SpanBreadcrumb(span, cfg.Pages.QuarterlyEnabled(), cfg.Pages.WeeklyEnabled()),
								"HeadingMOS":   day.HeadingMOS("", ""),
								"SideQuarters": year.SideQuarters(day.Quarter()),
								"SideMonths":   year.SideMonths(day.Month()),
								"Extra":        cal.SpanPrevNext(*span[0], *span[len(span)-1]).WithTopRightCorner(cfg.ClearTopRightCorner),
								"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, week, 0),
							},
						})

						continue
					}

					modules = append(modules, page.Module{
						Cfg: cfg,
						Tpl: tpls[0],
						Body: map[string]interface{}{
							"Year":         year,
							"Quarter":      quarter,
							"Month":        month,
							"Week":         week,
							"Day":          *day,
							"Breadcrumb":   day.Breadcrumb(prefix, leaf, cfg.ClearTopRightCorner && len(leaf) > 0, cfg.Pages.QuarterlyEnabled(), cfg.Pages.WeeklyEnabled()),
							"HeadingMOS":   day.HeadingMOS(prefix, leaf),
							"SideQuarters": year.SideQuarters(day.Quarter()),
							"SideMonths":   year.SideMonths(day.Month()),
							"Extra":        day.PrevNextSkipping(prefix, skip).WithTopRightCorner(cfg.ClearTopRightCorner),
							"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, week, 0),
						},
					})
				}
			}
		}

		return modules, nil
	}
}

// monthDays lists the month's days in order.
func monthDays(month *cal.Month) cal.Days {
	days := make(cal.Days, 0, 31)

	for _, week := range month.Weeks {
		for i := range week.Days {
			if !week.Days[i].Time.IsZero() {
				days = append(days, &week.Days[i])
			}
		}
	}

	return days
}

func weekOf(month *cal.Month, day *cal.Day) *cal.Week {
	for _, week := range month.Weeks {
		for i := range week.Days {
			if &week.Days[i] == day {
				return week
			}
		}
	}

	return nil
}
