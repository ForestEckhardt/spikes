package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type BuildpackTOML struct {
	Buildpack struct {
		Name string `toml:"name"`
	} `toml:"buildpack"`
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

func main() {
	var buildpackDir string
	flag.StringVar(&buildpackDir, "buildpack-dir", "", "buildpack directory")
	flag.Parse()

	if buildpackDir == "" {
		log.Fatal("required --buildpack-dir to be set")
	}

	file, err := os.Open(filepath.Join(buildpackDir, "buildpack.toml"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bpTOML BuildpackTOML

	err = toml.NewDecoder(file).Decode(&bpTOML)
	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("# %s\n\n## Environment Variable Configuration\n", bpTOML.Buildpack.Name)

	for _, e := range bpTOML.Metadata.Configurations {
		str := fmt.Sprintf("### %s\n%s\n", e.Name, e.Description)

		if e.Default != "" {
			str = fmt.Sprintf("%sDefault Value: `%s`\n", str, e.Default)
		}

		if e.Build {
			str = fmt.Sprintf("%sThis environment variable is used during build\n", str)
		}

		if e.Launch {
			str = fmt.Sprintf("%sThis environment variable is used during launch\n", str)
		}

		output += str + "\n"
	}

	output += "## Additional Reference Information\n"

	ref, err := os.ReadFile(filepath.Join(buildpackDir, "reference-doc.md"))
	if err != nil {
		log.Fatal(err)
	}

	output += string(ref)

	fmt.Println(output)
}
