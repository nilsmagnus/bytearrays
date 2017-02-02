package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
	"strings"
	"encoding/hex"
	"regexp"
)

func decodeLongFromHex(hex string) uint64 {
	purified := strings.TrimLeft(strings.Replace(hex, "\\x", "", 100), "0")
	reg := regexp.MustCompile("[^a-fA-F0-9]")
	for _, nonHexLetter := range reg.FindAllString(purified, 100) {
		purified = strings.Replace(purified, string(nonHexLetter), fmt.Sprintf("%x", nonHexLetter), 1)
	}
	res, err := strconv.ParseUint(fmt.Sprint("0x", purified), 0, 64)
	if err != nil {
		println(err.Error())
		return 0
	}
	return res
}

func decodeFromStringHex(hexValue string) string {
	purified := strings.TrimLeft(strings.Replace(hexValue, "\\x", "", 100), "0")
	bs, err := hex.DecodeString(purified)
	if err != nil {
		return hexValue
	}
	return string(bs)
}

func encodeFromLong(longValue int64) string {
	prettyLongBytes := ""
	long2Hex := fmt.Sprintf("%016X", longValue)
	for k, v := range long2Hex {
		if k%2 == 1 {
			prettyLongBytes = fmt.Sprintf("%s%s", prettyLongBytes, string(v))
		} else {
			prettyLongBytes = fmt.Sprintf("%s\\x%s", prettyLongBytes, string(v))
		}
	}

	return prettyLongBytes
}

func encodeFromString(stringValue string) string {
	var prettyLongBytes = ""

	var long2Hex = fmt.Sprintf("%016X", stringValue)
	for k, v := range long2Hex {
		if k%2 == 1 {
			prettyLongBytes = fmt.Sprintf("%s%s", prettyLongBytes, string(v))
		} else {
			prettyLongBytes = fmt.Sprintf("%s\\x%s", prettyLongBytes, string(v))
		}
	}

	return prettyLongBytes
}

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
	fmt.Println("Get Initial State")
	return gr.State{
		"longValue":   "1477161227964",
		"stringValue": "A nice string",
		"stringAsHex": "\\x61",
		"longAsHex":   "\\x00\\x00\\x01\\x57\\xED\\xAB\\x96\\xBC",
	}
}

// Implements the Renderer interface.
func (c rawValue) Render() gr.Component {
	longValue, err := strconv.ParseInt(c.State().String("longValue"), 10, 64)
	longAsHex := c.State().String("longAsHex")
	stringAsHex := c.State().String("stringAsHex")
	stringValue := c.State().String("stringValue")

	if err != nil {
		println(err.Error())
		longValue = 0
	}

	prettyLongBytes := encodeFromLong(longValue)
	prettyStringBytes := encodeFromString(stringValue)
	decodedStringValue := decodeFromStringHex(stringAsHex)
	decodedLongValue := decodeLongFromHex(longAsHex)

	elem := el.Div(
		el.Div(
			gr.Text(fmt.Sprintf("Current time in millis: %d (%v)", time.Now().UnixNano()/1000000, time.Now())),
		),
		el.Break(),
		el.Div(
			gr.Text("Long to hbase hex:"),
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
			gr.Text(fmt.Sprintf("=> %s", prettyLongBytes)),
		),
		el.Break(),
		el.Div(
			gr.Text("Hbase hex to long:"),
			el.Input(
				gr.Style("width", "400px"),
				gr.Prop("type", "text"),
				gr.Prop("value", longAsHex),
				evt.Change(func(event *gr.Event) {
					newValue := event.TargetValue()
					c.SetState(gr.State{"longAsHex": newValue})
				}),
			),
			gr.Text(fmt.Sprintf("=> %d", decodedLongValue)),
		),
		el.Break(),
		el.Div(
			gr.Text("String to hex:"),
			el.Input(
				gr.Style("width", "400px"),
				gr.Prop("type", "text"),
				gr.Prop("value", stringValue),
				evt.Change(func(event *gr.Event) {
					newValue := event.TargetValue()
					c.SetState(gr.State{"stringValue": newValue})
				}),
			),
			gr.Text(fmt.Sprintf("=> %v", prettyStringBytes)),
		),
		el.Break(),
		el.Div(
			gr.Text("Hex to string:"),
			el.Input(
				gr.Style("width", "400px"),
				gr.Prop("type", "text"),
				gr.Prop("value", stringAsHex),
				evt.Change(func(event *gr.Event) {
					newValue := event.TargetValue()
					c.SetState(gr.State{"stringAsHex": newValue})
				}),
			),
			gr.Text(fmt.Sprintf("=> %v", decodedStringValue)),
		),
	)

	return elem
}

// Implements the ShouldComponentUpdate interface.
func (c rawValue) ShouldComponentUpdate(next gr.Cops) bool {
	return c.State().HasChanged(next.State, "longValue") ||
		c.State().HasChanged(next.State, "longAsHex") ||
		c.State().HasChanged(next.State, "stringAsHex") ||
		c.State().HasChanged(next.State, "stringValue")
}
