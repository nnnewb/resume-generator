{{ define "experience" }}
<!-- show work experience -->
<h2>工作经历</h2>
<div class="container">
  {{ range .WorkExperience }}
  <h3>
    {{ .Organization }} <small>{{ .Title }}</small>
  </h3>

  <p class="end-xs text-gray">{{ .TimeRange.From }} - {{ .TimeRange.To }}</p>

  {{ range .Projects }}
  <div class="row">
    <p>任职期间负责项目{{.Name}}，工作职责包括{{ .Responsibility}}。</p>
  </div>
  <div class="row">
    {{ .Introduction | Markdown | Unescape }}
  </div>
  <div class="row">
    {{ .Achievements | Markdown | Unescape }}
  </div>
  {{ end }}
  {{ end }}
</div>
{{ end }}
