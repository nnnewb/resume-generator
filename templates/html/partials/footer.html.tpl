{{ define "footer" }}
<!-- say thanks -->
<h2>致谢</h2>
<p>
  {{ .Thanks | Markdown | Unescape }}
</p>
{{ end }}
