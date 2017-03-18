# udp-to-tcp

`udp-to-tcp` connects to a remote TCP connect, listens for UDP packets and sends them to said TCP connection.

This application is not maintained, and as used as for a single specific purpose and it's served its purpose. It
was used to assist in adding some reliability to video stream over a slightly lossy link due to intermittent congestion.

Note, many others may use `netcat` or similar applications to achieve this functionality.

Further investigation needed to determine exact behaviour of TCP connection, such as behaviour with Nagle's algorithm.

# Usage

In the following example, this application listens on port udp/44001, sending to some destination over tcp/44011.
`netcat` then listens on tcp/44011 and sends to the final destination on udp/44001.

UDP to TCP:
```
while true; do
    udp-to-tcp -listen :44001 -backend <tcp_destination>:44011;
    sleep 1;
done
```

TCP to UDP:
```
nc -k -l 0.0.0.0 44011 | nc -u <udp_destination> 44001
```
