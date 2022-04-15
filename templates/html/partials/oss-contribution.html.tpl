{{ define "oss-contributions" }}
<!-- show open source projects -->
<h2>开源项目</h2>
<div class="container">
  {{ range $chunk := chunk 2 .CommunityContribution }}
  <div class="row">
    {{ range $contribution := $chunk }}
    <div class="col-lg-6 col-md-6 col-sm-6 col-xs-6 oss-contribution">
      <p class="oss-contribution-name">
        <a href="{{ .Link }}">{{ .Name }}</a>
      </p>
      <p class="oss-contribution-stack text-right">{{ .Stack }}</p>
      <p class="oss-contribution-introduction">{{ .Introduction }}</p>
    </div>
    {{ end }}
  </div>
  {{ end }}
</div>
{{ end }}
