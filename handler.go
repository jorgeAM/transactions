package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/jorgeAM/kata-transactions/application"
	"github.com/jorgeAM/kata-transactions/infrastructure"
)

func createUser(db *sqlx.DB) http.HandlerFunc {
	srv := application.NewCreateUser(
		infrastructure.NewPostgresUserRepository(db),
		infrastructure.NewPostgresNotificationRepository(db),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		if err := srv.Exec(r.Context()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		w.Write([]byte("ok"))
	}
}
