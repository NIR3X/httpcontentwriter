# HTTP Content Writer - Custom http.ResponseWriter for Dynamic Content-Type

The `httpcontentwriter` package provides a custom `http.ResponseWriter` implementation that allows you to set the content type on the first write. This can be useful in scenarios where you want to dynamically determine the content type based on the content being written.

## Installation

You can install the package using the `go get` command:

```bash
go get -u github.com/NIR3X/httpcontentwriter
```

## Usage

```go
package main

import (
	"net/http"
	"github.com/NIR3X/httpcontentwriter"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a new HttpContentWriter instance
	contentWriter := httpcontentwriter.NewHttpContentWriter(w)

	// Your logic here, writing content to the response using contentWriter
	contentWriter.Write([]uint8("Hello, World!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

## License
[![GNU AGPLv3 Image](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.html)  

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
