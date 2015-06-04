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
		"Mais c'est pas du tout ça le probléme, en fait […]",
	}

	personalAnswers := []string{
		"C'est moi le plus fort, je suis SURPUISSANT !",
		"Je parle fort… J'ai du charisme, c'est tout !",
		"Yves, Yves, YVES ! Laisse nous parler",
		"J'ai un super nom pour le deuxiéme DC ! WATTOO WATTOO",
		"Attends, y a mon nom sur un t-shirt s'il te plait !",
	}

	openstackRegexp, _ := regexp.Compile("openstack|keystone|deploy|datacenter")
	technicalPoint := openstackRegexp.FindString(line.Text())
	if technicalPoint != "" {
		conn.Privmsg(channelName, technicalAnswers[rand.Intn(len(technicalAnswers))])
	} else if rand.Intn(20) < 2 {
		conn.Privmsg(channelName, personalAnswers[rand.Intn(len(personalAnswers))])
	}
}

func main() {
	// Gentux (for concurency tests)
	//go connect(channelName, "GentuxBot", handleMessageGuilhem)
	// Guilhem
	connect(channelName, "GuilhemBot", handleMessageGuilhem)
}
