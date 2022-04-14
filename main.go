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
	"github.com/nnnewb/resume-generator/pkg/livepreview"
	"gopkg.in/yaml.v3"
)

func main() {
	var (
		templateDirPath string
		inputYAML       string
	)

	flag.StringVar(&templateDirPath, "template", "", "模板文件夹路径")
	flag.Parse()
	if templateDirPath == "" {
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

	templateDirPath = filepath.Clean(templateDirPath)
	if _, err := os.Stat(filepath.Join(templateDirPath, "index.html.tpl")); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s/index.html.tpl: 文件不存在", templateDirPath)
			return
		}
		log.Fatal(err)
	}

	inputYAML = filepath.Clean(inputYAML)
	if _, err := os.Stat(inputYAML); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s: 文件不存在", inputYAML)
			return
		}
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := livepreview.WatchDir(templateDirPath, watcher); err != nil {
		log.Fatal(err)
	}

	if err = watcher.Add(inputYAML); err != nil {
		log.Fatal(err)
	}

	// redirect to /resume
	http.Handle("^/$", http.RedirectHandler("/resume", http.StatusPermanentRedirect))

	// serve websocket
	http.Handle("/ws", livepreview.MakeWebSocketHandler(watcher))

	// serve template static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(templateDirPath, "static")))))

	// render index.html.tpl with
	http.Handle("/resume", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl := template.New("index.html")

		bytes, err := os.ReadFile(filepath.Join(templateDirPath, "index.html.tpl"))
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
