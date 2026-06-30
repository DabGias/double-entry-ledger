<h1 align="center">Double-Entry-Ledger</h1>

# 🇧🇷 pt-br

> [!NOTE]
> Esse projeto é uma solução de um desafio proposto no vídeo [Criando um Ledger](https://youtu.be/ktM-ocowE4Q).

Esse projeto é um double-entry-ledger, sistema de bookkeeping onde todas as transações são registradas com, pelo menos, dois lançamentos, por exemplo, uma transação em que R$5.000,00 são debitados de uma conta obrigatoriamente deve conter um lançamento indicando onde os mesmos R$5.000,00 foram creditados, assim, garantindo um maior controle e facilitando a localização de fraudes.

Desenvolvi essa solução para aprimorar meus conhecimentos em Go, comecei a aprender a linguagem há algum tempo e pensei que um projeto nesse nível de dificuldade seria ótimo para aprimorar o que eu tinha aprendido até agora e, também, aprender coisas novas, além disso, achei que Go seria uma linguagem ótima para um projeto simples.

## Rodando o projeto

```bash
git clone https://github.com/DabGias/double-entry-ledger

cd double-entry-ledger

go run .
```

## Endpoints

### Contas

```bash
# GET

curl localhost:8000/accounts

curl localhost:8000/accounts/{id}

# POST

curl -X POST localhost:8000/accounts -d '{
    "name": "account1",
    "direction": "debit",
}'
```

### Transações

```bash
# GET

curl localhost:8000/transactions

curl localhost:8000/transactions/{id}

# POST

curl -X POST localhost:8000/transactions -d '{
    "name": "transaction1",
    "entries": [
        {
            "direction": "debit",
            "account_id": "{account ID}",
            "amount": 100
        },
        {
            "direction": "credit",
            "account_id": "{account ID}",
            "amount": 100
        }
    ]
}'
```

# 🇺🇸 en-us

> [!NOTE]
> This project is a solution for the challenge proposed in the video [Creating a Ledger](https://youtu.be/ktM-ocowE4Q)

This project is a double-entry-ledger, a bookkeeping principal which saves records of every transaction using, at least, two entries, for example, when USD$5.000,00 are spent there must be another entry record stating where this USD$5.000,00 where credited, this method of bookkeeping ensures a better control over the financial records and eases error and fraud detection.

I developed this solution with the objective to learn and practice Go, as I just started to learn the language, and thought that the difficult of this project would be a great way to improve my skills and learn new concepts, and I also thought that Go would be a great tool to develop this project because of the simplicity that Go provides to the developer.

## Running the project

```bash
git clone https://github.com/DabGias/double-entry-ledger

cd double-entry-ledger

go run .
```

## Endpoints

### Accounts

```bash
# GET

curl localhost:8000/accounts

curl localhost:8000/accounts/{id}

# POST

curl -X POST localhost:8000/accounts -d '{
    "name": "account1",
    "direction": "debit",
}'
```

### Transactions

```bash
# GET

curl localhost:8000/transactions

curl localhost:8000/transactions/{id}

# POST

curl -X POST localhost:8000/transactions -d '{
    "name": "transaction1",
    "entries": [
        {
            "direction": "debit",
            "account_id": "{account ID}",
            "amount": 100
        },
        {
            "direction": "credit",
            "account_id": "{account ID}",
            "amount": 100
        }
    ]
}'
```
