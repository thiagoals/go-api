package response

import(
	"net/http"
	"encoding/json"
)

func Send(response http.ResponseWriter, message interface{}, status int){
    result, _:= json.Marshal(message)
	response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(status)
    response.Write(result)
}