package cfg

import "embed"

//go:embed config.toml
var ConfigFile embed.FS
