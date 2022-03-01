package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lack/yamltrim"

	"gopkg.in/yaml.v3"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var original interface{}
	err = yaml.Unmarshal(input, &original)
	if err != nil {
		panic(err)
	}

	trimmed := yamltrim.YamlTrim(original)
	trimmedBytes, err := yaml.Marshal(trimmed)
	fmt.Printf("%s", trimmedBytes)
}
