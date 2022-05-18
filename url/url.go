package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://livesms.cisoftccampus.in//OFFICE/SMS_LiveSms/LiveSmsView?LIVESMSID%3D1876"

func main() {
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme, result.Host)
	qparams := result.Query()
	fmt.Println(qparams)
	// '&' represents refrence it takes the exact value rather than a copy of the const
	partsOfURL := &url.URL{
		Scheme:  "http",
		Host:    "livesms.cisoftccampus.in",
		Path:    "/OFFICE/SMS_LiveSms",
		RawPath: "LiveSmsView?LIVESMSID%3D1876",
	}
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
