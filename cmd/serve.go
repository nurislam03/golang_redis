package cmd

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nurislam03/golang_redis/internal/api"
)

var (
	templatePort int
)

func newServeCmd() *cobra.Command {

	// ServeCmd represents the serve command
	serveCmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start authorizer API Server",
		Long: `Start authorizer API Server 
with the provided configurations.`,
		Example: `$ go run main.go serve
or
$ go run main.go s`,
		Run: runServe,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			portStr := strconv.Itoa(templatePort)
			listener, err := net.Listen("tcp", ":"+portStr)
			if err != nil {
				return fmt.Errorf("port %s is not available", portStr)
			}

			listener.Close()
			return nil
		},
	}

	port, err := strconv.Atoi(viper.GetString("server.port"))
	if err != nil || port <= 0 {
		port = 8080
	}

	portDesc := fmt.Sprintf("Port on which the server will listen. Default port is %d.", port)
	serveCmd.PersistentFlags().IntVarP(&templatePort, "template_port", "p",
		port, portDesc,
	)
	viper.BindPFlag("template_port", serveCmd.PersistentFlags().Lookup("template_port"))

	return serveCmd
}

func runServe(cmd *cobra.Command, args []string) {
	server, err := api.NewServer(templatePort)
	if err != nil {
		log.Fatal(err)
	}
	server.Start()

}
