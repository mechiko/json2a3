package config

var TomlConfig = []byte(`
# This is a TOML document.

[layouts]
timelayout = "2006-01-02T15:04:05-0700"
timelayoutclear = "2006.01.02 15:04:05"
timelayoutday = "2006.01.02"
timelayoututc = "2006-01-02T15:04:05"

[src]
dbname = ''
driver = 'sqlite'
file = ''

[dst]
dbname = ''
driver = 'mssql'
file = ''
user = ''
pass = ''
port = '1433'
host = '127.0.0.1'

`)

type Configuration struct {
	File   string `mapstructure:"file"`
	DB     string `mapstructure:"db"`
	Order  string `mapstructure:"order"`
	Serial int    `mapstructure:"serial"`
	Gtin   string `mapstructure:"gtin"`

	Layouts LayoutConfiguration `mapstructure:"layouts" json:"layouts"` // описатели БД рефактор
}

type LayoutConfiguration struct {
	TimeLayout      string `mapstructure:"timelayout" json:"timelayout"`
	TimeLayoutClear string `mapstructure:"timelayoutclear" json:"timelayoutclear"`
	TimeLayoutDay   string `mapstructure:"timelayoutday" json:"timelayoutday"`
	TimeLayoutUTC   string `mapstructure:"timelayoututc" json:"timelayoututc"`
}
