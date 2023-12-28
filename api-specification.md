
- [awsx-api](#awsx-api)
- [getLandingZoneDetails subcommand details](#getlandingzonedetails-subcommand-details)
- [getElementDetails subcommand details](#getelementdetails-subcommand-details)
  
# awsx-api
AWSX is the modular command line CLI for for Appkube platform. All the subcommands are written as plugins for the main commands.Appkube job engine calls those CLI commands for all its supported automation jobs. The commands are also directly embeddded in AWSX-Api server. 
Please refer to Appkube Architecture diagram for details on how the appkube platform calls this commands.

Here follows the list of subcommands:
1. getLandingZoneDetails
2. getLandingZoneCompliance
3. getElementDetails
4. getCostDetails
5. getBusinessServiceDetails
6. getSlaDetails
7. provision
8. diaganose
9. health
10. secret

All the supported subcommands and there source code locations are as follows:

| S.No | Sub-command     |                Description                             |      Repository        |            functionalities                                                                                                         |
|------|-----------------------|----------------------------------------------------|-------------------|-------------------------------------------------------------------------------------------------------------------------|
| 1    | getLandingZoneDetails       | Collect Information about any specific landing zone  |    Percentage(%)  | 1. Get Elements Metadata,<br /> 2. Get List of every elements with their config infos <br /> 3. Get List of products and Environments for the landing Zone<br /> 4. Get the cost of the landing zone                                                                                           |
| 2    |  getLandingZoneCompliance  | Collect Information about any specific landing zone compliances and security      |     Bytes         | 1. Get overall account Compliance ,<br /> 2.Get the elementWise Compliance <br /> 3. Run Compliance on products and environments                                                             |
| 3    |  getElementDetails | Collect Information about specific cloud elements -Run Queries |     Bytes         | 1.EC2,<br /> 2.EKS <br /> 3. ECS <br /> 4.LAMBDA <br /> 5. API Gw <br /> 6.Load Balancer<br />|
| 4    |  getCostDetails | Collect Information about account and elements specific costs |     Bytes         | 1.Total Account <br /> 2.Product and Envwise <br /> 3. Element Wise <br /> 4.Spikes and Trends <br /> 5. App/Data/Nw Service wise Costs <br />|
| 5    |  getBusinessServiceDetails | Collect Information about the business Services that is hosted on infrastructure |     Bytes         | 1.Account Wise infralist and corresponding services <br /> 2.Product and Env wise Service Lists along with hosted infrastructure <br /> |
| 6    |  getSlaDetails | Collect Information about Infra and Business Service Element Details SLA score |     Bytes     | 1.All Infra Elements <br /> 2.All Service Elements <br /> 3. Product and Env Wise Score <br /> 4.Product / Env / Business/ Common Services Score <br /> 5. App/Data/Nw Service Score <br />|
| 7    |  provision | Provision infra and App Services on AWS |     Bytes     | 1.landingZone <br /> 2.Product Enclave <br /> 3. Cluster <br /> 4.Product and Services in product enclave  <br /> |
| 8    |  diaganose | Run Diagnostics on  infra and App Services on AWS |     Bytes     | 1.Infra Elements<br /> 2.Service Elements <br />  |
| 9    |  health | Run Diagnostics on  infra and App Services on AWS |     Bytes     | 1.Infra Elements<br /> 2.Service Elements <br />  |
| 10    |  secret | secret management of infra and service elements |     Bytes     | 1.Infra Elements<br /> 2.Service Elements <br />  |


Please refer the specification of subcommands for every subcommands input/ output / algos.

# getLandingZoneDetails subcommand details

| S.No | CLI Spec|  Description                           
|------|----------------|----------------------|
| 1    | awsx --vaultURL=vault.synectiks.net getLandingZoneDetails --zoneId="1234" --getElementsMetadata  | This will get the metadata from AWS config service |
| 2    | awsx --vaultURL=vault.synectiks.net getLandingZoneDetails --zoneId="1234" --getElementsDetails  | This will get all the element details from AWS config service |
| 3    | awsx --vaultURL=vault.synectiks.net getLandingZoneDetails --zoneId="1234" --getProductDetails  | This will get all the product and environment details that is hosted in that landing zone |
| 4    | awsx --vaultURL=vault.synectiks.net getLandingZoneDetails --zoneId="1234" --getCostDetails  | This will get all the cost details for the landing zone |

# getElementDetails subcommand details

This subcommand will need to take care for all the cloud elements and for every element, we need to support the composite 
method like network_utilization_panel. So , we can keep a single repo for the subcommand and keep separate folders for the different element handlers.

| S.No | CLI Spec|  Description                           
|------|----------------|----------------------|
| 1    | awsx --vaultURL=vault.synectiks.net getElementDetails --elementId="1234" --elementType=EC2 --query="ec2-config-data"  | This will get the specific EC2 instance config data |
| 2    | awsx --vaultURL=vault.synectiks.net getElementDetails --elementId="1234" --elementType=EC2 --query="cpu_utilization_panel"  | This will get the specific EC2 instance cpu utilization panel data in hybrid structure |
| 3    | awsx --vaultURL=vault.synectiks.net getElementDetails --elementId="1234" --elementType=EC2 --query="storage_utilization_panel" | This will get the specific EC2 instance storage utilization panel data in hybrid structure|
| 4    | awsx --vaultURL=vault.synectiks.net getElementDetails --elementId="1234" --elementType=EC2 --query="network_utilization_panel"  | This will get the specific EC2 instance network utilization panel data in hybrid structure |
