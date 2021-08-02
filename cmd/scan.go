package cmd

import (
	"fmt"
	"context"
	"github.com/spf13/cobra"
	googlesearch  "dorkscout/search"

)

var ctx = context.Background()

var target string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(target)
		example()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().StringVarP(&target, "target", "t", "" , "")
}

func example() {

	ctx := context.Background()
	//var result []googlesearch.Result
	//result, _ = googlesearch.Search(ctx, url, googlesearch.SearchOptions{Limit: 100, OverLimit: true})
	//fmt.Println(result)
	/*for i := 0; i < len(result); i++ {
        fmt.Println(result[i])
    }*/
	googlesearch.Test()
	for i := 0; i < 4; i++ {
        result, err := googlesearch.Search(ctx, "aa", googlesearch.SearchOptions{Limit: 1, ProxyAddr: "socks5://127.0.0.1:9050" })
		fmt.Println(result)
		fmt.Println(err)
    }
	//log.Println(result[0].URL)
	//log.Println(re.FindAllString(result[0].Title, -1)[0])
}