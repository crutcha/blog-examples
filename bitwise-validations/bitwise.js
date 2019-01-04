var examples = [
    '192.168.1.1/24',
    '192.168.1.4/30'
]

examples.forEach(function(ipAddr) {
    var validIP = ipAddr.match(/^(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\/(8|[1-2][1-9]|30)$/);

    if (validIP) {
        // Capture each octet and bitwise shift
        var o1 = (+validIP[1] << 24) >>> 0;
        var o2 = (+validIP[2] << 16) >>> 0;
        var o3 = (+validIP[3] << 8) >>> 0;
        var o4 = (+validIP[4]) >>> 0;
        var ipAsInt = (o1 + o2 + o3 + o4) >>> 0;
        var mask = -1 << (32 - validIP[5]);

        var start = (ipAsInt & mask) >>> 0;
        var end = (ipAsInt | ~mask) >>> 0;
        var result = start < ipAsInt && ipAsInt < end;
        console.log('---');
        console.log('Checking IP: ' + ipAddr);
        console.log('start: ' + start);
        console.log('end : ' + end);
        console.log('ip : ' + ipAsInt);
        console.log('result: ' + result);
    } else {
        alert(ipAddr + ' is not a valid IPv4 address.');
    }
});
