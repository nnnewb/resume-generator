{{ define "layout" }}
<div class="container">
    {{/* 姓名联系方式 */}}
    {{ template "header" . }}
    {{/* 介绍 */}}
    {{ template "introduction" . }}
    {{/* 工作经验 */}}
    {{ template "experience" . }}
    {{/* 开源项目 */}}
    {{ template "oss-contributions" . }}
    {{/* 研究成果、博客文章 */}}
    {{ template "researches" . }}
    {{/* 职业技能 */}}
    {{ template "skills" . }}
    {{/* 教育经历 */}}
    {{ template "education" . }}
    {{/* 感谢或者别的什么 */}}
    {{ template "footer" . }}
</div>
{{ end }}