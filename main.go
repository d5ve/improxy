package main

import "fmt"

/*

A cookie/js-free proxy for imgur links, inspired by https://imgin.voidnet.tech/

Just replace the imgur hostname with the hostname of improxy in your browser.

- Support basic imgur paths like:
	- /gallery/YUJYQ
	- /a/YUJYQ
	- /YUJYQ.jpg
- Work by fetching and parsing ingur HTML to get download links for the media.
- Then download the media, temporarily cache it, and serve simple HTML responses embedding the media.
- Require no cookies or javascript to function.

Steps:
	1. Parse imgur path from request.
	2. Fetch the HTML from imgur.
	3. Parse the HTML to get media download links.
	4. Download media and cache locally.
	5. Return a HTML response doc that embeds the local media.
	6. Serve local media requests.

Notes:
- The imgur HTML currently holds a JSON document which seems to enumerate all
the media, and all their titles/descriptions and other metadata. Parsing this
JSON looks the best first step towards getting a list of media to download.
- The proxy needs a way of mapping the ordered list of media and metadata with
the cached files, so the generated HTML response shows things in order and uses
the correct descriptions etc. Perhaps store a subset of the imgur JSON in the
cache as well, with some standard imgur path to filename mapping.
*/

func main() {
	fmt.Println("improxy")
}
