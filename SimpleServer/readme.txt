Here are 3 version of the simple server programm.

1. Build programm 
$ go build server#.go


"Server1.go" just return to us path of URL, that was usage
for request to server

./server1 http://localhost:8000
URL.PATH - "/"


"Server2.go" just doing the same as "server1.go" but now
can return the count of reques
Server has 2 handler so , the requested URL define, which is wil called:
request /count request counter
another - handler

For background mode just use:
$ ./server2 http://localhost:8000/count &

So , if you whant to kill a pid, just usage this comma
$ sudo netstat -tulpn | grep :8000
$ sudo kill -9 (pid)

"Server3.go" better than "Server2", because can report 
about headers and data

./server3 http://localhost:8000
web page on localhost:8000 looks like

GET / HTTP/1.1
Header["Sec-Ch-Ua-Mobile"] = ["?0"]
Header["Sec-Fetch-Mode"] = ["navigate"]
Header["Sec-Fetch-User"] = ["?1"]
Header["Sec-Ch-Ua-Platform"] = ["\"Windows\""]
Header["Upgrade-Insecure-Requests"] = ["1"]
Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"]
Header["Accept-Encoding"] = ["gzip, deflate, br"]
Header["Accept-Language"] = ["en-GB,en-US;q=0.9,en;q=0.8,ru;q=0.7"]
Header["Sec-Ch-Ua"] = ["\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\""]
Header["Sec-Fetch-Dest"] = ["document"]
Header["Connection"] = ["keep-alive"]
Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"]
Header["Sec-Fetch-Site"] = ["none"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:34334",