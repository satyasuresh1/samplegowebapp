package main

import(
	"text/template";"os"
	"log"
)

func main(){
	a:=10
	tpl,err := template.ParseFiles("index.html")
	if err!=nil{
		log.Fatal(err)
	}
	nf,err := os.Create("blog.html")
	if err!=nil{
		log.Fatal(err)
	}
	defer nf.Close()
	err = tpl.ExecuteTemplate(nf,"index.html",a)
	if err!=nil{
		log.Fatal(err)
	}
}