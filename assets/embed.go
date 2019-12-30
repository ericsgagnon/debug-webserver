// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(http.Dir("./assets"), vfsgen.Options{
		Filename:     "assets/assets.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "assets",
	})
	if err != nil {
		log.Fatalf("vfsgen failed: %s\n\n", err)
	}

}
