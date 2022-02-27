package main

import "fmt"

/*

A cookie/js-free proxy for imgur links, inspired by https://imgin.voidnet.tech/

Just replace the imgur hostname with the hostname of improxy in your browser.

Basic plan is the same as https://imgin.voidnet.tech/

- Support basic imgur paths like:
  - /gallery/YUJYQ
  - /YUJYQ
  - /a/YUJYQ
- Work by fetching and parsing ingur HTML, then caching images and other media locally.
- Require no cookies or javascript to function.

*/

func main() {
	fmt.Println("improxy")
}
