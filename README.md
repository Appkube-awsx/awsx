- [awsx](#awsx)
- [How to build / Run / Debug](#how-to-build--run--debug)
- [Command structure](#command-structure)
- [How to write a plugin subcommand](#how-to-write-a-plugin-subcommand)
- [How to support flags and inputs arguments in subcommand](#how-to-support-flags-and-inputs-arguments-in-subcommand)
- [How to embed sub-command in GO CLI](#how-to-embed-sub-command-in-go-cli)
- [How to download sub-command/module from git](#how-to-download-sub-commandmodule-from-git)
- [How to run a specific version of  sub-command/module from main awsx](#how-to-run-a-specific-version-of--sub-commandmodule-from-main-awsx)
- [All subcommands plugins references](#all-subcommands-plugins-references)
- [Building Docker image and running it](#building-docker-image-and-running-it)
- [How SRE platform calls CLI/API](#how-sre-platform-calls-cliapi)

# awsx
AWSX is the modular command line CLI for for Appkube platform. All the subcommands are written as plugins for the main commands.
Appkube job engine calls those CLI commands for all its supported automation jobs. The commands are also directly embeddded in AWSX-Api
server. 

Please refer to Appkube Architecture diagram for details on how the appkube platform calls this commands.

Please refer the specification of subcommands for every subcommands input/ output / algos.

# How to build / Run / Debug 
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

# Command structure
    <main command> <persistant flags> <sub-command> <sub-command flags>
    e.g. awsx --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx  getElementDetails  --zone=us-west-2
    
    awsx:               It is the main parent command
    --vaultURL:         It is persistant flag which pass the vault url to the sub-command. Vault is a separate server which store user cloud credentials in key=value pair in encrypted format. \
    --accountId:        This persistant flag pass the AWS account id to the sub-command. Sub-command pass this account id as key to the vault serve to get user credentials \
    getElementDetails:  It is the sub-command of parent awsx
    --zone:             Sub-command flag which provide AWS region to the sub-command

    In Dev Env , you could run the awsx also as follows:

        go run ./main.go --vaultUrl=http://localhost:10000/vi/account --accountId=xxxxxxxxxx
        * On Linux or Mac:
            
            $ ./awsx --vaultUrl=http://localhost:10000/vi/account --accountId=xxxxxxxxxx
        * On Windows:
            $ awsx.exe --vaultUrl=http://localhost:10000/vi/account --accountId=xxxxxxxxxx


# How to write a plugin subcommand
The best way to start writing an subcommand is to follow the example subcommand -
https://github.com/Appkube-awsx/awsx-cloudelements

The process is as follows:
1. Create a independent git repo for the subcommnad like https://github.com/Appkube-awsx/awsx-cloudelements

2. Clone the repo to your local machine to add code git clone https://github.com/Appkube-awsx/awsx-cloudelements.git

3. Go in the awsx-cloudelements directory and execute the following commands
    a. initialize the project
        go mod init github.com/Appkube-awsx/awsx-cloudelements
    b. download the latest version of cobra cli
        go get github.com/spf13/cobra@latest
    c. install the Cobra cli
        go install github.com/spf13/cobra-cli@latest
    d. execute cobra-cli init command. This command will generate the application with the correct file structure and imports:
        cobra-cli init

4. The above command will create directory structure as below and generate the basic cli code in root.go and main.go

    awsx-cloudelements
        |
        |__cmd
            |__root.go
        |__main.go

        In the root.go you will find the code as below

            var rootCmd = &cobra.Command{
                Use:   "aws-cloudelements",
                Short: "A brief description of your application",
                Long: `A longer description that spans multiple lines and likely contains
                        examples and usage of using your application. For example:

                        Cobra is a CLI library for Go that empowers applications.
                        This application is a tool to generate the needed files
                        to quickly create a Cobra application.`,
                // Uncomment the following line if your bare application
                // has an action associated with it:
                Run: func(cmd *cobra.Command, args []string) {
                    fmt.Println("Calling aws-cloudelements")
                },
            }

            func Execute() {
                err := rootCmd.Execute()
                if err != nil {
                    log.Fatal("There was some error while executing the CLI: ", err)
                    os.Exit(1)
                }
            }

            func init() {
                
            }

            - In the Run inline function we should write our cli code. 
            - In our example we have written a fmt.Println("Calling aws-cloudelements")
            - When we execute this command, this message will be printed on console

        In main.go we should call the command. So the main.go should be as below:

            package main

            import "github.com/Appkube-awsx/awsx-cloudelements/cmd"

            func main() {
                cmd.Execute()
            }

5.  Run and test the code as follows:
            go run main.go
                - Program will print Calling aws-cloudelements on console 

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-cloudelements) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            awsx-cloudelements

6.  Publish the code in git so that other modules can download this code as dependency from git

        a. Commit and push the code
        b. Tag the code. Use the following git commands to tag it
            git tag "v1.0.0"
            git push --tags
    
        c. Developers interested in this module, import it by running the go get command as below
            go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.0
            
            In the above go get command (go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.0) we have specified the version (v1.0.0).
            
            This version is the git tag, what we specified in the git tag command earlier. 

# How to support flags and inputs arguments in subcommand 
    In GO it is possible to pass parameters or arguments to a program's main function using flags or command-line arguments .
    
    Flags:
        Cobra has two types of flags:
        * Persistent flags - available to the command it is assigned to, as well as all its sub-commands\
        * Local flags - only assigned to single command

        1. Declaring persistent flags in cobra command
            We need to write the code to declare flags in init() function of cobra command.  
        
            func init() {
                AwsxCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
                AwsxCmd.PersistentFlags().String("accountId", "", "aws account number")
            }
            - As per the above init function the awsx command accepts two persistant flags --vaultUrl and --accountId. 
            - Since these are persistant flags, both the flags will be available to all the subcommand of awsx command
        
        2. Declaring local flags in cobra command
            Same way local flags also get declared in init() function of cobra command.  
        
            func init() {
                AwsxCloudElementsCmd.Flags().String("zone", "", "aws region")
            }
            - The above init function is taken from getElementDetails sub-command which accepts single local flag --zone. 
            - Since this is a local flag, this will be available to getElementDetails sub-command only
            NOTE: getElementDetails is another independent cobra command which act as a sub-command of awsx as well
    
        3. Run the command with flags
            Open the terminal and run the command as below with flags
            
            awsx --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx  getElementDetails  --zone=us-west-2
                    
            - In the above command we have passed two flags --vaultURL and --accountId to awsx command and --zone to getElementDetails sub-command
              These flags are persistant flags and will be accessible in getElementDetails sub-command
            - One local flag --zone is paased in the getElementDetails sub-command and this will be available to this sub-command only
    
        4. Flag structure
            AwsxCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
            AwsxCloudElementsCmd.Flags().String("zone", "", "aws region")
            
            - As above Both the flags (persistant and local) declared as of type String. 
            - Both flags have three parameters
                        String("vaultUrl", "", "vault end point")
                        String("zone", "", "aws region")
                    * parameter 1 is the name of the flag. In our case; vaultUrl is the name of persistant flag and zone is the name of local flag
                    * parameter 2 is the default values. In our case it is empty. If we provide some default value to the second paramerter, command take it as flag value if we omit the flag from command line
                    * parameter 3 is the description of the flag
        
        5. Retrieve flag values in command
            var rootCmd = &cobra.Command{
                Use:   "command name",
                Short: "short description of command",
                Long:  `detailed/long description of command`,
                Run: func(cmd *cobra.Command, args []string) {
                    vaultUrl, _ := cmd.Flags().GetString("vaultUrl")
            
                    fmt.Println(vaultUrl)
                },
            }
            - In the above code cmd.Flags().GetString("vaultUrl") instruction is getting vaultUrl flag value and storing in vaultUrl variable
    
    Input Arguments:
        To retrieve the command line arguments, we need os package. The Args variable in the os package returns all the command line arguments as array of string
            - var Args []string: : Args hold the command-line arguments, starting with the program name.
        Example:
            1. create a go file main.go
            2. Add the below code to main.go
                    package main

                    import (
                    "fmt"
                    "os"
                    )
                    
                    func main() {
                        programName := os.Args[0]
                        secondArg := os.Args[1]
                        thirdArg := os.Args[2]
                        allArgs := os.Args[1:]
                    
                        fmt.Println(programName)
                        fmt.Println(secondArg)
                        fmt.Println(thirdArg)
                        fmt.Println(allArgs)
                    }
            3. build the program
                $ go build main.go
            4. Run its executable
                $ ./main first second third
            5. Output
                    ./main
                    first
                    second
                    [first second third]
    

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

# How to download sub-command/module from git

    Open command prompt and go to the directory where go.mod file available
    
    * To download any specific version of a module, run go get command with version number
        go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.1
    * To download latest version of a module, run go get command with @latest
        go get github.com/Appkube-awsx/awsx-cloudelements@latest
        - If we omit the version number or @latest, go get command downloads the latest version always
    *  After it we can do the code chage and run go build or go install on command prompt
    
# How to run a specific version of  sub-command/module from main awsx
    * make the change in go.mod file (in subcommand require section , you can import a specific version)
    * run go mod tidy command at command prompt. This will downlod the required modules in GOPATH

# All subcommands plugins references
All the supported subcommands and there source code locations are mentiioned in 

    https://github.com/AppkubeCloud/appkube-api-specs/blob/main/awsx-api.md

| S.No | Sub-command           | Description                                           | Repository                                  | Functionalities                                                                                                                                                                            | Specs Links |
|------|-----------------------|-------------------------------------------------------|---------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| 1    | getLandingZoneDetails | Collect Information about any specific landing zone  | Percentage(%)                               | 1. Get Elements Metadata, 2. Get List of every elements with their config infos, 3. Get List of products and Environments for the landing Zone, 4. Get the cost of the landing zone   |    |
| 2    | getLandingZoneCompliance | Collect Information about any specific landing zone compliances and security | Bytes                                 | 1. Get overall account Compliance, 2. Get the elementWise Compliance, 3. Run Compliance on products and environments                                                                       |  |
| 3    | getElementDetails     | Collect Information about specific cloud elements -Run Queries | Bytes                             | 1. EC2, 2. EKS, 3. ECS, 4. LAMBDA, 5. API Gw, 6. Load Balancer                                                                                                                           | [https://github.com/Appkube-awsx/awsx-getelementdetails/tree/main?tab=readme-ov-file#awsx-getelementdetails](https://github.com/Appkube-awsx/awsx-getelementdetails/tree/main?tab=readme-ov-file#awsx-getelementdetails)   |
| 4    | getCostDetails        | Collect Information about account and elements specific costs | Bytes                            | 1. Total Account, 2. Product and Envwise, 3. Element Wise, 4. Spikes and Trends, 5. App/Data/Nw Service wise Costs                                                                      | [https://github.com/Appkube-awsx/awsx-cost](https://github.com/Appkube-awsx/awsx-cost)|
| 5    | getBusinessServiceDetails | Collect Information about the business Services that is hosted on infrastructure | Bytes             | 1. Account Wise infralist and corresponding services, 2. Product and Env wise Service Lists along with hosted infrastructure                                                             |    |
| 6    | getSlaDetails         | Collect Information about Infra and Business Service Element Details SLA score | Bytes                       | 1. All Infra Elements, 2. All Service Elements, 3. Product and Env Wise Score, 4. Product / Env / Business/ Common Services Score, 5. App/Data/Nw Service Score                        |  |
| 7    | provision             | Provision infra and App Services on AWS               | Bytes                                       | 1. landingZone, 2. Product Enclave, 3. Cluster, 4. Product and Services in product enclave                                                                                               |    |
| 8    | diaganose             | Run Diagnostics on infra and App Services on AWS      | Bytes                                       | 1. Infra Elements, 2. Service Elements                                                                                                                                                   | |
| 9    | health                | Run Diagnostics on infra and App Services on AWS      | Bytes                                       | 1. Infra Elements, 2. Service Elements                                                                                                                                                   |    |
| 10   | secret                | Secret management of infra and service elements       | Bytes                                       | 1. Infra Elements, 2. Service Elements                                                                                                                                                   |  |


    
# Building Docker image and running it

--Build it 
docker build -t awsx .

--Run it
docker run -it awsx --help

# How SRE platform calls CLI/API 

From SUI , following is fired when we click element explorer

https://sre.synectiks.net/?elementType=EC2 , elementId=2435

It opens EC2 plugin inside SRE platform with Global Variable been equated with 2435

Panel Wise Query 

From DS UI , corresponding to elementId, elementType and list of supported Queries is shown.

/awsx-api/getSupportedQueries?elementType=EC2

It get response like 
1. cpu_utilization_panel 
2. storage_utilization_panel
3. network_utilization_panel
4. ec2-config-data

From DS UI , you select cpu_utilization_panel

DS UI then fire the queries to DS backend as follows:

elementType=EC2 , elementId=2435 , query=cpu_utilization_panel , --timeRange={}

DS backend get the value of customerRoleArn for the elementId from cmdb

DS backend get the value of instanceId for the elementId from cmdb

DS backend then calls the API 

/awsx-api/getQueryOutput? elementType=EC2, elementId="1234" , query=cpu_utilization_panel, --timeRange={}

API server will call the awsx cli command as 

awsx --vaultURL=vault.synectiks.net getElementDetails --elementId="1234" --elementType=EC2 --query="cpu_utilization_panel" --timeRange={}

Following Output is passed back from awsx to api to datasource to dashboard layer
{
	CurrentUsage:25,
	AverageUsage:30,
	MaxUsage:40
}

The dashboard then associate the output data to the panel.
