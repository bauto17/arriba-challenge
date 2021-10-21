# ARRIBA CHALLENGE

## Use
the service is running on this [url](https://arriba-challenge-production.up.railway.app)

### List all accounts 
```
curl --location --request GET 'https://arriba-challenge-production.up.railway.app/account'
```

### Create account
```
curl --location --request POST 'https://arriba-challenge-production.up.railway.app/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"The Name"
}'
```

### Deposit
```
curl --location --request POST 'https://arriba-challenge-production.up.railway.app/account/:account_id/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 50
}'
```

### Withdraw
```
curl --location --request POST 'https://arriba-challenge-production.up.railway.app/account/:account_id/withdraw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 50
}'
```

### Buy
```
curl --location --request POST 'https://arriba-challenge-production.up.railway.app/account/:account_id/buy' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 5,
    "currency": "BTC"
}'
```

### Vender
```
curl --location --request POST 'https://arriba-challenge-production.up.railway.app/account/:account_id/vender' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 5,
    "currency": "ETH"
}'
```