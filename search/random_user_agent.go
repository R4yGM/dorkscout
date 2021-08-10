package search

import (
	"fmt"
	"math/rand"
	"strings"
)

var uaGens = []func() string{
	genFirefoxUA,
	genChromeUA,
	genEdgeUA,
	genOperaUA,
}

var ffVersions = []float32{
	// 2020
	72.0,
	73.0,
	74.0,
	75.0,
	76.0,
	77.0,
	78.0,
	79.0,
	80.0,
	81.0,
	82.0,
	83.0,
	84.0,

	// 2021
	85.0,
	86.0,
	87.0,
}

var chromeVersions = []string{
	// 2020
	"79.0.3945.117",
	"79.0.3945.130",
	"80.0.3987.106",
	"80.0.3987.116",
	"80.0.3987.122",
	"80.0.3987.132",
	"80.0.3987.149",
	"80.0.3987.163",
	"80.0.3987.87",
	"81.0.4044.113",
	"81.0.4044.122",
	"81.0.4044.129",
	"81.0.4044.138",
	"81.0.4044.92",
	"83.0.4103.106",
	"83.0.4103.116",
	"83.0.4103.97",
	"84.0.4147.105",
	"84.0.4147.125",
	"84.0.4147.135",
	"85.0.4183.102",
	"85.0.4183.121",
	"85.0.4183.83",
	"86.0.4240.111",
	"86.0.4240.183",
	"86.0.4240.198",
	"86.0.4240.75",

	// 2021
	"87.0.4280.141",
	"87.0.4280.66",
	"87.0.4280.88",
	"88.0.4324.146",
	"88.0.4324.182",
	"88.0.4324.190",
	"89.0.4389.114",
	"89.0.4389.90",
	"90.0.4430.72",
}

var edgeVersions = []string{
	"79.0.3945.74,79.0.309.43",
	"80.0.3987.87,80.0.361.48",
	"84.0.4147.105,84.0.522.50",
	"89.0.4389.128,89.0.774.77",
	"90.0.4430.72,90.0.818.39",
}

var operaVersions = []string{
	"2.7.62 Version/11.00",
	"2.2.15 Version/10.10",
	"2.9.168 Version/11.50",
	"2.2.15 Version/10.00",
	"2.8.131 Version/11.11",
	"2.5.24 Version/10.54",
}



var osStrings = []string{
	// MacOS - High Sierra
	"Macintosh; Intel Mac OS X 10_13",
	"Macintosh; Intel Mac OS X 10_13_1",
	"Macintosh; Intel Mac OS X 10_13_2",
	"Macintosh; Intel Mac OS X 10_13_3",
	"Macintosh; Intel Mac OS X 10_13_4",
	"Macintosh; Intel Mac OS X 10_13_5",
	"Macintosh; Intel Mac OS X 10_13_6",

	// MacOS - Mojave
	"Macintosh; Intel Mac OS X 10_14",
	"Macintosh; Intel Mac OS X 10_14_1",
	"Macintosh; Intel Mac OS X 10_14_2",
	"Macintosh; Intel Mac OS X 10_14_3",
	"Macintosh; Intel Mac OS X 10_14_4",
	"Macintosh; Intel Mac OS X 10_14_5",
	"Macintosh; Intel Mac OS X 10_14_6",

	// MacOS - Catalina
	"Macintosh; Intel Mac OS X 10_15",
	"Macintosh; Intel Mac OS X 10_15_1",
	"Macintosh; Intel Mac OS X 10_15_2",
	"Macintosh; Intel Mac OS X 10_15_3",
	"Macintosh; Intel Mac OS X 10_15_4",
	"Macintosh; Intel Mac OS X 10_15_5",
	"Macintosh; Intel Mac OS X 10_15_6",
	"Macintosh; Intel Mac OS X 10_15_7",

	// MacOS - Big Sur
	"Macintosh; Intel Mac OS X 11_0",
	"Macintosh; Intel Mac OS X 11_0_1",
	"Macintosh; Intel Mac OS X 11_1",
	"Macintosh; Intel Mac OS X 11_2",
	"Macintosh; Intel Mac OS X 11_2_1",
	"Macintosh; Intel Mac OS X 11_2_2",
	"Macintosh; Intel Mac OS X 11_2_3",

	// Windows
	"Windows NT 10.0; Win64; x64",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",

	// Linux
	"X11; Linux x86_64",
}


func genFirefoxUA() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}

func genEdgeUA() string {
	version := edgeVersions[rand.Intn(len(edgeVersions))]
	chromeVersion := strings.Split(version, ",")[0]
	edgeVersion := strings.Split(version, ",")[1]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Edg/%s", os, chromeVersion, edgeVersion)
}

func genOperaUA() string {
	version := operaVersions[rand.Intn(len(operaVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Opera/9.80 (%s; U; en) Presto/%s", os, version)
}
