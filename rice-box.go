package forms

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `baseform.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("<form{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\r\n\t{{ range .fields}}{{ .Render \"\" }}{{end}}\r\n</form>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/button.html`,
		FileModTime: time.Unix(1510041316, 0),
		Content:     string("{{ define \"main\"}}<button type=\"{{.type}}\" name=\"{{.name}}\" class='btn {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}{{if eq .type \"submit\"}}btn-default{{end}}'{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</button>{{end}}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/cron.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{ define \"main\"}}\r\n{{- if .id}}{{set . \"widget_id\" .id}}{{else}}{{generateID | set . \"widget_id\" }} {{end}}\r\n<script>\r\n      if (tpt_form_callbacks == null) {\r\n        tpt_form_callbacks = new Array();\r\n      }\r\n      tpt_form_callbacks.push(function () {\r\n            $('#{{.widget_id}}_cornExpression a').click(function (e) {\r\n                e.preventDefault();\r\n                $(this).tab('show')\r\n            });\r\n    //        function offset(){\r\n    //            var ele = $(\"#offset\");\r\n    //            return ele.val()?\"_\"+ele.val():\"\"\r\n    //        }\r\n            var cornExpression = [];\r\n            $(\"#{{.widget_id}}_month select\").on(\"change\", function () {\r\n                var month_monthDay = $(\"#{{.widget_id}}_month_monthDay\").val(),\r\n                        month_hour = $(\"#{{.widget_id}}_month_hour\").val(),\r\n                        month_minute = $(\"#{{.widget_id}}_month_minute\").val(),\r\n                        cornExpression = [\"0\", month_minute, month_hour, month_monthDay, \"*\", \"?\" ];\r\n                $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\r\n            });\r\n            $(\"#{{.widget_id}}_week select\").on(\"change\", function () {\r\n                var week_weekDay = $(\"#{{.widget_id}}_week_weekDay\").val(),\r\n                        week_hour = $(\"#{{.widget_id}}_week_hour\").val(),\r\n                        week_minute = $(\"#{{.widget_id}}_week_minute\").val(),\r\n                        cornExpression = [\"0\", week_minute, week_hour, \"?\", \"*\", week_weekDay ];\r\n                $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\r\n            });\r\n            $(\"#{{.widget_id}}_day select\").on(\"change\", function () {\r\n                var day_hour = $(\"#{{.widget_id}}_day_hour\").val(),\r\n                        day_minute = $(\"#{{.widget_id}}_day_minute\").val(),\r\n                        cornExpression = [\"0\", day_minute, day_hour, \"*\", \"*\", \"?\" ];\r\n                $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\r\n            });\r\n            $(\"#{{.widget_id}}_hour select\").on(\"change\", function () {\r\n                var hour_hour = $(\"#{{.widget_id}}_hour_hour\").val(),\r\n                        cornExpression = [\"0\", \"0\", \"0/\" + hour_hour, \"*\", \"*\", \"?\" ];\r\n                $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\r\n            });\r\n            $(\"#{{.widget_id}}_offset\").on(\"change\", function () {\r\n                $(\".tab-content .active select\").trigger(\"change\")\r\n            });\r\n            $(\"#{{.widget_id}}_cornExpression\").on(\"change\", \"select\", function () {\r\n                $(\"[name='{{.name}}']\").trigger(\"change\")\r\n            });\r\n        });\r\n</script>\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label{{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\" \r\n    {{- if .id}} for=\"{{.id}}\" {{end}}>{{ if .label }}{{.label}}{{end}}</label>\r\n    <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n        <div class=\"list-group-item\">\r\n            <ul class=\"nav nav-tabs\" role=\"tablist\" id=\"{{.widget_id}}_cornExpression\">\r\n                <li role=\"presentation\" class=\"active\"><a href=\"#{{.widget_id}}_month\" style=\"color:#555\">每月</a></li>\r\n                <li role=\"presentation\"><a href=\"#{{.widget_id}}_week\" style=\"color:#555\">每周</a></li>\r\n                <li role=\"presentation\"><a href=\"#{{.widget_id}}_day\"  style=\"color:#555\">每日</a></li>\r\n                <li role=\"presentation\"><a href=\"#{{.widget_id}}_hour\" style=\"color:#555\">每时</a></li>\r\n            </ul>\r\n\r\n            <div class=\"tab-content\">\r\n                <div role=\"tabpanel\" class=\"tab-pane active\" id=\"{{.widget_id}}_month\">\r\n                    <span>每月</span>\r\n                    <select id=\"{{.widget_id}}_month_monthDay\">\r\n                        <option value=\"1\">1日</option>\r\n                        <option value=\"2\">2日</option>\r\n                        <option value=\"3\">3日</option>\r\n                        <option value=\"4\">4日</option>\r\n                        <option value=\"5\">5日</option>\r\n                        <option value=\"6\">6日</option>\r\n                        <option value=\"7\">7日</option>\r\n                        <option value=\"8\">8日</option>\r\n                        <option value=\"9\">9日</option>\r\n                        <option value=\"10\">10日</option>\r\n                        <option value=\"11\">11日</option>\r\n                        <option value=\"12\">12日</option>\r\n                        <option value=\"13\">13日</option>\r\n                        <option value=\"14\">14日</option>\r\n                        <option value=\"15\">15日</option>\r\n                        <option value=\"16\">16日</option>\r\n                        <option value=\"17\">17日</option>\r\n                        <option value=\"18\">18日</option>\r\n                        <option value=\"19\">19日</option>\r\n                        <option value=\"20\">20日</option>\r\n                        <option value=\"21\">21日</option>\r\n                        <option value=\"22\">22日</option>\r\n                        <option value=\"23\">23日</option>\r\n                        <option value=\"24\">24日</option>\r\n                        <option value=\"25\">25日</option>\r\n                        <option value=\"26\">26日</option>\r\n                        <option value=\"27\">27日</option>\r\n                        <option value=\"28\">28日</option>\r\n                        <option value=\"29\">29日</option>\r\n                        <option value=\"30\">30日</option>\r\n                        <option value=\"30\">31日</option>\r\n                    </select>\r\n                    <span>的</span>\r\n                    <select id=\"{{.widget_id}}_month_hour\">\r\n                        <option value=\"0\">0时</option>\r\n                        <option value=\"1\">1时</option>\r\n                        <option value=\"2\">2时</option>\r\n                        <option value=\"3\">3时</option>\r\n                        <option value=\"4\">4时</option>\r\n                        <option value=\"5\">5时</option>\r\n                        <option value=\"6\">6时</option>\r\n                        <option value=\"7\">7时</option>\r\n                        <option value=\"8\">8时</option>\r\n                        <option value=\"9\">9时</option>\r\n                        <option value=\"10\">10时</option>\r\n                        <option value=\"11\">11时</option>\r\n                        <option value=\"12\">12时</option>\r\n                        <option value=\"13\">13时</option>\r\n                        <option value=\"14\">14时</option>\r\n                        <option value=\"15\">15时</option>\r\n                        <option value=\"16\">16时</option>\r\n                        <option value=\"17\">17时</option>\r\n                        <option value=\"18\">18时</option>\r\n                        <option value=\"19\">19时</option>\r\n                        <option value=\"20\">20时</option>\r\n                        <option value=\"21\">21时</option>\r\n                        <option value=\"22\">22时</option>\r\n                        <option value=\"23\">23时</option>\r\n                    </select>\r\n                    <span>:</span>\r\n                    <select id=\"{{.widget_id}}_month_minute\">\r\n                        <option value=\"0\">0分</option>\r\n                        <option value=\"1\">1分</option>\r\n                        <option value=\"2\">2分</option>\r\n                        <option value=\"3\">3分</option>\r\n                        <option value=\"4\">4分</option>\r\n                        <option value=\"5\">5分</option>\r\n                        <option value=\"6\">6分</option>\r\n                        <option value=\"7\">7分</option>\r\n                        <option value=\"8\">8分</option>\r\n                        <option value=\"9\">9分</option>\r\n                        <option value=\"10\">10分</option>\r\n                        <option value=\"11\">11分</option>\r\n                        <option value=\"12\">12分</option>\r\n                        <option value=\"13\">13分</option>\r\n                        <option value=\"14\">14分</option>\r\n                        <option value=\"15\">15分</option>\r\n                        <option value=\"16\">16分</option>\r\n                        <option value=\"17\">17分</option>\r\n                        <option value=\"18\">18分</option>\r\n                        <option value=\"19\">19分</option>\r\n                        <option value=\"20\">20分</option>\r\n                        <option value=\"21\">21分</option>\r\n                        <option value=\"22\">22分</option>\r\n                        <option value=\"23\">23分</option>\r\n                        <option value=\"24\">24分</option>\r\n                        <option value=\"25\">25分</option>\r\n                        <option value=\"26\">26分</option>\r\n                        <option value=\"27\">27分</option>\r\n                        <option value=\"28\">28分</option>\r\n                        <option value=\"29\">29分</option>\r\n                        <option value=\"30\">30分</option>\r\n                        <option value=\"31\">31分</option>\r\n                        <option value=\"32\">32分</option>\r\n                        <option value=\"33\">33分</option>\r\n                        <option value=\"34\">34分</option>\r\n                        <option value=\"35\">35分</option>\r\n                        <option value=\"36\">36分</option>\r\n                        <option value=\"37\">37分</option>\r\n                        <option value=\"38\">38分</option>\r\n                        <option value=\"39\">39分</option>\r\n                        <option value=\"40\">40分</option>\r\n                        <option value=\"41\">41分</option>\r\n                        <option value=\"42\">42分</option>\r\n                        <option value=\"43\">43分</option>\r\n                        <option value=\"44\">44分</option>\r\n                        <option value=\"45\">45分</option>\r\n                        <option value=\"46\">46分</option>\r\n                        <option value=\"47\">47分</option>\r\n                        <option value=\"48\">48分</option>\r\n                        <option value=\"49\">49分</option>\r\n                        <option value=\"50\">50分</option>\r\n                        <option value=\"51\">51分</option>\r\n                        <option value=\"52\">52分</option>\r\n                        <option value=\"53\">53分</option>\r\n                        <option value=\"54\">54分</option>\r\n                        <option value=\"55\">55分</option>\r\n                        <option value=\"56\">56分</option>\r\n                        <option value=\"57\">57分</option>\r\n                        <option value=\"58\">58分</option>\r\n                        <option value=\"59\">59分</option>\r\n                    </select>\r\n                </div>\r\n                <div role=\"tabpanel\" class=\"tab-pane\" id=\"{{.widget_id}}_week\">\r\n                    <span>每周:</span>\r\n                    <select id=\"{{.widget_id}}_week_weekDay\">\r\n                        <option value=\"SUN\">星期日</option>\r\n                        <option value=\"MON\">星期一</option>\r\n                        <option value=\"TUE\">星期二</option>\r\n                        <option value=\"WED\">星期三</option>\r\n                        <option value=\"THU\">星期四</option>\r\n                        <option value=\"FIR\">星期五</option>\r\n                        <option value=\"SAT\">星期六</option>\r\n                    </select>\r\n                    <span>的</span>\r\n                    <select id=\"{{.widget_id}}_week_hour\">\r\n                        <option value=\"0\">0时</option>\r\n                        <option value=\"1\">1时</option>\r\n                        <option value=\"2\">2时</option>\r\n                        <option value=\"3\">3时</option>\r\n                        <option value=\"4\">4时</option>\r\n                        <option value=\"5\">5时</option>\r\n                        <option value=\"6\">6时</option>\r\n                        <option value=\"7\">7时</option>\r\n                        <option value=\"8\">8时</option>\r\n                        <option value=\"9\">9时</option>\r\n                        <option value=\"10\">10时</option>\r\n                        <option value=\"11\">11时</option>\r\n                        <option value=\"12\">12时</option>\r\n                        <option value=\"13\">13时</option>\r\n                        <option value=\"14\">14时</option>\r\n                        <option value=\"15\">15时</option>\r\n                        <option value=\"16\">16时</option>\r\n                        <option value=\"17\">17时</option>\r\n                        <option value=\"18\">18时</option>\r\n                        <option value=\"19\">19时</option>\r\n                        <option value=\"20\">20时</option>\r\n                        <option value=\"21\">21时</option>\r\n                        <option value=\"22\">22时</option>\r\n                        <option value=\"23\">23时</option>\r\n                    </select>\r\n                    <span>:</span>\r\n                    <select id=\"{{.widget_id}}_week_minute\">\r\n                        <option value=\"0\">0分</option>\r\n                        <option value=\"1\">1分</option>\r\n                        <option value=\"2\">2分</option>\r\n                        <option value=\"3\">3分</option>\r\n                        <option value=\"4\">4分</option>\r\n                        <option value=\"5\">5分</option>\r\n                        <option value=\"6\">6分</option>\r\n                        <option value=\"7\">7分</option>\r\n                        <option value=\"8\">8分</option>\r\n                        <option value=\"9\">9分</option>\r\n                        <option value=\"10\">10分</option>\r\n                        <option value=\"11\">11分</option>\r\n                        <option value=\"12\">12分</option>\r\n                        <option value=\"13\">13分</option>\r\n                        <option value=\"14\">14分</option>\r\n                        <option value=\"15\">15分</option>\r\n                        <option value=\"16\">16分</option>\r\n                        <option value=\"17\">17分</option>\r\n                        <option value=\"18\">18分</option>\r\n                        <option value=\"19\">19分</option>\r\n                        <option value=\"20\">20分</option>\r\n                        <option value=\"21\">21分</option>\r\n                        <option value=\"22\">22分</option>\r\n                        <option value=\"23\">23分</option>\r\n                        <option value=\"24\">24分</option>\r\n                        <option value=\"25\">25分</option>\r\n                        <option value=\"26\">26分</option>\r\n                        <option value=\"27\">27分</option>\r\n                        <option value=\"28\">28分</option>\r\n                        <option value=\"29\">29分</option>\r\n                        <option value=\"30\">30分</option>\r\n                        <option value=\"31\">31分</option>\r\n                        <option value=\"32\">32分</option>\r\n                        <option value=\"33\">33分</option>\r\n                        <option value=\"34\">34分</option>\r\n                        <option value=\"35\">35分</option>\r\n                        <option value=\"36\">36分</option>\r\n                        <option value=\"37\">37分</option>\r\n                        <option value=\"38\">38分</option>\r\n                        <option value=\"39\">39分</option>\r\n                        <option value=\"40\">40分</option>\r\n                        <option value=\"41\">41分</option>\r\n                        <option value=\"42\">42分</option>\r\n                        <option value=\"43\">43分</option>\r\n                        <option value=\"44\">44分</option>\r\n                        <option value=\"45\">45分</option>\r\n                        <option value=\"46\">46分</option>\r\n                        <option value=\"47\">47分</option>\r\n                        <option value=\"48\">48分</option>\r\n                        <option value=\"49\">49分</option>\r\n                        <option value=\"50\">50分</option>\r\n                        <option value=\"51\">51分</option>\r\n                        <option value=\"52\">52分</option>\r\n                        <option value=\"53\">53分</option>\r\n                        <option value=\"54\">54分</option>\r\n                        <option value=\"55\">55分</option>\r\n                        <option value=\"56\">56分</option>\r\n                        <option value=\"57\">57分</option>\r\n                        <option value=\"58\">58分</option>\r\n                        <option value=\"59\">59分</option>\r\n                    </select>\r\n\r\n                </div>\r\n                <div role=\"tabpanel\" class=\"tab-pane\" id=\"{{.widget_id}}_day\">\r\n                    <span>每天 的</span>\r\n                    <select id=\"{{.widget_id}}_day_hour\">\r\n                        <option value=\"0\">0时</option>\r\n                        <option value=\"1\">1时</option>\r\n                        <option value=\"2\">2时</option>\r\n                        <option value=\"3\">3时</option>\r\n                        <option value=\"4\">4时</option>\r\n                        <option value=\"5\">5时</option>\r\n                        <option value=\"6\">6时</option>\r\n                        <option value=\"7\">7时</option>\r\n                        <option value=\"8\">8时</option>\r\n                        <option value=\"9\">9时</option>\r\n                        <option value=\"10\">10时</option>\r\n                        <option value=\"11\">11时</option>\r\n                        <option value=\"12\">12时</option>\r\n                        <option value=\"13\">13时</option>\r\n                        <option value=\"14\">14时</option>\r\n                        <option value=\"15\">15时</option>\r\n                        <option value=\"16\">16时</option>\r\n                        <option value=\"17\">17时</option>\r\n                        <option value=\"18\">18时</option>\r\n                        <option value=\"19\">19时</option>\r\n                        <option value=\"20\">20时</option>\r\n                        <option value=\"21\">21时</option>\r\n                        <option value=\"22\">22时</option>\r\n                        <option value=\"23\">23时</option>\r\n                    </select>\r\n                    <span>:</span>\r\n                    <select id=\"{{.widget_id}}_day_minute\">\r\n                        <option value=\"0\">0分</option>\r\n                        <option value=\"1\">1分</option>\r\n                        <option value=\"2\">2分</option>\r\n                        <option value=\"3\">3分</option>\r\n                        <option value=\"4\">4分</option>\r\n                        <option value=\"5\">5分</option>\r\n                        <option value=\"6\">6分</option>\r\n                        <option value=\"7\">7分</option>\r\n                        <option value=\"8\">8分</option>\r\n                        <option value=\"9\">9分</option>\r\n                        <option value=\"10\">10分</option>\r\n                        <option value=\"11\">11分</option>\r\n                        <option value=\"12\">12分</option>\r\n                        <option value=\"13\">13分</option>\r\n                        <option value=\"14\">14分</option>\r\n                        <option value=\"15\">15分</option>\r\n                        <option value=\"16\">16分</option>\r\n                        <option value=\"17\">17分</option>\r\n                        <option value=\"18\">18分</option>\r\n                        <option value=\"19\">19分</option>\r\n                        <option value=\"20\">20分</option>\r\n                        <option value=\"21\">21分</option>\r\n                        <option value=\"22\">22分</option>\r\n                        <option value=\"23\">23分</option>\r\n                        <option value=\"24\">24分</option>\r\n                        <option value=\"25\">25分</option>\r\n                        <option value=\"26\">26分</option>\r\n                        <option value=\"27\">27分</option>\r\n                        <option value=\"28\">28分</option>\r\n                        <option value=\"29\">29分</option>\r\n                        <option value=\"30\">30分</option>\r\n                        <option value=\"31\">31分</option>\r\n                        <option value=\"32\">32分</option>\r\n                        <option value=\"33\">33分</option>\r\n                        <option value=\"34\">34分</option>\r\n                        <option value=\"35\">35分</option>\r\n                        <option value=\"36\">36分</option>\r\n                        <option value=\"37\">37分</option>\r\n                        <option value=\"38\">38分</option>\r\n                        <option value=\"39\">39分</option>\r\n                        <option value=\"40\">40分</option>\r\n                        <option value=\"41\">41分</option>\r\n                        <option value=\"42\">42分</option>\r\n                        <option value=\"43\">43分</option>\r\n                        <option value=\"44\">44分</option>\r\n                        <option value=\"45\">45分</option>\r\n                        <option value=\"46\">46分</option>\r\n                        <option value=\"47\">47分</option>\r\n                        <option value=\"48\">48分</option>\r\n                        <option value=\"49\">49分</option>\r\n                        <option value=\"50\">50分</option>\r\n                        <option value=\"51\">51分</option>\r\n                        <option value=\"52\">52分</option>\r\n                        <option value=\"53\">53分</option>\r\n                        <option value=\"54\">54分</option>\r\n                        <option value=\"55\">55分</option>\r\n                        <option value=\"56\">56分</option>\r\n                        <option value=\"57\">57分</option>\r\n                        <option value=\"58\">58分</option>\r\n                        <option value=\"59\">59分</option>\r\n                    </select>\r\n                </div>\r\n                <div role=\"tabpanel\" class=\"tab-pane\" id=\"{{.widget_id}}_hour\">\r\n                    <span>每隔</span>\r\n                    <select id=\"{{.widget_id}}_hour_hour\">\r\n                        <option value=\"1\">1时</option>\r\n                        <option value=\"2\">2时</option>\r\n                        <option value=\"3\">3时</option>\r\n                        <option value=\"4\">4时</option>\r\n                        <option value=\"5\">5时</option>\r\n                        <option value=\"6\">6时</option>\r\n                        <option value=\"7\">7时</option>\r\n                        <option value=\"8\">8时</option>\r\n                        <option value=\"9\">9时</option>\r\n                        <option value=\"10\">10时</option>\r\n                        <option value=\"11\">11时</option>\r\n                        <option value=\"12\">12时</option>\r\n                        <option value=\"13\">13时</option>\r\n                        <option value=\"14\">14时</option>\r\n                        <option value=\"15\">15时</option>\r\n                        <option value=\"16\">16时</option>\r\n                        <option value=\"17\">17时</option>\r\n                        <option value=\"18\">18时</option>\r\n                        <option value=\"19\">19时</option>\r\n                        <option value=\"20\">20时</option>\r\n                        <option value=\"21\">21时</option>\r\n                        <option value=\"22\">22时</option>\r\n                        <option value=\"23\">23时</option>\r\n                    </select>\r\n                    <span>小时</span>\r\n                </div>\r\n            </div>\r\n\r\n            <div style=\"height: 10px\"></div>\r\n            <div>\r\n                <input type=\"text\" readonly name=\"{{.name}}\" class=\"{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n                       {{- if .id}} id=\"{{.id}}\" {{end}}\r\n                       {{- if .params}}\r\n                         {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}\r\n                         {{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n                         {{- end}}\r\n                         {{- range $v :=.tags}} {{$v}}{{end}}\r\n                         {{- if .value}} value=\"{{.value}}\" {{end}}>\r\n                {{- if or .helptext .errors }}\r\n                <span class=\"help-block\">\r\n                  {{- if .helptext}}{{ .helptext }}{{end}}\r\n                  {{- if .errors}}\r\n                  <ul>\r\n                    {{- range .errors }}\r\n                    <li>{{.}}</li>\r\n                    {{- end}}\r\n                  </ul>\r\n                  {{- end}}\r\n                </span>\r\n                {{- end}}\r\n            </div>\r\n        </div>\r\n    </div>\r\n</div>\r\n{{- end}}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/date.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("\r\n{{- define \"main\"}}\r\n{{- append .ctx_parent \"moreScripts\" \"/public/js/plugins/datepicker/bootstrap-datepicker.js\"}}\r\n{{- append .ctx_parent \"moreLazyStyles\" \"/public/css/plugins/datepicker/datepicker3.css\"}}\r\n{{- unique .ctx_parent \"moreScripts\" -}}\r\n{{- unique .ctx_parent \"moreLazyStyles\" -}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n    {{- if .id}} for=\"{{.id}}\" {{end}}>\r\n    {{- if .label -}}{{.label}}{{- end -}}\r\n  </label>\r\n  \r\n  {{if .id}}{{else}}{{generateID | set . \"widget_id\" }} {{end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\" name=\"{{.name}}-form-group\" id=\"{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}}\">\r\n    <div class=\"input-group date\">\r\n      <span class=\"input-group-addon\"><i class=\"fa fa-calendar\"></i></span>\r\n      <input type=\"text\" class=\"form-control {{ if .classes }}{{range .classes}} {{.}} {{end}}{{end}}\"\r\n            name=\"{{.name}}\"\r\n            {{- if .id}} id=\"{{.id}}\" {{end -}}\r\n            {{- if .params -}} \r\n              {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}} \r\n            {{- end -}}\r\n            {{- range $v :=.tags}} {{$v}} {{end -}}\r\n            {{- if .value -}} value=\"{{form_date .value}}\" {{end}}>\r\n      {{- if or .helptext .errors -}}\r\n      <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n          {{- if .errors}}\r\n            <ul>\r\n            {{- range .errors }}\r\n              <li>{{.}}</li>\r\n            {{- end -}}\r\n            </ul>\r\n          {{- end}}\r\n    </span>\r\n    {{- end}}\r\n    </div>\r\n\r\n    <script>\r\n      if (tpt_form_callbacks == null) {\r\n        tpt_form_callbacks = new Array();\r\n      }\r\n      tpt_form_callbacks.push(function () {\r\n            $('#{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}} .input-group.date').datepicker({\r\n                todayBtn: \"linked\",\r\n                todayHighlight: false,\r\n                keyboardNavigation: true,\r\n                forceParse: true,\r\n                calendarWeeks: true,\r\n                autoclose: true,\r\n                format: \"yyyy-mm-dd\"\r\n            });\r\n        });\r\n    </script>\r\n\r\n  </div>\r\n</div>\r\n{{- end}}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/datetime.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("\r\n{{- define \"main\"}}\r\n{{- append .ctx_parent \"moreScripts\" \"/public/js/plugins/datepicker/bootstrap-datepicker.js\"}}\r\n{{- append .ctx_parent \"moreScripts\" \"/public/js/plugins/clockpicker/clockpicker.js\"}}\r\n{{- append .ctx_parent \"moreLazyStyles\" \"/public/css/plugins/datepicker/datepicker3.css\"}}\r\n{{- append .ctx_parent \"moreLazyStyles\" \"/public/css/plugins/clockpicker/clockpicker.css\"}}\r\n{{- unique .ctx_parent \"moreScripts\" -}}\r\n{{- unique .ctx_parent \"moreLazyStyles\" -}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n\t<label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n\t\t\t\t {{- if .id}} for=\"{{.id}}\" {{end}}>\r\n\t\t{{- if .label -}}{{.label}}{{- end -}}\r\n\t</label>\r\n\r\n\t{{if .id}}{{else}}{{generateID | set . \"widget_id\" }} {{end}}\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\" name=\"{{.name}}-form-group\" id=\"{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}}\">\r\n\r\n\t\t<div class=\"col-lg-6\">\r\n\t\t<div class=\"input-group date\">\r\n\t\t\t<span class=\"input-group-addon\"><i class=\"fa fa-calendar\"></i></span>\r\n\t\t\t<input type=\"text\" class=\"form-control {{ if .classes }}{{range .classes}} {{.}} {{end}}{{end}}\"\r\n\t\t\t\t\t\t name=\"{{.name}}date\"\r\n\t\t\t\t\t\t {{- if .id}} id=\"{{.id}}\" {{end -}}\r\n\t\t\t\t\t\t {{- if .params -}}\r\n\t\t\t\t\t\t {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}}\r\n\t\t\t\t\t\t {{- end -}}\r\n\t\t\t\t\t\t {{- range $v :=.tags}} {{$v}} {{end -}}\r\n\t\t\t\t\t\t {{- if .value -}} value=\"{{form_date .value}}\" {{end}}>\r\n\t\t\t{{- if or .helptext .errors -}}\r\n\t\t\t<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n          {{- if .errors}}\r\n            <ul>\r\n            {{- range .errors }}\r\n              <li>{{.}}</li>\r\n            {{- end -}}\r\n            </ul>\r\n          {{- end}}\r\n    </span>\r\n\t\t\t{{- end}}\r\n\t\t</div>\r\n\t\t</div>\r\n\r\n\t\t<div class=\"col-lg-6\">\r\n\t\t<div class=\"input-group clockpicker\" data-autoclose=\"true\">\r\n\t\t\t<input type=\"text\" class=\"form-control\" value=\"{{ if .value }}{{ form_time .value }}{{ end }}\" name=\"{{.name}}time\">\r\n\t\t\t<span class=\"input-group-addon\">\r\n\t\t\t\t<span class=\"fa fa-clock-o\"></span>\r\n\t\t\t</span>\r\n\t\t</div>\r\n\t\t</div>\r\n\t\t<input type=\"hidden\" name=\"{{.name}}\" value=\"{{ form_date_and_time .value }}\">\r\n\t\t<script>\r\n        if (tpt_form_callbacks == null) {\r\n            tpt_form_callbacks = new Array();\r\n        }\r\n        tpt_form_callbacks.push(function () {\r\n            $('#{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}} input').on(\"change\",function () {\r\n                var date = $('[name=\"{{ .name }}date\"]').val();\r\n                var time = $('[name=\"{{ .name }}time\"]').val();\r\n                if (date || time){\r\n                    $('[name=\"{{ .name }}\"]').val(date+\" \"+time);\r\n                }\r\n            });\r\n\r\n            $('#{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}} .input-group.date').datepicker({\r\n                todayBtn: \"linked\",\r\n                todayHighlight: false,\r\n                keyboardNavigation: true,\r\n                forceParse: true,\r\n                calendarWeeks: true,\r\n                autoclose: true,\r\n                format: \"yyyy-mm-dd\"\r\n            });\r\n            $('.clockpicker').clockpicker();\r\n        });\r\n\t\t</script>\r\n\r\n\t</div>\r\n</div>\r\n{{- end}}"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/time.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/generic.tmpl`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"generic\"}}<div class=\"form-group{{if .errors}} has-error{{end}}\">\n{{ if .label }}<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\n<input type=\"{{.type}}\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\n{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\n</div>\n{{end}}"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/hidden.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\" -}}\r\n<input type=\"hidden\" name=\"{{.name}}\"\r\n           {{- if .params}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{- end}} class=\"{{range .classes}}{{.}} {{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\r\n{{- end -}}"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/input.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filed := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/mapinput.html`,
		FileModTime: time.Unix(1506325202, 0),
		Content:     string("{{ define \"main\"}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{if .id}} for=\"{{.id}}\" {{end}}>{{ if .label }}{{.label}}{{end}}</label>\r\n  {{end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\r\n  {{- if or .helptext .errors }}\r\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n    {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}\r\n  </span>\r\n  {{- end}}\r\n  </div>\r\n</div>\r\n{{end}}"),
	}
	filef := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/number.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{- if .id}} for=\"{{.id}}\" {{end}}>{{- if .label }}{{.label}}{{- end}}</label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end}}\r\n           {{- if .params}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{- end}}\r\n           {{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{- end}}\r\n           {{- range $v :=.tags}} {{$v}}\r\n           {{- end}}\r\n           {{- if .value}} value=\"{{.value}}\"\r\n           {{- end}}>\r\n    {{- if or .helptext .errors }}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n            {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end -}}\r\n              </ul>\r\n            {{- end}}\r\n    </span>\r\n    {{- end}}\r\n  </div>\r\n</div>\r\n{{- end}}"),
	}
	fileg := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/range.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{- if .id}} for=\"{{.id}}\" {{end}}>{{if .label }}{{.label}}{{end}}</label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end}}\r\n           {{- if .params}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{- end}}\r\n           {{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{- end}}\r\n           {{- range $v :=.tags}} {{$v}}\r\n           {{- end}}\r\n           {{- if .value}} value=\"{{.value}}\"\r\n           {{- end}}>\r\n    {{- if or .helptext .errors }}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n      {{- if .errors}}\r\n      <ul>\r\n        {{- range .errors }}\r\n         <li>{{.}}</li>\r\n         {{- end}}\r\n      </ul>\r\n      {{- end}}\r\n    </span>\r\n    {{- end}}\r\n  </div>\r\n</div>\r\n{{- end}}"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/checkbox.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{-  define \"main\"}}\r\n{{- $p := . }}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group\">\r\n  {{- if not .nolabel -}}\r\n   <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\"></label>\r\n  {{- end -}}\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\r\n\t\t<div class=\"checkbox{{if .errors}} has-error{{end}}\">\r\n\t\t\t<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\">\r\n\t\t\t\t<input type=\"checkbox\" name=\"{{.name}}\"{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}\r\n\t\t\t\t{{- if toBoolean .value }} checked {{end}}\r\n\t\t\t\t{{- if .id}} id=\"{{.id}}\"{{end}}\r\n\t\t\t\t{{- if .params}}\r\n\t\t\t\t  {{- range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}\r\n\t\t\t\t{{- end}}\r\n\t\t\t\t{{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}\r\n\t\t\t\t{{- range $v := .tags}} {{$v}}{{end}}>\r\n\t\t\t\t{{.label}}\r\n\t\t\t</label>\r\n\t\t\t{{- if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{- end}}\r\n\t\t\t{{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\r\n\t\t</div>\r\n\t</div>\r\n</div>\r\n{{- end}}\r\n"),
	}
	filej := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/radiobutton.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}{{ $p := . }}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group\">\r\n  {{- if not .nolabel -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">\r\n    {{- if .label}}{{.label}}{{- end}}</label>\r\n  {{- end -}}\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\r\n  {{ range .choices }}\r\n    <label class=\"control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">\r\n      <input type=\"radio\" name=\"{{$p.name}}\"\r\n      {{- if $p.classes }} class=\"{{range $p.classes}}{{.}} {{end}}\"\r\n      {{- end}} value=\"{{.Value}}\" id=\"{{$p.id}}\"\r\n      {{- if $p.params}}\r\n        {{- range $k2, $v2 := $p.params}} {{$k2}}=\"{{$v2}}\"\r\n        {{- end}}\r\n      {{- end}}\r\n      {{- if $p.css}} style=\"{{range $k2, $v2 := .css}}{{$k2}}: {{$v2}}; {{end}}\"{{end}}\r\n      {{- if eq $p.value .Value}} checked{{end}}\r\n      {{- range $v2 := $p.tags}} {{$v2}}{{end}}>\r\n      {{.Label}}\r\n    </label>\r\n  {{- end}}\r\n\t</div>\r\n<!-- </div> -->\r\n\r\n</div>{{end}}\r\n"),
	}
	filek := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/select.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n      {{- if .id}} for=\"{{.id}}\" {{end -}}>\r\n    {{- if .label}}{{.label}}{{- end}}\r\n  </label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    \r\n    {{- if .preText}}<span>{{.preText}}</span>{{end -}}\r\n\r\n    <select name=\"{{.name}}\" class=\"form-control {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"\r\n            {{- if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end -}}\r\n            {{- if .css}}\r\n            style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n            {{- end -}}\r\n            {{- range $v :=.tags}} {{$v}}{{end}}>\r\n      {{- $p := . -}}\r\n      {{- range $v := .choices -}}\r\n        {{- if $v.Label -}}\r\n        <optgroup label=\"{{$v.Label}}\">\r\n        {{- end -}}\r\n\r\n        {{- range $v.Children -}}\r\n        <option value=\"{{.Value}}\"\r\n                {{- if strIn \"multiple\" $p.tags -}}\r\n                  {{- $id :=.Value -}}\r\n                  {{- range $k2, $p2 :=$p.multValues -}}\r\n                    {{- if eq $k2 $id}}selected{{end -}}\r\n                  {{- end -}}\r\n                {{- else -}}\r\n                  {{- if eq $p.value .Value}} selected{{end -}}\r\n                {{- end}}>{{raw .Label -}}\r\n        </option>\r\n        {{- end -}}\r\n        \r\n        {{- if $v.Label -}}\r\n        </optgroup>\r\n        {{- end -}}\r\n      {{- end -}}\r\n    </select>\r\n    {{- if .postText}}<span>{{.postText}}</span>{{end -}}\r\n\r\n    {{- if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end -}}\r\n  </div>\r\n</div>\r\n{{- end}}\r\n"),
	}
	filel := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/static.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"main\"}}<div class=\"form-group\">\n{{ if .label }}<label{{ if .labelClasses }} class=\"{{range .labelClasses}}{{.}} {{end}}\"{{end}}{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\n<p name=\"{{.name}}\" class=\"form-control-static {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</p>\n</div>{{end}}"),
	}
	filen := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/passwordinput.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\" {{if .id}} for=\"{{.id}}\" {{end}}>{{- if .label -}}{{.label}}{{- end -}}</label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"password\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}{{if .value}} value=\"{{.value}}\" {{end}}>\r\n    {{- if or .helptext .errors -}}\r\n    <span class=\"help-block\">\r\n      {{- if .helptext}}{{ .helptext }}{{end -}}\r\n      {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\r\n    </span>\r\n    {{- end -}}\r\n  </div>\r\n</div>{{end -}}"),
	}
	fileo := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textareainput.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{- if .id}} for=\"{{.id}}\" {{end}}>{{- if .label -}}{{.label}}{{- end}}</label>\r\n  {{- end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\r\n  {{- if or .helptext .errors -}}\r\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\r\n  </span>\r\n  {{- end -}}\r\n  </div>\r\n</div>\r\n{{- end -}}"),
	}
	filep := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textinput.html`,
		FileModTime: time.Unix(1512720743, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div {{if .id}}id='{{.id}}_div'{{end}} class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if not .nolabel -}}\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses -}}\r\n    {{range .labelClasses}} {{.}}{{end}}\r\n    {{- end}}\"\r\n  {{- if .id}} for=\"{{.id}}\" \r\n  {{- end -}}>{{- if .label -}}{{.label}}{{- end -}}</label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"text\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end}}\r\n           {{- if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}} {{end}}\r\n           {{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}\r\n           {{- range $v :=.tags}} {{$v}} {{end}}\r\n           {{- if .value}} value=\"{{.value}}\" {{end}}>\r\n    {{- if or .helptext .errors -}}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end -}}\r\n              </ul>{{end -}}\r\n    </span>\r\n    {{end -}}\r\n  </div>\r\n</div>\r\n{{- end -}}\r\n"),
	}
	fileq := &embedded.EmbeddedFile{
		Filename:    `bootstrapform.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("<form role=\"form\"{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\r\n\t{{ range .fields}}{{ .Render \"\" }}{{end}}\r\n</form>"),
	}
	filet := &embedded.EmbeddedFile{
		Filename:    `default/datetime/date.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("\r\n{{- define \"main\"}}\r\n{{- append .ctx_parent \"moreScripts\" \"/public/js/plugins/datepicker/bootstrap-datepicker.js\"}}\r\n{{- append .ctx_parent \"moreLazyStyles\" \"/public/css/plugins/datepicker/datepicker3.css\"}}\r\n{{- unique .ctx_parent \"moreScripts\" -}}\r\n{{- unique .ctx_parent \"moreLazyStyles\" -}}\r\n\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\" name=\"{{.name}}-form-group\" id=\"{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}}\">\r\n    <div class=\"input-group date\">\r\n      <span class=\"input-group-addon\"><i class=\"fa fa-calendar\"></i></span>\r\n      <input type=\"text\" class=\"form-control {{ if .classes }}{{range .classes}} {{.}} {{end}}{{end}}\"\r\n            name=\"{{.name}}\"\r\n            {{- if .id}} id=\"{{.id}}\" {{end -}}\r\n            {{- if .params -}} \r\n              {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}} \r\n            {{- end -}}\r\n            {{- range $v :=.tags}} {{$v}} {{end -}}\r\n            {{- if .value -}} value=\"{{form_date .value}}\" {{end}}>\r\n      {{- if or .helptext .errors -}}\r\n      <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n          {{- if .errors}}\r\n            <ul>\r\n            {{- range .errors }}\r\n              <li>{{.}}</li>\r\n            {{- end -}}\r\n            </ul>\r\n          {{- end}}\r\n    </span>\r\n    {{- end}}\r\n    </div>\r\n\r\n    <script>\r\n      if (tpt_form_callbacks == null) {\r\n        tpt_form_callbacks = new Array();\r\n      }\r\n      tpt_form_callbacks.push(function () {\r\n            $('#{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}} .input-group.date').datepicker({\r\n                todayBtn: \"linked\",\r\n                todayHighlight: false,\r\n                keyboardNavigation: true,\r\n                forceParse: true,\r\n                calendarWeeks: true,\r\n                autoclose: true,\r\n                format: \"yyyy-mm-dd\"\r\n            });\r\n        });\r\n    </script>\r\n  </div>"),
	}
	fileu := &embedded.EmbeddedFile{
		Filename:    `default/datetime/datetime.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("\r\n{{- define \"main\"}}\r\n{{- append .ctx_parent \"moreScripts\" \"/public/js/plugins/datepicker/bootstrap-datepicker.js\"}}\r\n{{- append .ctx_parent \"moreScripts\" \"/public/js/plugins/clockpicker/clockpicker.js\"}}\r\n{{- append .ctx_parent \"moreLazyStyles\" \"/public/css/plugins/datepicker/datepicker3.css\"}}\r\n{{- append .ctx_parent \"moreLazyStyles\" \"/public/css/plugins/clockpicker/clockpicker.css\"}}\r\n{{- unique .ctx_parent \"moreScripts\" -}}\r\n{{- unique .ctx_parent \"moreLazyStyles\" -}}\r\n\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\" name=\"{{.name}}-form-group\" id=\"{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}}\">\r\n\r\n\t\t<div class=\"col-lg-6\">\r\n\t\t<div class=\"input-group date\">\r\n\t\t\t<span class=\"input-group-addon\"><i class=\"fa fa-calendar\"></i></span>\r\n\t\t\t<input type=\"text\" class=\"form-control {{ if .classes }}{{range .classes}} {{.}} {{end}}{{end}}\"\r\n\t\t\t\t\t\t name=\"{{.name}}date\"\r\n\t\t\t\t\t\t {{- if .id}} id=\"{{.id}}\" {{end -}}\r\n\t\t\t\t\t\t {{- if .params -}}\r\n\t\t\t\t\t\t {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}}\r\n\t\t\t\t\t\t {{- end -}}\r\n\t\t\t\t\t\t {{- range $v :=.tags}} {{$v}} {{end -}}\r\n\t\t\t\t\t\t {{- if .value -}} value=\"{{form_date .value}}\" {{end}}>\r\n\t\t\t{{- if or .helptext .errors -}}\r\n\t\t\t<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n          {{- if .errors}}\r\n            <ul>\r\n            {{- range .errors }}\r\n              <li>{{.}}</li>\r\n            {{- end -}}\r\n            </ul>\r\n          {{- end}}\r\n    </span>\r\n\t\t\t{{- end}}\r\n\t\t</div>\r\n\t\t</div>\r\n\r\n\t\t<div class=\"col-lg-6\">\r\n\t\t<div class=\"input-group clockpicker\" data-autoclose=\"true\">\r\n\t\t\t<input type=\"text\" class=\"form-control\" value=\"{{ if .value }}{{ form_time .value }}{{ end }}\" name=\"{{.name}}time\">\r\n\t\t\t<span class=\"input-group-addon\">\r\n\t\t\t\t<span class=\"fa fa-clock-o\"></span>\r\n\t\t\t</span>\r\n\t\t</div>\r\n\t\t</div>\r\n\r\n\t\t<input type=\"hidden\" name=\"{{.name}}\" value=\"\">\r\n\t\t<script>\r\n        if (tpt_form_callbacks == null) {\r\n            tpt_form_callbacks = new Array();\r\n        }\r\n        tpt_form_callbacks.push(function () {\r\n            $('#{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}} input').on(\"change\",function () {\r\n                var date = $('[name=\"{{ .name }}date\"]').val();\r\n                var time = $('[name=\"{{ .name }}time\"]').val();\r\n                if (date || time){\r\n                    $('[name=\"{{ .name }}\"]').val(date+\" \"+time);\r\n                }\r\n            });\r\n\r\n            $('#{{if .id}}{{.id}}{{else}}{{.widget_id}}{{end}} .input-group.date').datepicker({\r\n                todayBtn: \"linked\",\r\n                todayHighlight: false,\r\n                keyboardNavigation: true,\r\n                forceParse: true,\r\n                calendarWeeks: true,\r\n                autoclose: true,\r\n                format: \"yyyy-mm-dd\"\r\n            });\r\n            $('.clockpicker').clockpicker();\r\n        });\r\n\t\t</script>\r\n\r\n\t</div>\r\n"),
	}
	filev := &embedded.EmbeddedFile{
		Filename:    `default/datetime/time.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filex := &embedded.EmbeddedFile{
		Filename:    `default/number/number.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end}}\r\n           {{- if .params}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{- end}}\r\n           {{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{- end}}\r\n           {{- range $v :=.tags}} {{$v}}\r\n           {{- end}}\r\n           {{- if .value}} value=\"{{.value}}\"\r\n           {{- end}}>\r\n    {{- if or .helptext .errors }}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n            {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end -}}\r\n              </ul>\r\n            {{- end}}\r\n    </span>\r\n    {{- end}}\r\n  </div>\r\n{{- end}}"),
	}
	filey := &embedded.EmbeddedFile{
		Filename:    `default/number/range.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\" -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end}}\r\n           {{- if .params}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{- end}}\r\n           {{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{- end}}\r\n           {{- range $v :=.tags}} {{$v}}\r\n           {{- end}}\r\n           {{- if .value}} value=\"{{.value}}\"\r\n           {{- end}}>\r\n    {{- if or .helptext .errors }}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n      {{- if .errors}}\r\n      <ul>\r\n        {{- range .errors }}\r\n         <li>{{.}}</li>\r\n         {{- end}}\r\n      </ul>\r\n      {{- end}}\r\n    </span>\r\n    {{- end}}\r\n  </div>\r\n{{- end}}"),
	}
	file10 := &embedded.EmbeddedFile{
		Filename:    `default/options/checkbox.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{-  define \"main\"}}\r\n{{- $p := . }}\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\r\n\t\t<div class=\"checkbox{{if .errors}} has-error{{end}}\">\r\n\t\t\t<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\">\r\n\t\t\t\t<input type=\"checkbox\" name=\"{{.name}}\"{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}\r\n\t\t\t\t{{- if toBoolean .value }} checked {{end}}\r\n\t\t\t\t{{- if .id}} id=\"{{.id}}\"{{end}}\r\n\t\t\t\t{{- if .params}}\r\n\t\t\t\t  {{- range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}\r\n\t\t\t\t{{- end}}\r\n\t\t\t\t{{- if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}\r\n\t\t\t\t{{- range $v := .tags}} {{$v}}{{end}}>\r\n\t\t\t\t{{.label}}\r\n\t\t\t</label>\r\n\t\t\t{{- if or .helptext .errors }}\r\n\t\t\t<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{- end}}\r\n\t\t\t{{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}\r\n\t\t\t</span>\r\n\t\t\t{{- end}}\r\n\t\t</div>\r\n\t</div>\r\n{{- end}}\r\n"),
	}
	file11 := &embedded.EmbeddedFile{
		Filename:    `default/options/radiobutton.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}{{ $p := . }}\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\r\n  {{ range .choices }}\r\n    <label class=\"control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">\r\n      <input type=\"radio\" name=\"{{$p.name}}\"\r\n      {{- if $p.classes }} class=\"{{range $p.classes}}{{.}} {{end}}\"\r\n      {{- end}} value=\"{{.Value}}\" id=\"{{$p.id}}\"\r\n      {{- if $p.params}}\r\n        {{- range $k2, $v2 := $p.params}} {{$k2}}=\"{{$v2}}\"\r\n        {{- end}}\r\n      {{- end}}\r\n      {{- if $p.css}} style=\"{{range $k2, $v2 := .css}}{{$k2}}: {{$v2}}; {{end}}\"{{end}}\r\n      {{- if eq $p.value .Value}} checked{{end}}\r\n      {{- range $v2 := $p.tags}} {{$v2}}{{end}}>\r\n      {{.Label}}\r\n    </label>\r\n  {{- end}}\r\n\t</div>\r\n{{- end}}"),
	}
	file12 := &embedded.EmbeddedFile{
		Filename:    `default/options/select.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n    <select name=\"{{.name}}\" class=\"form-control {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"\r\n            {{- if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end -}}\r\n            {{- if .css}}\r\n            style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n            {{- end -}}\r\n            {{- range $v :=.tags}} {{$v}}{{end}}>\r\n      {{- $p := . -}}\r\n      {{- range $v := .choices -}}\r\n        {{- if $v.Label -}}\r\n        <optgroup label=\"{{$v.Label}}\">\r\n        {{- end -}}\r\n\r\n        {{- range $v.Children -}}\r\n        <option value=\"{{.Value}}\"\r\n                {{- if strIn \"multiple\" $p.tags -}}\r\n                  {{- $id :=.Value -}}\r\n                  {{- range $k2, $p2 :=$p.multValues -}}\r\n                    {{- if eq $k2 $id}}selected{{end -}}\r\n                  {{- end -}}\r\n                {{- else -}}\r\n                  {{- if eq $p.value .Value}} selected{{end -}}\r\n                {{- end}}>{{raw .Label -}}\r\n        </option>\r\n        {{- end -}}\r\n        \r\n        {{- if $v.Label -}}\r\n        </optgroup>\r\n        {{- end -}}\r\n      {{- end -}}\r\n    </select>\r\n    {{- if .postText}}<span>{{.postText}}</span>{{end -}}\r\n\r\n    {{- if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end -}}\r\n  </div>\r\n{{- end }}"),
	}
	file14 := &embedded.EmbeddedFile{
		Filename:    `default/text/passwordinput.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n    <input type=\"password\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}{{if .value}} value=\"{{.value}}\" {{end}}>\r\n    {{- if or .helptext .errors -}}\r\n    <span class=\"help-block\">\r\n      {{- if .helptext}}{{ .helptext }}{{end -}}\r\n      {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\r\n    </span>\r\n    {{- end -}}\r\n{{- end -}}\r\n\r\n"),
	}
	file15 := &embedded.EmbeddedFile{
		Filename:    `default/text/textareainput.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n<textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\r\n  {{- if or .helptext .errors -}}\r\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\r\n  </span>\r\n  {{- end}}\r\n  {{- end}}\r\n"),
	}
	file16 := &embedded.EmbeddedFile{
		Filename:    `default/text/textinput.html`,
		FileModTime: time.Unix(1512704849, 0),
		Content:     string("{{- define \"main\"}}\r\n    <input type=\"text\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end -}}\r\n           {{- if .params -}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}}\r\n           {{- end -}}\r\n           {{- if .css -}}\r\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{- end -}}\r\n           {{- range $v :=.tags}} {{$v -}}\r\n           {{- end -}}\r\n           {{- if .value -}}\r\n           value=\"{{.value}}\"\r\n           {{end}}>\r\n    {{- if or .helptext .errors -}}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end -}}\r\n              </ul>{{end -}}\r\n    </span>\r\n    {{- end}}\r\n{{- end}}\r\n\r\n\r\n"),
	}
	file17 := &embedded.EmbeddedFile{
		Filename:    `fieldset.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("<fieldset{{if .classes }} class=\"{{range $v := .classes}} {{$v}}{{end}}\"{{end}}{{ if .tags}}{{range $v := .tags}} {{$v}}{{end}}{{end}}>\n\t{{range .fields}}{{ .Render }}{{end}}\n</fieldset>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2,  // baseform.html
			fileq,  // bootstrapform.html
			file17, // fieldset.html

		},
	}
	dir3 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // bootstrap3/button.html
			file5, // bootstrap3/cron.html
			filea, // bootstrap3/generic.tmpl
			fileb, // bootstrap3/hidden.html
			filec, // bootstrap3/input.html
			filed, // bootstrap3/mapinput.html
			filel, // bootstrap3/static.html

		},
	}
	dir6 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/datetime`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file7, // bootstrap3/datetime/date.html
			file8, // bootstrap3/datetime/datetime.html
			file9, // bootstrap3/datetime/time.html

		},
	}
	dire := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/number`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filef, // bootstrap3/number/number.html
			fileg, // bootstrap3/number/range.html

		},
	}
	dirh := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/options`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filei, // bootstrap3/options/checkbox.html
			filej, // bootstrap3/options/radiobutton.html
			filek, // bootstrap3/options/select.html

		},
	}
	dirm := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/text`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filen, // bootstrap3/text/passwordinput.html
			fileo, // bootstrap3/text/textareainput.html
			filep, // bootstrap3/text/textinput.html

		},
	}
	dirr := &embedded.EmbeddedDir{
		Filename:   `default`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dirs := &embedded.EmbeddedDir{
		Filename:   `default/datetime`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filet, // default/datetime/date.html
			fileu, // default/datetime/datetime.html
			filev, // default/datetime/time.html

		},
	}
	dirw := &embedded.EmbeddedDir{
		Filename:   `default/number`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filex, // default/number/number.html
			filey, // default/number/range.html

		},
	}
	dirz := &embedded.EmbeddedDir{
		Filename:   `default/options`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file10, // default/options/checkbox.html
			file11, // default/options/radiobutton.html
			file12, // default/options/select.html

		},
	}
	dir13 := &embedded.EmbeddedDir{
		Filename:   `default/text`,
		DirModTime: time.Unix(1512704849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file14, // default/text/passwordinput.html
			file15, // default/text/textareainput.html
			file16, // default/text/textinput.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir3, // bootstrap3
		dirr, // default

	}
	dir3.ChildDirs = []*embedded.EmbeddedDir{
		dir6, // bootstrap3/datetime
		dire, // bootstrap3/number
		dirh, // bootstrap3/options
		dirm, // bootstrap3/text

	}
	dir6.ChildDirs = []*embedded.EmbeddedDir{}
	dire.ChildDirs = []*embedded.EmbeddedDir{}
	dirh.ChildDirs = []*embedded.EmbeddedDir{}
	dirm.ChildDirs = []*embedded.EmbeddedDir{}
	dirr.ChildDirs = []*embedded.EmbeddedDir{
		dirs,  // default/datetime
		dirw,  // default/number
		dirz,  // default/options
		dir13, // default/text

	}
	dirs.ChildDirs = []*embedded.EmbeddedDir{}
	dirw.ChildDirs = []*embedded.EmbeddedDir{}
	dirz.ChildDirs = []*embedded.EmbeddedDir{}
	dir13.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1512704849, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":                    dir1,
			"bootstrap3":          dir3,
			"bootstrap3/datetime": dir6,
			"bootstrap3/number":   dire,
			"bootstrap3/options":  dirh,
			"bootstrap3/text":     dirm,
			"default":             dirr,
			"default/datetime":    dirs,
			"default/number":      dirw,
			"default/options":     dirz,
			"default/text":        dir13,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"baseform.html":                       file2,
			"bootstrap3/button.html":              file4,
			"bootstrap3/cron.html":                file5,
			"bootstrap3/datetime/date.html":       file7,
			"bootstrap3/datetime/datetime.html":   file8,
			"bootstrap3/datetime/time.html":       file9,
			"bootstrap3/generic.tmpl":             filea,
			"bootstrap3/hidden.html":              fileb,
			"bootstrap3/input.html":               filec,
			"bootstrap3/mapinput.html":            filed,
			"bootstrap3/number/number.html":       filef,
			"bootstrap3/number/range.html":        fileg,
			"bootstrap3/options/checkbox.html":    filei,
			"bootstrap3/options/radiobutton.html": filej,
			"bootstrap3/options/select.html":      filek,
			"bootstrap3/static.html":              filel,
			"bootstrap3/text/passwordinput.html":  filen,
			"bootstrap3/text/textareainput.html":  fileo,
			"bootstrap3/text/textinput.html":      filep,
			"bootstrapform.html":                  fileq,
			"default/datetime/date.html":          filet,
			"default/datetime/datetime.html":      fileu,
			"default/datetime/time.html":          filev,
			"default/number/number.html":          filex,
			"default/number/range.html":           filey,
			"default/options/checkbox.html":       file10,
			"default/options/radiobutton.html":    file11,
			"default/options/select.html":         file12,
			"default/text/passwordinput.html":     file14,
			"default/text/textareainput.html":     file15,
			"default/text/textinput.html":         file16,
			"fieldset.html":                       file17,
		},
	})
}
