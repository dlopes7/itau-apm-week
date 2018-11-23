#!/usr/bin/env bash


NUMBER_OF_CONTAINERS=5

for i in $(seq 0 $(($NUMBER_OF_CONTAINERS - 1))); do
    echo Creating container $i

    docker run -d -p 808$i:8080 --name shopizer_$i -v /home/david/Documents/projects/java/itau-apm-week/java-agent:/opt/appdynamics/java-agent -e "CATALINA_OPTS=-javaagent:/opt/appdynamics/java-agent/javaagent.jar -Dappdynamics.agent.applicationName=Shopizer_'$HOSTNAME'_$i -Dappdynamics.agent.tierName=front -Dappdynamics.agent.nodeName='$HOSTNAME'_$i -Dappdynamics.controller.hostName=platform-davidlopessmallcon-gokodoek.srv.ravcloud.com -Dappdynamics.controller.port=8090 -Dappdynamics.agent.accountAccessKey=ravello" dlopes7/shopizer
done

