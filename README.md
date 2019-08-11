# BabyWAF (Rudimentary Web Application Firewall)

The following is a rudimentary implementation of an HTTP proxy with malware scan capabilities. The HTTP proxy is expected to load heavy loads of traffic and execute some operations against each request. If the request is found to be malicious by an internal web service the proxy returns an error to the client, otherwise the request is forwarded to the destination server.

## URL Lookup Service

We have an HTTP proxy that is scanning traffic looking for malware URLs. Before allowing HTTP connections to be made, this proxy asks a service that maintains several databases of malware URLs if the resource being requested is known to contain malware.

Write a small web service in the language/framework your choice, that responds to GET requests where the caller passes in a URL and the service responds with some information about the URL. The GET request look like this:

```
GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
```

The caller wants to know if it is safe to access that URL or not. As the implementer, you get to choose the response format and structure. These lookups are blocking users from accessing the URL until the caller receives a response from your service.

## Considerations

Give some thought to the following:

- The size of the URL list could grow infinitely, how might you scale this beyond the memory capacity of this VM? Bonus if you implement this;
- The number of requests may exceed the capacity of this VM, how might you solve that? Bonus if you implement this;
- What are some strategies you might use to update the service with new URLs? Updates may be as much as 5 thousand URLs a day with updates arriving every 10 minutes;
- Bonus points if you containerize the app.

## Initial Thoughts

- Use a [Bloom Filter](https://en.wikipedia.org/wiki/Bloom_filter) to quickly determine if a URL is benign, in which case we can finish the operation fast, otherwise run the malware identifier against it;
- LRU cache for the most common malicious URLs. Keep track of how many hits other malicious URLs get. Update the cache every few minutes with the list of URLs with more hits;
- Assume the proxy is internal which means implementing a rate limiter is not an option;
- Worst case scenario is, every URL (including query params) is different and malicious;
- Because the query string is also taken in consideration, we may be able to split the malware check in subsets of the URL. For example, check if the hostname is blacklisted, check if the hostname+port is blacklisted, if not then proceed to check the entire URL against the malware database. We could use a Bloom Filter here as well for the hostname;
- Can we use HAProxy to manage a swarm of web services? This way we can distribute the load and offer high availability. However, this doesn’t help to address the memory consumption on each server;
