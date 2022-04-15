{{ define "skills" }}
<!-- skill stack -->
<h2>技术栈</h2>
<div class="container">
  <div class="row">
    <div class="col-lg-4 col-md-4 col-sm-4 col-xs-4">
      <ul>
        {{ range slice .Me.Skills 0 (div (len .Me.Skills) 3)}}
        <li>
          {{ . }}
        </li>
        {{ end }}
      </ul>
    </div>
    <div class="col-lg-4 col-md-4 col-sm-4 col-xs-4">
      <ul>
        {{ range slice .Me.Skills (div (len .Me.Skills) 3) (mul 2 (div (len .Me.Skills) 3)) }}
        <li>
          {{ . }}
        </li>
        {{ end }}
      </ul>
    </div>
    <div class="col-lg-4 col-md-4 col-sm-4 col-xs-4">
      <ul>
        {{ range slice .Me.Skills (mul 2 (div (len .Me.Skills) 3)) (len .Me.Skills) }}
        <li>
          {{ . }}
        </li>
        {{ end }}
      </ul>
    </div>
  </div>
</div>
{{ end }}
