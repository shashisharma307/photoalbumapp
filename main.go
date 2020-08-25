package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"photoalbum/routes"
)

const PORT = ":8080"
func main()  {
	fmt.Println()
	fmt.Println(os.Getwd())
	fmt.Println("Listening Port %s\n", PORT)

	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))



}
//
//func init(){
//	if ok:= config.InitDB(); !ok{
//		fmt.Println("error creating connection")
//	}
//
//	fmt.Println("connection established")
//
//	if ok:= config.TestConnection(); ok{
//		fmt.Println(ok)
//	}else{
//		fmt.Println(ok)
//	}
//}


