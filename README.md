<!--
	mdparser library Project
	Copyright (C) 2021 ALiwoto
	This file is subject to the terms and conditions defined in
	file 'LICENSE', which is part of the source code.
-->

# mdparser
mdparser for telegram bot API

<hr/>

## How to use
First you need to get the package:
`go get -u github.com/ALiwoto/mdparser`

then you can use the package like this:

```go
import "github.com/ALiwoto/mdparser"



func sendMessage(msg Message) {
	md := mdparser.GetBold("This is a message").AppendNormal(":\n")
	md = md.AppendItalic("Italic\n")
	md = md.AppendMono("Mono space\n")
	md = md.AppendHyperLink("text", "https://google.com")

	msg.Reply(md.ToString(), options{version: "MarkdownV2"})
}


```
