package main

import irc "github.com/fluffle/goirc/client"
import "math/rand"
import "regexp"

const channelName = "#poc_beta_test"

// Guilhem Bot core
func handleMessageGuilhem(conn *irc.Conn, line *irc.Line) {
	technicalAnswers := []string{
		"Mais ! C'est de la merde !",
		"On doit le réecrire en GO ça !",
		"Mais c'est pas du tout ça le problème, en fait […]",
		"C'est fait roulé sous les aisselles",
		"I have no *FRAKING* idea!",
		"NAN ! Je suis en copie, je suis perché !",
		"je veux pas devenir un SPOF",
		"'Tain IRC ! Protocole de merde ! Je vais mettre en place un XM Peuh Peuh",
		"Les mecs de chez [...] ont rien compris...",
		"Non mais c'est sooooo 2011",
	}

	personalAnswers := []string{
		"Mais j'suis pas super fort, j'suis mieux qu'ça même. J'suis surpuissant.",
		"Je parle fort… J'ai du charisme, c'est tout !",
		"Yves, Yves, YVES ! Laisse nous parler",
		"J'ai un super nom pour le deuxième DC ! WATTOO WATTOO",
		"Attends, y'a mon nom sur un t-shirt s'il-te-plaît !",
		"J'me vois pas me priver de bonheur parce que les autres veulent une vie de merde",
		"«Vous ne pouvez *résoudre* un problème avec le *même* type de *pensée* qui a créé le problème» A.Einstein",
		"Il faut que je me sente... Important",
	}

	openstackRegexp, _ := regexp.Compile("jarvis|openstack|keystone|deploy|datacenter|bug|meeting|horizon|glance|cinder|contrail|coreos|dc|nova|neutron")
	technicalPoint := openstackRegexp.FindString(line.Text())
	randomLine := rand.Intn(20)

	if technicalPoint != "" {
		conn.Privmsg(
			channelName,
			technicalAnswers[rand.Intn(len(technicalAnswers))],
		)
	} else if randomLine < 3 {
		conn.Privmsg(
			channelName,
			personalAnswers[rand.Intn(len(personalAnswers))],
		)
	}
}

func main() {
	// Gentux (for concurency tests)
	//go connect(channelName, "GentuxBot", handleMessageGuilhem)
	// Guilhem
	connect(channelName, "GuilhemBot", handleMessageGuilhem)
}
