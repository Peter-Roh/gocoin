package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Peter-Roh/gocoin/blockchain"
)

const (
	port        string = ":4000"
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

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := homeData{
		PageTitle: "Home",
		Blocks:    blockchain.GetBlockchain().AllBlocks(),
	}
	templates.ExecuteTemplate(w, "home", data)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	data := addData{
		PageTitle: "Add",
	}

	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", data)
	case "POST":
		r.ParseForm()
		blockData := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(blockData)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
	fmt.Printf("Listening on http://localhost%s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
