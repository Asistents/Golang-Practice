Just a Simple Programm of Url request.
In case of bad HTTP request, programm let us what happend.

Here is a simple usage example:

1. Build programm
$ go build url.go

2. Call the executable file with argc
$ ./url http://ya.ru

3. Now you are not needed to use Prefix http:// in your argc.
New Call of executable file looks like
$ ./url ya.ru

4. Now you can see a status code of HTTP request