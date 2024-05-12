// server.go

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// type Film struct {
// 	Title    string
// 	Director string
// 	Choice   string
// }

type Game struct {
	Player1Choice      string `json:"player1Choice"`
	Player2Choice      string `json:"player2Choice"`
	Winner             string `json:"winner"`
	isPlayer2_computer string `json:"isComputer"`
}

var db *sql.DB

// var computerChoice string
var isPlayer2_computer bool

func main() {
	fmt.Println("Go app...")

	//dbConnection()

	// var err error
	// db, err = sql.Open("mysql", "root:adminPWD@tcp(127.0.0.1:3306)/game_db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Connected to database")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("----- in h1 handler ------")

		fmt.Println("Go app...")
		tmpl := template.Must(template.ParseFiles("player-computer.html"))
		tmpl.Execute(w, nil)
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-game/", playHandler)

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func playHandler(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(1 * time.Second)
	fmt.Println("----- in h2 handler ------")

	log.Print("HX-Request received : ", r.Header.Get("HX-Request"))

	player2Choice := []string{"rock", "paper", "scissors"}[rand.Intn(3)]
	// isPlayer2_computer := true
	player1Choice := r.PostFormValue("submit")

	winner := getWinner(player1Choice, player2Choice)
	//log.Println("Player choice:", player1Choice, ", Computer choice :", player2Choice, ", Winner :", winner, ", isPlayer2_computer :", isPlayer2_computer)
	game := Game{
		Player1Choice: player1Choice,
		Player2Choice: player2Choice,
		Winner:        winner,
	}
	log.Println(game)
	//insertGame(game, isPlayer2_computer)
	//json.NewEncoder(w).Encode(game)

	// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	// tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl := template.Must(template.ParseFiles("player-computer.html"))
	tmpl.ExecuteTemplate(w, "game-list-element", game)
	//tmpl.ExecuteTemplate(w, "game-list-element", Game{Player1Choice: player1Choice, Player2Choice: player2Choice, Winner: winner})
}

// func playHandler1(w http.ResponseWriter, r *http.Request) {
// 	log.Print("HX-Request received : ")
// 	log.Println(r.Header.Get("HX-Request"))
// 	var isPlayer2_computer bool
// 	player2Choice := []string{"rock", "paper", "scissors"}[rand.Intn(3)]
// 	isPlayer2_computer = true
// 	player1Choice := r.PostFormValue("submit")
// 	//	log.Println("player choice: ", player1Choice, "Computer choice : ", player2Choice)
// 	winner := getWinner(player1Choice, player2Choice)
// 	log.Println("Player choice:", player1Choice, ", Computer choice :", player2Choice, ", Winner :", winner, ", isPlayer2_computer :", isPlayer2_computer)
// 	game := Game{
// 		Player1Choice:      player1Choice,
// 		Player2Choice:      player2Choice,
// 		Winner:             winner,
// 		isPlayer2_computer: "true",
// 	}
// 	//insertGame(game, isPlayer2_computer)
// 	json.NewEncoder(w).Encode(game)
// 	//display in browser
// 	// htmlStr := fmt.Sprintf("<tr><td>%s</td><td>%s </td><td> %s</td></tr> ", player1Choice, player2Choice, winner)
// 	// templ, _ := template.New("t").Parse(htmlStr)
// 	// templ.Execute(w, nil)
// }

func dbConnection() {

	var err error
	db, err = sql.Open("mysql", "root:adminPWD@tcp(127.0.0.1:3306)/game_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
}

func getWinner(choice1, choice2 string) string {
	if choice1 == choice2 {
		return "draw"
	}
	if (choice1 == "rock" && choice2 == "scissors") || (choice1 == "scissors" && choice2 == "paper") || (choice1 == "paper" && choice2 == "rock") {
		return "player1"
	}
	return "player2"
}

func insertGame(game Game, isComputer bool) {
	stmt, err := db.Prepare("INSERT INTO games(player1_choice, player2_choice, winner, isPlayer2_computer) VALUES(?, ?, ?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(game.Player1Choice, game.Player2Choice, game.Winner, isComputer)
	if err != nil {
		log.Fatal(err)
	}
}
