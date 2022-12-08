package helper

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"snake_ladder_game/inputs"
)

func base_dir() string {

	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return mydir
}

func InputReader(filename string) *inputs.RawInputs{
	var path string = base_dir()
	path = filepath.Join(path, filename)
	
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return inputs.InitInput(string(content))
}
