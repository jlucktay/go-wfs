///////////////////////////////////////////////////////////////////////////////
//
// The MIT License (MIT)
// Copyright (c) 2018 Jivan Amara
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.
//
///////////////////////////////////////////////////////////////////////////////

// go-wfs project root.go

package wfs3

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
)

// checkOnly indicates that the caller doesn't care about the content, only the contentId
// contentId is a string that changes as the content changes, useful for ETag & caching.
func Root(serveAddress string, checkOnly bool) (content *RootContent, contentId string) {
	hasher := fnv.New64()
	contentId = hex.EncodeToString(hasher.Sum([]byte(serveAddress)))
	if checkOnly {
		return nil, contentId
	}

	apiUrl := fmt.Sprintf("http://%v/api", serveAddress)
	conformanceUrl := fmt.Sprintf("http://%v/conformance", serveAddress)
	collectionsUrl := fmt.Sprintf("http://%v/collections", serveAddress)
	selfUrl := fmt.Sprintf("http://%v/", serveAddress)

	content = &RootContent{
		Links: []*Link{
			{
				Href: selfUrl,
				Rel:  "self",
			},
			{
				Href: apiUrl,
				Rel:  "service",
			},
			{
				Href: conformanceUrl,
				Rel:  "conformance",
			},
			{
				Href: collectionsUrl,
				Rel:  "data",
			},
		},
	}

	return content, contentId
}
