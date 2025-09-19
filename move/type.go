package move

import (
	"json2a3/domain"

	"github.com/upper/db/v4"
)

type move struct {
	domain.Apper
	Src                    db.Session
	Dst                    db.Session
	OrderMarkCodes         map[int64]int64 // отображение старых ID но новые
	OrderMarkAggregation   map[int64]int64
	OrderMarkAtk           map[int64]int64
	OrderMarkCommissioning map[int64]int64
	OrderMarkConnectingTap map[int64]int64
	OrderMarkUtilisation   map[int64]int64
}

func New(app domain.Apper, src, dst db.Session) *move {
	return &move{
		Apper:                  app,
		Src:                    src,
		Dst:                    dst,
		OrderMarkCodes:         make(map[int64]int64),
		OrderMarkAggregation:   make(map[int64]int64),
		OrderMarkAtk:           make(map[int64]int64),
		OrderMarkCommissioning: make(map[int64]int64),
		OrderMarkConnectingTap: make(map[int64]int64),
		OrderMarkUtilisation:   make(map[int64]int64),
	}
}
