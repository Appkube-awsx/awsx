




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