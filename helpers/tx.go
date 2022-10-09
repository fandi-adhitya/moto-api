package helpers

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicError(errRollback.Error)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit.Error)
	}
}
