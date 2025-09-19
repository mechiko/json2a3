package main

import (
	"fmt"
	"json2a3/domain"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/upper/db/v4"
)

// берет из конфига настройки ДБ src
func openDbSrc(app domain.Apper) (src db.Session, err error) {
	dbi := &domain.DbInfo{
		File:   app.Options().DB,
		Driver: "sqlite",
	}
	src, err = dbi.Connect()
	return src, err
}

func InsertBatchTx(app domain.Apper, dbs db.Session, jsonData *Codes) (count int64, err error) {
	count = 0
	err = dbs.Tx(func(tx db.Session) (err error) {
		count, err = moveBatch(app, tx, jsonData)
		return err
	})
	return count, err
}

func moveBatch(app domain.Apper, dst db.Session, jsonData *Codes) (int64, error) {
	batchSize := 10
	batch := dst.SQL().InsertInto("order_mark_codes_serial_numbers").Columns("id_order_mark_codes", "gtin", "serial_number", "code", "block_id", "status").Batch(batchSize)
	// i := 1
	var inserted int64
	go func() {
		defer batch.Done()
		order := app.Options().Order
		max := app.Options().Serial
		gtin := app.Options().Gtin
		for i, code := range jsonData.Codes {
			serial := SerialStr(i+1, max)
			batch.Values(order, gtin, serial, code, jsonData.BlockId, "Получен")
			n := atomic.AddInt64(&inserted, 1)
			if (n % 10000) == 0 {
				app.Logger().Infof("I:%011d", n)
			}
		}
	}()
	err := batch.Wait()
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}
	return atomic.LoadInt64(&inserted), err
}

func SerialStr(i int, max int) string {
	out := strconv.Itoa(i)
	outLen := len(out)
	if outLen > max {
		return out[:max]
	}
	pad := max - outLen
	pad0 := strings.Repeat("0", pad)
	return pad0 + out
}
