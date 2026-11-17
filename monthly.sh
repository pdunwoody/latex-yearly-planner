#!/usr/bin/env bash

# Builds the planner as one PDF per month instead of one PDF for the whole year.
#
# Each month's file contains the annual overview, that month's calendar page and
# that month's daily pages. Links to another month (the annual page, the
# previous/next month arrows, the side tabs, the day before the 1st...) open the
# other month's PDF, so keep all twelve files together in one folder and do not
# rename them.
#
#   PLANNER_YEAR=2027 \
#   CFG="cfg/base.yaml,cfg/template_breadcrumb.yaml,cfg/paul_breadcrumb_manta.yaml" \
#   ./monthly.sh
#
# Optional variables:
#   NAME    file name prefix (default: planner)  -> planner-2027-01.pdf ...
#   MONTHS  which months to build (default: "1 2 3 4 5 6 7 8 9 10 11 12")
#   OUTDIR  where the PDFs go (default: monthly)
#   PASSES  xelatex passes per month (default: 2)

set -eo pipefail

: "${CFG:?set CFG to a comma-separated list of config files}"

export PLANNER_YEAR="${PLANNER_YEAR:-$(date +%Y)}"
export PLANNER_SPLIT_PREFIX="${NAME:-planner}"
MONTHS="${MONTHS:-1 2 3 4 5 6 7 8 9 10 11 12}"
OUTDIR="${OUTDIR:-monthly}"
PASSES="${PASSES:-2}"

mkdir -p out "${OUTDIR}"

if [ -z "$PLANNERGEN_BINARY" ]; then
  PLANNERGEN_BINARY="out/plannergen"
  go build -o "${PLANNERGEN_BINARY}" ./cmd/plannergen
fi

nakedname=$(echo "${CFG}" | rev | cut -d, -f1 | cut -d'/' -f 1 | cut -d'.' -f 2-99 | rev)

for month in ${MONTHS}; do
  mm=$(printf '%02d' "${month}")
  target="${OUTDIR}/${PLANNER_SPLIT_PREFIX}-${PLANNER_YEAR}-${mm}.pdf"

  echo "📅 ${PLANNER_YEAR}-${mm}"

  PLANNER_MONTH="${month}" "${PLANNERGEN_BINARY}" --config "${CFG}"

  if [ -n "${TRANSLATION}" ]; then
    python3 translate.py ${TRANSLATION}
  fi

  for _ in $(seq 1 "${PASSES}"); do
    xelatex \
      -file-line-error \
      -interaction=nonstopmode \
      -halt-on-error \
      -output-directory=./out \
      "out/${nakedname}.tex" >"out/${nakedname}.${mm}.log" 2>&1 ||
      { echo "xelatex failed for ${PLANNER_YEAR}-${mm}, see out/${nakedname}.${mm}.log"; exit 1; }
  done

  cp "out/${nakedname}.pdf" "${target}"
  echo "   created ${target}"
done
