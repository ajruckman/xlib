package xlib

const emailContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.subject}}</title>

    <style>
        div {
            font-family: monospace;
        }

        h5 {
            font-family: sans-serif;
            margin: 0;
            padding: 0;
        }

        table {
            border-collapse: collapse;
            font-family: monospace;
        }

        td {
            border: 1px solid black;
            padding: 3px 7px;
        }

        td:first-child {
            white-space: nowrap;
        }
    </style>
</head>

<body class="display">

<h5>Hostname:</h5>
<div>{{.hostname}}</div>
<br>
<h5>Error:</h5>
<div>{{.error}}</div>
<br>
<h5>Stacktrace:</h5>
<div>{{.stacktrace}}</div>
<br>
<h5>Meta:</h5>
<br>

{{if .meta}}
    <table>
        <tbody>
        {{range $k, $v := .meta}}
            <tr>
                <td><div>{{$k}}</div></td>
                <td><div>{{$v}}</div></td>
            </tr>
        {{end}}
        </tbody>
    </table>
{{end}}

</body>
</html>`
