package cmd

import (
	"fmt"
	"os"
	"log"
	"github.com/spf13/cobra"
	"path/filepath"
	"io/ioutil"
	"time"
)


var (
		dir string
	
		deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {

			if _, err := os.Stat(dir); os.IsNotExist(err) {

				log.Fatal(fmt.Sprintf("Error : the directory '%s' does not exist ",dir))

			}

			delete_old_dorks()

			now := time.Now()

			ioutil.WriteFile(fmt.Sprintf("%s/.dorkscout",outputDir), []byte(fmt.Sprintf("dorks deletion finished at %d",now.Unix(),int(dorksCount))), 0644)

			log.Println("All dorks list files are deleted")
		},
	}
)

func init() {

	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "" , "")
	deleteCmd.MarkPersistentFlagRequired("dir")

}

func delete_old_dorks() {

	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".dorkscout" {
				if err = os.Remove(dir+"/"+file.Name()); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}