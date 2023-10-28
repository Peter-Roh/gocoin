// Package explorer provides explorer of the blockchain.
package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Peter-Roh/gocoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

type addData struct {
	PageTitle string
}

var templates *template.Template

func handleHome(rw http.ResponseWriter, r *http.Request) {
	data := homeData{
		PageTitle: "Home",
		Blocks:    blockchain.GetBlockchain().AllBlocks(),
	}
	templates.ExecuteTemplate(rw, "home", data)
}

func handleAdd(rw http.ResponseWriter, r *http.Request) {
	data := addData{
		PageTitle: "Add",
	}

	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", data)
	case "POST":
		r.ParseForm()
		blockData := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(blockData)
		http.Redirect(rw, r, "/", http.StatusMovedPermanently)
	}
}

// Start starts an explorer at port
func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", handleHome)
	handler.HandleFunc("/add", handleAdd)
	fmt.Printf("Listening on http://localhost:%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
