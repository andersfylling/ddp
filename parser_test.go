package ddp

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func update() {
	if err := Clone(os.Stdout); err != nil {
		panic(err)
	}
}

const tableRegexpString = `^(\|[^\n]+\|\r?\n)((?:\|:?[-]+:?)+\|)(\n(?:\|[^\n]+\|\r?\n?)*)?$`
const tableRegexpString2 = `|(?:([^\r\n|]*)\|)+\r?\n\|(?:(:?-+:?)\|)+\r?\n(\|(?:([^\r\n|]*)\|)+\r?\n)+`

// thanks to https://github.com/fletcher/MultiMarkdown/blob/ef1038d4812322550801e9f3a2fc45a403f2a656/Utilities/table_cleanup.pl#L30
const regexpTableRow = `([^\n]*?\|[^\n]*?\n)`
const regexpTableArea = `^\#\#\#\#\#\# ([a-zA-Z0-9 ]+)\n*(?:\n?[^\n]*?\|[^\n]*?\n)+`

func TestNextSubSection(t *testing.T) {
	data, err := ioutil.ReadFile(DiscordAPIDocsPath + "/docs/resources/Guild.md")
	if err != nil {
		t.Fatal(err)
	}
	data2, err := ioutil.ReadFile(DiscordAPIDocsPath + "/docs/topics/Permissions.md")
	if err != nil {
		t.Fatal(err)
	}
	data = append(data, data2...)

	var processed uint
	for {
		sType, start, length, err := nextSubSection(processed, data)
		if err != nil {
			t.Error(err)
			break
		}
		processed = start + length
		if sType != typeConstants {
			continue
		}

		fmt.Println(string(data[start : start+length]))
		fmt.Println("==================")
	}
}

func TestNextSubSection2(t *testing.T) {
	data, err := ioutil.ReadFile(DiscordAPIDocsPath + "/docs/resources/Guild.md")
	if err != nil {
		t.Fatal(err)
	}
	data2, err := ioutil.ReadFile(DiscordAPIDocsPath + "/docs/topics/Permissions.md")
	if err != nil {
		t.Fatal(err)
	}
	data = append(data, data2...)

	rgxp := regexp.MustCompile(regexpTableArea)
	matches := rgxp.FindAll(data, -1)
	for _, match := range matches {
		fmt.Println(string(match))
	}
	//
	// var processed uint
	// for {
	// 	sType, start, length, err := nextSubSection(processed, data)
	// 	if err != nil {
	// 		t.Error(err)
	// 		break
	// 	}
	// 	processed = start + length
	// 	if sType != typeConstants {
	// 		continue
	// 	}
	//
	// 	fmt.Println(string(data[start : start+length]))
	// 	fmt.Println("==================")
	// }
}
