package ark

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func (h *Handler) logWatch(messageChan chan Message) {
	for {
		select {
		case event, ok := <-h.watcher.Events:
			if !ok {
				close(messageChan)
				h.Close()
				return
			}

			if event.Op&fsnotify.Write == 0 {
				continue
			}

			line, err := h.readNewLine(h.currentLogPath)
			if err != nil {
				log.Printf("ReadLineError: %s\n", err.Error())
				continue
			}

			message, err := NewMessageFromLine(line)
			if err != nil && err != illigalFormatError {
				log.Printf("NewMessageFromLine: %s\n", err.Error())
				continue
			}

			messageChan <- *message

		case err, ok := <-h.watcher.Errors:
			log.Printf("Error at waching files: %v\n", err.Error())
			if !ok {
				close(messageChan)
				h.Close()
				return
			}
		}
	}
}
