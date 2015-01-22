package main

import("net/http")
import("log")
import("bytes")
import("encoding/base64")
import("io")

func do(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	switch action {
		case "encode":
			buf := bytes.NewBuffer([]byte(r.FormValue("c")))
			result := base64.StdEncoding.EncodeToString(buf.Bytes())
			i, err := w.Write([]byte(result))
			if err != nil {
				panic(err)
			}
			log.Printf("%d bytes written to client\n", i)
		case "decode":
			rdr := bytes.NewReader([]byte(r.FormValue("c")))
			result := base64.NewDecoder(base64.StdEncoding, rdr)
			i, err := io.Copy(w, result)
			if err != nil {
				panic(err)
			}
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
	  <input type="submit" name="action" value="encode">
	  <input type="submit" name="action" value="decode">
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
