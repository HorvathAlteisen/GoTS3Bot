package ts3

// Bot stores a new Bot instance
type Bot struct {
	query *ServerQuery
}

// NewBot creates a new Bot instance
func NewBot(query *ServerQuery) Bot {

	bot := new(Bot)

	bot.query = query

	return *bot

} /*

// Login logs a user in to the ServerQuery
func (bot *Bot) Login(user string, password string) (string, error) {

	return bot.query.SendCommand(fmt.Sprintf("login %s %s", user, password))

}

// Use tell the vserver which Vserver to use
func (bot *Bot) Use(vserver byte) (string, error) {

	return bot.query.SendCommand(fmt.Sprintf("use %d", vserver))
}

// Close closes the ServerQuery connection and sends a ServerQuery "quit" command
func (bot *Bot) Close() error {

	bot.query.SendCommand("quit")
	return bot.query.Close()

}*/
