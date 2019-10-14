package ddp

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func update() {
	if err := Clone(os.Stdout); err != nil {
		panic(err)
	}
}

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
