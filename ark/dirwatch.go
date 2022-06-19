package ark

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func (h *Handler) dirWatch(messageChan chan Message) {
	for {
		select {
		case event, ok := <-h.dirWatcher.Events:
			if !ok {
				close(messageChan)
				h.Close()
				return
			}

			if event.Op&fsnotify.Create == 0 {
				continue
			}

			if !strings.Contains(event.Name, "ServerGame.") && !strings.HasSuffix(event.Name, ".log") {
				continue
			}

			h.watcher.Remove(h.currentLogPath)
			h.currentLogPath = event.Name
			h.watcher.Add(h.currentLogPath)

		case err, ok := <-h.dirWatcher.Errors:
			if !ok {
				close(messageChan)
				h.Close()
				return
			}

			if err != nil {
				log.Printf("Error at waching directory: %v\n", err.Error())
				return
			}
		}
	}
}
