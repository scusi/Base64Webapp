//
package main

import("encoding/base64")
import("log")
import("bytes")
import("flag")
import("os")
import("io")

var infile string
var enctype string

func init(){
	flag.StringVar(&infile, "in", "", "input file")
	flag.StringVar(&enctype, "type", "Std", "Type aof base64 encoding 'URL' or 'Std'")
}

func check(err error){
	if err != nil {
		panic(err)
	}
}

func DecoderType(enctype string, data bytes.Buffer)(io.Reader){
	rdr := bytes.NewReader(data.Bytes())
	if enctype == "URL" {
		log.Printf("useing 'URL' encoding\n")
		return base64.NewDecoder(base64.URLEncoding, rdr)
	} else {
		log.Printf("useing 'Std' encoding\n")
		return base64.NewDecoder(base64.StdEncoding, rdr)
	}
}

func main(){
	flag.Parse()
	data := bytes.NewBuffer([]byte(""))
	n, err := data.ReadFrom(os.Stdin)
	check(err)
	log.Printf("%d bytes read from stdin\n", n)
	b64dec := DecoderType(enctype, *data)
	//io.Copy(b64dec, os.Stdin)
	i, err := io.Copy(os.Stdout, b64dec)
	check(err)
	log.Printf("%d bytes written to stdout\n", i)
}
