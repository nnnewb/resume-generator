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
    任职期间负责项目{{.Name}}，工作职责包括{{ .Responsibility}}。
  </div>
  <div class="row">{{ .Introduction }}</div>
  <div class="row">{{ .Achievements }}</div>
  {{ end }}
  {{ end }}
</div>
{{ end }}
