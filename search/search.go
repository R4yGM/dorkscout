// Copyright 2020-21 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"golang.org/x/time/rate"
	"errors"
	"log"
	"net/http"
)
//
// See: https://github.com/rocketlaunchr/google-search#warning-warning
var ErrBlocked = errors.New("google block")

// RateLimit sets a global limit to how many requests to Google Search can be made in a given time interval.
// The default is unlimited (but obviously Google Search will block you temporarily if you do too many
// calls too quickly).
//
// See: https://godoc.org/golang.org/x/time/rate#NewLimiter
var RateLimit = rate.NewLimiter(rate.Inf, 0)
// Result represents a single result from Google Search.
type Result struct {

	// Rank is the order number of the search result.
	Rank int `json:"rank"`

	// URL of result.
	URL string `json:"url"`

	// Title of result.
	Title string `json:"title"`

	// Description of the result.
	Description string `json:"description"`
}

// GoogleDomains represents localized Google homepages. The 2 letter country code is based on ISO 3166-1 alpha-2.
//
// See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
var GoogleDomains = map[string]string{
	"us":  "https://www.google.com/search?q=",
	"ac":  "https://www.google.ac/search?q=",
	"ad":  "https://www.google.ad/search?q=",
	"ae":  "https://www.google.ae/search?q=",
	"af":  "https://www.google.com.af/search?q=",
	"ag":  "https://www.google.com.ag/search?q=",
	"ai":  "https://www.google.com.ai/search?q=",
	"al":  "https://www.google.al/search?q=",
	"am":  "https://www.google.am/search?q=",
	"ao":  "https://www.google.co.ao/search?q=",
	"ar":  "https://www.google.com.ar/search?q=",
	"as":  "https://www.google.as/search?q=",
	"at":  "https://www.google.at/search?q=",
	"au":  "https://www.google.com.au/search?q=",
	"az":  "https://www.google.az/search?q=",
	"ba":  "https://www.google.ba/search?q=",
	"bd":  "https://www.google.com.bd/search?q=",
	"be":  "https://www.google.be/search?q=",
	"bf":  "https://www.google.bf/search?q=",
	"bg":  "https://www.google.bg/search?q=",
	"bh":  "https://www.google.com.bh/search?q=",
	"bi":  "https://www.google.bi/search?q=",
	"bj":  "https://www.google.bj/search?q=",
	"bn":  "https://www.google.com.bn/search?q=",
	"bo":  "https://www.google.com.bo/search?q=",
	"br":  "https://www.google.com.br/search?q=",
	"bs":  "https://www.google.bs/search?q=",
	"bt":  "https://www.google.bt/search?q=",
	"bw":  "https://www.google.co.bw/search?q=",
	"by":  "https://www.google.by/search?q=",
	"bz":  "https://www.google.com.bz/search?q=",
	"ca":  "https://www.google.ca/search?q=",
	"kh":  "https://www.google.com.kh/search?q=",
	"cc":  "https://www.google.cc/search?q=",
	"cd":  "https://www.google.cd/search?q=",
	"cf":  "https://www.google.cf/search?q=",
	"cat": "https://www.google.cat/search?q=",
	"cg":  "https://www.google.cg/search?q=",
	"ch":  "https://www.google.ch/search?q=",
	"ci":  "https://www.google.ci/search?q=",
	"ck":  "https://www.google.co.ck/search?q=",
	"cl":  "https://www.google.cl/search?q=",
	"cm":  "https://www.google.cm/search?q=",
	"cn":  "https://www.google.cn/search?q=",
	"co":  "https://www.google.com.co/search?q=",
	"cr":  "https://www.google.co.cr/search?q=",
	"cu":  "https://www.google.com.cu/search?q=",
	"cv":  "https://www.google.cv/search?q=",
	"cy":  "https://www.google.com.cy/search?q=",
	"cz":  "https://www.google.cz/search?q=",
	"de":  "https://www.google.de/search?q=",
	"dj":  "https://www.google.dj/search?q=",
	"dk":  "https://www.google.dk/search?q=",
	"dm":  "https://www.google.dm/search?q=",
	"do":  "https://www.google.com.do/search?q=",
	"dz":  "https://www.google.dz/search?q=",
	"ec":  "https://www.google.com.ec/search?q=",
	"ee":  "https://www.google.ee/search?q=",
	"eg":  "https://www.google.com.eg/search?q=",
	"es":  "https://www.google.es/search?q=",
	"et":  "https://www.google.com.et/search?q=",
	"fi":  "https://www.google.fi/search?q=",
	"fj":  "https://www.google.com.fj/search?q=",
	"fm":  "https://www.google.fm/search?q=",
	"fr":  "https://www.google.fr/search?q=",
	"ga":  "https://www.google.ga/search?q=",
	"gb":  "https://www.google.co.uk/search?q=",
	"ge":  "https://www.google.ge/search?q=",
	"gf":  "https://www.google.gf/search?q=",
	"gg":  "https://www.google.gg/search?q=",
	"gh":  "https://www.google.com.gh/search?q=",
	"gi":  "https://www.google.com.gi/search?q=",
	"gl":  "https://www.google.gl/search?q=",
	"gm":  "https://www.google.gm/search?q=",
	"gp":  "https://www.google.gp/search?q=",
	"gr":  "https://www.google.gr/search?q=",
	"gt":  "https://www.google.com.gt/search?q=",
	"gy":  "https://www.google.gy/search?q=",
	"hk":  "https://www.google.com.hk/search?q=",
	"hn":  "https://www.google.hn/search?q=",
	"hr":  "https://www.google.hr/search?q=",
	"ht":  "https://www.google.ht/search?q=",
	"hu":  "https://www.google.hu/search?q=",
	"id":  "https://www.google.co.id/search?q=",
	"iq":  "https://www.google.iq/search?q=",
	"ie":  "https://www.google.ie/search?q=",
	"il":  "https://www.google.co.il/search?q=",
	"im":  "https://www.google.im/search?q=",
	"in":  "https://www.google.co.in/search?q=",
	"io":  "https://www.google.io/search?q=",
	"is":  "https://www.google.is/search?q=",
	"it":  "https://www.google.it/search?q=",
	"je":  "https://www.google.je/search?q=",
	"jm":  "https://www.google.com.jm/search?q=",
	"jo":  "https://www.google.jo/search?q=",
	"jp":  "https://www.google.co.jp/search?q=",
	"ke":  "https://www.google.co.ke/search?q=",
	"ki":  "https://www.google.ki/search?q=",
	"kg":  "https://www.google.kg/search?q=",
	"kr":  "https://www.google.co.kr/search?q=",
	"kw":  "https://www.google.com.kw/search?q=",
	"kz":  "https://www.google.kz/search?q=",
	"la":  "https://www.google.la/search?q=",
	"lb":  "https://www.google.com.lb/search?q=",
	"lc":  "https://www.google.com.lc/search?q=",
	"li":  "https://www.google.li/search?q=",
	"lk":  "https://www.google.lk/search?q=",
	"ls":  "https://www.google.co.ls/search?q=",
	"lt":  "https://www.google.lt/search?q=",
	"lu":  "https://www.google.lu/search?q=",
	"lv":  "https://www.google.lv/search?q=",
	"ly":  "https://www.google.com.ly/search?q=",
	"ma":  "https://www.google.co.ma/search?q=",
	"md":  "https://www.google.md/search?q=",
	"me":  "https://www.google.me/search?q=",
	"mg":  "https://www.google.mg/search?q=",
	"mk":  "https://www.google.mk/search?q=",
	"ml":  "https://www.google.ml/search?q=",
	"mm":  "https://www.google.com.mm/search?q=",
	"mn":  "https://www.google.mn/search?q=",
	"ms":  "https://www.google.ms/search?q=",
	"mt":  "https://www.google.com.mt/search?q=",
	"mu":  "https://www.google.mu/search?q=",
	"mv":  "https://www.google.mv/search?q=",
	"mw":  "https://www.google.mw/search?q=",
	"mx":  "https://www.google.com.mx/search?q=",
	"my":  "https://www.google.com.my/search?q=",
	"mz":  "https://www.google.co.mz/search?q=",
	"na":  "https://www.google.com.na/search?q=",
	"ne":  "https://www.google.ne/search?q=",
	"nf":  "https://www.google.com.nf/search?q=",
	"ng":  "https://www.google.com.ng/search?q=",
	"ni":  "https://www.google.com.ni/search?q=",
	"nl":  "https://www.google.nl/search?q=",
	"no":  "https://www.google.no/search?q=",
	"np":  "https://www.google.com.np/search?q=",
	"nr":  "https://www.google.nr/search?q=",
	"nu":  "https://www.google.nu/search?q=",
	"nz":  "https://www.google.co.nz/search?q=",
	"om":  "https://www.google.com.om/search?q=",
	"pa":  "https://www.google.com.pa/search?q=",
	"pe":  "https://www.google.com.pe/search?q=",
	"ph":  "https://www.google.com.ph/search?q=",
	"pk":  "https://www.google.com.pk/search?q=",
	"pl":  "https://www.google.pl/search?q=",
	"pg":  "https://www.google.com.pg/search?q=",
	"pn":  "https://www.google.pn/search?q=",
	"pr":  "https://www.google.com.pr/search?q=",
	"ps":  "https://www.google.ps/search?q=",
	"pt":  "https://www.google.pt/search?q=",
	"py":  "https://www.google.com.py/search?q=",
	"qa":  "https://www.google.com.qa/search?q=",
	"ro":  "https://www.google.ro/search?q=",
	"rs":  "https://www.google.rs/search?q=",
	"ru":  "https://www.google.ru/search?q=",
	"rw":  "https://www.google.rw/search?q=",
	"sa":  "https://www.google.com.sa/search?q=",
	"sb":  "https://www.google.com.sb/search?q=",
	"sc":  "https://www.google.sc/search?q=",
	"se":  "https://www.google.se/search?q=",
	"sg":  "https://www.google.com.sg/search?q=",
	"sh":  "https://www.google.sh/search?q=",
	"si":  "https://www.google.si/search?q=",
	"sk":  "https://www.google.sk/search?q=",
	"sl":  "https://www.google.com.sl/search?q=",
	"sn":  "https://www.google.sn/search?q=",
	"sm":  "https://www.google.sm/search?q=",
	"so":  "https://www.google.so/search?q=",
	"st":  "https://www.google.st/search?q=",
	"sv":  "https://www.google.com.sv/search?q=",
	"td":  "https://www.google.td/search?q=",
	"tg":  "https://www.google.tg/search?q=",
	"th":  "https://www.google.co.th/search?q=",
	"tj":  "https://www.google.com.tj/search?q=",
	"tk":  "https://www.google.tk/search?q=",
	"tl":  "https://www.google.tl/search?q=",
	"tm":  "https://www.google.tm/search?q=",
	"to":  "https://www.google.to/search?q=",
	"tn":  "https://www.google.tn/search?q=",
	"tr":  "https://www.google.com.tr/search?q=",
	"tt":  "https://www.google.tt/search?q=",
	"tw":  "https://www.google.com.tw/search?q=",
	"tz":  "https://www.google.co.tz/search?q=",
	"ua":  "https://www.google.com.ua/search?q=",
	"ug":  "https://www.google.co.ug/search?q=",
	"uk":  "https://www.google.co.uk/search?q=",
	"uy":  "https://www.google.com.uy/search?q=",
	"uz":  "https://www.google.co.uz/search?q=",
	"vc":  "https://www.google.com.vc/search?q=",
	"ve":  "https://www.google.co.ve/search?q=",
	"vg":  "https://www.google.vg/search?q=",
	"vi":  "https://www.google.co.vi/search?q=",
	"vn":  "https://www.google.com.vn/search?q=",
	"vu":  "https://www.google.vu/search?q=",
	"ws":  "https://www.google.ws/search?q=",
	"za":  "https://www.google.co.za/search?q=",
	"zm":  "https://www.google.co.zm/search?q=",
	"zw":  "https://www.google.co.zw/search?q=",
}

// SearchOptions modifies how the Search function behaves.
type SearchOptions struct {

	// CountryCode sets the ISO 3166-1 alpha-2 code of the localized Google Search homepage to use.
	// The default is "us", which will return results from https://www.google.com.
	CountryCode string

	// LanguageCode sets the language code.
	// Default: en
	LanguageCode string

	// Limit sets how many results to fetch (at maximum).
	Limit int

	// Start sets from what rank the new result set should return.
	Start int

	// UserAgent sets the UserAgent of the http request.
	// Default: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	UserAgent string

	// OverLimit searches for more results than that specified by Limit.
	// It then reduces the returned results to match Limit.
	OverLimit bool

	// ProxyAddr sets a proxy address to avoid IP blocking.
	ProxyAddr string
}

// Search returns a list of search results from Google.
func Search(ctx context.Context, searchTerm string, opts ...SearchOptions) ([]Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if err := RateLimit.Wait(ctx); err != nil {
		return nil, err
	}


	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.WithTransport(&http.Transport{ 
		DisableKeepAlives: true, 
	})
	
	if len(opts) == 0 {
		opts = append(opts, SearchOptions{})
	}

	if opts[0].UserAgent == "" {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	} else {
		c.UserAgent = opts[0].UserAgent
	}

	var lc string
	if opts[0].LanguageCode == "" {
		lc = "en"
	} else {
		lc = opts[0].LanguageCode
	}

	results := []Result{}
	var rErr error
	rank := 1

	c.OnRequest(func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			r.Abort()
			rErr = err
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		
		rErr = err
	})

	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		sel := e.DOM

		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("div > div > a > h3").Text())

		descText := strings.TrimSpace(sel.Find("div > div > div > span > span").Text())

		if linkText != "" && linkText != "#" {
			result := Result{
				Rank:        rank,
				URL:         linkText,
				Title:       titleText,
				Description: descText,
			}
			results = append(results, result)
			rank += 1
		}
	})

/*	c.OnResponse(func(r *colly.Response) {
		log.Println("OnResponse")
		log.Println("r.Request.ProxyURL", r.Request.ProxyURL) 
		// log.Println("OnResponse Visited", r.Request.URL)

		//log.Println(string(r.Body[:]))
		proxy := r.Ctx.Get("proxy")// alway is same one, the Last proxy
		fmt.Println("OnResponse proxy:", proxy)
		fmt.Println("------------")
	})*/
	limit := opts[0].Limit
	if opts[0].OverLimit {
		limit = int(float64(opts[0].Limit) * 1.5)
	}

	url := url(searchTerm, opts[0].CountryCode, lc, limit, opts[0].Start)

	if opts[0].ProxyAddr != "" {
		rp, err := proxy.RoundRobinProxySwitcher(opts[0].ProxyAddr)
		if err != nil {
			return nil, err
		}
		c.SetProxyFunc(rp)
	}

	c.Visit(url)

	if rErr != nil {
		if strings.Contains(rErr.Error(), "Too Many Requests") {
			return nil, ErrBlocked
		}
		return nil, rErr
	}

	// Reduce results to max limit
	if opts[0].Limit != 0 && len(results) > opts[0].Limit {
		return results[:opts[0].Limit], nil
	}

	return results, nil
}

func url(searchTerm string, countryCode string, languageCode string, limit int, start int) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	countryCode = strings.ToLower(countryCode)

	var url string

	if googleBase, found := GoogleDomains[countryCode]; found {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", googleBase, searchTerm, languageCode)
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", googleBase, searchTerm, languageCode, start)
		}
	} else {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", GoogleDomains["us"], searchTerm, languageCode)
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", GoogleDomains["us"], searchTerm, languageCode, start)
		}
	}

	if limit != 0 {
		url = fmt.Sprintf("%s&num=%d", url, limit)
	}

	return url
}

func Test(){
	url := "https://httpbin.org/ip"
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
	)
	c.UserAgent = "curl/7.54.0"
	c.WithTransport(&http.Transport{ 
		DisableKeepAlives: true, 
	})

	c.OnRequest(func(r *colly.Request) {
		proxy := r.Ctx.Get("proxy")
		c.SetProxy(proxy) //Not working when colly.Async(true),
		log.Println("OnRequest proxy:", proxy)
	})
	c.OnResponse(func(r *colly.Response) {
		log.Println("OnResponse")
		log.Println("r.Request.ProxyURL", r.Request.ProxyURL) 
		// log.Println("OnResponse Visited", r.Request.URL)

		log.Println(string(r.Body[:]))
		proxy := r.Ctx.Get("proxy")// alway is same one, the Last proxy
		fmt.Println("OnResponse proxy:", proxy)
		fmt.Println("------------")
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("OnError ", r.StatusCode, err)
		proxy := r.Ctx.Get("proxy")
		fmt.Println("OnError proxy:", proxy)

		fmt.Println("------------")
	})
	proxy_list := []string{"socks5://127.0.0.1:9050"}
	for idx, proxy := range proxy_list {
		fmt.Println(idx, proxy)
		var ctx = colly.NewContext()
		ctx.Put("proxy", proxy)
		c.Request("GET", url, nil, ctx, nil) 
	}
	c.Wait()
}