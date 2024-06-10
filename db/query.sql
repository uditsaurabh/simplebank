-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateAccount :one
INSERT INTO
    accounts (
        owner,
        balance,
        currency,
        country_code
    )
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET
    owner = $2,
    balance = $3,
    currency = $4,
    country_code = $5
WHERE
    id = $1 RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET
    balance = balance + sqlc.arg(amount)
WHERE
    id = sqlc.arg (id) RETURNING *;
-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateEntry :one
INSERT INTO
    entries (account_id, amount)
VALUES ($1, $2) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO
    transfers (
        from_account_id,
        to_account_id,
        amount
    )
VALUES ($1, $2, $3) RETURNING *;