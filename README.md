# HTTP Access Log Analyzer

There is a large amount of HTTP requests made to the NFIS servers from all around the world, including requests that are from outside of Canada. 

Given the Apache HTTP access log, the program will scan through the log and: <br>
  - find the top 10 cities outside of Canada where the requests to the NFIS servers are being made from; <br>
  - find the number of HTTP requests for those cities;<br>
  - find the city that is the furthest away from Victoria, BC, of the top 10 cities
