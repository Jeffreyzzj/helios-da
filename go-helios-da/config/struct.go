package config

type TomlConfig struct {
	Title            string
	HeliosInitConfig HeliosInitConfig `toml:"helios_init_config"`
}

type HeliosInitConfig struct {
	Title        string
	IndexConfigs map[string]IndexConf `toml:"index_conf"`
}

type IndexConf struct {
	Conf     string `toml:"conf"`
	DataConf string `toml:"data_conf"`
}
