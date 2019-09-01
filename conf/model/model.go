package model

type Config struct {
	Info struct {
		Name string
		URI  string
	} `toml:"-" json:"-" yaml:"-"`

	General struct {
		LoggerLevel string `toml:"logger_level"`
	} `toml:"general"`

	Listener []struct {
		// Typ: eg. [http, https, socks]
		Typ string `toml:"typ"`
		// Addr: e.g. [":8080", "[::1]:8080", "192.168.1.23:8080"]
		Addr string `toml:"addr"`
		// Params of listener: e.g. {"auth_type": "basic", "password": "password", "user": "user name"}
		Params map[string]string `toml:"params"`
	} `toml:"listener"`

	Controller struct {
		// Addr: e.g. [":8080", "[::1]:8080", "192.168.1.23:8080"]
		Addr string `toml:"addr"`
		// Params of listener: e.g. {"auth_type": "basic", "password": "password", "user": "user name"}
		Params map[string]string `toml:"params"`
	}

	DNS struct {
		IncludeSystem bool     `toml:"include_system"`
		Servers       []string `toml:"servers"`
		TimeoutSec    int      `toml:"timeout_sec"`
		Mapping       []struct {
			Domain string   `toml:"domain"`
			IP     []string `toml:"ip"`
			Server []string `toml:"server"`
		} `toml:"mapping"`
	} `toml:"dns"`

	Server map[string]struct {
		Typ    string            `toml:"typ"`
		Host   string            `toml:"host"`
		Port   int               `toml:"port"`
		Params map[string]string `toml:"params"`
	} `toml:"server"`

	ServerGroup map[string]struct {
		// Typ: e.g. ["rtt", "select"]
		Typ string `toml:"typ"`
		// Servers: e.g. in {Server..., ServerGroup...}
		Servers []string          `toml:"servers"`
		Params  map[string]string `toml:"params"`
	} `toml:"server_group"`

	Rule []struct {
		Typ    string            `toml:"typ"`
		Value  string            `toml:"value"`
		Proxy  string            `toml:"proxy"`
		Params map[string]string `toml:"params"`
	} `toml:"rule"`

	Include []struct {
		// e.g. {typ = local, path = "/User/root/config/server.toml"}
		Typ    string            `toml:"typ"`
		Params map[string]string `toml:"params"`
	} `toml:"include"`

	Filter []struct {
		Typ    string            `toml:"typ"`
		Params map[string]string `toml:"params"`
	} `toml:"filter"`

	Stream struct {
		Before []struct {
			Typ    string            `toml:"typ"`
			Params map[string]string `toml:"params"`
		} `toml:"before"`
		After []struct {
			Typ    string            `toml:"typ"`
			Params map[string]string `toml:"params"`
		} `toml:"after"`
	} `toml:"stream"`

	Plugins map[string]map[string]string `json:"plugins"`
}
