package api

import "gopkg.in/telebot.v3"

type Route struct {
	Endpoint interface{}
	Handler  telebot.HandlerFunc
	Name     string
}

// type RouteHandler func(c telebot.Context) error

func (h *Handlers) GetRoutes() []Route {
	return []Route{
		{
			Endpoint: "/start",
			Handler:  h.StartHandler,
			Name:     "start",
		},
		{
			Endpoint: "/help",
			Handler:  h.HelpHandler,
			Name:     "help",
		},
	}
}

func (h *Handlers) SetupHandlers(bot *telebot.Bot) {
	routes := h.GetRoutes()

	for _, route := range routes {
		bot.Handle(route.Endpoint, route.Handler)
		h.logger.Info("handler registered",
			"route", route.Name,
			"endpoint", route.Endpoint)
	}

	//кнопки
	bot.Handle(telebot.OnCallback, h.CallbackHandler)
}

// func Init() {

// 	HandleRoute("start", "/start", h.StartHandler)
// }

// func HandleRoute(name, path string, handler RouteHandler) {

// }
