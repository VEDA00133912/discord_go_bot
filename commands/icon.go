package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	botRouter "github.com/maguro-alternative/discord_go_bot/bot_handler/bot_router"
)

func IconCommand() *botRouter.Command {
	return &botRouter.Command{
		Name:        "icon",
		Description: "指定したユーザーのアイコンを表示します",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "アイコンを表示するユーザー",
				Required:    false,
			},
		},
		Executor: handleIcon,
	}
}

func handleIcon(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// オプションからユーザーを取得、指定がなければコマンド実行者
	var user *discordgo.User
	options := i.ApplicationCommandData().Options
	if len(options) > 0 && options[0].Name == "user" && options[0].UserValue(s) != nil {
		user = options[0].UserValue(s)
	} else {
		user = i.Member.User
	}

	if user == nil {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "ユーザーが見つかりませんでした",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			fmt.Printf("error responding to icon command: %v\n", err)
		}
		return
	}

	avatarURL := user.AvatarURL("2048") // AvatarURLの第一引数が画像サイズらしい
	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**<@%s> のアイコン**", user.ID),
		Image: &discordgo.MessageEmbedImage{
			URL: avatarURL,
		},
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		fmt.Printf("error responding to icon command: %v\n", err)
	}
}