curl -X POST localhost:8000/accounts -d '{
    "name": "conta1",
    "direction": "debit",
    "id": "7f8b59c4-3fab-48c0-8deb-74db964882c1"
}'

curl -X POST localhost:8000/accounts -d '{
    "name": "conta2",
    "direction": "credit",
    "id": "5dbae912-bfeb-486f-81c0-28ac0e03d2e8"
}'

curl -X POST localhost:8000/transactions -d '{
    "name": "transacao1",
    "entries": [
        {
            "direction": "debit",
            "account_id": "7f8b59c4-3fab-48c0-8deb-74db964882c1",
            "amount": 100
        },
        {
            "direction": "credit",
            "account_id": "5dbae912-bfeb-486f-81c0-28ac0e03d2e8",
            "amount": 100
        }
    ]
}'
