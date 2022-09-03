package main

import(
	"fmt"
	"flag"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/fatih/color"
	"sync"
	"bufio"
	"regexp"
	"time"
	"strings"
)

type Regexes struct {
    Regexes []Regex `json:"regexes"`
}

type Regex struct {
   Name string `json:"name"`
   Regex string `json:"regex"`
}

var concurrency int
var silent bool
var outputDir string
var customReg string
var regList string
var allReg bool

func Banner() {
	fmt.Println(`
   ______      __  ___                           
  / ____/___  /  |/  /_  ______ _____ ____  _____
 / / __/ __ \/ /|_/ / / / / __ \/ __ \/ _ \/ ___/
/ /_/ / /_/ / /  / / /_/ / /_/ / /_/ /  __/ /    
\____/\____/_/  /_/\__,_/\__, /\__, /\___/_/     
                        /____//____/                                                                          
`)
	fmt.Println("  https://github.com/viper0x\n")
}


func parseArguments() {
    flag.IntVar(&concurrency, "concurrency", 25, "The concurrency level")
    flag.IntVar(&concurrency, "c", 25, "The concurrency level")

	flag.BoolVar(&silent, "silent", false, "Show results only without printing banner")
	flag.BoolVar(&silent, "s", false, "Show results only without printing banner")

	flag.StringVar(&customReg, "regex", customReg, "Use custom regex instead of using regex.json list")
	flag.StringVar(&customReg, "r", customReg, "Use custom regex instead of using regex.json list")

	flag.StringVar(&regList, "rL", "regex.json", "Use another json regex list instead of regex.json")
	flag.StringVar(&regList, "regex-list", "regex.json", "Use another json regex list instead of regex.json")

	flag.BoolVar(&allReg, "a", false, "Will check for all regexes including regexes named (Credentials Disclosure)")
	flag.BoolVar(&allReg, "all", false, "Will check for all regexes including regexes named (Credentials Disclosure)")


    flag.Parse()
}

func matchContent(name string, regex string, target string, content string) {
	re1, _ := regexp.Compile(regex)
	match := re1.FindString(content)

	if match != "" {
		fmt.Printf("[%s] [%s] %s [%s]\n", color.CyanString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString(name), target, color.YellowString(match))
	}
    
}

func main() {

	flag.Usage = func() {
		h := []string{
			"Usage:",
			"  [stdin] | gomugger [options]\n",
			"Options:",
			"  -c, --concurrency <val>     The concurrency level (default 25)",
			"  -s, --silent                Show results only without printing banner",
			"  -r, --regex <regex>         Use custom regex instead of using regex.json list",
			"  -rL, --regex-list <file>    Use another json regex list instead of regex.json (default regex.json)",
			"  -a, --all <file>            Will check for all regexes including regexes named (Credentials Disclosure)",
			"  -h, --help                  Display help\n",
			"Examples:",
			"  cat targets.txt | gomugger",
			"  echo https://example.com | gomugger",
			"  echo https://example.com | waybackurls | gomugger",
			"  cat targets.txt | gomugger -r <REGEX>",
			"  cat targets.txt | gomugger -rL <FILE.json>\n",
		}
		Banner()
		fmt.Fprintf(os.Stderr, "%s", strings.Join(h, "\n"))
	}

	parseArguments()

	if !silent {
		Banner()
	}

	jobs := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		go func() {

			for target := range jobs {
				response, err := http.Get(target)

				if err != nil {
					continue
				}

				content, _ := ioutil.ReadAll(response.Body)

				if customReg != "" {
				matchContent("Custom Regex", customReg, target, string(content))

				} else {

					file, err := os.Open(regList)

					if err != nil {
						fmt.Println(err)
					}

					defer file.Close()

					byteResult, _ := ioutil.ReadAll(file)

					var regexes Regexes

					json.Unmarshal(byteResult, &regexes)

					for i := 0; i < len(regexes.Regexes); i++ {
						if !allReg {
							if regexes.Regexes[i].Name == "Credentials Disclosure" {
								continue
							}
						}
						matchContent(regexes.Regexes[i].Name, regexes.Regexes[i].Regex, target, string(content))
					}
				}
			}

			wg.Done()
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
                jobs <- sc.Text()
        }

        close(jobs)
		wg.Wait()
}
