package main

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"lean/config"
	_ "lean/config"
	"lean/dao"
	"log"
	"os"
)

func main() {

	var cfg, err = config.Load()
	if err != nil {
		log.Fatalln("Could not get config.", err)
		return
	}

	dbConfig := (*cfg).Db

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Dbname,
		dbConfig.Port)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Connection failed.", err)
		return
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln("Close failed", err)
			return
		}
	}(db)

	questions := dao.NewQuestion(db)
	scanner := bufio.NewScanner(os.Stdin)

	userScore := 0
	for i := 0; i < len(*questions); i++ {
		fmt.Println((*questions)[i].Text)
		scanner.Scan()
		in := scanner.Text()
		if (*questions)[i].Answers[(*questions)[i].Correct] == in {
			userScore += 1
		}
	}

	fmt.Printf("Your score is %d\n", userScore)
}
