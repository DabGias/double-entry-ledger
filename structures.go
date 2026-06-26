package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var DIRECTIONS map[string]struct{} = map[string]struct{}{
	"debit": {},
	"credit": {},
}

type account struct {
	Id uuid.UUID `json:"id"`
	Balance float64 `json:"balance"`
	Direction string `json:"direction"`
	Name string `json:"name"`
}

func (a account) String() string {
	return fmt.Sprintf("{ \"id\": \"%s\", \"name\": \"%s\", \"balance\": %f, \"direction\": \"%s\" }", a.Id.String(), a.Name, a.Balance, a.Direction)
}

func newAccount(data []byte) (account, error) {
	acc := account{}

	err := json.Unmarshal(data, &acc)

	if err != nil {
		return account{}, err
	}

	if acc.Id == uuid.Nil {
		acc.Id, err = uuid.NewRandom()

		if err != nil {
			return account{}, fmt.Errorf("Error generating account ID! (%s)", err)
		}
	}

	if acc.Direction == "" {
		return account{}, errors.New("Account direction is required to create an account!")
	} else if _, ok := DIRECTIONS[acc.Direction]; !ok {
		return account{}, fmt.Errorf("Account direction \"%s\" is invalid!", acc.Direction)
	}

	return acc, nil
}

func findAccount(accId uuid.UUID, accounts []*account) (*account, error) {
	for _, acc := range accounts {
		if acc.Id == accId {
			return acc, nil
		}
	}

	return &account{}, fmt.Errorf("Account with ID \"%s\" not found!", accId)
}

type entry struct {
	Id uuid.UUID `json:"id"`
	AccountId uuid.UUID `json:"account_id"`
	Direction string `json:"direction"`
	Amount float64 `json:"amount"`
}

func (e entry) String() string {
	return fmt.Sprintf("{ \"id\": \"%s\", \"account_id\": \"%s\", \"direction\": \"%s\", \"amount\": %f }", e.Id.String(), e.AccountId, e.Direction, e.Amount)
}

func newEntry(data []byte) (entry, error) {
	ent := entry{}

	err := json.Unmarshal(data, &ent)

	if err != nil {
		return entry{}, err
	}

	if ent.AccountId == uuid.Nil {
		return entry{}, errors.New("Account ID must not be null!")
	}

	if ent.Id == uuid.Nil {
		ent.Id, err = uuid.NewRandom()

		if err != nil {
			return entry{}, fmt.Errorf("Error generating entry ID! (%s)", err)
		}
	}

	if ent.Amount <= 0 {
		return entry{}, errors.New("Entry amount must be greater than 0!")
	}

	if ent.Direction == "" {
		return entry{}, errors.New("Entry direction is required to create an entry!")
	} else if _, ok := DIRECTIONS[ent.Direction]; !ok {
		return entry{}, fmt.Errorf("Entry direction \"%s\" is invalid!", ent.Direction)
	}

	return ent, nil
}

type transaction struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Entries []entry `json:"entries"`
}

func (t transaction) String() string {
	return fmt.Sprintf("{ \"id\": \"%s\", \"name\": \"%s\", \"entries\": %v }", t.Id.String(), t.Name, t.Entries)
}

func newTransaction(data []byte) (transaction, error) {
	trnsc := transaction{}

	err := json.Unmarshal(data, &trnsc)

	if err != nil {
		return transaction{}, err
	}

	if len(trnsc.Entries) == 0 {
		return transaction{}, errors.New("Transaction must contain at least one entry!")
	}

	if trnsc.Id == uuid.Nil {
		trnsc.Id, err = uuid.NewRandom()

		if err != nil {
			return transaction{}, fmt.Errorf("Error generating transaction ID! (%s)", err)
		}
	}

	for i, ent := range trnsc.Entries {
		data, err := json.Marshal(ent)

		if err != nil {
			return transaction{}, fmt.Errorf("Error marshling transaction entries! (%s)", err)
		}

		e, err := newEntry(data)

		if err != nil {
			return transaction{}, err
		}

		trnsc.Entries[i] = e
	}

	return trnsc, nil
}
