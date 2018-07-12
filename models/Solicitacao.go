package models

import (
	db "database/sql"
	"time"
)

type Solicitacao struct{
	Id int `required="true" json:"id" description:"Identificação da solicitação"`
	Titulo string `required="true" json:"nome" description:"Título da solicitação"`
	Descricao string `required="true" json:"descricao" description:"Descricao da solicitação"`
	Data time.Time `required="true" json:"data" description:"Data da solicitação"`
}

func (solicitacao Solicitacao) GetSolicitacaoById(db *db.DB, id int) (*Solicitacao, error){
    err := db.QueryRow("SELECT id, titulo, descricao, data from solicitacoes where solicitacoes.id = ?", id).Scan(&solicitacao.Id, &solicitacao.Titulo, &solicitacao.Descricao, &solicitacao.Data)
	return &solicitacao, err
}

func (solicitacao Solicitacao) SaveSolicitacao(db *db.DB, titulo string, descricao string) (bool){
	statement, err := db.Prepare("INSERT INTO solicitacoes(titulo,descricao,data) values (?,?,GETDATE())")
	_, err = statement.Exec(statement, titulo, descricao)
	if err!=nil{
		return false;
	}else{
		return true;
	}
}