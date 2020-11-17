## Vercel multiple headers

### Local example

- Start the server
```bash
go run main.go
```

- On a second terminal

With one accept header:
```bash
$ curl -H "Accept: application/json" 127.0.0.1:8000

User-Agent [curl/7.64.1]
Accept [application/json]

Tuesday, 17-Nov-20 22:09:28 CET
```

With Accept headers with two values:
```bash
$ curl -H "Accept: application/json,application/yaml" 127.0.0.1:8000

Accept [application/json,application/yaml]
User-Agent [curl/7.64.1]

Tuesday, 17-Nov-20 22:10:28 CET
```

With two Accept headers:
```bash
$ curl -H "Accept: application/json" -H "Accept: application/yaml" 127.0.0.1:8000

User-Agent [curl/7.64.1]
Accept [application/json application/yaml]

Tuesday, 17-Nov-20 22:10:54 CET
```


### On vercel

With one accept header:
```bash
$ curl -H "Accept: application/json" https://vercel-multi-headers.vercel.app/

X-Vercel-Id [cdg1::qtfj2-1605647119502-6019402895c6]
Accept [application/json]
X-Vercel-Forwarded-For [79.90.184.110]
Host [vercel-multi-headers.vercel.app]
X-Real-Ip [79.90.184.110]
X-Forwarded-Proto [https]
X-Vercel-Deployment-Url [vercel-multi-headers-fzggxqr6r.vercel.app]
X-Forwarded-For [79.90.184.110]
User-Agent [curl/7.64.1]
X-Forwarded-Host [vercel-multi-headers.vercel.app]

Tuesday, 17-Nov-20 21:05:19 UTC
```

With Accept headers with two values:
```bash
$ curl -H "Accept: application/json,application/yaml" https://vercel-multi-headers.vercel.app/

X-Forwarded-For [79.90.184.110]
X-Forwarded-Host [vercel-multi-headers.vercel.app]
X-Forwarded-Proto [https]
X-Vercel-Deployment-Url [vercel-multi-headers-7is10xscg.vercel.app]
User-Agent [curl/7.64.1]
X-Vercel-Id [cdg1::2xlnf-1605647308504-3506963ee299]
Host [vercel-multi-headers.vercel.app]
X-Real-Ip [79.90.184.110]
Accept [application/json,application/yaml]
X-Vercel-Forwarded-For [79.90.184.110]

Tuesday, 17-Nov-20 21:08:28 UTC
```

With two Accept headers:
```bash
$ curl -H "Accept: application/json" -H "Accept: application/yaml" https://vercel-multi-headers.vercel.app/

{"error": {"code": "502", "message": "An error occurred with this application."}}
```

## RFC
[HTTP 1.1 Section 4.2](https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2)

> Multiple message-header fields with the same field-name MAY be present in a message if and only if the entire field-value for that header field is defined as a comma-separated list [i.e., #(values)]. It MUST be possible to combine the multiple header fields into one "field-name: field-value" pair, without changing the semantics of the message, by appending each subsequent field-value to the first, each separated by a comma. The order in which header fields with the same field-name are received is therefore significant to the interpretation of the combined field value, and thus a proxy MUST NOT change the order of these field values when a message is forwarded.
