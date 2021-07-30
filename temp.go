package main



//sudo curl 'https://www.exploit-db.com/google-hacking-database?draw=1&columns%5B0%5D%5Bdata%5D=date&columns%5B0%5D%5Bname%5D=date&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=url_title&columns%5B1%5D%5Bname%5D=url_title&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=false&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=cat_id&columns%5B2%5D%5Bname%5D=cat_id&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=false&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=author_id&columns%5B3%5D%5Bname%5D=author_id&columns%5B3%5D%5Bsearchable%5D=false&columns%5B3%5D%5Borderable%5D=false&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=desc&start=0&length=520&search%5Bvalue%5D=&search%5Bregex%5D=false&author=&category=&_=1627574522843' -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: it-IT,it;q=0.8,en-US;q=0.5,en;q=0.3' --compressed -H 'Referer: https://www.exploit-db.com/google-hacking-database' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Cookie: XSRF-TOKEN=eyJpdiI6IlVzMnFXOGlWVkl0dG9EVTZ3S2dERlE9PSIsInZhbHVlIjoiUG9NczR5MTRaOFBrUUZBUm5mN3dHZGthdVJmbTFiQXcwSWQzb2szellJUEU0aFBqbGhKVzV5WEpqaTZxRm1mciIsIm1hYyI6IjIzNTk5ZGQ3NGY2ZDI1ZWM2MTZiMzNiY2I3ZDkwZmNhODYxM2M0MGQwZTBjYmJhOTE0OTljMTBhODMyNTcxNzEifQ%3D%3D; exploit_database_session=eyJpdiI6Ik9OR0trNmRIWGk5OXpMUmczaXFcL3lRPT0iLCJ2YWx1ZSI6InJZRTVESkh2RWVkZGJwbk1XYlwvc05HZmlkNUR3Q2VIMnVlT3lva2xoT3F4Z1BBS0hLdmVFaHRkSFpISndqSWlyIiwibWFjIjoiMmJjYWRkYTA4MzY5ZjI0ZDZmZjFmZGI0YTMzOTJhYjMzMzE1YWZhYTRkNmJhOGI2ZjkwZGIzNTRlZjFjOGY1MyJ9; CookieConsent={stamp:%279W9ylgZmitjrUHYjJwUViYiY6PVGxc/c9i7spqEr8QOFxDz9SrXiDQ==%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1627564471793%2Cregion:%27it%27}; _ga=GA1.3.1314249609.1627564471; _gid=GA1.3.1617558237.1627564471' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'TE: trailers' -o a.json



//sudo curl 'https://www.exploit-db.com/google-hacking-database?draw=1&columns%5B0%5D%5Bdata%5D=date&columns%5B0%5D%5Bname%5D=date&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=url_title&columns%5B1%5D%5Bname%5D=url_title&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=false&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=cat_id&columns%5B2%5D%5Bname%5D=cat_id&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=false&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=author_id&columns%5B3%5D%5Bname%5D=author_id&columns%5B3%5D%5Bsearchable%5D=false&columns%5B3%5D%5Borderable%5D=false&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=desc&start=0&length=6550&search%5Bvalue%5D=&search%5Bregex%5D=false&author=&category=&_=1627574522843' -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Language: it-IT,it;q=0.8,en-US;q=0.5,en;q=0.3' --compressed -H 'Referer: https://www.exploit-db.com/google-hacking-database' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Cookie: XSRF-TOKEN=eyJpdiI6IlVzMnFXOGlWVkl0dG9EVTZ3S2dERlE9PSIsInZhbHVlIjoiUG9NczR5MTRaOFBrUUZBUm5mN3dHZGthdVJmbTFiQXcwSWQzb2szellJUEU0aFBqbGhKVzV5WEpqaTZxRm1mciIsIm1hYyI6IjIzNTk5ZGQ3NGY2ZDI1ZWM2MTZiMzNiY2I3ZDkwZmNhODYxM2M0MGQwZTBjYmJhOTE0OTljMTBhODMyNTcxNzEifQ%3D%3D; exploit_database_session=eyJpdiI6Ik9OR0trNmRIWGk5OXpMUmczaXFcL3lRPT0iLCJ2YWx1ZSI6InJZRTVESkh2RWVkZGJwbk1XYlwvc05HZmlkNUR3Q2VIMnVlT3lva2xoT3F4Z1BBS0hLdmVFaHRkSFpISndqSWlyIiwibWFjIjoiMmJjYWRkYTA4MzY5ZjI0ZDZmZjFmZGI0YTMzOTJhYjMzMzE1YWZhYTRkNmJhOGI2ZjkwZGIzNTRlZjFjOGY1MyJ9; CookieConsent={stamp:%279W9ylgZmitjrUHYjJwUViYiY6PVGxc/c9i7spqEr8QOFxDz9SrXiDQ==%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1627564471793%2Cregion:%27it%27}; _ga=GA1.3.1314249609.1627564471; _gid=GA1.3.1617558237.1627564471' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'TE: trailers' -o a.json




/*categories = {
    1: "Footholds",
    2: "File Containing Usernames",
    3: "Sensitives Directories",
    4: "Web Server Detection",
    5: "Vulnerable Files",
    6: "Vulnerable Servers",
    7: "Error Messages",
    8: "File Containing Juicy Info",
    9: "File Containing Passwords",
    10: "Sensitive Online Shopping Info",
    11: "Network or Vulnerability Data",
    12: "Pages Containing Login Portals",
    13: "Various Online devices",
    14: "Advisories and Vulnerabilities",
}*/


import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"github.com/buger/jsonparser"
	"strings"
	"os"
	"github.com/spf13/cobra"
)

func main(){ 

	//install_dorks()
}


func install_dorks(){

	req, err := http.NewRequest("GET", "https://www.exploit-db.com/google-hacking-database?draw=1&columns%5B0%5D%5Bdata%5D=date&columns%5B0%5D%5Bname%5D=date&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=url_title&columns%5B1%5D%5Bname%5D=url_title&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=false&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=cat_id&columns%5B2%5D%5Bname%5D=cat_id&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=false&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=author_id&columns%5B3%5D%5Bname%5D=author_id&columns%5B3%5D%5Bsearchable%5D=false&columns%5B3%5D%5Borderable%5D=false&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=0&order%5B0%5D%5Bdir%5D=desc&start=0&length=6550&search%5Bvalue%5D=&search%5Bregex%5D=false&author=&category=&_=1627574522843", nil)
	if err != nil {
		// handle err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "it-IT,it;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Referer", "https://www.exploit-db.com/google-hacking-database")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "XSRF-TOKEN=eyJpdiI6IlVzMnFXOGlWVkl0dG9EVTZ3S2dERlE9PSIsInZhbHVlIjoiUG9NczR5MTRaOFBrUUZBUm5mN3dHZGthdVJmbTFiQXcwSWQzb2szellJUEU0aFBqbGhKVzV5WEpqaTZxRm1mciIsIm1hYyI6IjIzNTk5ZGQ3NGY2ZDI1ZWM2MTZiMzNiY2I3ZDkwZmNhODYxM2M0MGQwZTBjYmJhOTE0OTljMTBhODMyNTcxNzEifQ%3D%3D; exploit_database_session=eyJpdiI6Ik9OR0trNmRIWGk5OXpMUmczaXFcL3lRPT0iLCJ2YWx1ZSI6InJZRTVESkh2RWVkZGJwbk1XYlwvc05HZmlkNUR3Q2VIMnVlT3lva2xoT3F4Z1BBS0hLdmVFaHRkSFpISndqSWlyIiwibWFjIjoiMmJjYWRkYTA4MzY5ZjI0ZDZmZjFmZGI0YTMzOTJhYjMzMzE1YWZhYTRkNmJhOGI2ZjkwZGIzNTRlZjFjOGY1MyJ9; CookieConsent={stamp:%279W9ylgZmitjrUHYjJwUViYiY6PVGxc/c9i7spqEr8QOFxDz9SrXiDQ==%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1627564471793%2Cregion:%27it%27}; _ga=GA1.3.1314249609.1627564471; _gid=GA1.3.1617558237.1627564471")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
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

		//fmt.Println(strings.Join(v[:], ""))
		fmt.Println(category)

		if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
			// path/to/whatever does not exist
		}else{


		}


	}, "data")
}

