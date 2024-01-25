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
	tag := htmlmaker.New("p").AddClasses([]string{"class1", "class2", "class3"}).AddStyle("color", "red").SetId("id1").SetBody("Hello")

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}

func TestNestedTags(t *testing.T) {
	toGet := "<div><p>Hello</p></div>"
	tag := htmlmaker.New("div")
	tag.AddChild(htmlmaker.New("p").SetBody("Hello"))

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}

func TestComplex(t *testing.T) {
	toGet := "<div><a href=\"https://google.com\" class=\"class1 class2\">Google</a><p>Hello</p></div>"

	tag := htmlmaker.New("div")
	tag.AddChild(htmlmaker.New("a").AddAttr("href", "https://google.com").AddClasses([]string{"class1", "class2"}).SetBody("Google"))
	tag.AddChild(htmlmaker.New("p").SetBody("Hello"))

	if tag.Convert() != toGet {
		t.Errorf("Expected %s, got %s", toGet, tag.Convert())
	}
}
