package cmd

import (
	"fmt"
	"context"
	"github.com/spf13/cobra"
	googlesearch  "github.com/rocketlaunchr/google-search"
	"log"
	"regexp"
	"net/url"
)

var ctx = context.Background()


var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
		example()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func example() {
	var re = regexp.MustCompile(`(?m)(.*\))`)
	song := "tachikovsky symphony no.5"
	name := url.QueryEscape(song)
	ctx := context.Background()
	url := "site:imslp.org+" + name
	log.Println(url)
	var result []googlesearch.Result
	result, _ = googlesearch.Search(ctx, url)
	log.Println(result[0].URL)
	log.Println(re.FindAllString(result[0].Title, -1)[0])
}