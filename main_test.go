package main

import (
	"fmt"
	"testing"
)

func TestImgurGet(t *testing.T) {

	{
		path := "/49jzlTB"
		got := ImgurGet(path)
		expected := 1
		// fmt.Printf("%v\n", got)
		if len(got.Media) != expected {
			t.Errorf("Expected len=%v but got len=%v", expected, len(got.Media))
		}

	}
	{
		path := "/gallery/kaOZU"
		got := ImgurGet(path)
		expected := 10
		// fmt.Printf("%v\n", got)
		if len(got.Media) != expected {
			t.Errorf("Expected len=%v but got len=%v", expected, len(got.Media))
		}
	}
	{
		path := "/a/Bh7Sw"
		got := ImgurGet(path)
		expected := 1
		fmt.Printf("%v\n", got)
		if len(got.Media) != expected {
			t.Errorf("Expected len=%v but got len=%v", expected, len(got.Media))
		}
	}
}
