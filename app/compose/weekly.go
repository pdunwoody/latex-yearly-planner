package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Weekly(cfg config.Config, tpls []string) (page.Modules, error) {
	modules := make(page.Modules, 0, 53)
	year := cal.NewYear(cfg.WeekStart, cfg.Year)

	for _, week := range year.Weeks {
		if !weekIncluded(cfg, week) {
			continue
		}

		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Year":         year,
				"Week":         week,
				"Breadcrumb":   week.Breadcrumb(cfg.Pages.QuarterlyEnabled()),
				"HeadingMOS":   week.HeadingMOS(),
				"SideQuarters": year.SideQuarters(week.Quarters.Numbers()...),
				"SideMonths":   year.SideMonths(week.Months.Months()...),
				"Extra":        week.PrevNext().WithTopRightCorner(cfg.ClearTopRightCorner),
				"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, nil, 0),
			},
		})
	}

	return modules, nil
}

// weekIncluded keeps a week when any of its days falls in the month being built.
func weekIncluded(cfg config.Config, week *cal.Week) bool {
	if !cfg.SplitByMonth() {
		return true
	}

	for _, day := range week.Days {
		if day.Time.Year() == cfg.Year && cfg.IncludesMonth(day.Time.Month()) {
			return true
		}
	}

	return false
}
