package routes

import (
	"net/http"
	"tetris/game"
	"tetris/game/actions"
	"tetris/views"
)

var games = make(map[string]*game.Game)

func Root(w http.ResponseWriter, _ *http.Request) {
	// TODO: Middleware to stop longest match
	g := game.NewGame()
	// TODO: Game instance per ip or session
	games["ip"] = g
	views.Render(w, "index", g)
}

func Tick(w http.ResponseWriter, r *http.Request) {
	game := games["ip"]
	// TODO: Middleware to handle bad requests
	if game == nil {
		w.WriteHeader(400)
		return
	}

	action := r.Header["Action"]

	if len(action) == 0 {
		game.Tick(actions.Down)
	} else {
		game.Tick(actions.Action(action[0]))
	}

	if game.GameOver {
		w.WriteHeader(286) // Special status code that stops HTMX polling
	}

	views.Render(w, "state", game)
}

func Restart(w http.ResponseWriter, _ *http.Request) {
	game := games["ip"]
	// TODO: This too
	if game == nil {
		w.WriteHeader(400)
		return
	}

	game.Restart()
	views.Render(w, "game", game)
}
