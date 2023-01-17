package main

import (
	"github.com/Appkube-awsx/awsx-cloudelements/cmd"
	"github.com/spf13/cobra"
	"log"
	"os"
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
	//AwsxCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	//AwsxCmd.PersistentFlags().String("ac", "", "aws account number")
	//AwsxCmd.PersistentFlags().String("region", "", "aws region")
	//AwsxCmd.PersistentFlags().String("accessKey", "", "aws access key")
	//AwsxCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	//AwsxCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")

}

func main() {
	cmd.Execute()
}
