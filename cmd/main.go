package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg(("Error loading .env file"))
	}
}

// FIXME:
func logMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Debug().Str("path", r.URL.String())
	}
}

func main() {
	// db, err := rdb.NewDB()
	// if err != nil {
	// 	log.Fatal().Err(err)
	// }

	// taskRepo := repository.NewTask(db)
	// userRepo := repository.NewUser(db)
	// taskUsecase := usecase.NewTask(taskRepo, userRepo)
	// userUsecase := usecase.NewUser(userRepo)

	address := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal().Err(http.ListenAndServe(address+":"+port, nil)).Send()
}
