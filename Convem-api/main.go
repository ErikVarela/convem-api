package main //Onde a API será executada

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./banco.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS accounts (
			id TEXT PRIMARY KEY,
			holder TEXT,
			balance INTEGER
			)`)
	statement.Exec()

	db.Exec("INSERT OR IGNORE INTO accounts (id, holder, balance) VALUES ('1', 'Candidato Convem', 100000)")
	db.Exec("INSERT OR IGNORE INTO accounts (id, holder, balance) VALUES ('2', 'Empresa Parceira', 50000)")
}

type Account struct {
	ID      string `json:"id"`
	Holder  string `json:"holder"`
	Balance int64  `json:"balance"`
}

/*
	var accounts = []Account{
		{ID: "1", Holder: "Candidato Convem", Balance: 100000},
		{ID: "2", Holder: "Empresa Parceira", Balance: 50000},
	}
*/
func getBalance(context *gin.Context) {
	id := context.Param("id") //Pega o ID da URL e mostra
	var acc Account

	query := "SELECT id, holder, balance FROM accounts WHERE id =?"
	err := db.QueryRow(query, id).Scan(&acc.ID, &acc.Holder, &acc.Balance)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Conta não encontrada!"})
		return
	}

	/*
		for _, acc := range accounts {
			if acc.ID == id {
				//Se achar retorna a conta correspondente a busca...
				context.IndentedJSON(http.StatusOK, acc)
				return
			}
		}
	*/
	context.IndentedJSON(http.StatusOK, acc)

}

func transferMoney(context *gin.Context) {
	var req struct {
		FromID string `json:"from_id"`
		ToID   string `json:"to_id"`
		Amount int64  `json:"amount"`
	}

	if err := context.BindJSON(&req); err != nil {
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", req.Amount, req.FromID)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", req.Amount, req.ToID)
	if err != nil {
		tx.Rollback()
		return
	}

	/*
		for i := range accounts {
			if accounts[i].ID == request.FromID {
				accounts[i].Balance -= request.Amount //Manda o dinheiro
			}
			if accounts[i].ID == request.ToID {
				accounts[i].Balance += request.Amount //Recebe o dinheiro
			}
		}*/
	tx.Commit()
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Transferência realizada com sucesso!"})
}

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()
	router.GET("/accounts/:id/balance", getBalance)
	router.POST("/transfer", transferMoney)
	router.Run("localhost:8080")
}
