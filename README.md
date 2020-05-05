# Harness-F5Utility


## About

The purpose of this utility is to help automate steps on F5 for the purposes
of Blue/Green and Canary deployments .

I-rule definition for a canary deployment is required for the utility to work 

built and tested against BIGIP-15.1.0.2-0.0.9

## Flags

-poolname = Big Ip target poolname

-action = enable pool , disablepool or updateirule

-username = admin user

-password = admin password

-host = Big IP ip address port combination i.e. 10.1.1.170:8443

#### Example pool enable/disable 

-poolname=PoolA -action=disablepool -username=admin -password=******* -host=10.1.1.170:8443

#### Example I-rule update 

-irulename=CanaryRule -action=updateirule -username=admin -password=******* -host=10.1.1.170:8443 -irulepayload="{\"apiAnonymous\": \"when HTTP_REQUEST {\n    log local0. \\"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\\" \n   HTTP::respond 200 content \\"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\\"\n}\"}"

#### Example Pool query update (Blue/Green)

Note this querys the first node in a given pool and returns true or false for enabled/disabled 



#### A note on Escaping the JSON payload for an I-rule

On the shell step / commandline ecapsulate the JSON string with ' to escape the special characters and make it a string literal

i.e.

-irulepayload='{"apiAnonymous": "when HTTP_REQUEST {\n    log local0. \"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\" \n   HTTP::respond 200 content \"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\"\n}"}'


## Implementation in Harness

1. Create a delegate profile to pull go_build_main_go binary from the repo


2. Example Shell script step in Harness

    Harness Variables :

    ${workflow.variables.irulename} = CanaryRule

    ${workflow.variables.action} = updateirule

    ${workflow.variables.username} = admin

    ${workflow.variables.password} = *****

    ${workflow.variables.host} = 10.1.1.170:8443

    ${workflow.variables.irulepayload} = '{"apiAnonymous": "when HTTP_REQUEST {\n    log local0. \"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\" \n   HTTP::respond 200 content \"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\"\n}"}'

Harness shell script command with variables :

    Irule 

    cd /Users/gregorykroon/go/src/F5Utility/build
    ./go_build_main_go -irulename=${workflow.variables.irulename} -action=${workflow.variables.action} -username=${workflow.variables.username} -password=${workflow.variables.password} -host=${workflow.variables.host} -irulepayload=${workflow.variables.irulepayload}

    Enable pool

    cd /Users/gregorykroon/go/src/F5Utility/build
    ./go_build_main_go -poolname=PoolA -action=enablepool  -username=${workflow.variables.username} -password=${workflow.variables.password} -host=${workflow.variables.host}

    Disable pool

    cd /Users/gregorykroon/go/src/F5Utility/build
    ./go_build_main_go -poolname=PoolA -action=disablepool  -username=${workflow.variables.username} -password=${workflow.variables.password} -host=${workflow.variables.host}
