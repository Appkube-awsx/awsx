
## AWSX-CLI Documentation
### 1. AWSX ECS CPU Utilization Panel  

### Overview

cli to monitors AWS resources using cloudwatch metric queries. It is written in Go and customizable with various parameters.

### Prerequisites
- Go installed.

### Command Details
```
go run awsx-getelementdetails.go --zone=us-east-1 --externalId=DJ6@a8hzG@xkFwSvLmkSR5SN --crossAccountRoleArn=arn:aws:iam::657907747545:role/CrossAccount --elementType="ContainerInsights --clusterName="myclustTT" --query="cpu_utilization_panel"
```
### Output
```
{"AverageUsage":20.0125,"CurrentUsage":4,"MaxUsage":23.075}
```
### Command Parameter:
- --crossAccountRoleArn: AWS IAM role ARN for cross-account access.
- --elementType: Specify the type of the AWS resource element that you want to monitor. In the provided example, it is set to "ContainerInsights."
- --clusterName: Provide the clusterName of the AWS resource you want to monitor. In the example, it is set to "myclustTT."
- --query: Define the metric query to be executed on the specified AWS resource. In the example, it is set to "node_cpu_utilization."
- --timeRange: Set the time range for the metric query. Replace the "{}" placeholder with the desired time range.
    
### Logic to get GLOBAL_AWS_SECRETS (access/secret key) in cli: 
        Since we are only passing crossAccountRoleArn, we need GLOBAL_AWS_SECRETS (access/secret key) from vault. It can be retrieved by two ways explaind below: 
            1. make vault call with static key (GLOBAL_AWS_SECRETS)
            2. If vault is not available, get the GLOBAL_AWS_SECRETS from environment variable
            3. If environment variable is not available, get the GLOBAL_AWS_SECRETS from command line
            4. If command line variable is not available, get the GLOBAL_AWS_SECRETS from cmd database 
            5. If GLOBAL_AWS_SECRETS not found in cmd database, program should exit with error - clien connection could not be established. access/secret key not found

## AWSX (Parent Command) 

### 1. AWSX ECS CPU Utilization Panel  
Aws extension CLI \
awsx is the main parent command.

#### Command structure
    <main command> <persistant flags> <sub-command> <sub-command flags>
    e.g. awsx --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx  getElementDetails  --zone=us-west-2
    
    awsx:               It is the main parent command
    --vaultURL:         It is persistant flag which pass the vault url to the sub-command. Vault is a separate server which store user cloud credentials in key=value pair in encrypted format. \
    --accountId:        This persistant flag pass the AWS account id to the sub-command. Sub-command pass this account id as key to the vault serve to get user credentials \
    getElementDetails:  It is the sub-command of parent awsx
    --zone:             Sub-command flag which provide AWS region to the sub-command

#### How to download sub-command/module from git
    Open command prompt and go to the directory where go.mod file available
    * To download any specific version of a module, run go get command with version number
        go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.1
    * To download latest version of a module, run go get command with @latest
        go get github.com/Appkube-awsx/awsx-cloudelements@latest
        - If we omit the version number or @latest, go get command downloads the latest version always

#### How to embed sub-command in GO CLI code
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

#### Compile GO program: go build
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

#### Install GO program: go install    
    * The go install command behaves almost identically to go build, but instead of leaving the executable in the current directory, or a directory specified by the -o flag, it places the executable into the $GOPATH/bin directory.
    * The benefit of go install is that the executable binary can be executed outside of the source directory. 
      $ go install
        NOTE - The go install command does not support the -o flag, so it will use the default name described earlier to name the executable.

#### Compile and run the go program with go run command
    * Open command prompt and go to the project directory where a go file having main() function available
        cd /opt/mycode/awsx
    * Run the program with go run command
        go run ./main.go
        - In our example, our main() function is available in main.go
        - We can provide program arguments or flags to the main() function
            go run ./main.go arg1 arg2
            go run ./main.go --flag1=value --flag2=value



#### [Type of flags in Cobra CLI](https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1)
    Cobra has two types of flags:

    *    Persistent flags - available to the command it is assigned to, as well as all its sub-commands
    *    Local flags - only assigned to a specific command

### Command Details
```
awsx cpu_utilization_panel --zone=us-east-1 --externalId=DJ6@a8hzG@xkFwSvLmkSR5SN --crossAccountRoleArn=arn:aws:iam::657907747545:role/CrossAccount --elementType="ContainerInsights" --clusterName="myclustTT" --query="node_cpu_utilization"
```
### Output
```
{"AverageUsage":20.0125,"CurrentUsage":4,"MaxUsage":23.075}
```
### Command Parameter:
- --crossAccountRoleArn: AWS IAM role ARN for cross-account access.
- --elementType: Specify the type of the AWS resource element that you want to monitor. In the provided example, it is set to "ContainerInsights."
- --clusterName: Provide the clusterName of the AWS resource you want to monitor. In the example, it is set to "myclustTT".
- --query: Define the metric query to be executed on the specified AWS resource. In the example, it is set to "node_cpu_utilization."
- --timeRange: Set the time range for the metric query. Replace the "{}" placeholder with the desired time range.
    
### Logic to get GLOBAL_AWS_SECRETS (access/secret key) in cli: 
        Since we are only passing crossAccountRoleArn, we need GLOBAL_AWS_SECRETS (access/secret key) from vault. It can be retrieved by two ways explaind below: 
            1. make vault call with static key (GLOBAL_AWS_SECRETS)
            2. If vault is not available, get the GLOBAL_AWS_SECRETS from environment variable
            3. If environment variable is not available, get the GLOBAL_AWS_SECRETS from command line
            4. If command line variable is not available, get the GLOBAL_AWS_SECRETS from cmd database 
            5. If GLOBAL_AWS_SECRETS not found in cmd database, program should exit with error - clien connection could not be established. access/secret key not found


## AWSX-API
### 1. AWSX ECS CPU Utilization Panel API
## Overview
This Go code defines an HTTP handler function for retrieving CPU utilization metrics for ECS cluster in the AWS cloud. The API supports both direct authentication using AWS credentials and cross-account authentication.

## API Endpoint
- **Endpoint:** `/awsx-api/getQueryOutput`
- **HTTP Method:** `GET`
This markdown file contains all api document Order-wise how does flow works of ECS CPU Utilization Panel

	baseMetricUrl:
		http://localhost:7000

## API Reference
The ECS CPU Utilization Panel is organized around REST. Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.

## Errors

ECS CPU Utilization Panel uses conventional HTTP response codes to indicate the success or failure of an API request. In general: Codes in the 2xx range indicate success. Codes in the 4xx range indicate an error that failed given the information provided (e.g., a required parameter was omitted, a charge failed, etc.). Codes in the 5xx range indicate an error with Stripe's servers (these are rare).

Some 4xx errors that could be handled programmatically (e.g., a card is declined) include an error code that briefly explains the error reported.

 ## HTTPS STATUS CODE SUMMRY

Code   | Summary
------------- | -------------
200 - OK  | Everything worked as expected.
400 - Bad Request  | The request was unacceptable, often due to missing a required parameter.
401 - Unauthorized | No valid API key provided.
402 - Request Failed | The parameters were valid but the request failed.
403 - Forbidden | The API key doesn't have permissions to perform the request.
404 - Not Found | The requested resource doesn't exist.
409 - Conflict | The request conflicts with another request (perhaps due to using the same idempotent key)
429 - Too Many Requests | Too many requests hit the API too quickly. We recommend an exponential backoff of your requests.
500, 502, 503, 504 - Server Errors | Something went wrong on Stripe's end. (These are rare.)

## Request Parameters
- `zone`: AWS region or availability zone.
- `cloudElementId`: ID of the AWS cloud element (optional).
- `cloudElementApiUrl`: URL for AWS cloud element API (optional).
- `cluster`: Name of the ECS cluster.
- `elementType`: Type of the AWS element (e.g., ContainerInsights).
- `query`: Query string for metric data.
- `startTime`: Start time for the metric data retrieval (optional).
- `endTime`: End time for the metric data retrieval (optional).
- `statistic`: Statistic types for metric data (comma-separated, optional).
- `crossAccountRoleArn`: Cross-account role ARN for authentication (if using cross-account authentication).
- `externalId`: External ID for cross-account authentication.

## Curl Command 
```
curl --location http://localhost:7000/awsx-api/getQueryOutput?zone=us-east-1&externalId=DJ6@a8hzG@xkFwSvLmkSR5SN&crossAccountRoleArn=arn:aws:iam::657907747545:role/CrossAccount&elementType=ContainerInsights&clusterName=myclustTT&query=cpu_utilization_panel&responseType=json
```

## Output
```
{
    "AverageUsage": 14.461583914520872,
    "CurrentUsage": 10,
    "MaxUsage": 20.62193904108335
}

```


## Appkube-Platform (Grafana) Documentation


#### Purpose

The `testAppkubeCputUtilization` function is a Go function designed to query CPU utilization data using the AppKube API and Infinity client.

#### Parameters

- `zone` (string): The AWS region/zone to query (e.g., "us-east-1").
- `externalId` (string): External identifier for authentication.
- `crossAccountRoleArn` (string): ARN (Amazon Resource Name) of the cross-account IAM role.
- `elementType` (string): Type of the AWS resource (e.g., "ContainerInsights").
- `clusterName` (string): Name of the specific AWS cluster.
- `query` (string): Query type, in this case, "node_cpu_utilization".
- `statistic` (string): Statistic type, in this case, "SampleCount".

#### Usage

```
// Example Usage
testAppkubeCputUtilization()
```
#### Dependencies
- Infinity package for creating an Infinity client.
- The pluginhost and backend packages for querying data and plugin context.

#### Error Handling

If an error occurs while creating the Infinity client, an error message will be printed to the console.

#### Notes

- Ensure that the necessary dependencies are installed before using the function.
- This function assumes a specific JSON format for the query and sends it to the AppKube API.

Example JSON Query
```
{
    "type": "appkube-api",
    "source": "url",
    "productId": 1,
    "environmentId": 2,
    "moduleId": 2,
    "serviceId": 2,
    "serviceType": "java app service",
    "zone": "us-east-1",
    "externalId": "657907747545",
    "crossAccountRoleArn": "arn:aws:iam::657907747545:role/CrossAccount",
    "elementType": "ContainerInsights",
    "clusterName": "myclustTT",
    "query": "node_cpu_utilization",
    "statistic": "SampleCount"
}

```

#### Response

The function prints the response frames obtained from the AppKube API.

```
fmt.Println("Response: ", res.Frames)

```
## Acknowledgements

 - [Awesome Readme Templates](https://awesomeopensource.com/project/elangosundar/awesome-README-templates)
 - [Awesome README](https://github.com/matiassingers/awesome-readme)
 - [How to write a Good readme](https://bulldogjob.com/news/449-how-to-write-a-good-readme-for-your-github-project)


## API Reference

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.


## Appendix

Any additional information goes here

