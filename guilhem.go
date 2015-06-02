package main

import "crypto/tls"
import "fmt"
import irc "github.com/fluffle/goirc/client"
import "math/rand"

const (
	channelName = "#poc_beta_guilhem"
	channelNick = "GuilhemBot"
)

func main() {
	cfg := irc.NewConfig(channelNick)
	cfg.SSL = true
	cfg.Server = "irc.freenode.net:7000"
	//cfg.Pass = "secret"
	//cfg.SSLConfig = &tls.Config{InsecureSkipVerify: true}
	cfg.NewNick = func(n string) string { return n + "^" }
	c := irc.Client(cfg)

	// Add handlers to do things here!
	// e.g. join a channel on connect.
	c.HandleFunc("connected", joinChannel)

	c.HandleFunc("privmsg", handleMessage)

	// And a signal on disconnect
	quit := make(chan bool)
	c.HandleFunc("disconnected",
		func(conn *irc.Conn, line *irc.Line) { quit <- true })

	// Tell client to connect.
	if err := c.Connect(); err != nil {
		fmt.Println("Connection error: %s\n", err)
	} else {
		fmt.Println("Connected")
	}

	// Wait for disconnect
	<-quit
}

func joinChannel(conn *irc.Conn, line *irc.Line) {
	conn.Join(channelName)
}

func handleMessage(conn *irc.Conn, line *irc.Line) {
	//rand.Seed(42)
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
