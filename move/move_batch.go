package move

import (
	"fmt"
	"sync/atomic"

	"github.com/upper/db/v4"
)

func (m *move) MoveBatchTx(table string, fields []string) (count int64, err error) {
	if len(fields) == 0 {
		return 0, fmt.Errorf("fields must not be empty")
	}
	count = 0
	err = m.Dst.Tx(func(tx db.Session) (err error) {
		count, err = m.moveBatch(tx, table, fields)
		return err
	})
	return count, err
}

func makeVals(record map[string]interface{}, flds []string) ([]interface{}, error) {
	if len(flds) == 0 {
		return nil, fmt.Errorf("fields must not be empty")
	}
	out := make([]interface{}, 0, len(flds))
	for _, fname := range flds {
		if f, ok := record[fname]; ok {
			out = append(out, f)
		} else {
			return nil, fmt.Errorf("поле %s не найдено", fname)
		}
	}
	return out, nil
}

func (m *move) moveBatch(dst db.Session, table string, fields []string) (int64, error) {
	var record map[string]interface{}
	batchSize := len(fields)
	resSrc := m.Src.Collection(table).Find()
	defer resSrc.Close()
	batch := dst.SQL().InsertInto(table).Columns(fields...).Batch(batchSize)
	// i := 1
	var inserted int64
	go func() {
		defer batch.Done()
		for resSrc.Next(&record) {
			vals, err := makeVals(record, fields)
			if err != nil {
				// m.Logger().Errorf("ошибка batch копирования %s %d", table, i)
				m.Logger().Errorf("ошибка batch копирования %s: %v", table, err)
				continue
			}
			batch.Values(vals...)
			n := atomic.AddInt64(&inserted, 1)
			if (n % 100000) == 0 {
				m.Logger().Infof("I:%011d", n)
			}
		}
	}()
	err := batch.Wait()
	if err == nil {
		if e := resSrc.Err(); e != nil {
			err = fmt.Errorf("source iteration error: %w", e)
		}
	}
	return atomic.LoadInt64(&inserted), err
}
