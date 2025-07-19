package common

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/pelletier/go-toml"
)

type Names struct {
	Stars []string `toml:"stars"`
}

var Lib Names
var reloads int = 0

func LoadData(ord int) Names {
	file, err := os.Open("common/names.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, &Lib)
	if err != nil {
		panic(err)
	}
	if ord > 0 {
		for i, n := range Lib.Stars {
			Lib.Stars[i] = fmt.Sprintf("%s-i", n)
		}
	}
	return Lib
}

func GetStarName() string {
	if len(Lib.Stars) == 0 {
		LoadData(reloads)
	}
	i := rand.Intn(len(Lib.Stars))
	name := Lib.Stars[i]
	Lib.Stars = append(Lib.Stars[:i], Lib.Stars[i+1:]...)
	return name
}
