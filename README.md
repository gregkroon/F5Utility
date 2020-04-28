# Harness-F5Utility


## About

The purpose of this utility is to help automate steps on F5 for the purposes
of Blue/Green and Canary deployments .

I-rule definition for a canary deployment is required for the utility to work 

## Flags

-poolname

-action

-username

-password

-host 

#### Example pool disable 

-poolname PoolA -action disablepool -username admin -password ******* -host 10.1.1.170:8443

#### Example I-rule update 

-irulename CanaryRule -action updateirule -username admin -password ******* -host 10.1.1.170:8443 -irulepayload="{\"apiAnonymous\": \"when HTTP_REQUEST {\n    log local0. \\"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\\" \n   HTTP::respond 200 content \\"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\\"\n}\"}"


#### A note on Escaping the JSON payload for an I-rule

#### Example

Original

{"apiAnonymous": "when HTTP_REQUEST {\n    log local0. \"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\" \n   HTTP::respond 200 content \"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\"\n}"}

On the shell step / commandline ecapsulate this string with '' to escape the special characters and make it  string literal

i.e.

./go_build_main_go -irulename CanaryRule -action updateirule -username admin -password ****** -host 10.1.1.170:8443 -irulepayload='{"apiAnonymous": "when HTTP_REQUEST {\n    log local0. \"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\" \n   HTTP::respond 200 content \"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\"\n}"}'


## Implementation in Harness

1. Create a delegate profile called golang with the commands below 
   and add it to your delegate 

sudo apt-get update
sudo apt-get -y upgrade
cd /tmp
wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
sudo tar -xvf go1.14.2.linux-amd64.tar.gz
sudo mv go /usr/local


2. Example Shell script step in Harness

Harness Variables :

${workflow.variables.irulename} = CanaryRule
${workflow.variables.action} = updateirule
${workflow.variables.username} = admin
${workflow.variables.password} = *****
${workflow.variables.host} = 10.1.1.170:8443
${workflow.variables.irulepayload} = '{"apiAnonymous": "when HTTP_REQUEST {\n    log local0. \"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\" \n   HTTP::respond 200 content \"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\"\n}"}'

Shell script command with variables :


cd /Users/gregorykroon/go/src/F5Utility/build
./go_build_main_go -irulename ${workflow.variables.irulename} -action ${workflow.variables.action} -username ${workflow.variables.username} -password ${workflow.variables.password} -host ${workflow.variables.host} -irulepayload=${workflow.variables.irulepayload}