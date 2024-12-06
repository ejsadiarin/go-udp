# Go UDP Implementation for Learning Purposes

- client and server can send data to each other (connection is bidirectional) through the socket (port)
- in UDP, there is no handshakes involved and data is sent directly to the remote address

## How it works?

1. Server establishes socket with a fixed port by creating a `listener`
2. Client connects to server by targeting the server port
3. Client port (dynamic, ephemeral) is established by the OS automatically
4. To read data from the socket, a buffer `buf` (slice) with a fixed size of 1024 bytes is created (in both client and server)
    - `buf` --> to store the data read from the socket
    - then convert the bytes to a type (e.g. string) by type conversion from bytes to string, T
        - like `string(buf[:data_size])`
5. When writing data to UDP socket, `[]bytes(msg)` are sent
    - `msg` --> in this case is a string, so string is encoded before sending to the UDP socket
    - all data sent must first be encoded (converted to bytes)

## UDP Server
- fixed port
- should only be one instance since the port is fixed
- creates a socket by creating a `listener`
- needs

## UDP Client
- dynamic, ephemeral port assigned by the OS automatically
- can create many client instances --> different ports each time
- needs

