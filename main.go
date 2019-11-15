package main

import (
		"fmt"
		"html/template"
		"net/http"
		)
		
func index(w http.ResponseWriter, r *http.Request) {
			tpl, _ := template.ParseFiles("index.html")
			data := map[string]string{
			"Title": "Go Store :)",
			}
			w.WriteHeader(http.StatusOK)
			tpl.Execute(w,data)
			}
			
func main(){
			http.HandleFunc("/",index)
			fmt.Println("SErver is up and listening to port 8080. Brabo demais")
			http.ListenAndServe(":8080", nil)
			}