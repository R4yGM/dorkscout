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

        tmpl, _ := template.ParseFiles("results/template.tpl")

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