package main

import (
	"github.com/Appkube-awsx/awsx-getelementdetails/handler/EC2"
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-cloudelements/cmd"
	"github.com/spf13/cobra"
)

var AwsxCmd = &cobra.Command{
	Use:   "awsx",
	Short: "awsx main command",
	Long:  `awsx main command`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := AwsxCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application. Global means available in child/sub commands

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.appconfig.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	AwsxCmd.AddCommand(cmd.AwsxCloudElementsCmd)
	AwsxCmd.AddCommand(EC2.CpuUtilizationPanelCmd)
}

func main() {
	cmd.Execute()
}
