<!--
	mdparser library Project
	Copyright (C) 2021-2022 ALiwoto
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
	md := mdparser.GetBold("This is a message")
	md.Normal(":\n")
	md.Italic("Italic\n")
	md.Mono("Mono space\n")
	md.CodeBlockLang("go", "fmt.Println(\"hello\")")
	md.HyperLink("text", "https://google.com")

	msg.Reply(md.ToString(), options{version: "MarkdownV2"})
}


```

All instance methods mutate the current markdown. If you need a copy first, clone explicitly:

```go
base := mdparser.GetNormal("hello")
copy := base.Clone().Bold(" world")
```
