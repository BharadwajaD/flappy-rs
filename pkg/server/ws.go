package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog/log"

	"github.com/bharadwajaD/flappy-go/pkg/game"
)

type wsrw struct{
    reader *wsutil.Reader
    writer *wsutil.Writer

    full_content string
}

func newwsrw(conn *net.Conn) wsrw {
    return wsrw{
        reader: wsutil.NewServerSideReader(*conn),
        writer: wsutil.NewWriter(*conn, ws.StateServerSide, ws.OpText),
    }
}


func (wsc wsrw) ReadString(delim byte) (string, error) {
    content := make([]byte, 2048)
    log.Debug().Msgf("Called ReadString")
    wsc.reader.NextFrame()
    if _, err := wsc.reader.Read(content); err != nil {
        if err != io.EOF{
            //TODO: EOF error handling
            return "", err
        }
    }

    wsc.full_content += string(content)
    idx := strings.Index(wsc.full_content, string(delim))
    if idx == -1 {
        return "", nil
    }

    msg := wsc.full_content[:idx+1]
    wsc.full_content = wsc.full_content[idx+1:]
    log.Debug().Msgf("msg: %s", msg)
    return msg, nil
}

func (wsc wsrw) WriteString(str string) error {
    _, err :=  wsc.writer.Write([]byte(str))
    wsc.writer.Flush()
    return err
}

type WSServer struct {
	host string
	port string
}

type HtmlGame struct{
    GameId int
}

func NewWSServer(config *Config) WSServer {

    //to load js and css files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
        templ, err := template.ParseFiles("./tmpl/index.html");
        dir , err := os.Getwd()
        log.Debug().Msgf("%+v %s\n", templ, dir)
        if err != nil{
            w.WriteHeader(404)
            w.Write([]byte(err.Error()))
        }

        templ.Execute(w, HtmlGame{})

	})

	return WSServer{
		host: config.Host,
		port: config.Port,
	}

}

func (s WSServer) Run(gameOpts *game.GameOpts) {

	client_id := 0
	ggame := game.NewGroupGame(gameOpts)

    /*
    * If using custom frontend then make a websocket connection to /game-ws and proceed
    */
	http.HandleFunc("/game-ws", func(w http.ResponseWriter, r *http.Request) {

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
            log.Fatal().Err(err)
		}


		log.Debug().Msgf("DEBUG:SERVER RUN: %+v", conn)
		newgame := game.NewGame(&ggame, gameOpts.Clone())
		log.Debug().Msgf("DEBUG:SERVER RUN: %+v", newgame)

		client_id++
		client := &Client{
			id:   "WS:" + strconv.Itoa(client_id),
            rw: newwsrw(&conn),
            conn: conn,
			game: newgame,
		}

		go client.handleRequest()
	})

	http.ListenAndServe(fmt.Sprintf("%s:%s", s.host, s.port), nil)

}
