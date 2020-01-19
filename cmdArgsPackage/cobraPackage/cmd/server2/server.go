package server2

import (
	"log"
	"testGoScripts/cmdArgsPackage/cobraPackage/server"

	"github.com/spf13/cobra"
)

var Server2Cmd = &cobra.Command{
	Use:   "server2",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		server.Serve2()
	},
}
