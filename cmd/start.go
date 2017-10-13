package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Run local web server",
	Long:  `Run local web server`,
	Run:   serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {

	r := gin.Default()
	r.LoadHTMLGlob("webclient/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
