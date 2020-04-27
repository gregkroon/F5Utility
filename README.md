# Harness-F5Utility


## About

The purpose of this utility is to automate the enable and disable of F5 Big IP pool members 
for the purposes of Blue/Green style deployments .

## Flags

-poolname

-action

-username

-password

-host 

Example pool disable 

-poolname PoolA -action disablepool -username admin -password ******* -host 10.1.1.170:8443

Example Irule update 

-irulename CanaryRule -action updateirule -username admin -password ******* -host 10.1.1.170:8443 -irulepayload="{\"apiAnonymous\": \"when HTTP_REQUEST {\n    log local0. \\"[IP::client_addr]:[TCP::client_port]: Connected to [virtual name] [IP::local_addr]:[TCP::local_port]\\" \n   HTTP::respond 200 content \\"Connected to [virtual name] [IP::local_addr]:[TCP::local_port] from [IP::client_addr]:[TCP::client_port]\\"\n}\"}"