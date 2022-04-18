{{ define "experience" }}
<!-- show work experience -->
<h2>工作经历</h2>
<div class="container">
  {{ range .WorkExperience }}
  <h3>{{ .Organization }} <small>{{ .Title }}</small><span class="time-range">{{ .TimeRange.From }} ~ {{ .TimeRange.To }}</span></h3>
  <div>
    {{ .Introduction | Markdown | Unescape }}
  </div>
  {{ end }}
</div>
{{ end }}
