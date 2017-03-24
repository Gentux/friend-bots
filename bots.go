package main

import irc "github.com/fluffle/goirc/client"
import "math/rand"
import "regexp"

// Guilhem Bot core
func handleMessageGuilhem(conn *irc.Conn, line *irc.Line) {
	technicalAnswers := []string{
		"Mais ! C'est de la merde !",
		"Laisse tomber, de toutes manières on va le réécrire en GO...",
		"On doit le réecrire en GO ça !",
		"Mais c'est pas du tout ça le problème, en fait […]",
		"C'est fait roulé sous les aisselles",
		"I have no *FRAKING* idea!",
		"NAN ! Je suis en copie, je suis perché !",
		"je veux pas devenir un SPOF",
		"'Tain IRC ! Protocole de merde ! Je vais mettre en place un XM Peuh Peuh",
		"Non mais ça c'est soooo 2011 !",
		"Les mecs de chez […] ont rien compris, j'te fouterai tout le monde à la porte moi.",
		"Désolé d'avoir raison",
		"Je ne veux pas d'un consensus mou",
	}

	personalAnswers := []string{
		"Mais j'suis pas super fort, j'suis mieux qu'ça même. J'suis surpuissant.",
		"Je parle fort… J'ai du charisme, c'est tout !",
		"Yves, Yves, YVES ! Laisse nous parler",
		"Yves, Yves, YVES ! Après le daily stp !",
		"J'ai un super nom pour le deuxième DC ! WATTOO WATTOO",
		"Attends, y'a mon nom sur un t-shirt s'il-te-plaît !",
		"J'me vois pas me priver de bonheur parce que les autres veulent une vie de merde",
		"«Vous ne pouvez *résoudre* un problème avec le *même* type de *pensée* qui a créé le problème» A.Einstein",
		"Il faut que je me sente… Important",
		"Oula oui, au moins ça !",
		"«rm -Rf de-toute-l-équipe», oui !",
		"Nan, mais j'illumine la tour eiffel depuis chez moi !",
		"Si tu penses pas que tes exceptionel, personne le penseras pour toi !",
	}

	openstackRegexp, _ := regexp.Compile("(?i)jarvis|openstack|keystone|deploy|datacenter|bug|meeting|horizon|glance|cinder|contrail|coreos|dc|nova|neutron")
	technicalPoint := openstackRegexp.FindString(line.Text())
	randomLine := rand.Intn(40)

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

// ygbo Bot core
func handleMessageYgbot(conn *irc.Conn, line *irc.Line) {
	technicalAnswers := []string{
		"Ça a été codé avec le cul!",
		"C'est fait par un mec qu'a le QI d'un anus artificiel!",
		"J'y crois pas! jamais vu une horreur pareille!",
		"Une GUI pensée avec un QI minimum est sensé être plus simple que la CLI, si la CLI est plus simple, c'est que la target de la GUI est ratée!",
		"Les paquets Debian sont au développement ce que les cartons d'emballage sont aux meubles Ikéa.",
	}

	personalAnswers := []string{
		"Hein?",
		"Hein?",
		"Hein?",
		"Tu disait?",
	}

	personalAnswersFix := []string{
		"",
		"",
		"",
		"s/disait/disias/",
	}

	personalAnswersFixFix := []string{
		"",
		"",
		"",
		"t'ain d'clavire!",
	}

	openstackRegexp, _ := regexp.Compile("(?i)javascript|bug|git|kibana|jira|selenium|packaging|kde|unity")
	technicalPoint := openstackRegexp.FindString(line.Text())
	randomLine := rand.Intn(40)

	if technicalPoint != "" {
		randomAnswer := rand.Intn(len(technicalAnswers))
		conn.Privmsg(
			channelName,
			technicalAnswers[randomAnswer],
		)
	} else if randomLine < 3 {
		randomAnswer := rand.Intn(len(personalAnswers))
		conn.Privmsg(
			channelName,
			personalAnswers[randomAnswer],
		)
		if personalAnswersFix[randomAnswer] != "" {
			conn.Privmsg(
				channelName,
				personalAnswersFix[randomAnswer],
			)
			if personalAnswersFixFix[randomAnswer] != "" {
				conn.Privmsg(
					channelName,
					personalAnswersFixFix[randomAnswer],
				)
			}
		}
	}
}

// Carot Bot core
func handleMessageCarot(conn *irc.Conn, line *irc.Line) {
	technicalAnswers := []string{
		"La test suite est rouge, mais ça a marché",
		"Adil, le front est lent !",
		"J'peux pas valider la story, il manque un accent circonflexe...",
		"Chez moi ça marche, donc c'est bon",
		"C'est pas automatisable. On ne peut pas valider la story.",
		"Regarde dans 2 minutes",
		"C'est KO, mais c'est OK",
		"Vide ton cache",
		"Relance ton navigateur",
		"C'est pas un bug, c'est une feature...",
	}
	personalAnswers := []string{
		"J'vais t'peter la gueule",
		"Tu veux t'battre ?",
		"Oh, il est beau ton bleu !",
		"Chocolat !?",
		"Est-ce que t'es bon au baby ?",
		"Demain, j'ramène des cookies",
		"On va boire une bière vite fait ce soir ?",
		"Lâche pas les manettes !!",
		"Y a des croissants chez BSS :)",
		"Faut la rouler sous les pieds",
		"Tu me dois une bière !",
		"T'es pas sur l'IRC !",
		"C'est toi le [...] !",
	}

	randomLine := rand.Intn(40)
	openstackRegexp, _ := regexp.Compile("(?i)selenium|test suite|stg0|dev32|jira|dev37|bug|story|front|billing|CRM")
	personalRegexp, _ := regexp.Compile("(?i)biere|chocolat|battre|baby|foot")
	technicalPoint := openstackRegexp.FindString(line.Text())
	personalPoint := personalRegexp.FindString(line.Text())

	if technicalPoint != "" {
		conn.Privmsg(
			channelName,
			technicalAnswers[rand.Intn(len(technicalAnswers))],
		)
	} else if randomLine < 3 || personalPoint != "" {
		conn.Privmsg(
			channelName,
			personalAnswers[rand.Intn(len(personalAnswers))],
		)
	}

}

// VBAbot Bot core
func handleMessageVBAbot(conn *irc.Conn, line *irc.Line) {
	personalAnswers := []string{
		"ASV ?",
		"L'informatique dans les nuages…",
		"Jamais le vendredi !",
		"Oui ce serait effectivement mieux avec jabber",
		"Ce sera en salle cumulonimbus",
		"J'aime pas les cookies a la cachuete.",
		"Plus, mieux, mais moins cher !",
	}

	techRegexp, _ := regexp.Compile("(?i)opencontrail|contrail|neutron|django|nova|nagios|kibana|bss|prod|staging|integration|irc|bot")
	technicalPoint := techRegexp.FindString(line.Text())
	randomLine := rand.Intn(40)

	if technicalPoint != "" {
		conn.Privmsg(
			channelName,
			GetPipo(),
		)
	} else if randomLine < 3 {
		randomAnswer := rand.Intn(len(personalAnswers))
		conn.Privmsg(
			channelName,
			personalAnswers[randomAnswer],
		)
	}
}
