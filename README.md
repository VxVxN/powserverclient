To implement the PoW mechanism in this server, I used Partial Hash Inversion.

Reasons for choosing the algorithm:

1. **Simplicity:** The algorithm is very straightforward to implement. 
2. **Versatility:** It can be adapted to various needs. For example, the difficulty of the task can be adjusted by changing the number of leading zeros that the hash function must contain. This allows for load balancing and adaptation to changing conditions, such as an increase in the number of users. 
3. **Security:** Methods based on hash inversion provide a high level of security since they require significant time investment to find a suitable input, making them resistant to DoS attacks. The complexity of hashing also helps protect the server from potential attackers looking to overload the system.

To run:
```bash
  docker compose up --build
```