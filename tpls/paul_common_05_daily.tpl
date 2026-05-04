{{- $today := .Body.Day -}}

\begin{minipage}[t]{\dimexpr0.45\linewidth-0.5\myLenTriColSep\relax}
{{template "schedule.tpl" dict "Cfg" .Cfg "Day" .Body.Day}}
  \vspace{\dimexpr4mm+.3pt}

{{- if .Cfg.CalAfterSchedule -}}
{{- template "monthTabularV2.tpl" dict "Month" .Body.Month "Today" $today -}}
{{- end -}}
\end{minipage}%
\hspace{\myLenTriColSep}%
\begin{minipage}[t]{\dimexpr0.55\linewidth-0.5\myLenTriColSep\relax}
  \myUnderline{Top priorities\myDummyQ}
  \Repeat{4}{\myTodoLineGrayDouble}
  \vskip\dimexpr5.4mm
  \myUnderline{Short Tasks\myDummyQ}
  \Repeat{6}{\myTodoLineGrayDouble}
  \vskip\dimexpr5.4mm
  \myUnderline{Notes \hfill{}{{ $today.LinkLeaf "More" "More" }}}
  \myPaulGrid{\myNumDailyNotes}{16}
\end{minipage}
\par\pagebreak