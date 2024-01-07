package htmlmaker

type TagType string

// HtmlTag represents an HTML tag with its attributes, class names, children, body, and ID.
// It is a abstract representation of an HTML tag.
// It contains an recursive function to convert itself into a string.
// So it sort of acts as a poorly implemented dom tree.
// Also has the required methods to add attributes, class names, children, body, and ID.
type HtmlTag struct {
	TagType    TagType
	Attrs      map[string]string
	ClassNames []string
	Children   []HtmlTag
	Body       string
	Id         string
}

// New creates a new HtmlTag with the given tag type.
// Tag type can be anything, but it is recommended to use the standard HTML tags.
func New(tag TagType) HtmlTag {
	return HtmlTag{
		TagType: tag,
		Attrs:   make(map[string]string),
	}
}

// AddChild adds a child to the HtmlTag.
func (t *HtmlTag) AddChild(tag HtmlTag) {
	t.Children = append(t.Children, tag)
}

// AddAttr adds an attribute to the HtmlTag.
// The key is the attribute name, and the value is the attribute value.
func (t *HtmlTag) AddAttr(key, value string) {
	t.Attrs[key] = value
}

// AddClass adds a class name to the HtmlTag.
func (t *HtmlTag) AddClass(className string) {
	t.ClassNames = append(t.ClassNames, className)
}

func (t *HtmlTag) AddStyle(key, value string) {
	t.Attrs["style"] = key + ": " + value + ";"
}

func (t *HtmlTag) AddClasses(classNames []string) {
	t.ClassNames = append(t.ClassNames, classNames...)
}

// SetBody sets the body of the HtmlTag.
func (t *HtmlTag) SetBody(body string) {
	t.Body = body
}

// SetId sets the ID of the HtmlTag.
func (t *HtmlTag) SetId(id string) {
	t.Id = id
}

// partialConvert converts the HtmlTag into a string without the body or children.
func (t *HtmlTag) partialConvert() string {
	var toReturn string
	toReturn += "<" + string(t.TagType)
	for key, value := range t.Attrs {
		toReturn += " " + key + "=\"" + value + "\""
	}

	// add class names
	if t.Id != "" {
		toReturn += " id=\"" + t.Id + "\""
	}

	if t.ClassNames != nil {
		toReturn += " class=\""
		for _, className := range t.ClassNames {
			toReturn += className + " "
		}
		toReturn = toReturn[:len(toReturn)-1]

		toReturn += "\""
	}

	toReturn += ">"

	return toReturn
}

// Convert converts the HtmlTag into a string.
// This is a recursive function that calls itself on all of its children.
// Partially constructs the DOM tree.
// Returns the string representation of the HtmlTag.
func (t *HtmlTag) Convert() string {
	var toReturn string
	toReturn += t.partialConvert()
	toReturn += t.Body
	for _, child := range t.Children {
		toReturn += child.Convert()
	}
	toReturn += "</" + string(t.TagType) + ">"
	return toReturn
}
