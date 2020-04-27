# Harness-F5Utility


## About

The purpose of this utility is to help automate steps on F5 for the puposes
of Blue/Green and Canary deployments .

I-rule definition for a canary deployment is required for teh utility to work 

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



Utility version

"{\"apiAnonymous\": \"when HTTP_REQUEST {\n    log local0. \\"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\\" \n   HTTP::respond 200 content \\"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\\"\n}\"}"

## Implementation in Harness


