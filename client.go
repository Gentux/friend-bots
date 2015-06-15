package main

import "crypto/tls"
import "log"
import irc "github.com/gentux/goirc/client"

var channelName string

func connect(
	config configuration,
	channelNick string,
	callback func(conn *irc.Conn, line *irc.Line),
) {
	cfg := irc.NewConfig(channelNick)
	cfg.SSL = config.SSL
	cfg.SSLConfig = &tls.Config{InsecureSkipVerify: config.SSLInsecureSkipVerify}
	cfg.Server = config.Server
	cfg.Pass = config.Pass
	cfg.NewNick = func(n string) string { return n + "^" }
	c := irc.Client(cfg)

	channelName = config.ChannelName
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
		log.Fatalf("Connection error: %s\n", err)
	} else {
		log.Print("Connected")
	}

	// Wait for disconnect
	<-quit
}

func joinChannel(conn *irc.Conn, line *irc.Line) {
	conn.Join(channelName)
}
