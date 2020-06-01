# xmlrpc-scan

Scan urls or a single URL against XMLRPC wordpress issues.

usage:

* List of URLS
```bash
cat urls.txt | xmlrpcscan -server http://burpcollaborator.net
```

* Single URL
```bash
xmlrpscan -target https://target.com -server http://burpcollaborator.net
```

![](tool.gif)

