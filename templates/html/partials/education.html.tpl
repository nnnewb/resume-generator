{{ define "education" }}
<!-- education -->
<h2>教育经历</h2>
<div class="container">
  {{ range.Education.History }}
  <h3>学位</h3>
  <div class="row">
    <div class="col-lg-2 col-md-2 col-sm-2 col-xs-2">{{ .School }}</div>
    <div class="col-lg-2 col-md-2 col-sm-2 col-xs-2">{{ .Grade }}</div>
  </div>
  {{ end }}
  <h3>荣誉</h3>
  <ul>
    {{ range.Education.Achievements }}
    <li>{{ .Honor }}</li>
    {{ end }}
  </ul>
</div>
{{ end }}
