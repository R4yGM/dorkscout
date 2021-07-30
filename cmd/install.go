package cmd

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
	"github.com/buger/jsonparser"
	"strings"
	"os"
	"github.com/spf13/cobra"
	"time"
)

var (

	outputDir     string
	dorksCount	  int64

	installCmd = &cobra.Command{
		Use:   "install",
		Short: "installs a list of dorks from exploit-db.com",
		Long: `This command fetches and saves a list of dorks based on their category on different files,
these payloads are downloaded from exploit-db.com where then this lists will be saved in a directory that is going
to be passed with the -o or --output-dir flag, and now this lists can be used to start scanning with dorkscout`,
		Run: func(cmd *cobra.Command, args []string) {

			if _, err := os.Stat(outputDir); os.IsNotExist(err) {

				log.Fatal(fmt.Sprintf("Error : the directory '%s' does not exist ",outputDir))

			}

			delete_old_dorks()

			install_dorks()

			now := time.Now()

			ioutil.WriteFile(fmt.Sprintf("%s/.dorkscout",outputDir), []byte(fmt.Sprintf("dorks installation finished at %d with %d payloads",now.Unix(),int(dorksCount))), 0644)

			log.Println("Installation finished")
		},
	}
)

func init() {

	rootCmd.AddCommand(installCmd)

	installCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "" , "")
	installCmd.MarkPersistentFlagRequired("output-dir")

}

func install_dorks(){

	req, err := http.NewRequest("GET", "https://www.exploit-db.com/google-hacking-database?draw=1&columns%5B0%5D%5Bdata%5D=date&columns%5B0%5D%5Bname%5D=date&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=url_title&columns%5B1%5D%5Bname%5D=url_title&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=false&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=cat_id&columns%5B2%5D%5Bname%5D=cat_id&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=false&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=author_id&columns%5B3%5D%5Bname%5D=author_id&columns%5B3%5D%5Bsearchable%5D=false&columns%5B3%5D%5Borderable%5D=false&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=desc&start=0&length=6550&search%5Bvalue%5D=&search%5Bregex%5D=false&author=&category=&_=1627574522843", nil)

	if err != nil {
		// handle err
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "it-IT,it;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Referer", "https://www.exploit-db.com/google-hacking-database")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

	dorksCount, err = jsonparser.GetInt(body, "recordsTotal")

	re := strings.NewReplacer("<nil>", "", "</a>","")
	jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		payload, err := jsonparser.GetString(value, "url_title")
		category, err := jsonparser.GetString(value, "category","cat_title")

		if err != nil {

			log.Fatal(err)
		}
		
		v := []string{}
		v = strings.Split(re.Replace(payload), ">")
		v[0] = v[len(v)-1]
		v[len(v)-1] = "" 
		v = v[:len(v)-1]  


		file, err := os.OpenFile(fmt.Sprintf("%s/%s.dorkscout",outputDir,category), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			ioutil.WriteFile(fmt.Sprintf("%s/%s.dorkscout",outputDir,category), []byte(strings.Join(v[:], "")), 0644)
		}
		defer file.Close()
		if _, err := file.WriteString(fmt.Sprintf("%s\n",strings.Join(v[:], ""))); err != nil {
			//log.Fatal(err)
		}


	}, "data")
}
