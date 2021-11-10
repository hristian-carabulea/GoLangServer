package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// https://golang.org
// https://golang.org/doc/code
// to compile for Windows on Windows, open powershell as admin and run: 
//		$Env:GOOS = "windows"; $Env:GOARCH = "amd64"
//
// to compile for Linux on Windows, open powershell as admin and run: 
//		$Env:GOOS = "linux"; $Env:GOARCH = "amd64"
//
// To compile Windows 32-bit, enter in bash before building: export GOOS=windows GOARCH=386.
// To compile Windows 64-bit, enter in bash before building: export GOOS=windows GOARCH=amd64.
// Rename files from xxxx to xxxx.exe, and run them from command line or by double clicking in file manager.

// To compile Linux 32-bit, enter in bash before building: export GOOS=linux GOARCH=386. Then run ./program
// To compile Linux 64-bit, enter in bash before building: export GOOS=linux GOARCH=amd64. Then run ./program

// To compile Mac 32-bit, enter in bash before building: export GOOS=darwin GOARCH=386.
// To compile Mac 64-bit, enter in bash before building: export GOOS=darwin GOARCH=amd64.
// I never tried out the Mac environment.

// main will open a web server for html pages at the port 8889.
// Type in your browser, http://localhost:8889/
// If there is an index.html file in the WebContent directory, it will display automatically.
// Otherwise choose the file or directory, which you wish to open.

var flags *flag.FlagSet

//Separator ...
var Separator = "----------------------------------------"
var schema = "http"
var webroot = "."
var ip = "127.0.0.1"
var port = 8889
var cert = ""
var key = ""

// UsageHTTPServer ...
func UsageHTTPServer() {
	fmt.Println("usage:")
	fmt.Println(Separator)
	fmt.Println("--schema http|https, for example -schema http")
	fmt.Println("--cert /path/to/ssl.crt")
	fmt.Println("--key /path/to/ssl.key")
	fmt.Println("--webroot /path/to/your/http/webroot, for example -webroot ../../../MyWebsites")
	fmt.Println("--ip {ip-address}, for example -ip 127.0.0.1")
	fmt.Println("--port {port-number}, for example -port 8889")
	fmt.Println("")
}

func init() {
	flag.StringVar(&schema, "schema", "http", "schema default is \"http\"")
	flag.StringVar(&cert, "cert", "", "cert default is \"\"")
	flag.StringVar(&key, "key", "", "key default is \"\"")
	
	//flag.StringVar(&webroot, "webroot", ".", "webroot default is \".\"")
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip default is \"127.0.0.1\"")
	flag.IntVar(&port, "port", 8889, "port default is \"8889\"")
}

func main() {

	flag.Parse()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(webroot))))

	path, err := os.Executable()
	if err != nil {
			log.Println(err)
	}
	log.Println("starting the server in webroot <" + path + ">")
	log.Println("starting the server at address <" + ip + ":" + strconv.Itoa(port) + ">")

	if port == 80 {
		fmt.Println("you can call now: " + schema + "://" + ip)
	} else {
		fmt.Println("you can call now: " + schema + "://" + ip + ":" + strconv.Itoa(port))
	}

	switch schema {
	case "http":
		log.Fatal(http.ListenAndServe(ip+":"+strconv.Itoa(port), nil))
	case "https":
		log.Fatal(http.ListenAndServeTLS(ip+":"+strconv.Itoa(port), cert, key, nil))
	default:
		fmt.Println("the given schema <" + schema + "> is not supported now!")
		UsageHTTPServer()
		os.Exit(2)
	}
}
