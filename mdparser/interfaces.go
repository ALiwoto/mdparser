package mdparser

type WMarkDown interface {
	Append(md WMarkDown) WMarkDown
	ToString() string
	AppendNormal(v string) WMarkDown
	AppendBold(v string) WMarkDown
	AppendItalic(v string) WMarkDown
	AppendMono(v string) WMarkDown
	AppendHyperLink(text, url string) WMarkDown
	AppendMention(text string, id int64) WMarkDown
	getValue() string
	setValue(v string)
}
