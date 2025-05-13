package tasks

import (
	"Rise-Of-Kingdoms-AutoBot/bot_related"
	"Rise-Of-Kingdoms-AutoBot/pkg/filepath"
	"fmt"
	"time"

	"github.com/electricbubble/gadb"
)

type Task struct {
	Bot    *bot_related.Bot
	Device gadb.Device
	// GUI    *GuiDetector
}

func NewTask(bot *Bot) *Task {
	return &Task{
		Bot:    bot,
		Device: bot.Device,
		// GUI:    bot.Gui,
	}
}

func (t *Task) CallIdleBack() {
	t.SetText("call back idle commander")
	t.BackToMapGUI()

	for {
		if _, _, pos := t.GUI.CheckAny(filepath.HOLD_ICON_SMALL_IMAGE_PATH); pos != nil {
			x, y := pos[0], pos[1]
			t.Tap(x-10, y-10, 2)
			t.Tap(640, 360, 2)
		} else {
			break
		}

		if _, _, pos := t.GUI.CheckAny(filepath.RETURN_BUTTON_IMAGE_PATH); pos != nil {
			t.Tap(pos[0], pos[1], 1)
		} else {
			break
		}
	}
}

func (t *Task) HealTroops() {
	t.SetText("Heal Troops")
	healButton := [2]int{960, 590}

	t.BackToHomeGUI()
	t.HomeGUIFullView()
	t.Tap(t.Bot.BuildingPos["hospital"][0], t.Bot.BuildingPos["hospital"][1], 2)
	t.Tap(285, 20, 500)
	if _, _, pos := t.GUI.CheckAny(filepath.HEAL_ICON_IMAGE_PATH); pos != nil {
		t.Tap(pos[0], pos[1], 2)
		t.Tap(healButton[0], healButton[1], 2)
	}
	t.Tap(t.Bot.BuildingPos["hospital"][0], t.Bot.BuildingPos["hospital"][1], 4)
}

func (t *Task) Tap(x, y int, sleepMillis int, longPressMillis ...int) {
	if len(longPressMillis) > 0 && longPressMillis[0] > -1 {
		t.Device.RunShellCommand(fmt.Sprintf("input swipe %d %d %d %d %d", x, y, x, y, longPressMillis[0]))
		time.Sleep(time.Duration(longPressMillis[0]+200) * time.Millisecond)
	} else {
		t.Device.RunShellCommand(fmt.Sprintf("input tap %d %d", x, y))
		time.Sleep(time.Duration(sleepMillis) * time.Millisecond)
	}
}

func (t *Task) Swipe(x1, y1, x2, y2, times, duration int) {
	cmd := fmt.Sprintf("input swipe %d %d %d %d %d", x1, y1, x2, y2, duration)
	for i := 0; i < times; i++ {
		t.Device.RunShellCommand(cmd)
		time.Sleep(time.Duration(duration+200) * time.Millisecond)
	}
}

func (t *Task) BackToMapGUI() int {
	loop := 0
	for {
		guiName, pos := t.GetCurrentGUIName()
		switch guiName {
		case "MAP":
			return loop
		case "HOME":
			t.Tap(pos[0], pos[1], 500)
		default:
			t.Back()
		}
		loop++
		time.Sleep(500 * time.Millisecond)
	}
}

func (t *Task) BackToHomeGUI() int {
	loop := 0
	for {
		guiName, pos := t.GetCurrentGUIName()
		switch guiName {
		case "HOME":
			return loop
		case "MAP":
			t.Tap(pos[0], pos[1], 500)
		default:
			t.Back()
		}
		loop++
		time.Sleep(500 * time.Millisecond)
	}
}

func (t *Task) Back() {
	t.Device.RunShellCommand("input keyevent 4")
	time.Sleep(500 * time.Millisecond)
}

func (t *Task) SetText(text string) {
	fmt.Println("[Task] ", text)
}

func (t *Task) GetCurrentGUIName() (string, [2]int) {
	// Mock implementation
	return "UNKNOW", [2]int{0, 0}
}

func (t *Task) HomeGUIFullView() {
	t.Tap(60, 540, 500)
	t.Tap(1105, 200, 1000)
	t.Tap(1220, 35, 2000)
}

func (t *Task) HasBuff(checkingLocation string, buffImageProps filepath.ImageProps) bool {
	if checkingLocation == filepath.HOME {
		t.BackToHomeGUI()
	} else if checkingLocation == filepath.MAP {
		t.BackToMapGUI()
	} else {
		return false
	}

	return true
}

func (t *Task) UseItem(usingLocation string, itemImagePropsList []filepath.ImageProps) bool {
	if usingLocation == filepath.HOME {
		t.BackToHomeGUI()
	} else if usingLocation == filepath.MAP {
		t.BackToMapGUI()
	} else {
		return false
	}

	itemsIconPos := [2]int{930, 675}
	
	return true
}
