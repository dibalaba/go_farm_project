package managedb

import (

	"fmt"
    "database/sql"
    "log"
   // "github.com/dowta_dibalaba/model"
	
	//need it  don't why
    _ "github.com/go-sql-driver/mysql"
)

//Client for db
type Client struct{

	ID int `json:"id"`
   
	Surname string `json:"surname"`
   
	Firstname string `json:"firstname"`
   
	Phonenumber string `json:"phonenumber"`
   
	Address string `json:"address"`
   
	Credit int `json:"credit"`
   
	Debit int `json:"debit"`
   
   }

   

//InsertClient to connect to db
func (client *Client) InsertClient(){
    
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/AgbleHakilima")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    query := fmt.Sprintf("INSERT IGNORE INTO clients VALUES ( %x, '%s','%s','%x','%s', %x,%x)",client.ID, client.Surname, client.Firstname,client.Phonenumber,client.Address,client.Credit,client.Debit)
    fmt.Println(query)
    insert, err := db.Query(query)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}


//Investor type
type Investor struct{

   ID int
   
   FisrtName string

//reason why you joined
   Reason string

//My perspectives
   Vision string

   Position string

}
//InsertInvestor function to insert
func (investor *Investor) InsertInvestor(){

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/AgbleHakilima")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    query := fmt.Sprintf("INSERT IGNORE INTO clients VALUES ( %x, '%s','%s','%s', %s)",investor.ID, investor.FisrtName,investor.Reason,investor.Vision,investor.Position)
    fmt.Println(query)
    insert, err := db.Query(query)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()

}

func listAllInvestor(){

    //open connection
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/AgbleHakilima")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()


    query := "select * from investors"
    fmt.Println(query)

    list, err := db.Query(query)

    if err != nil {
        panic(err.Error())
    }
  
    //
    for list.Next() {
        var investor Investor
        
        // for each row, scan the result into our tag composite object
        err = list.Scan(&investor.ID, &investor.FisrtName,&investor.Position,&investor.Reason,&investor.Vision)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(investor.FisrtName)
    }

    // be careful deferring Queries if you are using transactions
    defer list.Close()

}