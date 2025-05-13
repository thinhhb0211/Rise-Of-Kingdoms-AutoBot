package bot_related

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/electricbubble/gadb"
)

type Bot struct {
	device gadb.Device
	// gui            *GuiDetector
	// config         *BotConfig
	resolution     [2]int
	currTask       tasks.TaskName
	roundCount     int
	currThread     *sync.WaitGroup
	stopChan       chan struct{}
	textUpdate     func(string)
	buildingUpdate func(map[string]interface{})
	configUpdate   func(map[string]interface{})

	// Tasks
	Tasks map[string]tasks.Task
}

func NewBot(device gadb.Device, config *BotConfig) *Bot {
	b := &Bot{
		device:     device,
		gui:        NewGuiDetector(device),
		config:     config,
		stopChan:   make(chan struct{}),
		currTask:   tasks.Break,
		Tasks:      make(map[string]tasks.Task),
		textUpdate: func(s string) {}, // default no-op
	}

	// Init tasks
	b.Tasks["collecting"] = NewCollectingTask(b)
	b.Tasks["alliance"] = NewAllianceTask(b)
	// ... other tasks
	return b
}

func (b *Bot) Start(taskFn func()) {
	// Stop existing thread
	b.Stop()

	var wg sync.WaitGroup
	wg.Add(1)
	b.currThread = &wg

	go func() {
		defer wg.Done()
		taskFn()
	}()

	go b.daemon(taskFn)
}

func (b *Bot) Stop() {
	close(b.stopChan)
	if b.currThread != nil {
		b.currThread.Wait()
	}
}

func (b *Bot) DoTaskLoop() {
	tasksList := []struct {
		Key        string
		EnableFunc func() bool
	}{
		{"collecting", b.config.EnableCollecting},
		{"alliance", b.config.EnableAlliance},
	}

	for {
		select {
		case <-b.stopChan:
			fmt.Println("Stopping task loop")
			return
		default:
			rand.Shuffle(len(tasksList), func(i, j int) {
				tasksList[i], tasksList[j] = tasksList[j], tasksList[i]
			})

			for _, t := range tasksList {
				if t.EnableFunc() {
					if task, ok := b.Tasks[t.Key]; ok {
						b.currTask = tasks.TaskName(task.Do())
					}
				}
			}

			b.roundCount++
			time.Sleep(2 * time.Second)
		}
	}
}

func (b *Bot) daemon(fn func()) {
	for {
		select {
		case <-b.stopChan:
			return
		default:
			time.Sleep(60 * time.Second)
			found := b.gui.CheckVerificationTitle()
			if found {
				ok := b.gui.CheckVerificationButton()
				if !ok {
					b.Stop()
					time.Sleep(1 * time.Second)
					b.Start(fn)
				}
			}
		}
	}
}
