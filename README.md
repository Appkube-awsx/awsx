# awsx
AWSX is the modular command line CLI for for Appkube platform. All the subcommands are written as plugins for the main commands.
Appkube job engine calls those CLI commands for all its supported automation jobs. The commands are also directly embeddded in AWSX-Api
server. 

Please refer to Appkube Architecture diagram for details on how the appkube platform calls this commands.

Please refer the specification of subcommands for every subcommands input/ output / algos.

# Command structure
    <main command> <persistant flags> <sub-command> <sub-command flags>
    e.g. awsx --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx  getElementDetails  --zone=us-west-2
    
    awsx:               It is the main parent command
    --vaultURL:         It is persistant flag which pass the vault url to the sub-command. Vault is a separate server which store user cloud credentials in key=value pair in encrypted format. \
    --accountId:        This persistant flag pass the AWS account id to the sub-command. Sub-command pass this account id as key to the vault serve to get user credentials \
    getElementDetails:  It is the sub-command of parent awsx
    --zone:             Sub-command flag which provide AWS region to the sub-command

# How to embed sub-command in GO CLI
    We can add as many sub-commands as required
    1. Create a main cobra command. We have created the parent AwsxCmd as below
        var AwsxCmd = &cobra.Command{
            Use:   "awsx",
            Short: "awsx main command",
            Long:  `awsx main command`,
            Run: func(cmd *cobra.Command, args []string) {
            },
        }
    2. In the init function add the sub-command to the parent command
        func init() {
            AwsxCmd.AddCommand(cmd.AwsxCloudElementsCmd)
        }

# [Type of flags in Cobra CLI](https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1)
    Cobra has two types of flags:

    *    Persistent flags - available to the command it is assigned to, as well as all its sub-commands
    *    Local flags - only assigned to a specific command