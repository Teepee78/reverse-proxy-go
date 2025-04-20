package config

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	Config string
}

func GetFlags() Flags {
	var configPath string

	// Define flags
	flag.StringVar(&configPath, "c", "", "Path to config file (short)")
	flag.StringVar(&configPath, "config", "", "Path to config file (long)")

	flag.Parse()

	validateConfig(&configPath)

	return Flags{Config: configPath}
}

func validateConfig(path *string) {
	if flag.Lookup("c").Value.String() != "" && flag.Lookup("config").Value.String() != "" {
		_, err := fmt.Fprintln(os.Stderr, "Do not use both -c and --config flags together.")
		if err != nil {
			return
		}
		os.Exit(1)
	}

	configFromPosArg := flag.Arg(0)

	if *path == "" {
		if configFromPosArg == "" {
			fmt.Println("No config file provided")
			os.Exit(1)
		}
		*path = configFromPosArg
	}

	fmt.Printf("Using config file: %s\n", *path)
}
