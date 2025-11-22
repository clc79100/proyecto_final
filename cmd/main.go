package main

import (
	"database/sql"
	"log"
	"net/http"

	"proyecto_final/cmd/handlers"
	"proyecto_final/cmd/repositories"
	"proyecto_final/cmd/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Password1@tcp(127.0.0.1:3306)/School")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	studentRepo := repositories.NewMySQLStudentRepository(db)
	studentHandler := handlers.NewStudentHandler(studentRepo)

	subjectRepo := repositories.NewMySQLSubjectRepository(db)
	subjectHandler := handlers.NewSubjectHandler(subjectRepo)

	gradeRepo := repositories.NewSQLGradeRepository(db)
	gradeHandler := handlers.NewGradeHandler(gradeRepo)

	router := routes.SetupRoutes(studentHandler, subjectHandler, gradeHandler)

	log.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
