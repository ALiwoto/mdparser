<!--
	mdparser library Project
	Copyright (C) 2021-2026 ALiwoto
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
	md.Styled("Bold + Italic\n", mdparser.StyleBold, mdparser.StyleItalic)
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

## Nesting Format Entities

Telegram officially documents nested message entities in the Bot API docs:

- [Bot API: Formatting options](https://core.telegram.org/bots/api#formatting-options)
- [Bot API: MessageEntity](https://core.telegram.org/bots/api#messageentity)
- [Bot API changelog: version 4.5](https://core.telegram.org/bots/api-changelog#version-4-5)

Important rules from the official docs:

- Nested entities are supported, but if two entities overlap, one must fully contain the other.
- `bold`, `italic`, `underline`, `strike-through`, and `spoiler` may be nested in each other.
- `code` and `pre` are special and have stricter nesting limits.
- Legacy `Markdown` is limited; use `MarkdownV2` if you need modern formatting behavior.

That means combined formatting is real and officially supported. For example, bold + italic on the same words is valid when represented according to Telegram's nesting rules.

## Migrating From Old Versions

Current rule:

- `Clone()` is the only copy operation.
- All other instance methods mutate the receiver.

Removed methods:

- `AppendThis`
- `AppendNormal`
- `AppendBold`
- `AppendItalic`
- `AppendMono`
- `AppendUnderline`
- `AppendStrike`
- `AppendHyperLink`
- `AppendMention`
- `AppendSpoiler`
- `ReplaceMdThis`
- `ReplaceMdThisN`
- `ElThis`
- `SpaceThis`
- `TabThis`

Replacement pattern:

- Use the non-`This` method directly for mutation.
- If old code expected copy-style behavior, call `Clone()` first.

Examples:

- `md.AppendBold("x")` -> `md.Clone().Bold("x")`
- `md.AppendBoldThis("x")` -> `md.Bold("x")`
- `md.AppendThis(other)` -> `md.Append(other)`
- `md.ReplaceMdThis(old, new)` -> `md.ReplaceMd(old, new)`
- `md.ElThis()` -> `md.El()`

Behavior changes:

- `Append`
- `ReplaceMd`
- `ReplaceMdN`
- `ReplaceToNew`
- `ReplaceToNewN`
- `El`
- `Space`
- `Tab`

These methods still exist, but they now mutate the current markdown. If older code relied on copy-style behavior, migrate to `md.Clone().SomeMethod(...)`.
