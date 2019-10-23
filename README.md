# Lightweight-Web-Server /LWS
A Lightweight Web Server implemented in GO to serve out html, wasm, js or any other content. Additional logic can be added as needed.

LWS can be used with <a href="https://github.com/maxng07/ted-gui"> TED-GUI </a> on either locally or in the Cloud, and recognises wasm and will set Content Type as "Application/WASM". As more functionality is needed, it can be added to LWS. 

The design goal for LWS is to keep it simple and lightweight Web Server.

# Usage
The binary and config.json must be located in the same directory. LWS binary will retrieve configuration informations from config.json and start the Web Server. External config file is used instead of flags/CLI as the design goal is to install LWS either locally where users may not be technical, or within a Centralised Server or Private Cloud to serve out simple web contents.

The essential config.json informations needed are
1. IP Address - set this using the loopback IP address, 127.0.0.1 if used locally.
2. TCP Port number
3. Default Directory for the Web Server to serve out contents (html, js, wasm, images)

This is a graphical view of FireFox Web Browser accessing TED-GUI off LWS. 
<img src="https://github.com/maxng07/Lightweight-Web-Server/blob/master/graphics/webserver2.png">
