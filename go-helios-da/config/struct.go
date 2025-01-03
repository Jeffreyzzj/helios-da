package config

type TomlConfig struct {
	Title            string
	Port             string           `toml:"port"`
	LogInfoPath      string           `toml:"log_level_info"`
	HeliosInitConfig HeliosInitConfig `toml:"helios_init_config"`
}

type HeliosInitConfig struct {
	Title        string
	IndexConfigs map[string]IndexConf `toml:"index_conf"`
	UpdateTime   int                  `toml:"update_time"`
}

type IndexConf struct {
	Conf     string `toml:"conf"`
	DataConf string `toml:"data_conf"`
}
