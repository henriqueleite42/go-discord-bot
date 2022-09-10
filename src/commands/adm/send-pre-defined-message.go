package adm

import (
	"context"
	"fmt"
	"strings"
	"time"

	"monetizzer/src/config/colors"
	"monetizzer/src/config/ids"
	"monetizzer/src/config/images"
	"monetizzer/src/config/urls"
	"monetizzer/src/types"
	"monetizzer/src/utils"

	"github.com/andersfylling/disgord"
)

type preDefinedMessage struct {
	name string
	description string
	channelId disgord.Snowflake
	messages []*disgord.CreateMessage
}

var spdm = []preDefinedMessage{
	{
		name: "welcome",
		description: "Welcome Message",
		channelId: ids.WelcomeChannelId,
		messages: []*disgord.CreateMessage{
			{
				Embeds: []*disgord.Embed{
					{
						Title: "Seja bem vindo a Monetizzer!",
						Description: fmt.Sprintf(`A Monetizzer é um serviço que te ajuda a monetizar suas comunidades!

Atualmente atuamos com **Bots de Discord**, que podem ser adicionado seguindo as instruções no canal <#%d>.`, ids.BotsChannelId),
						Color: colors.Monetizze(),
					},
				},
			},
			{
				Content: urls.InviteMonetizzerServer,
			},
		},
	},
	{
		name: "bots",
		description: "Bots Info",
		channelId: ids.BotsChannelId,
		messages: []*disgord.CreateMessage{
			{
				Embeds: []*disgord.Embed{
					{
						Title: "Veja nossos bots",
						Description: "Para adiciona-los, basta clicar no botão \"Adicionar ao servidor\"",
						Color: colors.Monetizze(),
					},
				},
			},
			{
				Embeds: []*disgord.Embed{
					{
						Title: "🔞 Maitê",
						Description: strings.Join([]string{
							"Maitê é sua melhor amiga para compra e venda de conteúdo adulto!",
							"Sua loja é compartilhada em todos os servidores que ela está presente, e ela cuida de todo o processo de venda pra você!",
						}, "\n\n"),
						Color: colors.Maite(),
						Thumbnail: &disgord.EmbedThumbnail{
							URL: images.MaiteIcon,
						},
					},
				},
				Components: []*disgord.MessageComponent{
					{
						Type: 1,
						Components: []*disgord.MessageComponent{
							{
								Type: 2,
								Label: "Adicionar ao servidor",
								Url: urls.InviteMaiteBot,
								Style: 5,
							},
						},
					},
				},
			},
		},
	},
	{
		name: "support",
		description: "Create ticket button",
		channelId: ids.SupportChannelId,
		messages: []*disgord.CreateMessage{
			{
				Embeds: []*disgord.Embed{
					{
						Title: "🎟 Suporte",
						Description: "Se você precisa de ajuda, você pode criar um ticket para entrar em contato com nosso suporte e nós faremos nosso melhor para ajuda-lo.",
						Color: colors.Monetizze(),
					},
				},
				Components: []*disgord.MessageComponent{
					{
						Type: 1,
						Components: []*disgord.MessageComponent{
							{
								Type: 2,
								CustomID: "CREATE_TICKET",
								Emoji: &disgord.Emoji{
									Name: "🎟",
								},
								Label: "Criar ticket",
								Style: 1,
							},
						},
					},
				},
			},
		},
	},
	{
		name: "rules",
		description: "Rules Message",
		channelId: ids.RulesChannelId,
		messages: []*disgord.CreateMessage{
			{
				Embeds: []*disgord.Embed{
					{
						Title: "Regras da Monetizzer",
						Description: `Diretrizes da Comunidade do Discord, Termos de Serviço da Monetizzer e Termos de Serviço da Discord deverão ser respeitadas.

Qualquer ato de violação dos termos do Discord, Diretrizes da Comunidade do Discord e Termos de Serviço da Monetizzer é proibido.

Você poderá ter sua conta banida do servidor e será adicionado a nossa blocklist sendo impedido de utilizar nossos serviços. Portanto, respeite-as.

Nota: Poderá, ou seja, cada caso será avaliado, dependendo do que infringir poderá ser banido.`,
						Footer: &disgord.EmbedFooter{
							Text: "Última atualização: 31 de Agosto de 2022",
						},
						Fields: []*disgord.EmbedField{
							{
								Name: "👥 Diretrizes da Comunidade",
								Value: "https://discord.com/guidelines",
							},
							{
								Name: "🛠️ Termos de Serviço",
								Value: "https://discord.com/terms",
							},
							{
								Name: "📃 Termos da Monetizzer",
								Value: "_EM BREVE_",
							},
						},
						Color: colors.Monetizze(),
					},
				},
			},
			{
				Embeds: []*disgord.Embed{
					{
						Title: "📝 Regras dos canais de texto e voz",
						Description: strings.Join([]string{
							"Nota: Cada canal tem sua respectiva função, portanto respeite-as.",
							"`1.0` **Tenha bom senso!** Ensinar outros usuários e/ou incentiva-las a quebrar as regras do servidor assim como os Termos e Serviços do Discord, abusos de API e demais afins é proibido. A mesma proibição vale para usuários que tentam contornar punições.",
							"`1.1` Qualquer ato de preconceito, perseguição, racismo, homofobia, ameaças, demais afins são proibidos, verifique também o seu microfone e tenha certeza que certos ruídos não venha a sair dele, caso o seu microfone esteja com problemas evite falar para não levar punição por irritar outros usuários. Seja amigável, não saia \"trollando\" o usuário caso ele(a) sinta-se desconfortável com suas atitudes.",
							"`1.2` Compartilhar qualquer conteúdo dentro dos canais do servidor NSFW (not safe for work), isso inclui links, vídeos, fotos, hentai e qualquer outro conteúdo que seja impróprio para menores de 18 anos de idade não será tolerado. O compartilhamento dos mesmos poderá implicar em um banimento permanente no servidor.",
							"`1.3` Floodar nos canais de texto e/ou enviar conteúdos considerável SPAM e/ou reagir nos canais excessivamente é proibido.",
							"`1.4` Qualquer ato de divulgação como canais de YouTube, noticiários, fake news, produtos, convite de outros servidores são proibidos. Tais compartilhamentos só serão autorizados com a breve autorização de um membro da Equipe.",
							"`1.5` Mencionar um membro e depois apagar a marcação (menção fantasma), o mesmo vale para cargos é proibido. Caso tenha marcado por engano, não delete a mensagem e diga que foi por engano, seja educado(a).",
							strings.Join([]string{
								"`1.6` Compartilhar informações pessoais dos membros sem a prévia e explicita autorização do mesmo não será tolerada. Antes de compartilhar qualquer informação verifique com o mesmo se tal ato é permitido.",
								"Nota: Toda a situação será analisada caso chegue as ocorrências até nós, porém, vale ressaltar que punições poderão ocorrer em ambos os lados de acordo com a analise feita pela equipe.",
							}, "\n"),
							strings.Join([]string{
								"`1.7` Não será tolerado usuários difamando serviços terceiros, sempre que precisa expressar qualquer opinião sobre, evite difamar se possível faça uma crítica construtiva.",
								"Nota: Se notarmos que você está indo longe demais e pedirmos para parar, apenas pare ou poderá ser punido por desrespeito a equipe.",
							}, "\n"),
							strings.Join([]string{
								"`1.8` Comércio de qualquer produto que não use nosso bot ou envio massivo de mensagens fora da Monetizzer (DM) não é permitido. Iremos analisar a situação e punição poderá ocorrer caso insista com os mesmos comportamentos.",
								"Nota: O mesmo vale para divulgações na DM dos usuários sem autorização caso seja reportado.",
							}, "\n"),
							strings.Join([]string{
								"`1.9` Assoprar, assobiar e demais sons que sejam prejudiciais para a comunicação dos outros membros nos canais de voz do servidor não é permitido.",
								"Nota: Caso não tenha algo para falar, mantenha o seu microfone desativado.",
							}, "\n"),
							"`1.10` Qualquer outro tipo de ato ou comportamento irregular, serão avaliados e terão as devidas punições.",
							strings.Join([]string{
								"**:man_judge: Regras complementares:**",
								"`2.0` Você terá seu apelido redefinido pela equipe caso o mesmo infrinja uma das regras a seguir:",
							}, "\n"),
							strings.Join([]string{
								"__Apelidos que serão considerados proibido:__",
								"`2.0.1` Apelidos NSFW (not safe for work) (+18).",
								"`2.0.2` Apelidos que personificam outros membros do servidor e/ou da equipe.",
								"`2.0.3` Apelidos onde grande parte dele esteja escrito em maiúsculo.",
								"`2.0.4` Apelidos muito longos ou que possuam uma sequência de vários emojis seguidos.",
								"`2.0.5` Apelidos que possam ofender qualquer usuário do servidor, ou qualquer outro serviço fora da Monetizzer.",
							}, "\n"),
							strings.Join([]string{
								"__Avatares que serão considerados proibido:__",
								"`2.0.6` Avatares NSFW (not safe for work) (+18).",
								"`2.0.7` Avatares que personificam a foto real (direitos autorais) de um dos membros do servidor.",
								"`2.0.8` Avatares que personificam ou induzem ao erro em relação ao membro ser ou não da equipe do servidor.",
								"`2.0.9` Avatares que personificam a foto real da Monetizzer (ou usuários que possivelmente tenta se passar pelos bots da Monetizzer).",
							}, "\n"),
						}, "\n\n"),
						Color: colors.Monetizze(),
					},
				},
			},
		},
	},
}

func sendPreDefinedMessage(pdmConfig preDefinedMessage) func(s disgord.Session, h *disgord.InteractionCreate) {
	return func(s disgord.Session, h *disgord.InteractionCreate) {
		s.SendInteractionResponse(context.Background(), h, &disgord.CreateInteractionResponse{
			Type: 5,
			Data: &disgord.CreateInteractionResponseData{
				Flags: 1 << 6,
			},
		})

		channel := s.Channel(pdmConfig.channelId)

		for _, msg := range pdmConfig.messages {
			channel.CreateMessage(msg)

			time.Sleep(time.Second / 2)
		}

		s.EditInteractionResponse(context.Background(), h, &disgord.Message{
			Embeds: []*disgord.Embed{
				{
					Title: "Done",
					Color: colors.Success(),
				},
			},
			Flags: 1 << 6,
		})
	}
}

func init() {
	for idx, value := range spdm {
		utils.RegisterSlashCommand(types.SlashCommand{
			ID: disgord.Snowflake(100 + idx),
			Name: fmt.Sprintf("spdm-%v", value.name),
			Description: value.description,
			DefaultPermission: false,
			Function: sendPreDefinedMessage(value),
		})
	}
}
