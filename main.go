package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

func main() {
	var (
		templateFilePath string
		inputYAML        string
	)

	flag.StringVar(&templateFilePath, "template", "", "模板文件路径")
	flag.Parse()
	if templateFilePath == "" {
		println("resume-generator -template <template.tpl> <resume.yaml>")
		flag.Usage()
		log.Fatal("必须提供 -template 选项")
	}

	inputYAML = flag.Arg(0)
	if inputYAML == "" {
		println("resume-generator <resume.yaml>")
		flag.Usage()
		log.Fatal("必须提供简历数据 yaml")
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
	defer watcher.Close()

	err = watcher.Add(inputYAML)
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(templateFilePath)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-watcher.Errors
			if err != nil {
				log.Printf("fsnotify 错误: %s", err.Error())
			}
		}
	}()

	http.Handle("/", http.RedirectHandler("/resume", http.StatusPermanentRedirect))
	http.Handle("/resume", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl := template.New("index.html")

		bytes, err := os.ReadFile(templateFilePath)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("template read error: %v", err)))
			return
		}

		tpl, err = tpl.Parse(string(bytes))
		if err != nil {
			w.Write([]byte(fmt.Sprintf("template parse error: %v", err)))
			return
		}

		yml, err := os.Open(inputYAML)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("yaml read error: %v", err)))
			return
		}

		data := make(map[string]interface{})
		err = yaml.NewDecoder(yml).Decode(&data)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("yaml decode error: %v", err)))
			return
		}

		err = tpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("template execute error: %v", err)))
			return
		}
	}))

	log.Printf("服务启动于 http://localhost:8889")
	err = http.ListenAndServe(":8889", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}
