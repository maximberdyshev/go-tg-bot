package telegram_client

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	ID   int    `json:"message_id"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type From struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}
