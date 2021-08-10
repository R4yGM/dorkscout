package results

import (
        "fmt"
        "os"
        "strings"
        "text/template"
        "time"
)

var (
        timeStart  time.Time
        timeFinish time.Time
)

var tpl = `<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Dorkscout results</title>
  <meta name="description" content="Dorkscout results">
  <meta name="author" content="Dorkscout">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <!--[if lt IE 9]>
<script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
<![endif]-->
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.6.3/css/font-awesome.min.css">
  <style>
    body {
      margin: 1.5em auto;
      font-family: "Helvetica Neue", sans-serif;
      color: #575553;
      max-width: 1000px;
      padding: 0 1em;
    }

    h1 {
      font-size: 1.75em;
      color: #003e54;
      font-weight: bold;
      line-height: 1.4;
      margin: 0 0 0.8em 0;
    }

    .item {
      padding: 1em;
      border-bottom: 1px solid #f0f0f0;
    }

    .dorkscout-item {
      background-color: #f0f0f0;
    }

    p {
      font-size: 1em;
      line-height: 1.6;
      margin: 0 0 0.3em 0;
    }

    p:last-of-type {
      margin: 0;
    }

    span.divider-dot {
      margin: 0 0.3em;
      color: #999794;
    }

    .handle {
      font-weight: bold;
      color: #000000;
    }

    a {
      color: #2d9eb2;
      text-decoration: none;
    }

    a:active,
    a:hover {
      color: #207180;
    }

    img {
      border-radius: 4px;
    }
  </style>
  <meta name="user-style-sheet" content="pdf.css">
</head>
<body>
  <h1 style="color:#000000">Dorkscout results {{.Target}}</h1>
  {{.Time}}
  {{.Next}}
</body>
</html>`

func InitHTMLResults(path string, target string) {

        timeStart = time.Now()

        vars := make(map[string]interface{})

        if target != "" {
                vars["Target"] = "for the website " + target
        } else {
                vars["Target"] = ""
        }
        vars["Next"] = "{{.Next}}"
        vars["Time"] = timeStart.Format(time.RFC850)

        tmpl, _ := template.New("template").Parse(tpl)

        file, _ := os.Create(path)
        defer file.Close()

        tmpl.Execute(file, vars)
}

func HTMLFinish(path string) {

        timeFinish := time.Now()
        vars := make(map[string]interface{})

        vars["Next"] = fmt.Sprintf("Finished scanning at %s, time elapsed %s", timeFinish.Format(time.RFC850), time.Since(timeStart))

        tmpl, _ := template.ParseFiles(path)

        file, _ := os.Create(path)
        defer file.Close()

        tmpl.Execute(file, vars)
}

func HTMLInject(text string, isLast bool, a string, path string, isFirst bool) {

        var result string

        switch true {
        case a == "link":
                if isLast {

                        t := strings.Split(text, ",")

                        result = fmt.Sprintf(`
                                  <div class="item dorkscout-item">
                                        <p>
                                        <i class="fa fa-link" aria-hidden="true"></i>
                                          <a class="handle">%s</a>
                                          <a href="%s">%s</a>

                                        </p>
                                  </div>`, t[1], t[0], t[0])

                        vars := make(map[string]interface{})

                        vars["NextResult"] = result
                        vars["Next"] = "{{.Next}}"

                        tmpl, _ := template.ParseFiles(path)

                        file, _ := os.Create(path)
                        defer file.Close()

                        tmpl.Execute(file, vars)

                } else if isFirst {

                        t := strings.Split(text, ",")

                        result = fmt.Sprintf(`
                                <div class="item">
                                  <div class="item dorkscout-item">
                                        <p>
                                        <i class="fa fa-link" aria-hidden="true"></i>
                                          <a class="handle">%s</a>
                                          <a href="%s">%s</a>

                                        </p>
                                  </div>{{.NextResult}}
                                </div>
                                </br></br>
                                {{.Next}}`, t[1], t[0], t[0])

                        vars := make(map[string]interface{})

                        vars["Next"] = result

                        tmpl, _ := template.ParseFiles(path)

                        file, _ := os.Create(path)
                        defer file.Close()

                        tmpl.Execute(file, vars)

                } else {

                        t := strings.Split(text, ",")

                        result = fmt.Sprintf(`
                                  <div class="item dorkscout-item">
                                        <p>
                                        <i class="fa fa-link" aria-hidden="true"></i>
                                          <a class="handle">%s</a>
                                          <a href="%s">%s</a>

                                        </p>
                                  </div>{{.NextResult}}`, t[1], t[0], t[0])

                        vars := make(map[string]interface{})

                        vars["NextResult"] = result
                        vars["Next"] = "{{.Next}}"

                        tmpl, _ := template.ParseFiles(path)

                        file, _ := os.Create(path)
                        defer file.Close()

                        tmpl.Execute(file, vars)

                }

        case a == "dork title":

                result = fmt.Sprintf(`
                        Results for dork <b>%s</b>
                        </br></br>
                        {{.Next}}`, text)

                vars := make(map[string]interface{})

                vars["Next"] = result

                tmpl, _ := template.ParseFiles(path)

                file, _ := os.Create(path)
                defer file.Close()

                tmpl.Execute(file, vars)

        case a == "file title":

                result = fmt.Sprintf(`
                        Started scanning with : <b>%s</b>
                        </br></br>
                        {{.Next}}`, text)

                vars := make(map[string]interface{})

                vars["Next"] = result

                tmpl, _ := template.ParseFiles(path)

                file, _ := os.Create(path)
                defer file.Close()

                tmpl.Execute(file, vars)

        case a == "nothing found":

                result = fmt.Sprintf(`
                        No results found for dork <b>%s</b>
                        </br></br>
                        {{.Next}}`, text)

                vars := make(map[string]interface{})

                vars["Next"] = result

                tmpl, _ := template.ParseFiles(path)

                file, _ := os.Create(path)
                defer file.Close()

                tmpl.Execute(file, vars)

        }

}