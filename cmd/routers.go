package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/spf13/cobra"

	"github.com/nurislam03/golang_redis/internal/api"
)

var (
	exportTo string
)

func newRoutersCmd() *cobra.Command {
	// routersCmd represents the routers command
	var routersCmd = &cobra.Command{
		Use:   "routers",
		Short: "Generate routers documentation",
		Run: func(cmd *cobra.Command, args []string) {
			router, _ := api.New()
			switch exportTo {
			case "json":
				genRoutesJSONDoc(router)
			case "md":
				genRoutesMarkdownDoc(router)
			default:
				printRoutes(router)
			}
		},
	}

	routersCmd.Flags().StringVarP(&exportTo, "export", "e", "cli",
		`create api routes to file as JSON or Markdown format. 
				Options are json | md | cli`,
	)

	return routersCmd
}

func printRoutes(router *chi.Mux) {
	fmt.Println("Printing available routes: ")
	docgen.PrintRoutes(router)
	fmt.Println("OK")
}

func genRoutesJSONDoc(router *chi.Mux) {
	fmt.Print("generating routes json file: ")
	jsonapi := docgen.JSONRoutesDoc(router)
	if err := ioutil.WriteFile("routes.json", []byte(jsonapi), 0644); err != nil {
		log.Println(err)

		return
	}
	fmt.Println("OK")
}

func genRoutesMarkdownDoc(router *chi.Mux) {
	fmt.Print("generating routes markdown file: ")
	md := docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
		ProjectPath: "github.com/nurislam03/golang_redis",
		Intro:       "template REST API.",
	})
	if err := ioutil.WriteFile("routes.md", []byte(md), 0644); err != nil {
		log.Println(err)
		return
	}

	fmt.Println("OK")
}
