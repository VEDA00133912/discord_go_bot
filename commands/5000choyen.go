package commands

import (
	"fmt"
	"net/url"

	"github.com/bwmarrin/discordgo"
	botRouter "github.com/maguro-alternative/discord_go_bot/bot_handler/bot_router"
)

const gosenChoyenAPIURL = "https://gsapi.cbrx.io/image"

func GosenChoyenCommand() *botRouter.Command {
	return &botRouter.Command{
		Name:        "5000choyen",
		Description: "5000兆円欲しい!!画像を生成します",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "top",
				Description: "上部に表示するテキスト",
				Required:    true,
				MaxLength:   30,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "bottom",
				Description: "下部に表示するテキスト",
				Required:    true,
				MaxLength:   30,
			},
		},
		Executor: handleGosenChoyen,
	}
}

func handleGosenChoyen(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.ApplicationCommandData().Name != "5000choyen" {
		return
	}

	topText := i.Interaction.ApplicationCommandData().Options[0].StringValue()
	bottomText := i.Interaction.ApplicationCommandData().Options[1].StringValue()

	topEncoded := url.QueryEscape(topText)
	bottomEncoded := url.QueryEscape(bottomText)

	responseURL := fmt.Sprintf("%s?top=%s&bottom=%s&type=png", gosenChoyenAPIURL, topEncoded, bottomEncoded)

	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: responseURL,
		},
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		fmt.Printf("error responding to 5000choyen command: %v\n", err)
	}
}