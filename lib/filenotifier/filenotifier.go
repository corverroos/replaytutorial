// filenotifier provides a reflex rsql events notifier using a local folder which provides responsive
// events between multiple local running processes.
package filenotifier

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
)

var folder = flag.String("notify_folder", "/tmp/replaynotify/", "folder to use for notification")

func New() (*Notifier, error) {
	err := os.MkdirAll(*folder, 0755)
	if err != nil {
		return nil, err
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = w.Add(*folder)
	if err != nil {
		return nil, err
	}

	n := &Notifier{}
	go func() {
		for {
			select {
			case err := <-w.Errors:
				log.Error(nil, errors.Wrap(err, "fsnotify"))
			case event := <-w.Events:
				if event.Op != fsnotify.Create {
					continue
				}

				n.mu.Lock()
				for _, ch := range n.chs {
					select {
					case ch <- struct{}{}:
					default:
					}
				}
				n.mu.Unlock()
			}
		}
	}()

	return n, nil
}

type Notifier struct {
	mu  sync.Mutex
	chs []chan struct{}
}

func (n *Notifier) C() <-chan struct{} {
	ch := make(chan struct{}, 1)
	n.mu.Lock()
	defer n.mu.Unlock()
	n.chs = append(n.chs, ch)
	return ch
}

func (n *Notifier) Notify() {
	now := fmt.Sprint(time.Now().UnixNano())
	file := path.Join(*folder, now)
	err := os.WriteFile(file, nil, 0644)
	if err != nil {
		panic(errors.Wrap(err, "notify to file", j.KS("file", file)))
	}
}
