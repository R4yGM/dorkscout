package cmd

import (
	"fmt"
	"os"
	"log"
	"github.com/spf13/cobra"
	"path/filepath"
	"io/ioutil"
	"time"
	"encoding/json"
	"strings"
	"strconv"
	"math"
)

type Info struct {
	Result  string `json:"result"`
	Timestamp string `json:"timestamp"`
	Payloads   int `json:"payloads"`
}

var (
		dir string
	
		deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "deletes all the .dorkscout files inside a given directory",
		Long: `Removes previous installed .dorkscout files in a directory passed throught the -d or --directory flag`,
		Run: func(cmd *cobra.Command, args []string) {

			if _, err := os.Stat(dir); os.IsNotExist(err) {

				log.Fatal(fmt.Sprintf("Error : the directory '%s' does not exist ",dir))

			}

			delete_old_dorks(dir)

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

func delete_old_dorks(dr string) {
	
	d, err := os.Open(dr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	defer d.Close()

	if _, err := os.Stat(dr + "/.dorkscout"); os.IsNotExist(err) {
		
	}else{
			jsonFile, err := os.Open(dr + "/.dorkscout")
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()
		
			byteValue, _ := ioutil.ReadAll(jsonFile)
	
			var Info Info

			json.Unmarshal(byteValue, &Info)

		if strings.Contains(Info.Result, "without any errors"){
			i, err := strconv.ParseInt(fmt.Sprintf("%v", Info.Timestamp), 10, 64)

			if err != nil {
				panic(err)
			}
			
			log.Println(fmt.Sprintf("Removing the previous installation dating back to %s containing %d payloads", TimeElapsed(time.Now(), time.Unix(i, 0), true), Info.Payloads ))
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".dorkscout" {
				if err = os.Remove(dr+"/"+file.Name()); err != nil {
					log.Fatal(err)
				}
				fmt.Println(fmt.Sprintf("[-] %s/%s",dr,file.Name()))
			}
		}
	}
}

func s(x float64) string {
	if int(x) == 1 {
		return ""
	}
	return "s"
}

func TimeElapsed(now time.Time, then time.Time, full bool) string {
	var parts []string
	var text string

	year2, month2, day2 := now.Date()
	hour2, minute2, second2 := now.Clock()

	year1, month1, day1 := then.Date()
	hour1, minute1, second1 := then.Clock()

	year := math.Abs(float64(int(year2 - year1)))
	month := math.Abs(float64(int(month2 - month1)))
	day := math.Abs(float64(int(day2 - day1)))
	hour := math.Abs(float64(int(hour2 - hour1)))
	minute := math.Abs(float64(int(minute2 - minute1)))
	second := math.Abs(float64(int(second2 - second1)))

	week := math.Floor(day / 7)

	if year > 0 {
		parts = append(parts, strconv.Itoa(int(year))+" year"+s(year))
	}

	if month > 0 {
		parts = append(parts, strconv.Itoa(int(month))+" month"+s(month))
	}

	if week > 0 {
		parts = append(parts, strconv.Itoa(int(week))+" week"+s(week))
	}

	if day > 0 {
		parts = append(parts, strconv.Itoa(int(day))+" day"+s(day))
	}

	if hour > 0 {
		parts = append(parts, strconv.Itoa(int(hour))+" hour"+s(hour))
	}

	if minute > 0 {
		parts = append(parts, strconv.Itoa(int(minute))+" minute"+s(minute))
	}

	if second > 0 {
		parts = append(parts, strconv.Itoa(int(second))+" second"+s(second))
	}

	if now.After(then) {
		text = " ago"
	} else {
		text = " after"
	}

	if len(parts) == 0 {
		return "just now"
	}

	if full {
		return strings.Join(parts, ", ") + text
	}
	return parts[0] + text
}