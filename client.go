package main

//import "crypto/tls"
import "fmt"
import irc "github.com/fluffle/goirc/client"

func connect(
	channelName string,
	channelNick string,
	callback func(conn *irc.Conn, line *irc.Line),
) {
	cfg := irc.NewConfig(channelNick)
	cfg.SSL = true
	//cfg.SSLConfig = &tls.Config{InsecureSkipVerify: true}
	cfg.Server = "irc.freenode.net:7000"
	//cfg.Pass = "secret"
	cfg.NewNick = func(n string) string { return n + "^" }
	c := irc.Client(cfg)

	// Add handlers to do things here!
	// e.g. join a channel on connect.
	c.HandleFunc("connected", joinChannel)

	c.HandleFunc("privmsg", callback)

	// And a signal on disconnect
	quit := make(chan bool)
	c.HandleFunc(
		"disconnected",
		func(conn *irc.Conn, line *irc.Line) { quit <- true },
	)

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
