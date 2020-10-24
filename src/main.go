package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// @author valor.

var (
	BaseDir           = "."
	TimeFmt           = "02-Jan-2006 15:04 MST[-07:00]"
	MaxFilenameLength = 50

	HttpPort int64 = 48080

	Version = "0.1.0"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)

func main() {
	stat()
	logger.Printf("gofs running at base dir [%s], listenning at [:%d]", BaseDir, HttpPort)
	run(HttpPort)
}

func help() string {
	return `Gofs - Go File Sharing

Usage:
  gofs [OPTION [argument]...]

Options:
  -h, --help             Show help (this message) and exit.
  -v, --version          Show version information.
  -d, --dir              Specify the base directory for file sharing.
  -f, --fmt              Specify the date format (go time format).
  -l, --len              Specify the longest file name that can be displayed.
  -p, --port             Specify the port of the HTTP server.

`
}

func version() string {
	return fmt.Sprintf("gofs v%s %s %s/%s\n", Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func stat() {
	args := os.Args[1:]
	length := len(args)

	for i, arg := range args {
		// -h, --help
		if strings.Compare(arg, "-h") == 0 || strings.Compare(arg, "--help") == 0 {
			print(help())
			os.Exit(0)
		}
		// -v, --version
		if strings.Compare(arg, "-v") == 0 || strings.Compare(arg, "--version") == 0 {
			print(version())
			os.Exit(0)
		}
		// -d, --dir
		if strings.Compare(arg, "-d") == 0 || strings.Compare(arg, "--dir") == 0 {
			if i >= length {
				panic("No argument for '-d, --dir'. Please use '-h' for details.\n")
			}
			i = i + 1
			BaseDir = args[i]
		}
		// -f, --fmt
		if strings.Compare(arg, "-f") == 0 || strings.Compare(arg, "--fmt") == 0 {
			if i >= length {
				panic("No argument for '-f, --fmt'. Please use '-h' for details.\n")
			}
			i = i + 1
			TimeFmt = args[i]
		}
		// -l, --len
		if strings.Compare(arg, "-l") == 0 || strings.Compare(arg, "--len") == 0 {
			if i >= length {
				panic("No argument for '-l, --len'. Please use '-h' for details.\n")
			}
			i = i + 1

			var err error
			MaxFilenameLength, err = strconv.Atoi(args[i])
			if err != nil {
				panic("Illegal number format for '-l, --len'. Please use '-h' for details.\n")
			}
		}
		// -p, --port
		if strings.Compare(arg, "-p") == 0 || strings.Compare(arg, "--port") == 0 {
			if i >= length {
				panic("No argument for '-p, --port'. Please use '-h' for details.\n")
			}
			i = i + 1

			var err error
			HttpPort, err = strconv.ParseInt(args[i], 10, 0)
			if err != nil {
				panic("Illegal number format for '-p, --port'. Please use '-h' for details.\n")
			}
			if HttpPort <= 1024 || HttpPort > 65535 {
				panic("Illegal number for HTTP Port. Please use '-h' for details.\n")
			}
		}
	}
}
