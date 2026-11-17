package compose

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Title(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	// A month file opens straight onto the calendar; no cover page.
	if cfg.SplitByMonth() {
		return page.Modules{}, nil
	}

	return page.Modules{{Cfg: cfg, Tpl: tpls[0]}}, nil
}
