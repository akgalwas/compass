# Credentials management in Provisioner

## Introduction

The goal of this document is to consider what should be next steps in Provisioner to improve credentials handling. 

## Curent state

- Credentials for GCP (Service Account Key) are extracted from a secret provided in ***CredentialsInput*** object (being a part of ***ProvisionRuntimeInput***)
-  Credentials for Gardener are extracted from two sources:
  - Gardener Kubeconfig is extracted from a secret provided in ***CredentialsInput*** object
  - Secrets to be used for provisioning are extracted from a secret specified in ***GardenerConfigInput*** 

## Questions to answer

- What changes to Provisioner API need to be introduced?
- Do we need to introduce a new micro service ([Hyperscaler Account Pool](./hyperscaler-account-pool-api-design.md)). 
  - Can we skip it in a short term and implement a component inside Provisioner? 
  - Is it vital in a long term?  

## Data needed for obtaining credentials

It was identified that the following data should be provided by the API user in Cockpit scenario:

- what type of Hyperscaler to use: e.g. GCP
- Account name (extracted from header)
- Credential Name

The following data need to be provided by the API user in SCP scenario:

- what type of Hyperscaler to use: e.g. Gardener-Azure

- Account name (extracted from Header)

- Credential Name : it should be discussed whether Environment Broker will have enough knowledge to pass credential name to the Provisioner ; there is a question whether HAP should be able to determine what is the Target Provider Secret Name 

## Reading credential data from Gardener

In the current solution kubeconfig associated with Gardener project is stored in a secret, and the secret name is passed    in ***Credentials Input***. 

What needs to be done is picking right Gardener project and credentials based on account name, type of provider and credential name. The first idea is to search credentials based on labels but the Gardener kubeconfig with appropriate priviliges is needed.   

## Proposed solution

### Proposed Provisioner API change

The following changes in Provisioner API would improve the design:

- ***secretName*** replaced with ***credentialName***.
- ***targetSecret*** removed from ***GardenerConfigInput***.

Current design is based on direct specifying a secret to be used. Specifying credential name instead is a step towards resolving credetnials in Provisioner (with HAP or internal component).

Additionally Provisioner should handle account header. 

In the first step the implementation will work as follows:

- For GCP ***credentialName*** will be treated as secret name in the cluster Provisioner is deployed (the logic of extracting secret data will remain the same)
- For Gardener account name extracted from header will be used for determining Gardener project name. Gardener secret name will be specified in ***credentialName***. We could consider using some default for credential name

In the next steps above logic will be changed (credential could be saved in some storage, functionality of adding credentials could be added).

### Introducing separate microservice for credentials management (HAP)

The microservice will be responsible for storing and returning credentials. The API schema is defined ([here](./hyperscaler-account-pool-api-design.md)). 

Pros

- Clear responsibilities : provisioner responsible for Kyma installation, HAP responsible for credentials management. Both services have different reason for change: we ned to change Provisioner code only if we change the way clusters are provisioned/deptovisioned. When we change the way credentials are extracted based on credential name there is no need to change Provisioner.

Cons

- Bigger implementation effort

### Introducing internal component 

Pros

- Smaller implementation effort

Cons

- Provisioner will have additional responsibility. If we will need to enhance credentials extraction mechanism (e.g. add functionality of adding new credentials) it will make Provisioner code more complex.

   

  

### 
