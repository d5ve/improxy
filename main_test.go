package main

import (
	"fmt"
	"testing"
)

func TestImgurGet(t *testing.T) {

	path := "/49jzlTB"
	got := ImgurGet(path)

	fmt.Println(got)

	path = "/gallery/kaOZU"
	got = ImgurGet(path)

	fmt.Println(got)
}
