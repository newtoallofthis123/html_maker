package htmlmaker_test

import (
	"testing"

	htmlmaker "github.com/newtoallofthis123/html_maker"
)

func TestBasicTag(t *testing.T) {
	toGet := "<p>Hello</p>"

	tag := htmlmaker.New("p")
	tag.SetBody("Hello")

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}

func TestClasses(t *testing.T) {
	toGet := "<p class=\"class1 class2 class3\">Hello</p>"
	tag := htmlmaker.New("p")
	tag.SetBody("Hello")
	tag.AddClass("class1")
	tag.AddClass("class2")
	tag.AddClass("class3")

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}

func TestFullTag(t *testing.T) {
	toGet := "<p style=\"color: red;\" id=\"id1\" class=\"class1 class2 class3\">Hello</p>"
	tag := htmlmaker.New("p")
	tag.SetBody("Hello")
	tag.AddClass("class1")
	tag.AddClass("class2")
	tag.AddClass("class3")
	tag.AddAttr("style", "color: red;")
	tag.SetId("id1")

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}

func TestNestedTags(t *testing.T) {
	toGet := "<div><p>Hello</p></div>"
	tag := htmlmaker.New("div")
	tag.AddChild(htmlmaker.New("p"))
	tag.Children[0].SetBody("Hello")

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}

func TestComplex(t *testing.T) {
	toGet := "<div><a href=\"https://google.com\" class=\"class1 class2\">Google</a><p>Hello</p></div>"

	tag := htmlmaker.New("div")
	tag.AddChild(htmlmaker.New("a"))
	tag.Children[0].SetBody("Google")
	tag.Children[0].AddAttr("href", "https://google.com")
	tag.Children[0].AddClasses([]string{"class1", "class2"})
	tag.AddChild(htmlmaker.New("p"))
	tag.Children[1].SetBody("Hello")

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}
