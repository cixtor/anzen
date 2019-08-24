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
- Google Safe Browsing API uses [gRPC Transcoding](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto) syntax for communication. For the sake of simplicity, this web service will return a JSON encoded object with some information associated to the URL.

## Google Safe Browsing

After some thinking, I realized that this exercise is exactly what [Google Safe Browsing](https://en.wikipedia.org/wiki/Google_Safe_Browsing) is. They also use a Bloom Filter in Google Chrome to determine when to send a URL to the web service for a full scan or not. So I think implementing the algorithm in this exercise makes more sense now.

Some of the original engineers and managers in the Safe Browsing team _—led by Niels Provos—_ have already moved to other companies or teams inside Google:

- [Niels Provos](https://www.linkedin.com/in/nielsprovos/), Head of Security at Stripe
- [Panos Mavrommatis](https://www.linkedin.com/in/panayiotismavrommatis/), Security Engineering Director at Google
- [Moheeb Abu Rajab](https://www.linkedin.com/in/moheeb/), Principal Engineer at Google
- [Noé Lutz](https://www.linkedin.com/in/noelutz/), Engineering Lead at Google AI
- [Nav Jagpal](https://www.linkedin.com/in/nav-jagpal-3972152/), Senior Staff Engineer at Google
- [Allison Miller](https://www.linkedin.com/in/allisonmiller/), SVP Engineering at Bank of America
- [Fabrice Jaubert](https://www.linkedin.com/in/fabrice-jaubert-40a651/), Senior Software Development Manager at Google
- [Stephan Somogyi](https://www.linkedin.com/in/stephan-somogyi-54618a1/), Product Lead, Android Platform Security
- [Emily Schechter](https://www.linkedin.com/in/emilyschechter/), Product Manager at Google Chrome
- [Brian Ryner](https://www.linkedin.com/in/brian-ryner-b0b226133/), Software Engineer at Google
- [Lucas Ballard](https://www.linkedin.com/in/lucas-ballard-b577889b/), Senior Staff Software Engineer at Google
- [Ian Fette](https://www.linkedin.com/in/ianfette/), Senior Director Of Engineering at Slack

Nav Jagpal is the only one who is still involved in the project, so I went ahead and contacted him on LinkedIn to see if he could give me some ideas of how the web service works and the scale at which they are operating. I was recommended to take a look at the implementation code in these two projects:

- https://github.com/google/safebrowsing
- https://github.com/google/webrisk
