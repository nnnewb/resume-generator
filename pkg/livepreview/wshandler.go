package livepreview

import (
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

type M map[string]interface{}

func WatchDir(path string, watcher *fsnotify.Watcher) error {
	return filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		err = watcher.Add(filepath.Join(path, d.Name()))
		if err != nil {
			return err
		}

		return nil
	})
}

func MakeWebSocketHandler(watcher *fsnotify.Watcher) http.Handler {
	var upgrader = websocket.Upgrader{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error during protocol upgrade: %s", err.Error())
			return
		}
		defer conn.Close()

		c := make(chan struct{}, 1)
		go func() {
			defer func() {
				c <- struct{}{}
			}()

			for {
				// 心跳频率 1s
				err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
				if err != nil {
					log.Fatal(err)
				}

				mt, _, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsCloseError(err) {
						log.Printf("%s: websocket closed", conn.RemoteAddr().String())
					}
					return
				}

				if mt == websocket.CloseMessage {
					log.Printf("%s: websocket closed", conn.RemoteAddr().String())
					return
				}
			}
		}()

		for {
			select {
			case <-c:
				log.Printf("%s: connection closed", conn.RemoteAddr().String())
				return
			case evt := <-watcher.Events:
				switch evt.Op {
				case fsnotify.Remove:
					log.Printf("%s: removed", evt.Name)
					err = conn.WriteJSON(M{"reload": true})
					if err != nil {
						log.Printf("%s: write json fail: %s", conn.RemoteAddr().String(), err.Error())
						return
					}
				case fsnotify.Rename:
					log.Printf("%s: renamed", evt.Name)
					err = conn.WriteJSON(M{"reload": true})
					if err != nil {
						log.Printf("%s: write json fail: %s", conn.RemoteAddr().String(), err.Error())
						return
					}
				case fsnotify.Create:
					log.Printf("%s: created", evt.Name)
					err = conn.WriteJSON(M{"reload": true})
					if err != nil {
						log.Printf("%s: write json fail: %s", conn.RemoteAddr().String(), err.Error())
						return
					}
				case fsnotify.Write:
					log.Printf("%s: write", evt.Name)
					err = conn.WriteJSON(M{"reload": true})
					if err != nil {
						log.Printf("%s: write json fail: %s", conn.RemoteAddr().String(), err.Error())
						return
					}
				case fsnotify.Chmod:
					log.Printf("%s: chmod", evt.Name)
					err = conn.WriteJSON(M{"reload": true})
					if err != nil {
						log.Printf("%s: write json fail: %s", conn.RemoteAddr().String(), err.Error())
						return
					}
				default:
					log.Printf("%s: unknown op %d", evt.Name, evt.Op)
					return
				}
			case err := <-watcher.Errors:
				log.Printf("fsnotify error: %s", err.Error())
				return
			default:
				time.Sleep(time.Second)
			}
		}
	})
}
