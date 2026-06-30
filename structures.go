package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

var DIRECTIONS map[string]struct{} = map[string]struct{}{
	"debit": {},
	"credit": {},
}

type money struct {
	Cents int64
}

func (m *money) UnmarshalJSON(data []byte) error {
	if num, err := strconv.ParseFloat(string(data), 64); err != nil {
		return err
	} else {
		m.Cents = int64(num * 1000)

		return nil
	}
}

func (m *money) Float() float32 {
	return float32(m.Cents) / 1000
}

type account struct {
	Id uuid.UUID `json:"id"`
	Balance money `json:"balance"`
	Direction string `json:"direction"`
	Name string `json:"name"`
}

func (a account) String() string {
	return fmt.Sprintf("{ \"id\": \"%s\", \"name\": \"%s\", \"balance\": %.3f, \"direction\": \"%s\" }", a.Id.String(), a.Name, a.Balance.Float(), a.Direction)
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

	if acc.Balance.Cents != 0 {
		return account{}, errors.New("Account balance can't be set on creation!")
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
	Amount money `json:"amount"`
}

func (e entry) String() string {
	return fmt.Sprintf("{ \"id\": \"%s\", \"account_id\": \"%s\", \"direction\": \"%s\", \"amount\": %.3f }", e.Id.String(), e.AccountId, e.Direction, e.Amount.Float())
}

type transaction struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Entries []*entry `json:"entries"`
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

	if len(trnsc.Entries) < 2 {
		return transaction{}, errors.New("Transaction must contain at least two entries!")
	}

	if trnsc.Id == uuid.Nil {
		trnsc.Id, err = uuid.NewRandom()

		if err != nil {
			return transaction{}, fmt.Errorf("Error generating transaction ID! (%s)", err)
		}
	}

	balance := 0

	for _, ent := range trnsc.Entries {
		if ent.Direction == "credit" {
			balance += int(ent.Amount.Cents)
		} else {
			balance -= int(ent.Amount.Cents)
		}

		if ent.AccountId == uuid.Nil {
			return transaction{}, errors.New("Account ID must not be null!")
		}

		if ent.Id == uuid.Nil {
			ent.Id, err = uuid.NewRandom()

			if err != nil {
				return transaction{}, fmt.Errorf("Error generating entry ID! (%s)", err)
			}
		}

		if ent.Amount.Cents <= 0 {
			return transaction{}, errors.New("Entry amount must be greater than 0!")
		}

		if ent.Direction == "" {
			return transaction{}, errors.New("Entry direction is required to create an entry!")
		} else if _, ok := DIRECTIONS[ent.Direction]; !ok {
			return transaction{}, fmt.Errorf("Entry direction \"%s\" is invalid!", ent.Direction)
		}
	}

	if balance != 0 {
		return transaction{}, errors.New("Transaction must be balanced!")
	}

	return trnsc, nil
}
