{{- /* Saturday and Sunday on one page: each half has a todo list and a 5mm grid */ -}}
{{- range $i, $day := .Body.Days }}
{{- if $i }}\vskip 5mm{{ end }}
\myUnderline{\textbf{ {{- $day.Time.Format "Monday, January 2" -}} }\myDummyQ}
\begin{minipage}[t]{\dimexpr0.45\linewidth-0.5\myLenTriColSep\relax}
  \vskip-\myLenLineHeightButLine
  \Repeat{17}{\myTodoLineGray}
\end{minipage}%
\hspace{\myLenTriColSep}%
\begin{minipage}[t]{\dimexpr0.55\linewidth-0.5\myLenTriColSep\relax}
  \vskip 0pt
  \myPaulGrid{17}{16}
\end{minipage}
\par
{{- end }}
\pagebreak
