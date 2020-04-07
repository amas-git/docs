package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

func init() {
	viper.SetDefault("author", "amas@gmail.com")
	viper.SetDefault("version", "v1.0.0")

	cmdTime.PersistentFlags().BoolP("verbose", "v", false, "show verbose info")
	viper.BindPFlag("verbose", cmdTime.PersistentFlags().Lookup("verbose"))
}

var cmdTime = &cobra.Command{
	Use:     "time",
	Aliases: []string{"t"},
	Short:   "Current Time",
	Long:    "Current Time with YYYY-mm-dd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Now is ", time.Now().Format("YYYY-MM-DD:HH:MM:SS"))
		if viper.GetBool("verbose") {
			fmt.Println("Author:", viper.GetString("author"))
			fmt.Println("Version:", viper.GetString("version"))
		}
	},
}

func main() {
	if err := cmdTime.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
