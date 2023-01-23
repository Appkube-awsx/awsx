# awsx
Aws extension CLI \
awsx is the main parent command.

# Command structure
    <main command> <persistant flags> <sub-command> <sub-command flags>
    e.g. awsx --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx  getElementDetails  --zone=us-west-2
    
    awsx:               It is the main parent command
    --vaultURL:         It is persistant flag which pass the vault url to the sub-command. Vault is a separate server which store user cloud credentials in key=value pair in encrypted format. \
    --accountId:        This persistant flag pass the AWS account id to the sub-command. Sub-command pass this account id as key to the vault serve to get user credentials \
    getElementDetails:  It is the sub-command of parent awsx
    --zone:             Sub-command flag which provide AWS region to the sub-command

# How to download sub-command/module from git
    Open command prompt and go to the directory where go.mod file available
    * To download any specific version of a module, run go get command with version number
        go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.1
    * To download latest version of a module, run go get command with @latest
        go get github.com/Appkube-awsx/awsx-cloudelements@latest
        - If we omit the version number or @latest, go get command downloads the latest version always

# How to embed sub-command in GO CLI code
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

# Compile GO program: go build
    * Open command prompt and go to the project directory where a go file having main() function available
        cd /opt/mycode/awsx
    
    * go build without argument
        The go build command compiles the packages, along with their dependencies.    
        $ go build
        - If we don't provide an argument to this command, go build will automatically compile go program in your current directory
        - When we run go build, by default Go decide the name of the generated executable by using the module created earlier. \    
          When the go mod init awsx command was run, it created the module with the name 'awsx', which is why the binary generated is named 'awsx' in turn.
        - From the command line in, run the new awsx executable to confirm that the code works.
            * On Linux or Mac:
                $ ./awsx
            * On Windows:
                $ awsx.exe

    * go build with argument
        -  With -o flag, Go can chooses a name of the executable 
            go build -o [optional file path]/<executable name>
        - To test this out, change the name of the to 'awsxcmd' and have it placed in a sub-folder called bin. You don’t have to create this folder; Go will do that on its own during the build process.
            $ go build -o bin/awsxcmd
                - The -o flag makes Go match the output of the command to whatever argument we chose. In this case, the result is a new executable named 'awsxcmd' in a sub-folder named 'bin'
                - To test the new executable, change into the new directory and run the binary
                    $ cd bin
                    $ ./awsxcmd

# Install GO program: go install    
    * The go install command behaves almost identically to go build, but instead of leaving the executable in the current directory, or a directory specified by the -o flag, it places the executable into the $GOPATH/bin directory.
    * The benefit of go install is that the executable binary can be executed outside of the source directory. 
      $ go install
        NOTE - The go install command does not support the -o flag, so it will use the default name described earlier to name the executable.

# Compile and run the go program with go run command
    * Open command prompt and go to the project directory where a go file having main() function available
        cd /opt/mycode/awsx
    * Run the program with go run command
        go run ./main.go
        - In our example, our main() function is available in main.go
        - We can provide program arguments or flags to the main() function
            go run ./main.go arg1 arg2
            go run ./main.go --flag1=value --flag2=value



# [Type of flags in Cobra CLI](https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1)
    Cobra has two types of flags:

    *    Persistent flags - available to the command it is assigned to, as well as all its sub-commands
    *    Local flags - only assigned to a specific command