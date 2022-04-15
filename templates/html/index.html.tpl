<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>简历</title>
    <script src="/static/js/live.js"></script>
    <link rel="stylesheet" href="/static/css/base.css" />
    <link rel="stylesheet" href="/static/css/theme.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/flexboxgrid@6.3.1/dist/flexboxgrid.min.css" />
  </head>
  <body>
    <div class="container">
      <!-- name card -->
      <div class="row">
        <div class="col-lg-6 col-md-6 col-sm-6 col-xs-6">
          <h1 class="name">
            {{ .Me.Name }}
            <small class="title">{{ .Me.Title }}</small>
          </h1>
        </div>
        <div class="col-lg-6 col-md-6 col-sm-6 col-xs-6">
          <p class="text-right">手机号码：{{ .Me.Contact.Phone }}</p>
          <p class="text-right">邮箱：{{ .Me.Contact.EMail }}</p>
          <p class="text-right">
            GitHub: &ensp;<a
              href="http://github.com/{{ .Me.SocialMedia.GitHub }}/"
              >{{ .Me.SocialMedia.Github }}</a
            >
          </p>
        </div>
      </div>

      <!-- introduction your self -->
      <h2>关于我</h2>
      <div class="row">
        {{ .Introduction }}
      </div>

      <!-- show work experience -->
      <h2>工作经历</h2>
      <div class="container">
        {{ range .WorkExperience }}
        <div class="row">
          <h3>
            {{ .Organization }} <small>{{ .Title }}</small>
          </h3>
        </div>

        <div class="row end-xs text-right text-gray">
          <p>{{ .TimeRange.From }} - {{ .TimeRange.To }}</p>
        </div>

        {{ range .Projects }}
        <div class="row">
          任职期间负责项目{{.Name}}，工作职责包括{{ .Responsibility}}。
        </div>
        <div class="row">{{ .Introduction }}</div>
        <div class="row">{{ .Achievements }}</div>
        {{ end }}
        {{ end }}
      </div>

      <!-- show open source projects -->
      <h2>开源项目</h2>
      <div class="container">
      {{ range $chunk := chunk 2 .CommunityContribution }}
        <div class="row">
        {{ range $contribution := $chunk }}
        <div class="col-lg-6 col-md-6 col-sm-6 col-xs-6 oss-contribution">
          <p class="oss-contribution-name"><a href="{{ .Link }}">{{ .Name }}</a></p>
          <p class="oss-contribution-stack text-right">{{ .Stack }}</p>
          <p class="oss-contribution-introduction">{{ .Introduction }}</p>
        </div>
        {{ end }}
        </div>
      {{ end }}
      </div>

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

      <!-- education -->
      <h2>教育经历</h2>
      <div class="container">
        {{ range .Education.History }}
        <h3>学位</h3>
        <div class="row">
          <div class="col-lg-2 col-md-2 col-sm-2 col-xs-2">{{ .School }}</div>
          <div class="col-lg-2 col-md-2 col-sm-2 col-xs-2">{{ .Grade }}</div>
        </div>
        {{ end }}
        <h3>荣誉</h3>
        <ul>
        {{ range .Education.Achievements }}
        <li>{{ .Honor }}</li>
        {{ end }}
        </ul>
      </div>

      <!-- say thanks -->
      <h2>致谢</h2>
      <p>
        {{ .Thanks }}
      </p>
    </div>
  </body>
</html>
