package routes

import (
	"log"
	"net/http"
	"tetris/game"
	"tetris/game/actions"
	"tetris/views"
)

var games = make(map[string]*game.Game)

func Root(w http.ResponseWriter, _ *http.Request) {
	log.Printf("Started a new Game")
	// TODO: Middleware to stop longest match
	g := game.NewGame()
	// TODO: Game instance per ip or session
	games["ip"] = g
	views.Render(w, "index", g, 200)
}

func Tick(w http.ResponseWriter, action string) {
	log.Printf("Tick got action %s", action)
	game := games["ip"]
	// TODO: Middleware to handle bad requests
	if game == nil {
		w.WriteHeader(400)
		return
	}

	if action == "" {
		game.Tick(actions.Down)
	} else {
		game.Tick(actions.Action(action))
	}

	if game.GameOver {
		// Special status code that stops HTMX polling
		views.Render(w, "state", game, 286)
		return
	}

	views.Render(w, "state", game, 200)
}

func Restart(w http.ResponseWriter, _ *http.Request) {
	game := games["ip"]
	// TODO: This too
	if game == nil {
		w.WriteHeader(400)
		return
	}

	game.Restart()
	views.Render(w, "game", game, 200)
}
