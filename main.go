package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/fsnotify/fsnotify"
	"github.com/nnnewb/resume-generator/pkg/livepreview"
	"github.com/nnnewb/resume-generator/pkg/utils"
	"gopkg.in/yaml.v2"
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
	log.Printf("使用模板: %s", templateDirPath)

	inputYAML = filepath.Clean(inputYAML)
	if _, err := os.Stat(inputYAML); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s: 文件不存在", inputYAML)
			return
		}
		log.Fatal(err)
	}
	log.Printf("使用简历数据: %s", inputYAML)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := livepreview.WatchDir(templateDirPath, watcher); err != nil {
		log.Fatalf("添加 fsnotify 监视失败, 错误 %+v", err)
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
		page, err := renderResumeTemplate(templateDirPath, inputYAML)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("render template error: %+v", err)))
			return
		}

		w.Write([]byte(page))
	}))

	log.Printf("服务启动于 http://localhost:8889")
	err = http.ListenAndServe(":8889", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}

func renderResumeTemplate(templateDirPath, resumeDataPath string) (string, error) {
	tb := utils.NewTemplateBuilder()
	tb.ExtendFuncs(sprig.HtmlFuncMap())
	tb.AddFunc("Markdown", utils.Markdown).AddFunc("Unescape", utils.Unescape)

	err := filepath.WalkDir(templateDirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".tpl" {
			tb.AddTemplateFile(path)
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	tb.SetMainTemplate("index.html.tpl")

	yml, err := os.Open(resumeDataPath)
	if err != nil {
		return "", err
	}

	data := make(map[string]interface{})
	err = yaml.NewDecoder(yml).Decode(&data)
	if err != nil {
		return "", err
	}

	tpl, err := tb.BuildHTMLTemplate()
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}
	err = tpl.Execute(&sb, data)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
