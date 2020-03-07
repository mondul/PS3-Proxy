# PS3-Proxy
PS3 proxy to allow PSN login on older firmwares

## How to use it:
* To detect the local IP and use the default port, run the binary with no arguments, and it will show your current local IP where you should point your PS3.
* To set a local IP and port, run the binary as `ps3-proxy IP PORT`. For example, `ps3-proxy 192.168.1.1 12000`.
* On your PS3 select _Settings_ ➡ _Network Settings_ ➡ _Internet Connection Settings_, press ✖ to continue, select _Custom_ method and continue until you get the _Proxy Server_ screen. Choose _Use_, continue, enter the IP address you got from the previous step and continue until you reach the save settings page. press ✖ to save settings and press it again if you want to to test the connection. You should now be able to connect to PSN.

## Thanks to:
[@elazarl](https://github.com/elazarl) for his goproxy library and examples. Without it this script would not be possible.
