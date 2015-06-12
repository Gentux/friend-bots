package main

import "io/ioutil"
import "log"
import "encoding/json"
import "os"
import "os/user"
import "path"

type configuration struct {
	Server                string
	Port                  int
	SSL                   bool
	SSLInsecureSkipVerify bool
	Pass                  string
	ChannelName           string
}

var (
	config     configuration
	configpath string
)

func configload() error {
	b, err := ioutil.ReadFile(configpath)
	if err == nil {
		return json.Unmarshal(b, &config)
	}
	return err
}

func configsave() error {
	j, err := json.MarshalIndent(config, "", "\t")
	if err == nil {
		return ioutil.WriteFile(configpath, j, 0600)
	}
	return err
}

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configpath = path.Join(usr.HomeDir, ".friendbotsrc")

	config.Server = "irc.freenode.net"
	config.Port = 7000
	config.SSL = true
	config.SSLInsecureSkipVerify = true
	config.Pass = "secret"
	config.ChannelName = "#chan"

	errConfigLoad := configload()
	if _, ok := errConfigLoad.(*os.PathError); ok {
		log.Printf("No config file found. Creating %s ...", configpath)
		if errConfigSave := configsave(); errConfigSave != nil {
			log.Fatal(errConfigSave)
		}
	} else if errConfigLoad != nil {
		log.Fatalf("could not parse %v: %v", configpath, errConfigLoad)
	}
}

func main() {
	// Yves
	go connect(config, "Ygbot", handleMessageYgbot)
	// Caro
	go connect(config, "Carot", handleMessageCarot)
	// Guilhem
	connect(config, "GuilhemBot", handleMessageGuilhem)
}
