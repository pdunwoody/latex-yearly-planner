{{- /* Breadcrumb header with the prev/next day links at normal size instead of \Large */ -}}
{\noindent\Large\renewcommand{\arraystretch}{\myNumArrayStretch}
{{- .Body.Breadcrumb -}}
\hfill%
{\normalsize\renewcommand{\arraystretch}{2.2}{{ .Body.Extra.Table false -}} }%
}
\myLineThick\medskip
