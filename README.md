urljoin
=======

Join URL parts into a single URL, correctly adding slashes. 
It doesn't add a trailing slash, but keeps one if present.

## Install

    go get github.com/ongoingio/urljoin
  
## Usage

    package main
    
    import (
        "fmt"
        "github.com/ongoingio/urljoin"
    )
    
    func main() {
        url := urljoin.Join("http://example.com", "foo")
        fmt.Println(url) // Output: http://example.com/foo
    }

## Tests

    go test
