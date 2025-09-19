package move

import (
	"fmt"

	"github.com/upper/db/v4"
)

func (m *move) MoveTableTx(table string) (count int64, err error) {
	count = 0
	err = m.Dst.Tx(func(tx db.Session) (err error) {
		count, err = m.MoveTable(tx, table)
		return err
	})
	return count, err
}

func (m *move) MoveTable(dst db.Session, table string) (int64, error) {
	var record map[string]interface{}
	resSrc := m.Src.Collection(table).Find()
	defer resSrc.Close()
	resDst := dst.Collection(table)
	countInsert := int64(0)
	for resSrc.Next(&record) {
		if _, err := resDst.Insert(record); err != nil {
			return countInsert, fmt.Errorf("insert into %s failed: %w", table, err)
		}
		countInsert++
	}
	if err := resSrc.Err(); err != nil {
		return countInsert, fmt.Errorf("read from %s failed: %w", table, err)
	}
	return countInsert, nil
}
