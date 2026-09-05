package Pkg

import (
	"fmt"
	"time"
)

type Event struct {
	Text string
	Err  string
	Time time.Time
}

func (a *Alltasks) Addevent(Text string, Err string) {
	eventt := Event{Text: Text, Err: Err, Time: time.Now()}
	a.Events = append(a.Events, eventt)

}

func (a *Alltasks) Allevents() {
	fmt.Println("All events:")
	for i := range a.Events {
		fmt.Printf(" %s\n Text: %s - Err: %s\n", a.Events[i].Time.Format("15:04 02.01.2006"), a.Events[i].Text, a.Events[i].Err)
		fmt.Println(" ===================================")
	}

}
