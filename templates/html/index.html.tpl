<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>简历</title>
    <script src="/static/js/live.js"></script>
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/flexboxgrid@6.3.1/dist/flexboxgrid.min.css"
    />
    <link rel="stylesheet" href="/static/css/base.css" />
  </head>
  <body>
    <div class="container">
      <!-- name card -->
      <div class="row">
        <div class="col-lg-6 col-md-6 col-sm-16 col-xs-12">
          <h2 class="name">
            {{ .Me.Name }}
            <small class="title">{{ .Me.Title }}</small>
          </h2>
        </div>
        <div class="col-lg-6 col-md-6 col-sm-16 col-xs-12">
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
      <div class="row">
        {{ .Introduction }}
      </div>

      <hr />

      <!-- show work experience -->
      <div class="container">
        {{ range .WorkExperience }}
        <!-- head line -->
        <div class="row">
          <h3>
            {{ .Organization }} <small>{{ .Title }}</small>
          </h3>
        </div>

        <div class="row end-xs text-right text-gray">
          <p>{{ .TimeRange.From }} - {{ .TimeRange.To }}</p>
        </div>

        <!-- project introduction -->
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
    </div>
  </body>
</html>
