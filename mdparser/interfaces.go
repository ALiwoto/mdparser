package mdparser

type WMarkDown interface {
	Append(md WMarkDown) WMarkDown
	AppendThis(md WMarkDown) WMarkDown
	ToString() string
	AppendNormal(v string) WMarkDown
	AppendNormalThis(v string) WMarkDown
	AppendBold(v string) WMarkDown
	AppendBoldThis(v string) WMarkDown
	AppendItalic(v string) WMarkDown
	AppendItalicThis(v string) WMarkDown
	AppendMono(v string) WMarkDown
	AppendMonoThis(v string) WMarkDown
	AppendHyperLink(text, url string) WMarkDown
	AppendHyperLinkThis(text, url string) WMarkDown
	AppendMention(text string, id int64) WMarkDown
	AppendMentionThis(text string, id int64) WMarkDown
	getValue() string
	setValue(v string)
}
