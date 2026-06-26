package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {
	accounts := []account{}
	entries := []entry{}
	transactions := []transaction{}

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong!")
	})

	mux.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, `{ "accounts": %v }`, accounts)
			case "POST":
				defer r.Body.Close()

				data, err := io.ReadAll(r.Body)

				if err != nil {
					fmt.Fprintf(w, `{ "error": %s, "status": %d }`, err, http.StatusBadRequest)
					return
				}

				acc, err := newAccount(data)

				if err != nil {
					fmt.Fprintf(w, `{ "error": %s, "status": %d }`, err, http.StatusBadRequest)
					return
				}

				accounts = append(accounts, acc)

				fmt.Fprintf(w, `{ "account": %s, "status": %d }`, acc, http.StatusCreated)
		}
	})

	mux.HandleFunc("/entries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, `{ "entries": %v }`, entries)
			case "POST":
				defer r.Body.Close()

				data, err := io.ReadAll(r.Body)

				if err != nil {
					fmt.Fprintf(w, `{ "error": %s, "status": %d }`, err, http.StatusBadRequest)
					return
				}

				ent, err := newEntry(data)

				if err != nil {
					fmt.Fprintf(w, `{ "error": %s, "status": %d }`, err, http.StatusBadRequest)
					return
				}

				entries = append(entries, ent)

				fmt.Fprintf(w, `{ "entry": %s, "status": %d }`, ent, http.StatusCreated)
		}
	})

	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, `{ "transactions": %v }`, transactions)
			case "POST":
				
		}
	})

	server := http.Server{
		Addr: ":8000",
		Handler: mux,
	}

	fmt.Printf("Trying to run server on localhost%s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error running HTTP server: %s\n", err)
		}
	}
}
