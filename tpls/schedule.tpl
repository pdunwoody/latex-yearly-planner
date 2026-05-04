\myUnderline{Schedule\textcolor{white}{g}}\vskip-\myLenLineThicknessDefault
{{range $hour := .Day.Hours .Cfg.Layout.Numbers.DailyBottomHour .Cfg.Layout.Numbers.DailyTopHour -}}
{{if and (ge $hour.Time.Hour $.Cfg.Layout.Numbers.DailyBottomHour) (le $hour.Time.Hour $.Cfg.Layout.Numbers.DailyTopHour) -}}
\myScheduleLineHeight%
{{if $.Cfg.AMPMTime -}}
\parbox{9mm}{\hfill\small {{- $hour.FormatHour $.Cfg.AMPMTime -}} }%
{{- else -}}
{\small {{- $hour.FormatHour $.Cfg.AMPMTime -}} }
{{- end -}}
{{- else -}}
\myScheduleLineHeight%
{{- end}}
\myLineLightGray\vskip\myLenScheduleLineHeight\myLineGray
{{- end}}
{{if $.Cfg.AddLastHalfHour}}\vskip\myLenScheduleLineHeight\vbox to 0pt{\myLineLightGray}{{end}}
