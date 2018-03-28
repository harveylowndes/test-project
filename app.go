package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Host: %s\n", html.EscapeString(req.Host))
	/*res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<html>
			<head>
			</head>
			<body>
				<h1>Hello World</h1>
				<p>How are you doing?</p>
			</body>
		</html>`)*/
}
