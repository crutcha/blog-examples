from ipaddress import IPv4Network, IPv4Address

class PrefixTrieNode(object):
    
    def __init__(self, prefix=''):
        self.left = None
        self.right = None
        self.prefix = prefix
        # Since there WILL be collisions, make this an array, it's up to the 
        # caller to handle logic on collisions
        self.value = []
    
    def __repr__(self):
        return '<PrefixTrieNode Prefix: {}*>'.format(self.prefix)

class PrefixTrie(object):
    
    def __init__(self):
        self.root = PrefixTrieNode()
    
    def _child_node_count(self, node):
        # These should cover base cases
        right_count = 0
        left_count = 0
        if node.left:
            left_count = 1 + self._child_node_count(node.left)
        if node.right:
            right_count = 1 + self._child_node_count(node.right)
        
        return left_count + right_count

    def __len__(self):
        return self._child_node_count(self.root)

    def _insert(self, bit_array, value):
        current_node = self.root
        prefix_len = len(bit_array)
        for i, bit in enumerate(bit_array):
            if int(bit) == 0:
                if not current_node.left:
                    current_node.left = PrefixTrieNode(current_node.prefix + bit)
                current_node = current_node.left
            elif int(bit) == 1:
                if not current_node.right:
                    current_node.right = PrefixTrieNode(current_node.prefix + bit)
                current_node = current_node.right

            if i == (prefix_len - 1):
                current_node.value.append(value)

    def insert(self, cidr, nexthop):
        '''
        Convert the hex output from EIP into binary equivalent, calculate
        network boundary based on bitwise shifts on wildcard mask, and 
        insert into trie at proper spot.
        '''

        network, mask = cidr.split('/')

        octets = network.split('.')
        oct1 = int(octets[0]) << 24
        oct2 = int(octets[1]) << 16
        oct3 = int(octets[2]) << 8
        oct4 = int(octets[3])

        boundary_int = oct1 + oct2 + oct3 + oct4
        boundary = '{0:032b}'.format(boundary_int)[:int(mask)]
        self._insert(boundary, nexthop)

    def _search(self, bit_array):
        '''
        Given IPv4 address in binary representation as bit array, search
        trie and return longest prefix match or None if not found.
        '''

        current_node = self.root
        current_longest = None
        for i, bit in enumerate(bit_array):
            if int(bit) == 0:
                current_node = current_node.left
            elif int(bit) == 1:
                current_node = current_node.right

            if not current_node:
                return current_longest.value
            else:
                current_longest = current_node
        

    def search(self, ip_addr):
        '''
        Given an IPv4 address as a string, convert to binary representation
        then search trie.
        '''

        # Let's not reinvent the wheel, ipaddress module is good
        # enough right here. This will also validate input IP.
        ip_object = IPv4Address(ip_addr)
        ip_bits = '{0:032b}'.format(int(ip_object))

        return self._search(ip_bits)
