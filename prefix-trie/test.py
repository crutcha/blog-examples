from ipaddress import IPv4Network, IPv4Address
import trie
import time
import random
import struct
import socket

ALL_ONES =  (2**32) - 1

with open('routes-clean.txt', 'r') as in_file:
    input = in_file.readlines()

input = [(e.split(',')[0], e.split(',')[1].strip('\n')) for e in input]

load_start = time.time()
test_trie = trie.PrefixTrie()
for entry in input:
    test_trie.insert(entry[0], entry[1])
load_end = time.time()

generated_ips = []
for i in range(0, 500):
    r = random.randint(1, 4294967295)
    t = struct.pack('!I', r)
    ips = socket.inet_ntoa(t)
    generated_ips.append(ips)

search_start = time.time()
answers = []
for i in generated_ips:
    result = test_trie.search(i)
    answers.append(result)
search_end = time.time()

naive_start = time.time()
naive_answers = []
for i, ip in enumerate(generated_ips):
    #test_ip = IPv4Address(ip)
    current_longest = None

    octets = ip.split('.')
    oct1 = int(octets[0]) << 24
    oct2 = int(octets[1]) << 16
    oct3 = int(octets[2]) << 8
    oct4 = int(octets[3])
    test_ip = oct1 + oct2 + oct3 + oct4

    for each in input:
        #test_net = IPv4Network(each[0])
        octets = each[0].split('/')[0].split('.')
        mask = int(each[0].split('/')[1])
        oct1 = int(octets[0]) << 24
        oct2 = int(octets[1]) << 16
        oct3 = int(octets[2]) << 8
        oct4 = int(octets[3])
        test_start = oct1 + oct2 + oct3 + oct4
        test_wildcard = ALL_ONES >> mask
        test_end = test_start | test_wildcard
        
        #if test_ip in test_net:
        #    current_longest = each
        #    print('{} - new lpm for {}: {}'.format(i, test_ip, test_net))

        if test_start < test_ip and test_ip < test_end:
            current_longest = each
            print('{} - new lpm for {}: {}'.format(i, ip, each[0]))

    if current_longest[0] != '0.0.0.0/0':
        naive_answers.append(current_longest[1])
    else:
        naive_answers.append([])

naive_end = time.time()

print('load time: {}'.format(load_end - load_start))
print('search time: {}'.format(search_end - search_start))
print('naive time: {}'.format(naive_end - naive_start))
