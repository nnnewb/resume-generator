package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var (
		templateFilePath string
		inputYAML        string
	)

	flag.StringVar(&templateFilePath, "template", "", "模板文件路径")
	flag.Parse()
	inputYAML = flag.Arg(0)
	if inputYAML == "" {
		log.Fatal("please provide rendering resume data in yaml format.")
		println("resume-generator <resume.yaml>")
		flag.Usage()
		return
	}

	templateFilePath = filepath.Clean(templateFilePath)
	if _, err := os.Stat(templateFilePath); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s: 文件不存在", templateFilePath)
			return
		}
	}

	inputYAML = filepath.Clean(inputYAML)
	if _, err := os.Stat(inputYAML); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s: 文件不存在", inputYAML)
			return
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	watcher.Add(inputYAML)
	watcher.Add(templateFilePath)

	http.Handle("/", http.RedirectHandler("/resume", 302))
	http.Handle("/resume", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
}
