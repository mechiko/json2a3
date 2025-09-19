package move

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/upper/db/v4"
)

func (m *move) Clear(dst db.Session, tables []string) (err error) {
	if len(tables) == 0 {
		m.Logger().Errorf("пустой список таблиц для сброса")
		return nil
	}
	// collect errors across tables to avoid dropping earlier failures
	var errs []error
	for _, table := range tables {
		delete := fmt.Sprintf("delete from %s;", table)
		_, err = dst.SQL().Exec(delete)
		if err != nil {
			errs = append(errs, err)
			m.Logger().Errorf("ошибка сброса таблицы %s delete %s", table, err.Error())
			continue
		}
		dbcc := fmt.Sprintf("DBCC CHECKIDENT (N'[dbo].[%s]', RESEED, 0);", table)
		_, err = dst.SQL().Exec(dbcc)
		if err != nil {
			errs = append(errs, err)
			m.Logger().Errorf("ошибка сброса таблицы %s dbcc %s", table, err.Error())
			continue
		}
	}
	return errors.Join(errs...)
}

func (m *move) ClearTx(tables []string) (err error) {
	if len(tables) == 0 {
		return fmt.Errorf("пустой список таблиц для сброса")
	}
	err = m.Dst.Tx(func(tx db.Session) (err error) {
		err = m.Clear(tx, tables)
		return err
	})
	return err
}
