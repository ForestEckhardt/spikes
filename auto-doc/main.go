package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
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

	c, err := os.ReadFile(filepath.Join(buildpackDir, "README.md"))
	if err != nil {
		log.Fatal(err)
	}

	content := string(c)

	re := regexp.MustCompile(`\n+## (.+)\n+`)

	headings := re.FindAllStringSubmatch(content, -1)

	yamlFile, err := os.Open(filepath.Join(buildpackDir, ".docs.yml"))
	if err != nil {
		log.Fatal(err)
	}
	defer yamlFile.Close()

	var docsYAML struct {
		Exclude []string `yaml:"exclude"`
	}

	err = yaml.NewDecoder(yamlFile).Decode(&docsYAML)
	if err != nil {
		log.Fatal(err)
	}

	for i, h := range headings {
		shouldExclude := func() bool {
			for _, e := range docsYAML.Exclude {
				if h[1] == e {
					return true
				}
			}
			return false
		}

		if shouldExclude() {
			continue
		}

		var tre *regexp.Regexp
		if i < len(headings)-1 {
			tre = regexp.MustCompile(fmt.Sprintf(`\n+(## %s\n+[\s\S]*\n+)## %s\n+`, h[1], headings[i+1][1]))
		} else {
			tre = regexp.MustCompile(fmt.Sprintf(`\n+(## %s\n+[\s\S]*)$`, h[1]))
		}
		temp := tre.FindAllStringSubmatch(string(content), -1)
		if len(temp) > 0 {
			output += temp[0][1]
		}
	}

	fmt.Println(output)
}
