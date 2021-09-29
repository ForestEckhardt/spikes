package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

func main() {
	file, err := os.Open("../ForestEckhardt/bellsoft-liberica/buildpack.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var envVars struct {
		Metadata struct {
			Configurations []struct {
				Name        string `toml:"name"`
				Description string `toml:"description"`
				Default     string `toml:"default"`
				Build       bool   `toml:"build"`
				Launch      bool   `toml:"launch"`
			} `toml:"configurations"`
		} `toml:"metadata"`
	}

	err = toml.NewDecoder(file).Decode(&envVars)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range envVars.Metadata.Configurations {
		str := fmt.Sprintf("%s\n%s\n", e.Name, e.Description)

		if e.Default != "" {
			str = fmt.Sprintf("%sDefault Value:%s\n", str, e.Default)
		}

		if e.Build {
			str = fmt.Sprintf("%sThis environment variable is used during build\n", str)
		}

		if e.Launch {
			str = fmt.Sprintf("%sThis environment variable is used during launch\n", str)
		}

		fmt.Println(str)
	}
}
