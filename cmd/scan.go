package cmd

import (
        "bufio"
        "context"
        res "github.com/R4yGM/dorkscout/results"
        googlesearch "github.com/R4yGM/dorkscout/search"
        "fmt"
        "github.com/spf13/cobra"
        "log"
        "math/rand"
        "os"
        "strconv"
        "strings"
        "time"
)

var ctx = context.Background()

var (
        text              []string
        isBlocked         bool
        Limit             int
        proxy             string
        dorklist          string
        target            string
        NormalResultsPath string
        HTMLResultsPath   string

        scanCmd = &cobra.Command{
                Use:   "scan",
                Short: "scans a specific website or all the websites it founds for a list of dorks",
                Long: `makes google searches with dorks in a given list that will then be parsed into different
readable formats such as HTML or .txt, this function contains also support proxy ip rotation to
avoid getting blocked or rate limited by google.
                                `,
                Run: func(cmd *cobra.Command, args []string) {

                        switch true {
                        case NormalResultsPath != "":
                                e := os.Remove(NormalResultsPath)
                                if e != nil {
                                        log.Fatal(e)
                                }
                        case HTMLResultsPath != "":
                                res.InitHTMLResults(HTMLResultsPath, target)
                        }

                        isBlocked = false

                        google_scan()

                },
        }
)

func init() {

        rootCmd.AddCommand(scanCmd)

        scanCmd.PersistentFlags().StringVarP(&proxy, "proxy", "x", "", "HTTP, HTTPS or SOCKS5 proxy to use in the requests")

        scanCmd.PersistentFlags().StringVarP(&dorklist, "dorklist", "d", "", "dorklists path separated by a comma")
        scanCmd.MarkPersistentFlagRequired("dorklist")

        scanCmd.PersistentFlags().IntVarP(&Limit, "Limit", "l", 100, "The limit of results you get from a single dork")

        scanCmd.PersistentFlags().StringVarP(&NormalResultsPath, "Output", "O", "", "file path where you want to save the results in a normal format")
        scanCmd.PersistentFlags().StringVarP(&HTMLResultsPath, "OutputHTML", "H", "", "file path where you want to save the results in a HTML format")

        scanCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "site or website to scan for dorks")

}

func scan(i int) {

        ctx := context.Background()

        if len(text) > i {
                payload := text[i]

                if strings.Contains(payload, ".dorkscout") {
                        if target != "" {
                                fmt.Println(fmt.Sprintf("Started scanning %s with %s\n=====================================", target, payload))
                        } else {
                                fmt.Println(fmt.Sprintf("Started scanning with %s\n=====================================", payload))
                        }
                        switch true {
                        case NormalResultsPath != "":
                                f, err := os.OpenFile(NormalResultsPath,
                                        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                                if err != nil {
                                        log.Println(err)
                                }
                                defer f.Close()
                                if _, err := f.WriteString("Started scanning with " + payload + "\n=====================================\n"); err != nil {
                                        log.Println(err)
                                }

                        case HTMLResultsPath != "":
                                res.HTMLInject(payload, false, "file title", HTMLResultsPath, false)
                        }
                        i = i + 1
                        payload = text[i]
                }

                if target != "" {
                        payload = payload + " site:" + target
                }
                result, err := googlesearch.Search(ctx, payload, googlesearch.SearchOptions{Limit: Limit, ProxyAddr: proxy})
                if len(result) == 0 {

                } else if len(result) > 0 {
                        fmt.Println("Results for : ", payload)

                        switch true {
                        case NormalResultsPath != "":
                                f, err := os.OpenFile(NormalResultsPath,
                                        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                                if err != nil {
                                        log.Println(err)
                                }
                                defer f.Close()
                                if _, err := f.WriteString(payload + "\n"); err != nil {
                                        log.Println(err)
                                }
                        case HTMLResultsPath != "":
                                res.HTMLInject(payload, false, "dork title", HTMLResultsPath, false)
                        }

                        for i, resu := range result {
                                switch true {
                                case NormalResultsPath != "":
                                        f, err := os.OpenFile(NormalResultsPath,
                                                os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                                        if err != nil {
                                                log.Println(err)
                                        }
                                        defer f.Close()
                                        if _, err := f.WriteString(resu.URL + "\n"); err != nil {
                                                log.Println(err)
                                        }
                                case HTMLResultsPath != "":
                                        n := strconv.Itoa(i + 1)
                                        if i == (len(result) - 2) {
                                                res.HTMLInject(resu.URL+","+n, true, "link", HTMLResultsPath, false)
                                        } else if i == 0 {
                                                res.HTMLInject(resu.URL+","+n, false, "link", HTMLResultsPath, true)
                                        } else {
                                                res.HTMLInject(resu.URL+","+n, false, "link", HTMLResultsPath, false)
                                        }
                                }

                                fmt.Println(resu.URL)
                        }
                        fmt.Println("=====================================")
                        if NormalResultsPath != "" {
                                f, err := os.OpenFile(NormalResultsPath,
                                        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                                if err != nil {
                                        log.Println(err)
                                }
                                defer f.Close()
                                if _, err := f.WriteString("=====================================\n"); err != nil {
                                        log.Println(err)
                                }
                        }
                }

                if err != nil {
                        if err.Error() == "google block" {
                                if isBlocked == false {
                                        log.Println("Google started blocking your requests but dorkscout will keep making requests")
                                        isBlocked = true
                                }
                                scan(i)
                                return
                        } else if strings.Contains(err.Error(), "context deadline exceeded") {
                                scan(i)
                                return
                        } else if strings.Contains(err.Error(), "didn't find any results") {

                                fmt.Println(payload, " nothing found")
                                fmt.Println("=====================================")

                        } else if strings.Contains(err.Error(), "proxyconnect") && strings.Contains(err.Error(), "connect: connection refused") {
                                log.Println(fmt.Sprintf("Can't connect to proxy : %s connection refused", proxy))
                                return
                        }
                }
                time.Sleep(time.Duration(rand.Intn(20-4)+4) * time.Second)
                scan(i + 1)
                return

        } else {
                fmt.Println("Finished scanning")
                switch true {
                case NormalResultsPath != "":
                        f, err := os.OpenFile(NormalResultsPath,
                                os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                        if err != nil {
                                log.Println(err)
                        }
                        defer f.Close()
                        if _, err := f.WriteString("Finished scanning\n"); err != nil {
                                log.Println(err)
                        }
                case HTMLResultsPath != "":
                        res.HTMLFinish(HTMLResultsPath)
                }
                return
        }

}

func google_scan() {

        dorks := strings.Split(dorklist, ",")

        for _, file_path := range dorks {
                file, err := os.Open(file_path)

                if err != nil {
                        log.Fatalf("failed to open")

                }

                scanner := bufio.NewScanner(file)

                scanner.Split(bufio.ScanLines)

                text = append(text, file_path)

                for scanner.Scan() {
                        text = append(text, scanner.Text())
                }

                file.Close()
        }

        scan(0)
}