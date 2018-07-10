package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/AddressBook/backend/data"
	"github.com/ComputePractice2018/AddressBook/backend/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
Небольшие подсказки:
- Поставьте пакеты go get -u github.com/jinzhu/gorm  и go get -u github.com/go-sql-driver/mysql
- ДЛя запуска mysql из Docker запустите: docker run -d --rm --name db -e MYSQL_RANDOM_ROOT_PASSWORD=yes -e MYSQL_DATABASE=addressbook -e MYSQL_USER=addressbook -e MYSQL_PASSWORD=SuperSecretPassword  -p 3306:3306 mysql:latest
- Для запуска бекенда - укажите правильный адрес соединения (db - IP docker'а)

*/

func main() {
	connection := flag.String("connection", "addressbook:SuperSecretPassword@tcp(db:3306)/addressbook", "mysql connection string")
	flag.Parse()

	contactList, err := data.NewDBContactList(*connection)
	defer contactList.Close()
	//contactList := data.NewContactList()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(contactList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
