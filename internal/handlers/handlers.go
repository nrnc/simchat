package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(jet.NewOSFileSystemLoader("./html"), jet.InDevelopmentMode())

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Home renders the home Page
func Home(rw http.ResponseWriter, r *http.Request) {
	err := renderPage(rw, "home.jet", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

type WebSocketConnection struct {
	*websocket.Conn
}

// WsJsonResponse defines the response sent back from websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

// WsEndPoint upgrades the end point to websocket
func WsEndPoint(rw http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("client connected to endpioint")
	var response WsJsonResponse
	response.Message = `<em><small>Connected to the server</small></em>`

	ws.WriteJSON(response)
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
