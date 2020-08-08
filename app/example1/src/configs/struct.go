package configs

const defaultFilename = "config"

type C struct {
	Key1   string   `yaml:"key1"`
	Key2   string   `yaml:"key2"`
	Groups []CGroup `yaml:"groups"`
}

type CGroup struct {
	Key1   string   `yaml:"key1"`
	Key2   string   `yaml:"key2"`
}

