package main

type Model interface {
	completion()
}

type GPT struct{}

func (g *GPT) completion() {}

type Claude struct{}

func (c *Claude) completion() {}

type Gemini struct{}

func (g *Gemini) completion() {}

type ChatBot struct {
	model Model
}

func NewChatBot() *ChatBot {
	return &ChatBot{
		model: &Gemini{},
	}
}

func (c *ChatBot) chat() {
	c.model.completion()
}

func init() {
	NewChatBot().chat()

}
