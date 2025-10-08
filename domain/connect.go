package domain

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mechiko/utility"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/sqlite"
)

const modError = "domain"

// func (d *dbs) IsConnected(info *DbInfo) (err error) {
func (d *DbInfo) Connect() (sess db.Session, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic: %v", modError, r)
		}
	}()
	switch strings.ToLower(d.Driver) {
	case "sqlite":
		if d.File == "" {
			return nil, fmt.Errorf("%s отсутствуют имя файла базы данных для sqlite", modError)
		}
		resultFilePath := filepath.Join(d.Path, d.File)
		if !utility.PathOrFileExists(resultFilePath) {
			return nil, fmt.Errorf("%s отсутствует файл базы данных %s для sqlite", modError, d.File)
		}
		// если указан не файл а путь к каталогу
		if st, statErr := os.Stat(resultFilePath); statErr != nil || !st.Mode().IsRegular() {
			return nil, fmt.Errorf("%s путь %s не является файлом sqlite", modError, resultFilePath)
		}
		uri := d.SqliteUri(resultFilePath)
		sess, err = sqlite.Open(uri)
		if err != nil {
			return nil, fmt.Errorf("%s ошибка подключения %w", modError, err)
		}
	case "mssql":
		if d.Name == "" {
			return nil, fmt.Errorf("%s отсутствуют имя базы данных для mssql", modError)
		}
		uri := d.MssqlUri()
		sess, err = mssql.Open(uri)
		if err != nil {
			return nil, fmt.Errorf("%s ошибка подключения: %w", modError, err)
		}
	default:
		return nil, fmt.Errorf("%s неизвестный драйвер %q", modError, d.Driver)
	}
	err = sess.Ping()
	if err != nil {
		_ = sess.Close() // best-effort cleanup
		return nil, fmt.Errorf("%s ошибка ping %w", modError, err)
	}
	// пинг успешен
	return sess, nil
}
