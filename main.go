package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	accounts := []*account{}
	transactions := []*transaction{}

	mux := http.NewServeMux()

	mux.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, "{ \"accounts\": %v }", accounts)
			case "POST":
				defer r.Body.Close()

				data, err := io.ReadAll(r.Body)

				if err != nil {
					fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
					return
				}

				acc, err := newAccount(data)

				if err != nil {
					fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
					return
				}

				accounts = append(accounts, &acc)

				fmt.Fprintf(w, "{ \"account\": \"%s\", \"status\": %d }", acc, http.StatusCreated)
		}
	})

	mux.HandleFunc("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				accountId, err := uuid.Parse(r.PathValue("id"))

				if err != nil {
					fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
					return
				}

				for _, acc := range accounts {
					if acc.Id == accountId {
						fmt.Fprintf(w, "{ \"account\": \"%s\", \"status\": %d }", acc, http.StatusOK)
						return
					}
				}

				fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", "Account not found!", http.StatusNotFound)
		}
	})

	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, "{ \"transactions\": %v }", transactions)
			case "POST":
				defer r.Body.Close()

				data, err := io.ReadAll(r.Body)

				if err != nil {
					fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
					return
				}

				trnsc, err := newTransaction(data)

				if err != nil {
					fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
					return
				}

				for _, ent := range trnsc.Entries {
					acc, err := findAccount(ent.AccountId, accounts)

					if err != nil {
						fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
						return
					}

					fmt.Printf("Entry: %s\n", ent)
					fmt.Printf("Account Before: %s\n", acc)

					if acc.Direction == ent.Direction {
						acc.Add(ent.Amount)
					} else {
						acc.Subtract(ent.Amount)
					}	

					fmt.Printf("Account After: %s\n", acc)
				}

				transactions = append(transactions, &trnsc)

				fmt.Fprintf(w, "{ \"transaction\": \"%s\", \"status\": %d }", trnsc, http.StatusCreated)
		}
	})

	mux.HandleFunc("/transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				trnscId, err := uuid.Parse(r.PathValue("id"))

				if err != nil {
					fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", err, http.StatusBadRequest)
					return
				}

				for _, t := range transactions {
					if t.Id == trnscId {
						fmt.Fprintf(w, "{ \"transaction\": \"%s\", \"status\": %d }", t, http.StatusOK)
						return
					}
				}

				fmt.Fprintf(w, "{ \"error\": \"%s\", \"status\": %d }", "Transaction not found!", http.StatusNotFound)
		}
	})

	server := http.Server{
		Addr: ":8000",
		Handler: mux,
	}

	fmt.Printf("Running server on localhost%s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error running HTTP server: %s\n", err)
		}
	}
}
