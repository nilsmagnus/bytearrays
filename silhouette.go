package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

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
	rawValue := c.State().String("rawValue")
	hexOfLong := c.State().String("hexOfLong")
	longValue, err := strconv.ParseUint(c.State().String("longValue"), 10, 64)

	if err != nil {
		longValue = 0
	}

	hexOfLong1 := strings.Replace(hexOfLong, "\\x", "", 20)
	hexOfLong1 = strings.Replace(hexOfLong1, "0", "", 20)
	longFromHex, err1 := strconv.ParseUint(hexOfLong1, 16, 64)

	if err1 != nil {
		longFromHex = 0
		println(err1)
	}

	var prettyStringBytes = ""

	for _, b := range []byte(rawValue) {
		prettyStringBytes = fmt.Sprintf("%s\\x%0X", prettyStringBytes, b)
	}

	var long2Hex = fmt.Sprintf("%X", longValue)

	var prettyLongBytes = ""
	for i := 0; i < 17; i++ {
		var v = ""
		if len(long2Hex) < i {
			v = string('0')
			println(v)
			if i == 0 {
				prettyLongBytes = fmt.Sprintf("%s%s", v, prettyLongBytes)
			} else if i%2 == 0 {
				prettyLongBytes = fmt.Sprintf("%s%s", v, prettyLongBytes)
			} else {
				prettyLongBytes = fmt.Sprintf("%s\\x%s", v, prettyLongBytes)
			}
		} else {
			r, _ := utf8.DecodeRuneInString(string(long2Hex[i]))
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				v = string(long2Hex[i])
				if i == 0 {
					prettyLongBytes = fmt.Sprintf("%s%s", prettyLongBytes, v)
				} else if i%2 == 0 {
					prettyLongBytes = fmt.Sprintf("%s%s", prettyLongBytes, v)
				} else {
					prettyLongBytes = fmt.Sprintf("%s\\x%s", prettyLongBytes, v)
				}
			}
		}
	}

	elem := el.Div(
		el.Div(
			gr.Text("String: "),
			el.Input(
				gr.Style("width", "400px"),
				gr.Prop("type", "text"),
				gr.Prop("value", rawValue),
				evt.Change(func(event *gr.Event) {
					newValue := event.TargetValue()
					c.SetState(gr.State{"rawValue": newValue})
				}),
			),
			el.Break(),
			gr.Text(prettyStringBytes),
		),
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

		el.Div(
			el.Break(),
			el.Break(),
			gr.Text("Hex to long"),
			el.Break(),
			gr.Text("Value:"),
			el.Input(
				gr.Style("width", "400px"),
				gr.Prop("type", "text"),
				gr.Prop("maxLength", "19"),
				gr.Prop("value", hexOfLong),
				evt.Change(func(event *gr.Event) {
					newValue := event.TargetValue()
					c.SetState(gr.State{"hexOfLong": newValue})
				}),
			),
			el.Break(),
			gr.Text(longFromHex),
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
