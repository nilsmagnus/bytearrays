package main

import (
	"fmt"
	"strconv"

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
		"rawValue":      "Some string",
		"showPlainText": false,
		"longValue":     "1477161227964",
		"hexOfLong":     "\\x00\\x00\x01\\x57\\xED\\xAB\\x96\\xBC",
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
			el.Break(),
			el.Break(),
			gr.Text("Long to hex "),
			el.Break(),
			gr.Text("Value:"),
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
			gr.Text(prettyLongBytes),
		),
	)

	return elem
}

// Implements the ShouldComponentUpdate interface.
func (c rawValue) ShouldComponentUpdate(next gr.Cops) bool {
	return c.State().HasChanged(next.State, "rawValue") ||
		c.State().HasChanged(next.State, "longValue") ||
		c.State().HasChanged(next.State, "showPlainText")

}
