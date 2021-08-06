package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
)

const token string = "ODcyNzMwOTg1OTcwNjYzNDU0.YQuIEQ.7cnjUQVo-vdYOKyyYHM2bmAvzG0"

var BotID string

var db *sql.DB

var today time.Weekday

var people []string

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Maquina01@"
	dbname   = "postgres"
)

func main() {

	var err error

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := s.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	s.AddHandler(messageHandler)

	err = s.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running!")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	<-make(chan struct{})
	return
}

/*func retrievePeople(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := db.QueryRow("SELECT people FROM weekly_availability WHERE week_days=?", today).Scan(&people)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Println(people)
}*/

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!today" {
		today := time.Now().Weekday()

		//http.HandleFunc("/retrieve", retrievePeople)

		userSql := "SELECT people FROM weekly_availability WHERE week_days=?"

		db.QueryRow(userSql, today.String()).Scan(&people)

		var peopleFull string = ""

		for _, str := range people {
			peopleFull += str + " "
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, today.String()+" schedule:\n "+peopleFull)

		//http.ListenAndServe(":8080", nil)

		//defer db.Close()

	}

	if m.Author.ID == BotID {
		return
	}

	fmt.Println(m.Content)
}
