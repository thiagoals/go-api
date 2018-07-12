package main

import(
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    models "APICeuma/models"
    con "APICeuma/connection"
    responseEntity "APICeuma/response"
)

func main() {
    /* Connecting with mysql */
    router := mux.NewRouter()
    router.HandleFunc("/solicitacao", GetSolicitacao).Methods("GET")
    router.HandleFunc("/solicitacaoById/{id}", GetSolicitacaoById).Methods("GET")
    router.HandleFunc("/solicitacaoByUserId/{id}", GetSolicitacaoByUserId).Methods("GET")
    router.HandleFunc("/saveSolicitacao/{title}/{description}", SaveSolicitacao).Methods("POST")
    log.Fatal(http.ListenAndServe(":8089", router))
}

func GetSolicitacao(response http.ResponseWriter, request *http.Request) {
    
}

func GetSolicitacaoById(response http.ResponseWriter, request *http.Request) {
    conn,_:= con.StablishConnection()
    defer conn.Close()
    vars := mux.Vars(request)
    id,_ := strconv.Atoi(vars["id"])
    solicitacao, err := models.Solicitacao{}.GetSolicitacaoById(conn, id)
    if err!=nil{
        responseEntity.Send(response, responseEntity.ErrorMessage{Status: http.StatusNotFound,Message: "Não foi possível encontrar a solicitação"}, http.StatusBadRequest)
    }else{
        responseEntity.Send(response, solicitacao, http.StatusOK)
    }
}

func GetSolicitacaoByUserId(response http.ResponseWriter, request *http.Request){

}

func SaveSolicitacao(response http.ResponseWriter, request *http.Request){
    conn,_:= con.StablishConnection()
    defer conn.Close()
    vars := mux.Vars(request)
    result := models.Solicitacao{}.SaveSolicitacao(conn,vars["title"],vars["description"])
    if result==false{
        responseEntity.Send(response, responseEntity.ErrorMessage{Status: http.StatusNotFound,Message: "Não foi possível inserir a sua solicitação"}, http.StatusBadRequest)
    }else{
        responseEntity.Send(response, result, http.StatusOK)
    }
}