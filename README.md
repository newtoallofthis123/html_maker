# HTML Maker

The enigmatic way to write HTML, now in Go.

> This library is inspired by [html_tag](https://github.com/newtoallofthis123/html_tag), which does a similar thing in Rust.

## Installation

```bash
go get github.com/newtoallofthis123/html_maker
```

## When and why to use this library

This library is basically useful if you want to avoid writing Raw HTML strings in your Go code.
Using this library to write HTML reduces the use of raw strings, replacing them with understandable way of writing HTML.

So for example, let's say you are using a library like [HTMX](https://htmx.org/) to make your website dynamic. You want to return some HTML from your Go server. You would probably write something like this:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<div><p>Hello</p></div>")
}
```

But with this library, you can write it like this:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    tag := htmlmaker.New("div")
    
    tag.AddChild(htmlmaker.New("p"))
    tag.Children[0].SetBody("Hello")
    
    fmt.Fprintf(w, tag.Convert())
}
```

Much more readable, right?

Know I know this would seem like a lot of work, but it's not. You can just create a function that returns a tag, and then use that function to return the tag's converted value.

```go
func Person(name string) *htmlmaker.Tag {
    tag := htmlmaker.New("div")
    
    tag.AddChild(htmlmaker.New("p"))
    tag.Children[0].SetBody("Hello " + name)
    
    return tag
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, Person("John").Convert())
}
```

So now you can just use the `Person` function to return the tag, and then use the tag's converted value to return the HTML.

Cool right?

## Usage

```go
package main

import (
    "fmt"
    maker "github.com/newtoallofthis123/html_maker"
)

func main() {
    // Create a new tag
    tag := htmlmaker.New("div")
    
    // Add a class to the tag
    tag.AddChild(htmlmaker.New("a"))
    
    // Manipulate the attributes of the tag
    tag.Children[0].SetBody("Google")
    tag.Children[0].AddAttr("href", "https://google.com")
    tag.Children[0].AddClasses([]string{"class1", "class2"})
    
    // Add a new tag to the tag
    tag.AddChild(htmlmaker.New("p"))
    // Add it's body
    tag.Children[1].SetBody("Hello")

    // Print the tag
    fmt.Println(tag.Convert())
}
```

Output:

```bash
<div><a class="class1 class2" href="https://google.com">Google</a><p>Hello</p></div>
```

## Mini Documentation

This library is so small and simple, yet I end up using quite a lot. So here's a mini documentation for ease of use.

### `New(tag string) *Tag`

Creates a new tag with the tag name `tag`.

### `AddChild(child *Tag)`

Adds a child to the tag.

### `SetBody(body string)`

Sets the body of the tag.

### `AddAttr(key string, value string)`

Adds an attribute to the tag.

### `AddClasses(classes []string)`

Adds classes to the tag.

### `AddClass(class string)`

Adds a class to the tag.

### `AddStyle(key string, value string)`

Adds a style to the tag.

### `Convert() string`

Converts the tag to a string.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
