![Grafana](docs/logo-horizontal.png)

##  Grafana Setup


#### Checkout Grafana

Appkube-plateform (Grafana)  `Clone the appkube-platform ` repository:

    git clone https://github.com/AppkubeCloud/Appkube-Platform.git

#### Checkout Appkube Cloud Datasource

1. Checkout appkube_cloud_datasource in grafana `(xformation/plugin directory)`.
```
https://github.com/AppkubeCloud/appkube-cloud-datasource.git
```
2. build appkube_cloud_datasource.

#### Checkout Out Mage

1.  Clone `mage ` repository:
```
https://github.com/magefile/mage
```

2. Installation: Mage has no dependencies outside the Go standard library, and builds with Go 1.7 and above (possibly even lower versions, but they're not regularly tested).
```
Using Go Modules

git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```


3. go to appkube_cloud_datasource directory
```
run below command
mage v
```

#### Build Appkube Plateform (Grafana)

1. tdm gcc 
 ```
install tdm_gcc (64 bit version 5.1.0)
```
2.  run wire command
```
wire gen -tags oss ./pkg/server ./pkg/cmd/grafana-cli/runner
```
3. setup  grafana
```
go run build.go setup
```
4. build  grafana
```
go run build.go build
```

#### Run Grafana
1. yarn install
```
yarn install
```
2. yarn dev
```
yarn dev
```
3. start Grafana server
Start grafana
```
.\bin\windows-amd64\grafana-server.exe
```

### Appkube Cloud Data Source calls AWX-API

#### Query Configuration

In this Go code snippet, a query configuration is being set up for the Appkube AWSX API. The query is intended to retrieve information related to AWS resources.

```
respType := strings.ToLower(query.ResponseType)
query.Type = "url"
query.Parser = "backend"
query.URL = "http://localhost:7000/awsx-api/getQueryOutput"
fmt.Println("Appkube awsx url :" + query.URL)

urlOptions := models.URLOptions{
    Method: "Get",
    Params: []models.URLOptionKeyValuePair{
        {"responseType", "frame"},
        {"zone", "us-east-1"},
        {"externalId", "DJ6@a8hzG@xkFwSvLmkSR5SN"},
        {"crossAccountRoleArn", "arn:aws:iam::657907747545:role/CrossAccount"},
        {"elementType", "AWS/EC2"},
        {"instanceID", "i-05e4e6757f13da657"},
        {"query", "cpu_utilization_panel"},
    },
}
```
#### Explanation
```
- The respType variable is set to the lowercased value of query.ResponseType.
- Various properties of the query object are configured, such as Type, Parser, and URL.
- The URL is set to "http://localhost:7000/awsx-api/getQueryOutput".
- A set of parameters is defined within the urlOptions struct for the HTTP GET request, including the response type, AWS zone, external ID, cross-account role ARN, element type, instance ID, and the specific query.
```

### AWX-API Intgrate The AWSX Cli

The `GetCpuUtilizationPanel` function is an HTTP handler written in Go. It processes HTTP requests, extracts parameters from the URL query, and performs some actions based on those parameters.

#### Parameters

- **region**: The region parameter obtained from the URL query.
- **cloudElementId**: Cloud element ID obtained from the URL query.
- **cloudElementApiUrl**: Cloud element API URL obtained from the URL query.
- **crossAccountRoleArn**: Cross-account role ARN obtained from the URL query.
- **externalId**: External ID obtained from the URL query.
- **responseType**: Response type obtained from the URL query.
- **filter**: Filter obtained from the URL query.
- **instanceID**: Instance ID obtained from the URL query.
- **elementType**: Element type obtained from the URL query.
- **startTime**: Start time obtained from the URL query.
- **endTime**: End time obtained from the URL query.

#### Execution

The function creates a `CommandParam` struct and populates it with the provided parameters. It then attempts to authenticate using the `DoAuthenticate` function.

If authentication is successful (`authFlag` is true), the function creates a `cobra.Command` and sets up persistent flags for additional parameters. These additional parameters are related to instance ID, element type, start time, end time, and response type.

#### Usage

The function is designed to be an HTTP endpoint for obtaining CPU utilization panel data. Clients can make HTTP requests with the required parameters in the query string to retrieve the desired information.

#### Example Request

```
GET /cpu-utilization-panel?zone=us-east-1&cloudElementId=123&cloudElementApiUrl=https://example.com/api&crossAccountRoleArn=arn:aws:iam::123456789012:role/CrossAccountRole&externalId=abcd1234&responseType=json&filter=example_filter&instanceID=i-1234567890abcdef0&elementType=example_type&startTime=2022-01-01T00:00:00Z&endTime=2022-01-02T00:00:00Z

```

###  AWSX Cli (Output)

### Overview

cli to monitors AWS resources using cloudwatch metric queries. It is written in Go and customizable with various parameters.

### Prerequisites
- Go installed.

### Command Details
```
go run awsx-getelementdetails.go --zone=us-east-1 --externalId=DJ6@a8hzG@xkFwSvLmkSR5SN --crossAccountRoleArn=arn:aws:iam::657907747545:role/CrossAccount --elementType="AWS/EC2" --instanceID="i-05e4e6757f13da657" --query="cpu_utilization_panel"
```
### Output
```
{"AverageUsage":20.0125,"CurrentUsage":4,"MaxUsage":23.075}
```
### Command Parameter:
- --crossAccountRoleArn: AWS IAM role ARN for cross-account access.
- --elementType: Specify the type of the AWS resource element that you want to monitor. In the provided example, it is set to "AWS/EC2."
- --instanceID: Provide the instance ID of the AWS resource you want to monitor. In the example, it is set to "i-05e4e6757f13da657."
- --query: Define the metric query to be executed on the specified AWS resource. In the example, it is set to "CPUUtilization."
- --timeRange: Set the time range for the metric query. Replace the "{}" placeholder with the desired time range.
    
### Logic to get GLOBAL_AWS_SECRETS (access/secret key) in cli: 
        Since we are only passing crossAccountRoleArn, we need GLOBAL_AWS_SECRETS (access/secret key) from vault. It can be retrieved by two ways explaind below: 
            1. make vault call with static key (GLOBAL_AWS_SECRETS)
            2. If vault is not available, get the GLOBAL_AWS_SECRETS from environment variable
            3. If environment variable is not available, get the GLOBAL_AWS_SECRETS from command line
            4. If command line variable is not available, get the GLOBAL_AWS_SECRETS from cmd database 
            5. If GLOBAL_AWS_SECRETS not found in cmd database, program should exit with error - clien connection could not be established. access/secret key not found

#### flags
```
--cloudElementId: Cloud element ID.
--cloudElementApiUrl: Cloud element API URL.
--vaultUrl: Vault endpoint.
--vaultToken: Vault token.
--accountId: AWS account number.
--zone: AWS region.
--accessKey: AWS access key.
--secretKey: AWS secret key.
--crossAccountRoleArn: AWS cross-account role ARN.
--externalId: AWS external ID.
--cloudWatchQueries: AWS CloudWatch metric queries.
--elementType: Element type.
--instanceID: Instance ID.
--query: Query.
--startTime: Start time for metrics.
--endTime: End time for metrics.
--responseType: Response type. Options: json/frame.
```
#### Subcommands

AwsxEc2CpuUtilizationCmd: Subcommand specific to AWS EC2 CPU utilization metrics.

#### Note
```
The AwsxEc2CpuUtilizationCmd subcommand is available for fetching CPU utilization metrics specifically for AWS EC2 instances.
The responseType flag allows you to specify the format of the output, with options for JSON or frame.
```