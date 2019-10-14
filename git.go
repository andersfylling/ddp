package ddp

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
)

const DiscordAPIDocsPath = "./discord-api-docs"

// Clone output such as os.Stdout
func Clone(output sideband.Progress) error {
	_, err := git.PlainClone(DiscordAPIDocsPath, false, &git.CloneOptions{
		URL:      "https://github.com/discordapp/discord-api-docs",
		Progress: output,
	})

	return err
}
