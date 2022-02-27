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
- I'm going to parse with regex initially and see how far that gets me.
	`<script>window.postDataJSON="$JSON_DOC"</script>`
- Maybe keep track of the imgur data version number and have a parser for each version.
	`data-release="imgur@2.0.8"`
- bash for extracting the direct image URL from a single-image page:
	curl -s https://imgur.com/49jzlTB \
		| perl -ne 'm{(https://i.imgur.com/[^"]+)}xms; print $1'
- bash for extracting the JSON document from a gallery page:
	curl -s https://imgur.com/gallery/kaOZU \
		| perl -ne 'm{<script>window.postDataJSON="(.*?)"</script>}xms; print $1' \
		| sed 's/\\"/"/g; s/\\\\"/\\"/g' \
		| sed "s/\\\'/'/g" \
		| jq .
*/

func main() {
	fmt.Println("improxy")
}
