package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

//SQLStoreは実際に接続する。mockする時はStoreを使用する
//Store provides all functions to execure db queries and transactions
type Store interface {
	Querier
}

//SQLStore provides all functions to execure SQL queries and transaction
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

//txはトランザクション
//ececTx executes a function within a database transacsion
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	//このNewはNewStoreとは別物（*sql.DB,*sql.Txはどちらもインターフェースを満たしている）
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v,rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
