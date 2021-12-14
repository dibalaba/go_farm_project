package main

import (
	
	"fmt"
	"github.com/dowta_dibalaba/managedb"
	"log"
	"database/sql"
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"

	 

)
//ListAllClient to have them all
func ListAllClient(w http.ResponseWriter, r *http.Request){
    var listOfClient []managedb.Client
    //open connection
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/AgbleHakilima")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()


    query  := "select * from clients"
    fmt.Println(query)

    list, err := db.Query(query)

    if err != nil {
        panic(err.Error())
    }
  
    //
    for list.Next() {

        var client managedb.Client
        
        // for each row, scan the result into our tag composite object
        err = list.Scan(&client.ID, &client.Firstname,&client.Surname,&client.Phonenumber,&client.Address,&client.Debit,&client.Credit)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
				// and then print out the tag's Name attribute
				listOfClient=append(listOfClient,client)		
        log.Printf(client.Surname)
    }

    // be careful deferring Queries if you are using transactions
    defer list.Close()
	json.NewEncoder(w).Encode(listOfClient)
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    
    myRouter.HandleFunc("/allClient", ListAllClient)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":9880", myRouter))
}

func main(){

fmt.Println("hello welcome to AgbleHakilima")

//ListAllClient function


// create a few clients
//var dummyClient =  managedb.Client{9,"adja","toto","36542258","begbegamey",0,0}
//var clientDePiment = managedb.Client{6,"komi","mamie"," 32145889","Adewui",0,0}

//insert the clients into db
//clientDePiment.InsertClient()
//dummyClient.InsertClient()


 
//fmt.Println(ListAllClient())



}