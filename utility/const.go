package utility

var FenArt = `
______ _____ _   _ ______ ___________  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓█████▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
|  ___|  ___| \ | || ___ \_   _| ___ \ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▒▒▒▒██▓▓▓▓▓▓▓▓▓▓▓▓▓▓
| |_  | |__ |  \| || |_/ / | | | |_/ / ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▒▒██░░██▓▓▓▓▓▓▓▓▓▓▓▓
|  _| |  __|| . ` + "`" + ` ||    /  | | |    /  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▒▒██░░██▓▓▓▓▓▓▓▓▓▓▓▓
| |   | |___| |\  || |\ \ _| |_| |\ \  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▒▒██░░░░██▓▓▓▓▓▓▓▓▓▓
\_|   \____/\_| \_/\_| \_|\___/\_| \_| ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██████▒▒██░░░░▒▒██▓▓▓▓▓▓▓▓
                                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▒▒▒▒▒▒░░░░░░░░▒▒▒▒██▓▓▓▓▓▓
______________________________________ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▒▒▒▒░░░░░░░░░░░░░░▒▒██▓▓▓▓▓▓
Git: github.com/RDLrpl/fenrir          ██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░▒▒██▓▓▓▓▓▓
Version: 2.0.0                         ██████░░░░░░░░░░░░░░░░████▒▒░░░░░░░░░░▒▒██▓▓▓▓
V|CN: Gleipnir                         ▓▓██▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒██▓▓▓▓
  __  __     __  __                    ▓▓▓▓▓▓██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒██▓▓
 /__)/  )/  /__)/__)/                  ▓▓▓▓▓▓░░░░░░▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░▒▒██▓▓
/ ( /(_/(__/ ( /   (__                 ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒██░░░░░░░░░░░░░░░░░░░░▒▒██▓▓
______________________________________ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒██░░░░░░░░░░░░░░░░░░░░▒▒██▓▓ 
`

var Usage = `
emmm
`

var FenrirCAAUpack = map[string]string{
	"windows-amd": "https://storage.googleapis.com/chromium-browser-snapshots/Win_x64/1638973/chrome-win.zip",
	"linux-amd":   "https://storage.googleapis.com/chromium-browser-snapshots/Linux_x64/1638973/chrome-linux.zip",
}

type Configuration struct {
	Messages map[string]string `toml:"Messages"`
	Telegram TelegramConfig    `toml:"Telegram"`
	Discord  DiscordConfig     `toml:"Discord"`
}

type TelegramConfig struct {
	Sessions string                     `toml:"sessions"`
	Targets  map[string]string          `toml:"targets"`
	Accounts map[string]TelegramAccount `toml:"accounts"`
}

type TelegramAccount struct {
	API_id     string `toml:"api_id"`
	API_hash   string `toml:"api_hash"`
	Number     string `toml:"number"`
	CloudPass  string `toml:"cloudpass"`
	Proxy      string `toml:"proxy"`
	ProxyPass  string `toml:"proxypass"`
	ProxyLogin string `toml:"proxylogin"`
	Targ       string `toml:"targ"`
	Marg       string `toml:"marg"`
}

type DiscordConfig struct {
	Targets  map[string][]string       `toml:"targets"`
	Accounts map[string]DiscordAccount `toml:"accounts"`
}

type DiscordAccount struct {
	Token      string `toml:"token"`
	Proxy      string `toml:"proxy"`
	ProxyPass  string `toml:"proxypass"`
	ProxyLogin string `toml:"proxylogin"`
	Targ       string `toml:"targ"`
	Marg       string `toml:"marg"`
}
