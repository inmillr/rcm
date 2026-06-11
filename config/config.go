package config

import "flag"

type Config struct {
	InputFile  string
	OutputFile string
}

func Load() Config {
	in := flag.String("in", "claims.csv", "input csv")
	out := flag.String("out", "output.csv", "output csv")

	flag.Parse()

	return Config{
		InputFile:  *in,
		OutputFile: *out,
	}
}
