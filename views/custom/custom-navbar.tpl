{{define "navbar"}}
<ul class="nav navbar-nav">
    <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
{{$UrlPath :=.UrlPath}}
{{range .Menus}}
    <li {{if eq .Path $UrlPath}}class="active"{{end}}><a href="{{.Path}}">{{.Name}}</a></li>

{{end}}
</ul>
<div class="pull-right">
    <ul class="nav navbar-nav">
    {{if .IsLogin}}
        <li><a href="#"><strong>{{.Username}}</strong></a></li>
        <li {{if eq "/password/modify" $UrlPath}}class="active"{{end}}><a href="/password/modify">修改密码</a></li>
        <li><a href="/logout">退出</a></li>
    {{else}}
        <li><a href="/login">登录</a></li>
    {{end}}
    </ul>
</div>
{{end}}