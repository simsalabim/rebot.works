package main

import (
	"flag"
	"fmt"
	"github.com/simsalabim/stockprice/finance"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var port = flag.String("p", "8000", "port")

func main() {
	flag.Parse()
	fmt.Println("Starting server on http://0.0.0.0:" + *port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+*port, nil))
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		http.ServeFile(writer, request, "public_html/index.html")
		return
	}

	if request.URL.Path == "/PoweredBy_200px-White_HorizText.png" {
		http.ServeFile(writer, request, "public_html/PoweredBy_200px-White_HorizText.png")
		return
	}

	if strings.HasPrefix(request.URL.Path, "/gif/") {
		writer.Write(imageContent(request.URL.Path[4:len(request.URL.Path)], "gif"))
		return
	}

	if strings.HasPrefix(request.URL.Path, "/image/") {
		writer.Write(imageContent(request.URL.Path[4:len(request.URL.Path)], "image"))
		return
	}

	if strings.HasPrefix(request.URL.Path, "/weather") {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(weatherJSON()))
		return
	}

	if strings.HasPrefix(request.URL.Path, "/grpn") {
		writer.Header().Set("Content-Type", "text/plain")
		grpn := finance.FindStockPrice("grpn")
		result := strconv.FormatFloat(grpn, 'G', -1, 64)
		writer.Write([]byte(result))
		return
	}

}
