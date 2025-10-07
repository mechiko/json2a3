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

func InsertBatchTx(app domain.Apper, dbs db.Session, jsonData *Codes, serialInit int64) (count int64, err error) {
	count = 0
	err = dbs.Tx(func(tx db.Session) (err error) {
		count, err = moveBatch(app, tx, jsonData, serialInit)
		return err
	})
	return count, err
}

func moveBatch(app domain.Apper, dst db.Session, jsonData *Codes, serialInit int64) (int64, error) {
	batchSize := 10
	batch := dst.SQL().InsertInto("order_mark_codes_serial_numbers").Columns("id_order_mark_codes", "gtin", "serial_number", "code", "block_id", "status").Batch(batchSize)

	var inserted int64
	go func() {
		defer func() {
			if r := recover(); r != nil {
				app.Logger().Errorf("panic in batch goroutine: %v", r)
			}
		}()
		defer batch.Done()
		order := app.Options().Order
		max := app.Options().Serial
		gtin := app.Options().Gtin
		for i, code := range jsonData.Codes {
			serial := SerialStr(i+int(serialInit), max)
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

func findSerialMax(app domain.Apper, dst db.Session, order string) (int64, error) {
	orderInt, err := strconv.Atoi(order)
	if err != nil {
		return 0, fmt.Errorf("order string error %w", err)
	}
	col := dst.Collection("order_mark_codes")
	res := col.Find("id", orderInt)
	var valOrder map[string]interface{}
	err = res.One(&valOrder)
	if err != nil {
		return 0, fmt.Errorf("find error %w", err)
	}
	var valSerial map[string]interface{}
	gtin := valOrder["gtin"].(string)
	app.Logger().Infof("order:%d gtin:%s", orderInt, gtin)
	ress := dst.SQL().Select(db.Raw("max(serial_number)")).From("order_mark_codes_serial_numbers").Where("gtin", gtin)
	err = ress.One(&valSerial)
	if err != nil {
		return 0, fmt.Errorf("find error %w", err)
	}
	if serial, ok := valSerial["max(serial_number)"].(string); ok {
		serialInt, err := strconv.ParseInt(serial, 10, 64)
		return serialInt, err
	}
	if serial, ok := valSerial["max(serial_number)"].(int64); ok {
		return serial, nil
	}
	// app.Logger().Infof("type max id %T", valSerial["max(serial_number)"])
	// return 0, nil
	return 0, fmt.Errorf("unexpected type for max(serial_number): %T", valSerial["max(serial_number)"])
}
