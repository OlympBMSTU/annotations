package main

import (
	"net/http"
	"log"

	"github.com/OlympBMSTU/annotations/config"
	"github.com/OlympBMSTU/annotations/controllers"

	// "github.com/OlympBMSTU/annotations/db"
	"github.com/jackc/pgx"
)

func Init() (*pgx.ConnPool, error) {
	_, err := config.GetConfigInstance()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// port, err := strconv.Atoi(conf.GetDBPort())
	// if err != nil {
	// 	log.Print(err)
	// 	return nil, err
	// }
	// connPoolConfig := pgx.ConnPoolConfig{
	// 	ConnConfig: pgx.ConnConfig{
	// 		Host:     conf.GetDBHost(),
	// 		User:     conf.GetDBUser(),
	// 		Port:     uint16(port),
	// 		Password: conf.GetDBPassword(),
	// 		Database: conf.GetDatabase(),
	// 	},
	// 	MaxConnections: 5,
	// }

	// pool, err := pgx.NewConnPool(connPoolConfig)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return nil, nil
}

func InitConfig() {
	http.Handle("/api/annotations/upload", http.HandlerFunc(controllers.UploadAnnotationHandler))
}


func main() {
	_, err := Init()
	if err != nil {
		log.Print(err)
	}

	conf, _ := config.GetConfigInstance() 

	// fmt.Println(pool)
	// service := db.AnnotationService{pool}'

	http.ListenAndServe(conf.GetListenerHost() + ":" + conf.GetListenerPort(), nil)
}
