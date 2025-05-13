package tasks

import (
	"time"
)

type Alliance struct {
	*Task
}

func NewAlliance(bot *Bot) *Alliance {
	return &Alliance{
		Task: NewTask(bot),
	}
}

func (a *Alliance) Do(nextTask string) string {
	a.SetText("Alliance", true)

	allianceBtnPos := [2]int{1030, 670}
	names := []string{"GIFTS", "TERRITORY", "TECHNOLOGY"}

	for _, name := range names {
		a.SetTextInsert("Open alliance")
		a.BackToHomeGui()
		a.MenuShouldOpen(true)
		a.Tap(allianceBtnPos[0], allianceBtnPos[1], 3*time.Second)

		switch name {
		case "GIFTS":
			a.SetTextInsert("Claim gift")
			giftsPos := [2]int{885, 560}
			ratePos := [2]int{930, 205}
			normalPos := [2]int{670, 205}
			claimAllPos := [2]int{1110, 205}
			treasure := [2]int{330, 410}

			a.Tap(giftsPos[0], giftsPos[1], 2*time.Second)

			// Claim rate gift
			a.SetTextInsert("Claim rate gift")
			a.Tap(ratePos[0], ratePos[1], 1*time.Second)
			for i := 0; i < 20; i++ {
				_, _, pos := a.Gui.CheckAny(filepath.GIFTS_CLAIM_BUTTON_IMAGE_PATH)
				if pos == nil {
					break
				}
				a.Tap(pos[0], pos[1], 500*time.Millisecond)
			}

			// Claim normal gift
			a.SetTextInsert("Claim normal gift")
			a.Tap(normalPos[0], normalPos[1], 1*time.Second)
			a.Tap(claimAllPos[0], claimAllPos[1], 1*time.Second)

			// Claim treasure
			a.Tap(treasure[0], treasure[1], 1*time.Second)

		case "TERRITORY":
			a.SetTextInsert("Claim resource")
			territoryPos := [2]int{885, 405}
			claimPos := [2]int{1020, 140}
			a.Tap(territoryPos[0], territoryPos[1], 2*time.Second)
			a.Tap(claimPos[0], claimPos[1], 1*time.Second)

		case "TECHNOLOGY":
			a.SetTextInsert("Donate technology")
			technologyPos := [2]int{760, 560}
			a.Tap(technologyPos[0], technologyPos[1], 5*time.Second)

			_, _, recommendPos := a.Gui.CheckAny(filepath.TECH_RECOMMEND_IMAGE_PATH)
			if recommendPos != nil {
				a.Tap(recommendPos[0], recommendPos[1]+60, 1*time.Second)
				_, _, donateBtnPos := a.Gui.CheckAny(filepath.TECH_DONATE_BUTTON_IMAGE_PATH)
				if donateBtnPos != nil {
					for i := 0; i < 20; i++ {
						a.Tap(donateBtnPos[0], donateBtnPos[1], 30*time.Millisecond)
					}
				}
			} else {
				a.SetTextInsert("Cannot found Officer's Recommendation")
			}
		}
	}

	return nextTask
}
