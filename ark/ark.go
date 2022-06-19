package ark

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
)

type Handler struct {
	dirpath        string
	currentLogPath string

	watcher    *fsnotify.Watcher
	dirWatcher *fsnotify.Watcher
}

func New(serverLogDirPath string) *Handler {
	return &Handler{
		dirpath: serverLogDirPath,
	}
}

func (h *Handler) Close() error {
	return h.watcher.Close()
}

func (h *Handler) Start() (chan Message, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.Wrap(err, "NewWatcher")
	}

	dirWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.Wrap(err, "NewWatcher")
	}

	h.watcher = watcher
	h.dirWatcher = dirWatcher

	dirWatcher.Add(h.dirpath)

	var messageChan = make(chan Message)

	logFile, err := h.findNewLog()
	if err != nil {
		return nil, errors.Wrap(err, "findNewLog")
	}

	h.currentLogPath = filepath.Join(h.dirpath, logFile)

	h.watcher.Add(h.currentLogPath)

	go h.dirWatch(messageChan)

	go h.logWatch(messageChan)

	return messageChan, nil
}
