package domain

import (
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/sqlite"
)

// host
// user
// password
// name имя файла без пути
// file полное имя файла и путь
// name имя БД
// driver драйвер
// Exists файл существует и только
type DbInfo struct {
	Host       string
	User       string
	Pass       string
	File       string
	Name       string
	Driver     string
	Connection string
	Exists     bool // только для sqlite делает поиск файла
	Path       string
}

func (d *DbInfo) MssqlUri() *mssql.ConnectionURL {
	uri := &mssql.ConnectionURL{
		User:     d.User,
		Password: d.Pass,
		Host:     d.Host,
		Database: d.Name,
		Options: map[string]string{
			"encrypt": "disable",
		},
	}
	if uri.Host == "" {
		uri.Host = "127.0.0.1:1433"
	}
	return uri
}

func (d *DbInfo) SqliteUri(file string) *sqlite.ConnectionURL {
	uri := &sqlite.ConnectionURL{
		Database: file,
		Options: map[string]string{
			"mode":          "rw",
			"_journal_mode": "DELETE",
		},
	}
	return uri
}
