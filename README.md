= FChk - File Check Daemon
The server is in fchk/fchkserver, client in fchk/fchkclient.
Uses go RPC and requires one TCP port to be opened through the firewall.

= Building server and client binaries
./build.sh w (w for windows, m for mac , l or any other char for linux)

Build scripts generate static binaries by disabling CGO.
