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
	// Hostname string `json:"hostname"`
	// HostPort string `json:"hostport"`
	// Output   string
	// Export   string
	// Browser  string `json:"browser"`
	File   string `mapstructure:"file"`
	DB     string `mapstructure:"db"`
	Order  string `mapstructure:"order"`
	Serial int    `mapstructure:"serial"`
	Gtin   string `mapstructure:"gtin"`

	// Application AppConfiguration    `json:"application"`
	Layouts LayoutConfiguration `json:"layouts"`
	// описатели БД рефактор
	// Src DatabaseConfiguration `json:"src"`
	// Dst DatabaseConfiguration `json:"dst"`
	// Config    DatabaseConfiguration `json:"config"`
	// AlcoHelp3 DatabaseConfiguration `json:"alcohelp3"`
	// TrueZnak  DatabaseConfiguration `json:"trueznak"`
	// SelfDB    DatabaseConfiguration `json:"selfdb"`
	// описание клиента ЧЗ
	// TrueClient TrueClientConfig `json:"trueclient"`
}

type LayoutConfiguration struct {
	TimeLayout      string `json:"timelayout"`
	TimeLayoutClear string `json:"timelayoutclear"`
	TimeLayoutDay   string `json:"timelayoutday"`
	TimeLayoutUTC   string `json:"timelayoututc"`
}

type DatabaseConfiguration struct {
	Connection string `json:"connection"`
	Driver     string `json:"driver"`
	DbName     string `json:"dbname"`
	File       string `json:"file"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Host       string `json:"host"`
	Port       string `json:"port"`
}

type AppConfiguration struct {
	// Pwd          string `json:"pwd"`
	// Console      bool   `json:"console"`
	// Disconnected bool   `json:"disconnected"`
	Fsrarid string `json:"fsrarid"`
	// DbType       string `json:"dbtype"`
	License string `json:"license"`
	// ScanTimer    int    `json:"scantimer"`
	StartPage string `json:"startpage"`
}

type TrueClientConfig struct {
	Test        bool   `json:"test"`
	StandGIS    string `json:"standgis"`
	StandSUZ    string `json:"standsuz"`
	TestGIS     string `json:"testgis"`
	TestSUZ     string `json:"testsuz"`
	TokenGIS    string `json:"tokengis"`
	TokenSUZ    string `json:"tokensuz"`
	AuthTime    string `json:"authtime"`
	LayoutUTC   string `json:"layoututc"`
	HashKey     string `json:"hashkey"`
	DeviceID    string `json:"deviceid"`
	OmsID       string `json:"omsid"`
	UseConfigDB bool   `json:"useconfigdb"`
}
