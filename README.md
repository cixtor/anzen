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
