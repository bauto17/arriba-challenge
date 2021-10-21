package repositories

const CreateTable = `
CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL,
  name varchar(45) NOT NULL, 
  usd  INTEGER NOT NULL,
  btc  INTEGER NOT NULL,
  eth  INTEGER NOT NULL,
  CONSTRAINT ch_positive_usd CHECK (usd>=0),
  CONSTRAINT ch_positive_btc CHECK (btc>=0),
  CONSTRAINT ch_positive_eth CHECK (eth>=0),
  PRIMARY KEY (id)
)
`

const CreateAccountQuery = `
INSERT INTO accounts (name, usd, btc, eth) VALUES ($1, 0, 0, 0) RETURNING id
`

const GetAccountsQuery = `
SELECT id, name, usd, btc, eth FROM accounts
`

const DepositQuery = `
UPDATE accounts SET usd = (usd+$1) WHERE id = $2;
`

const WithdrawQuery = `
UPDATE accounts SET usd = (usd-$1) WHERE id = $2;
`

const BuyQuery = `
UPDATE accounts SET usd = (usd-$1), :currency = (:currency+$2) WHERE id = $3;
`

const VenderQuery = `
UPDATE accounts SET usd = (usd+$1), :currency = (:currency-$2) WHERE id = $3;
`
