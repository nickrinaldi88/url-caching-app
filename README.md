# url-caching-app
## TODO
## find out why 'exit' doesn't break loop
## finish cache.go to make k,v pair
## test cache lib import
## test passing in arguments to cache go
## use http in fetcher

## keys to fetch + store in map

<!-- Fastly-Restarts
X-Served-By
Vary
Content-Length
Server
Etag
X-Content-Type-Options
Content-Language
Strict-Transport-Security
Age
X-Datadome
Content-Type
Set-Cookie
Date
Via
X-Cache
X-Is-Sg-Next-Page
Accept-Ranges
X-Timer
X-Cache-Hits
Cache-Control
Geo-Country-Code
X-Xss-Protection
Status Code: 200 OK -->

## fields and their use-cases

<!-- When caching URL metadata obtained from a GET request, you might want to consider storing key headers that can be useful for various purposes. Here are some common headers that could be practical to store in the cache:

Cache-Control:

This header provides directives for caching mechanisms in both requests and responses. It can be useful to determine the caching behavior specified by the server.
Expires:

Indicates the date and time after which the response is considered stale. This can be useful for understanding the freshness of the cached data.
Last-Modified:

Specifies the last modification date and time of the resource. It can be useful for conditional requests, allowing you to check if the resource has been modified since a certain date.
ETag:

Provides a unique identifier for the version of the resource. It can be used for efficient caching and conditional requests to check if the resource has changed.
Content-Type:

Specifies the media type of the resource. This can be helpful for handling different types of content appropriately.
Content-Length:

Indicates the size of the response body. This information can be useful for managing storage and understanding the payload size.
Pragma:

Used for backward compatibility with HTTP/1.0. It may include directives like "no-cache," which can influence caching behavior.
Allow:

Specifies the HTTP methods that are allowed for the resource. This can be useful for understanding the allowed actions on the resource.
Access-Control-Allow-Origin:

If you are dealing with cross-origin requests and need to implement CORS (Cross-Origin Resource Sharing), this header specifies which origins are permitted to access the resource.
Other custom headers:

Depending on the API and your specific use case, there might be custom headers containing relevant information specific to the resource or the application. These could include user-specific information, versioning details, or any other metadata relevant to your use case.
Remember to carefully review the API documentation to identify which headers are most relevant to your caching strategy and use case. Additionally, consider the privacy and security implications of caching certain information, especially if it involves sensitive data. -->