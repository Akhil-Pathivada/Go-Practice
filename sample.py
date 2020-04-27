import re
import os 

with open("/logs/logstash-plain.log" 'r') as f:
    for line in f:
        if "Ok" in line:
            print line