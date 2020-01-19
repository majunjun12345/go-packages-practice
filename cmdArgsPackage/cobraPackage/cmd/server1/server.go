package server1

import (
	"log"

	"testGoScripts/cmdArgsPackage/cobraPackage/server"

	"github.com/spf13/cobra"
)

var Server1Cmd = &cobra.Command{
	Use:   "server1",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		server.Serve1()

		// fmt.Println("test_backend:", viper.Get("test_backend"))
		// fmt.Println("PATH:", viper.Get("PATH"))
	},
}

// 这里可以初始化具体服务的配置
func init() {
	Server1Cmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	Server1Cmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.pem", "cert pem path")
	Server1Cmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert key path")
	Server1Cmd.Flags().StringVarP(&server.CertName, "cert-name", "", "grpc server name", "server's hostname")
}
