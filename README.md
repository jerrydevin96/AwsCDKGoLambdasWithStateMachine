# AWS CDK With Go
This is a sample project which deploys two lambda functions and a statemachine(step functions) which calls the two lambdas. The project is completely written in Go, with the CDK being implemented in Go as well.  

## Resources in the Project:
### Lambda One:
Lambda, which takes in a json with first name as input and returns a json with last name as output. It along with it's execution role is deployed as a separate stack through the CDK.  

`Input Json:`
```
    {
        "firstName": "any-name"
    }
```
   
### Lambda Two:
Lambda, which takes in a json with the last name  as input and prints the last name. It along with it's execution role is deployed as a separate stack through the CDK.  

`input Json`
```
    {
        "lastName": "any-name"
    }
```
  
### State Machine:
The statemachine simply links the two lambda's along with a small wait duration. The statemachine along with it's execution role is deployed as a separate stack through the cdk.

### Deployment:
This folder contains the code implementing the cdk. The code is modularized into packages corresponding to the resources being deployed. The templates package contains the template funcitons which creates the stacks for the lambdas and the statemachine. The resources folder contain the actual data that needs to be passed to create every stack.
  
## Deployment Steps:
1. Build the lambdas:   
`make buildlambdaone`   
`make buildlambdatwo`
2. Duild the cdk binary:  
`make buildcdk`
3. Download aws cdk (requires nodejs installation)  
`npm install -g aws-cdk`
4. Use aws cli to login to your aws account (admin access recommended)
5. Export environment variables     
`export ENV=envone || export ENV=envtwo` (the code here is designed to deploy to two separate AWS accounts)  
`export ACCTNO=<aws acct number>` (account number for your relevant env). Region defaults to `eu-west-1`. This can be changed from code.
6. Change to deployment folder
7. Booststrap your CDK environment (has to be done the first time)  
`cdk bootstrap`
8. Synthesize the deployment folder to generate cloudformation stacks  
`cdk synthesize` (this will generate `cdk.out` folder)
9. List stacks for the environment to verify  `cdk list`
10. Deploy all stacks or the preferred stack    
`cdk deploy --all` (deploys all stacks for the env)  
`cdk deploy <stack name from cdk list>` (deploys specified stack for the env)
11. Check stack and resource creation in AWS console