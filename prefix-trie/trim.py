# Trim output from JunOS BGP output and save to file where each line consists
# of a comma seperated value of CIDR address and next-hop.

import re

cidr_regex = ('^(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)'
              '\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d'
              ')\/\d+')
ip_regex   = ('(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)'
              '\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)')

output = []
with open('route.txt', 'r') as rfile:
    to_parse = rfile.readlines()

    current_pos = 0
    end_pos = len(to_parse)
    while current_pos < end_pos:
        ip_matched = re.search(cidr_regex, to_parse[current_pos])
        if ip_matched:
            # Next hop will be 2 lines down in JunOS CLI output for BGP. This wont catch
            # every single route but it will catch any route that's sourced from BGP.
            current_pos += 2
            next_hop_matched = re.search(ip_regex, to_parse[current_pos])
            #import pdb; pdb.set_trace()
            if next_hop_matched:
                output.append(ip_matched.group(0)+ ',' + next_hop_matched.group(0) + '\n')
                # We can skip 2 ahead if we matched both
                current_pos += 1
        else:
            current_pos += 1

with open('routes-clean.txt', 'w') as out_file:
    out_file.writelines(output)

