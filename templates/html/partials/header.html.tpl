{{ define "header" }}
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
{{ end }}