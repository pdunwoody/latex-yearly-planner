package compose

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

// CrossMonthLinks returns LaTeX that registers, for a one-month build, every
// monthly/daily page target that lives in another month's file. The document
// preamble then turns \hyperlink calls to those targets into links that open
// the other file at that page, instead of links to a page that isn't there.
func CrossMonthLinks(cfg config.Config) string {
	if !cfg.SplitByMonth() {
		return ""
	}

	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	sb := strings.Builder{}

	for _, ref := range year.MonthAndDayRefs("", "More", "Reflect") {
		if cfg.IncludesMonth(ref.Month) {
			continue
		}

		sb.WriteString(`\myRemoteTarget{` + ref.Ref + `}{` + cfg.MonthFile(ref.Month) + "}\n")
	}

	return sb.String()
}
