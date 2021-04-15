package context_strengthen

import (
    `context`
    `database/sql`
)

var txKey = &struct {}{}


func CreateTransactionContext(parent context.Context, tx *sql.Tx) context.Context{
    return context.WithValue(parent, txKey, tx)
}

