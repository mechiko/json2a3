package app

import (
	"context"
	"fmt"

	"json2a3/config"
	"json2a3/domain"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// type IApp interface {
// 	Options() *config.Configuration
// 	SaveOptions(string, any) error
// 	Logger() *zap.SugaredLogger
// }

type app struct {
	ctx           context.Context
	uuid          string // идентификатор для уникальности формы
	config        *config.Config
	options       *config.Configuration // копия config.Configuration
	loger         *zap.SugaredLogger
	pwd           string
	startTime     time.Time
	endTime       time.Time
	output        string
	dbSelfPath    string
	defaultDbPath string
}

var _ domain.Apper = (*app)(nil)

// const modError = "app"

func New(cfg *config.Config, logger *zap.SugaredLogger, pwd string) *app {
	newApp := &app{}
	newApp.pwd = pwd
	newApp.loger = logger
	newApp.config = cfg
	newApp.options = cfg.Configuration()
	newApp.uuid = uuid.New().String()
	return newApp
}

func (a *app) Pwd() string {
	return a.pwd
}

func (a *app) Config() *config.Config {
	return a.config
}

func (a *app) Logger() *zap.SugaredLogger {
	return a.loger
}

// выдаем адрес структуры опций программы чтобы править по месту
func (a *app) Options() *config.Configuration {
	return a.options
}

// записываем ключ и его значение только в пакет config
// и Options
// изменения не записываются в файл конфигурации
func (a *app) SetOptions(key string, value any) error {
	a.config.SetInConfig(key, value)
	a.options = a.config.Configuration()
	return nil
}

// записываем файл конфигурации состояние конфигурации
func (a *app) SaveOptions() error {
	if err := a.config.Save(); err != nil {
		return fmt.Errorf("save all in config error %w", err)
	}
	return nil
}

func (a *app) ConfigPath() string {
	if a.config != nil {
		return a.config.ConfigPath()
	}
	return ""
}

func (a *app) DefaultDbPath() string {
	return a.defaultDbPath
}

func (a *app) SetDefaultDbPath(path string) {
	a.defaultDbPath = path
}

func (a *app) LogPath() string {
	if a.config != nil {
		return a.config.LogPath()
	}
	return ""
}
