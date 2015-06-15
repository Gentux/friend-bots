package main

import irc "github.com/gentux/goirc/client"
import "math/rand"
import "regexp"

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
		"Non mais ça c'est soooo 2011 !",
		"Les mecs de chez […] ont rien compris, j'te fouterai tout le monde à la porte moi.",
		"Désolé d'avoir raison",
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
	}

	openstackRegexp, _ := regexp.Compile("jarvis|openstack|keystone|deploy|datacenter|bug|meeting|horizon|glance|cinder|contrail|coreos|dc|nova|neutron")
	technicalPoint := openstackRegexp.FindString(line.Text())
	randomLine := rand.Intn(30)

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
		"C'est fait par un mec qu'a le QI d'un annus artificiel!",
		"J'y crois pas! jamais vu une horreur pareille!",
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

	openstackRegexp, _ := regexp.Compile("javascript|JavaScript|bug|git")
	technicalPoint := openstackRegexp.FindString(line.Text())
	randomLine := rand.Intn(30)

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
	answers := []string{
		"J'vais t'peter la gueule",
		"Tu veux t'battre ?",
		"Oh, il est beau ton bleu !",
		"Chocolat !?",
		"Est-ce que t'es bon au baby ?",
		"Demain, j'ramène des cookies",
		"On va boire une bière vite fait ce soir ?",
		"Lâche pas les manettes !!",
	}

	randomLine := rand.Intn(30)
	if randomLine < 3 {
		conn.Privmsg(channelName, answers[rand.Intn(len(answers))])
	}
}

// VBAbot Bot core
func handleMessageVBAbot(conn *irc.Conn, line *irc.Line) {
	technicalAnswers := []string{
		"Un truc codé avec emacs par un manchot !",
		"Evidemment c est encore la faute aux sysadmins :(",
		"Ils auraient du le faire en cobol !",
		"Cette bouse me reveille toutes les nuits !",
		"As tu teste en dev ? integration ? staging ?",
		"Dans le doute reboot ?",
		"Et si on le recodait en LISP ?",
	}

	personalAnswers := []string{
		"ASV ?",
		"L'informatique dans les nuages..",
		"Jamais le vendredi",
		"R IP",
		"Oui ce serait effectivement mieux avec jabber",
		"Ce sera en salle cumulonimbus",
	}

	techRegexp, _ := regexp.Compile("opencontrail|contrail|neutron|django|nova|nagios|kibana|bss|prod|staging|integration|irc|bot")
	technicalPoint := techRegexp.FindString(line.Text())
	randomLine := rand.Intn(30)

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
	}
}
