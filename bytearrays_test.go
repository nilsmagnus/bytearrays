package main

import "testing"

func TestConvertLongToHex(t *testing.T){
	hex:=encodeFromLong(int64(123))

	if hex != "" {
		t.Error("Hex is not what I expected, but it was ", hex)
	}
}