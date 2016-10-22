package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
)

func main() {
	component := gr.New(new(rawValue))
	gr.RenderLoop(func() {
		component.Render("hbaseutils", gr.Props{})
	})
}

type rawValue struct {
	*gr.This
}

// Implements the StateInitializer interface.
func (c rawValue) GetInitialState() gr.State {
	return gr.State{
		"longValue": "1477161227964",
	}
}

// Implements the Renderer interface.
func (c rawValue) Render() gr.Component {
	longValue, err := strconv.ParseUint(c.State().String("longValue"), 10, 64)

	if err != nil {
		longValue = 0
	}

	var long2Hex = fmt.Sprintf("%016X", longValue)

	var prettyLongBytes = ""

	for k, v := range long2Hex {
		if k%2 == 1 {
			prettyLongBytes = fmt.Sprintf("%s%s", prettyLongBytes, string(v))
		} else {
			prettyLongBytes = fmt.Sprintf("%s\\x%s", prettyLongBytes, string(v))
		}
	}

	elem := el.Div(
		el.Div(
			gr.Text(fmt.Sprint("Current time in millis: ", time.Now().UnixNano()/1000000)),
		),
		el.Div(
			el.Break(),
			gr.Text("Enter long value:"),
			el.Input(
				gr.Style("width", "400px"),
				gr.Prop("type", "text"),
				gr.Prop("maxLength", "19"),
				gr.Prop("value", longValue),
				evt.Change(func(event *gr.Event) {
					newValue := event.TargetValue()
					c.SetState(gr.State{"longValue": newValue})
				}),
			),
			el.Break(),
			gr.Text(fmt.Sprintf("Result: %s", prettyLongBytes)),
		),
	)

	return elem
}

// Implements the ShouldComponentUpdate interface.
func (c rawValue) ShouldComponentUpdate(next gr.Cops) bool {
	return c.State().HasChanged(next.State, "longValue")
}
