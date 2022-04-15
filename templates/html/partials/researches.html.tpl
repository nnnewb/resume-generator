{{ define "researches" }}
<!-- show your blog article or researches -->
<h2>技术文章</h2>
<div class="container">
  <div class="row">
    <div class="col-lg-6 col-md-6 col-sm-6 col-xs-6">
      <ul>
        {{ range slice .Researches 0 (div (len .Researches) 2)}}
        <li>
          <a href="{{ .Link }}">{{ .Title }}</a>
        </li>
        {{ end }}
      </ul>
    </div>
    <div class="col-lg-6 col-md-6 col-sm-6 col-xs-6">
      <ul>
        {{ range slice .Researches (div (len .Researches) 2) (len .Researches)}}
        <li>
          <a href="{{ .Link }}">{{ .Title }}</a>
        </li>
        {{ end }}
      </ul>
    </div>
  </div>
</div>
{{ end }}
