package main

import irc "github.com/fluffle/goirc/client"
import "math/rand"

const channelName = "#poc_beta_test"

func handleMessageGuilhem(conn *irc.Conn, line *irc.Line) {
	answers := []string{
		"Mais ! C'est de la merde !",
		"On doit le réecrire en GO ça !",
		"C'est moi le plus fort, je suis SURPUISSANT !",
		"Je parle fort… J'ai du charisme, c'est tout !",
		"Yves, Yves, YVES ! Laisse nous parler",
		"Mais c'est pas du tout ça le probléme, en fait […]",
		"J'ai un super nom pour le deuxiéme DC ! WATTOO WATTOO",
		"Attends, y a mon nom sur un t-shirt s'il te plait !",
	}
	if rand.Intn(10) < 2 {
		conn.Privmsg(channelName, answers[rand.Intn(len(answers))])
	}
}

func main() {
	// Guilhem
	go connect(channelName, "GuilhemBot", handleMessageGuilhem)

	// Gentux (for concurency tests)
	//connect(channelName, "GentuxBot", handleMessageGuilhem)
}
