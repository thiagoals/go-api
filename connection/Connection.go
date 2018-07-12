package connection

import(
    db "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func StablishConnection()(*db.DB, error){
    condb, errdb := db.Open("mysql","root:@/solicitacoes?parseTime=true")
    if errdb != nil {
        log.Fatal(errdb.Error())
    }
    return condb, errdb;
}