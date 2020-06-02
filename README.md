# xmlrpc-scan

Scan urls or a single URL against XMLRPC wordpress issues.

usage:

##### Install

- Download from releases: https://github.com/nullfil3/xmlrpc-scan/releases/tag/v1

- Or Compiling by yourself

```bash
git clone https://github.com/nullfil3/xmlrpc-scan.git
cd xmlrpc-scan
go build -o xmlrpcscan
./xmlrpcscan
```

##### Usage

* List of URLS
```bash
cat urls.txt | xmlrpcscan -server http://burpcollaborator.net
```

* Single URL
```bash
xmlrpcscan -target https://target.com -server http://burpcollaborator.net
```

![](tool.gif)

#Todo
--> Tracking where ssrf request come from, when vulnerable.
