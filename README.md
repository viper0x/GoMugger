# GoMugger

![text-1662149638843](https://user-images.githubusercontent.com/43915187/188277728-fefe46d4-6848-4cce-aaee-115037648b7d.png)

A fast tool written in Golang used to check for sensitive/juicy information within web pages content.
The regexes list will be updated whenever I find/community shares interesting regexes.
There's a very good ready regexes file and most of them has been collected by me from different repositories (big shout-out to the people who created/posted them).

## Installation
Using Go ([Go 1.16+ compiler](https://golang.org/doc/install) should be installed & configured!):
```
$ go get -u github.com/viper0x/gomugger && wget https://raw.githubusercontent.com/viper0x/gomugger/main/regex.json
```
If the above didn't work:
```
$ go install github.com/viper0x/gomugger@latest && wget https://raw.githubusercontent.com/viper0x/gomugger/main/regex.json
```

Or by manual building:
```
$ git clone https://github.com/viper0x/gomugger
$ cd gomugger
$ go build .
$ cat targets.txt | ./gomugger
```
NOTE: make sure you are running the tool in the same directory where `regex.json` exist. Or you can use `-rL <path>` instead.
## Usage
```
$ gomugger -h

   ______      __  ___                           
  / ____/___  /  |/  /_  ______ _____ ____  _____
 / / __/ __ \/ /|_/ / / / / __ \/ __ \/ _ \/ ___/
/ /_/ / /_/ / /  / / /_/ / /_/ / /_/ /  __/ /    
\____/\____/_/  /_/\__,_/\__, /\__, /\___/_/     
                        /____//____/                                                                          

  https://github.com/viper0x

Usage:
  [stdin] | gomugger [options]

Options:
  -c, --concurrency <val>     The concurrency level (default 25)
  -s, --silent                Show results only without printing banner
  -r, --regex <regex>         Use custom regex instead of using regex.json list
  -rL, --regex-list <file>    Use another json regex list instead of regex.json (default regex.json)
  -h, --help                  Display help

Examples:
  cat targets.txt | gomugger
  echo https://example.com | gomugger
  echo https://example.com | waybackurls | gomugger
  cat targets.txt | gomugger -r <REGEX>
  cat targets.txt | gomugger -rL <FILE.json>
```

## Concurrency
You can set the concurrency level with the `-c` flag:
```
$ cat targets.txt | gomugger -c 50
```

## Append/Use different regexes
### Appending new regex:
You only have to edit the `regex.json` file and add your regex using the following format:
```
{
 "name": "<NEW REGEX NAME>",
 "regex": "<REGEX>"
}
```

### Using custom regex against targets:
```
$ cat targets.txt | gomugger -r "REGEX"
```

### Using different regexes list:
Create `.json` file and add regexes using the following format. Then you can use `-rL` followed by the json file name to use the custom created list:
```
{
  "regexes": [
    {
      "name": "<REGEX NAME>",
      "regex": "<REGEX>"
    },
		
    {
      "name": "<REGEX NAME>",
      "regex": "<REGEX>"
    }
  ]
}
```
