package main

import (
	"encoding/json"
	"errors"
	"flag"
	"json2a3/app"
	"json2a3/config"
	"json2a3/zaplog"
	"os"
	"path/filepath"

	"github.com/upper/db/v4"
)

func main() {
	flag.Parse()
	if *file == "" {
		errMessageExit(nil, "ошибка", errors.New("не указан файл json"))
	}
	if *dbA3 == "" {
		errMessageExit(nil, "ошибка", errors.New("не указан файл db"))
	}
	if *order == "" {
		errMessageExit(nil, "ошибка", errors.New("не указан номер заказа"))
	}
	if *gtin == "" {
		errMessageExit(nil, "ошибка", errors.New("не указан gtin товара"))
	}
	if *serial == 0 {
		errMessageExit(nil, "ошибка", errors.New("не указана длина серийного номера для ТГ товара"))
	}

	cfg, err := config.New("", false)
	if err != nil {
		errMessageExit(nil, "ошибка конфигурации", err)
	}

	var logsOutConfig = map[string][]string{
		"logger": {"stdout", filepath.Join(cfg.LogPath(), config.Name)},
	}
	zl, err := zaplog.New(logsOutConfig, config.Mode == "development")
	if err != nil {
		errMessageExit(nil, "ошибка создания логера", err)
	}

	lg, err := zl.GetLogger("logger")
	if err != nil {
		errMessageExit(nil, "ошибка получения логера", err)
	}
	loger := lg.Sugar()
	loger.Debug("zaplog started")
	loger.Infof("mode = %s", config.Mode)
	if cfg.Warning() != "" {
		loger.Infof("pkg:config warning %s", cfg.Warning())
	}

	errProcessExit := func(title string, err error) {
		errMessageExit(loger, title, err)
	}
	// создаем приложение с опциями из конфига и логером основным
	app := app.New(cfg, loger, ".")
	if err := app.SetOptions("file", *file); err != nil {
		errProcessExit("ошибка установки опции file", err)
	}
	if err := app.SetOptions("db", *dbA3); err != nil {
		errProcessExit("ошибка установки опции db", err)
	}
	if err := app.SetOptions("order", *order); err != nil {
		errProcessExit("ошибка установки опции order", err)
	}
	if err := app.SetOptions("serial", *serial); err != nil {
		errProcessExit("ошибка установки опции serial", err)
	}
	if err := app.SetOptions("gtin", *gtin); err != nil {
		errProcessExit("ошибка установки опции gtin", err)
	}
	// убираем log upper slow query
	db.LC().SetLevel(db.LogLevelError)
	src, err := openDbSrc(app)
	if err != nil {
		errProcessExit("ошибка открытия БД источник", err)
	}
	defer src.Close()

	jsonBytes, err := os.ReadFile(*file)
	if err != nil {
		errProcessExit("ошибка открытия JSON", err)
	}
	var Codes Codes
	err = json.Unmarshal(jsonBytes, &Codes)
	if err != nil {
		errProcessExit("ошибка JSON", err)
	}
	count, err := InsertBatchTx(app, src, &Codes)
	if err != nil {
		errProcessExit("ошибка вставки кодов в БД", err)
	}
	loger.Infof("вставлено %d кодов в заказ", count)
	// завершаем все логи
	zl.Shutdown()
}
