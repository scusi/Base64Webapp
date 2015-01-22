package main

import("net/http")
import("log")
import("bytes")
import("encoding/base64")
import("io")

func check(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func do(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	switch action {
		case "encode":
			buf := bytes.NewBuffer([]byte(r.FormValue("c")))
			result := ""
			if r.FormValue("type") == "URL" {
			    result = base64.URLEncoding.EncodeToString(buf.Bytes())
				log.Println("URL Encoding choosen")
			} else {
			    result = base64.StdEncoding.EncodeToString(buf.Bytes())
				log.Println("Std Encoding choosen")
			}
			i, err := w.Write([]byte(result))
			check(err, w)
			log.Printf("%d bytes written to client\n", i)
		case "decode":
			rdr := bytes.NewReader([]byte(r.FormValue("c")))
			var result io.Reader
			if r.FormValue("type") == "URL" {
				result = base64.NewDecoder(base64.URLEncoding, rdr)
				log.Println("URL Decoding choosen")
			} else {
				result = base64.NewDecoder(base64.StdEncoding, rdr)
				log.Println("Std Decoding choosen")
			}
			i, err := io.Copy(w, result)
			check(err, w)
			log.Printf("%d bytes written to client\n", i)
	}
}

func form(w http.ResponseWriter, r *http.Request) {
	f := `<html>
	<head>
	 <title>Base64 Encoding/Decoding</title>
	</head>
	<body>
	 <h1>Base64 Encoding/Decoding</h1>
	 <form action="/do" method="POST">
	  <textarea name="c" rows="20" cols="80"></textarea><br/>
	  Type: <input type="radio" name="type" value="Stdandard" checked> Standard | <input type="radio" name="type" value="URL"> URL 
	  <input type="submit" name="action" value="encode">
	  <input type="submit" name="action" value="decode">
	  <input type="reset" value="reset">
	 </form>
	</body>
</html>`
	rdr := bytes.NewReader([]byte(f))
	io.Copy(w, rdr)
}

func main() {
	http.HandleFunc("/", form)
	http.HandleFunc("/do", do)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
