<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>简历</title>
    <script src="/static/js/live.js"></script>
  </head>
  <body>
    <div class="page">
      <div class="head">
        <div class="head-left">
          <h2 class="name">
            {{ .Me.Name }}
            <small class="title">{{ .Me.Title }}</small>
          </h2>
        </div>
        <div class="head-right">
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
    </div>
  </body>
  <style>
    .text-right {
      text-align: right;
    }
  </style>
  <style>
    .page {
      margin: auto;
      max-width: 1080px;
    }
    .head {
      height: 155px;
      width: 100%;
      display: flex;
    }
    .head-left {
      width: 50%;
      padding: 0;
      margin: 0;
    }
    .head-right {
      padding-top: 2.5em;
      width: 50%;
      vertical-align: top;
      line-height: 1em;
    }
    .name {
      font-size: 42px;
    }
    .name > .title {
      color: gray;
    }
  </style>
</html>
