package main

import("io")
import("log")
import("bytes")
import("strings")
import("net/http")
import("io/ioutil")
import("html/template")
import("encoding/base64")

type Result struct {
	Data	string
	Type	string
	Action	string
}

const formsrc = `<html>
	<head>
	 <title>Base64 Encoding/Decoding</title>
	</head>
	<body>
	 <h1>Base64 Encoding/Decoding</h1>
	 <form action="/do" method="POST">
	  <textarea name="c" rows="20" cols="80">{{if .}}{{.Data}}{{end}}</textarea><br/>
	  Base64 Type: <input type="radio" name="t" value="std" checked> Standard <input type="radio" name="type" value="url"> URL 
	  <input type="submit" name="a" value="encode">
	  <input type="submit" name="a" value="decode">
	  <input type="reset" value="reset">
	 </form>
	</body>
</html>`

var tmpl = template.Must(template.New("form").Parse(formsrc))

func check(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
	}
}

func do(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s %s %s '%s'\n", r.RemoteAddr, r.Method, r.URL, r.Proto, r.Referer(), r.UserAgent())
	action := r.FormValue("a")
    enctype := r.FormValue("t")
    enctype = strings.ToLower(enctype)
	switch action {
		case "encode":
            log.Printf("%s action: %s, type: %s\n", r.RemoteAddr, action, enctype)
			buf := bytes.NewBuffer([]byte(r.FormValue("c")))
			result := ""
			if enctype == "url" {
			    result = base64.URLEncoding.EncodeToString(buf.Bytes())
			} else {
			    result = base64.StdEncoding.EncodeToString(buf.Bytes())
			}
			res := Result{result, r.FormValue("type"), action}
			tmpl.Execute(w, res)
		case "decode":
            log.Printf("%s action: %s, type: %s\n", r.RemoteAddr, action, enctype)
			rdr := bytes.NewReader([]byte(r.FormValue("c")))
			var result io.Reader
			if enctype == "url" {
				result = base64.NewDecoder(base64.URLEncoding, rdr)
			} else {
				result = base64.NewDecoder(base64.StdEncoding, rdr)
			}
			resb, err := ioutil.ReadAll(result)
			if err != nil {
				http.Error(w, err.Error(), 500)
				log.Println(err)
				break
			}
			res := Result{string(resb), r.FormValue("t"), action}
			tmpl.Execute(w, res)
		default:
			tmpl.Execute(w, "")
	}
}

func main() {
	http.HandleFunc("/", do)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
