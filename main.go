package main
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os/signal"
	"syscall"
	"os"
	"strings"
	"github.com/fatih/color"
	"time"
	"math/rand"
	"github.com/lxi1400/gotitle"


)


// shitty code but dont care

func afkchecked(author, content string) () {
	currentTime := time.Now()
	color.Yellow(fmt.Sprintf("[%s] Replied to AFK Check from: %s\nMessage: %s", currentTime.Format("3:4 PM"), author, content))
}




func randomresponse() string {
	rand.Seed(time.Now().Unix())
    return responses[rand.Intn(len(responses))]
}
  

func banner() {
	fmt.Print("\033[H\033[2J")
	color.Red(` 
		    :::     :::    ::: ::::::::::: ::::::::           :::     :::::::::: :::    ::: 
		  :+: :+:   :+:    :+:     :+:    :+:    :+:        :+: :+:   :+:        :+:   :+:  
		 +:+   +:+  +:+    +:+     +:+    +:+    +:+       +:+   +:+  +:+        +:+  +:+   
		+#++:++#++: +#+    +:+     +#+    +#+    +:+      +#++:++#++: :#::+::#   +#++:++    
		+#+     +#+ +#+    +#+     +#+    +#+    +#+      +#+     +#+ +#+        +#+  +#+   
		#+#     #+# #+#    #+#     #+#    #+#    #+#      #+#     #+# #+#        #+#   #+#  
		###     ###  ########      ###     ########       ###     ### ###        ###    ### 
                                                   
                                                   
	`)

}
func main() {
	title.SetTitle("Auto AFK | Login")
	banner() 
	fmt.Print("~$ Insert Token > ")
	fmt.Scan(&token)
	banner()
	fmt.Print("~$ Insert Channel ID to read/send messages to > ")
	fmt.Scan(&channelid)
	banner()
	selfbot, err := discordgo.New(token)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	
	selfbot.AddHandler(messageCreate)
	selfbot.AddHandler(func(dg *discordgo.Session, event *discordgo.Ready) {
		username, _ := dg.User("@me")
		title.SetTitle(fmt.Sprintf("Auto AFK | Logged in as %s", username))
		color.Blue(fmt.Sprintf("Logged in as %s", username))
	})
	err = selfbot.Open()
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	selfbot.Close()
}



func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID != channelid {
		return
	}
	for _, text := range afktext {
		if strings.Contains(strings.ToLower(m.Content), text) {
			if len(m.Mentions) == 0 {
				return
			}
			if m.Mentions[0].ID != s.State.User.ID {
				return
			}
			time.Sleep(2 *time.Second)
			s.ChannelMessageSend(m.ChannelID, randomresponse())
			afkchecked(m.Author.Username, m.Content)

		}
	}
}



var (
	token string
	channelid string
	responses =  []string{"im here", "im here son", "what do you want", "what",}
	afktext = []string{"afk check", "are you afk", "are you there", "afk chck", "afk checrk", "afk chrck",}
)
