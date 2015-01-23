// install: go install Base64StdinEncode.go
// usage:   cat infile | Base64StdinEncode -type=URL > outfile
package main

import("encoding/base64")
import("log")
import("bytes")
import("flag")
import("os")

var infile string
var enctype string

func init(){
	flag.StringVar(&infile, "in", "", "input file")
	flag.StringVar(&enctype, "type", "Std", "Type of Base 64 encoding to use: 'URL' or 'Std'")
}

func check(err error){
	if err != nil {
		panic(err)
	}
}

func main(){
	flag.Parse()
	//data, err := ioutil.ReadFile(os.Stdin)
	data := bytes.NewBuffer([]byte(""))
	n, err := data.ReadFrom(os.Stdin)
	check(err)
	log.Printf("%d bytes read from stdin\n", n)
	b64String := ""
	if enctype == "URL" {
		b64String = base64.URLEncoding.EncodeToString(data.Bytes())
		log.Printf("Useing 'URL' encoding\n")
	} else {
		b64String = base64.StdEncoding.EncodeToString(data.Bytes())
		log.Printf("Useing 'Std' encoding\n")
	}
	w, err := os.Stdout.Write([]byte(b64String+"\n"))
	log.Printf("%d bytes written to stdout\n", w)
	check(err)
}
