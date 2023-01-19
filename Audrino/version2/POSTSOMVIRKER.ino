#include <SPI.h>
#include <Ethernet.h>

byte mac2[] = { 0xA8, 0x61, 0x0A, 0xAE, 0x98, 0x02 }; // mac2 address of 

EthernetClient client2;

int    HTTP_PORT2   = 8080;
String HTTP_METHOD2 = "POST"; 
char   HOST_NAME2[] = "10.130.54.111";
String PATH_NAME2   = "/data";


void startpost(float temp, float hum) {
  
  
  Serial.begin(9600);

  // Serial.println(data2)

  String test = String(temp);
  String test2 = String(hum);


  String body = "\r\n{\"temp\":" + String(temp) + ",\r\n\"hum\":" + String(hum) + ",\r\n\"local_id\":2\r\n  \r\n}\r\n";

  Serial.println(body);


// invalid character 'h' after object key:value pair



  // initialize the Ethernet shield using DHCP:
  if (Ethernet.begin(mac2) == 0) {
    Serial.println("Failed to obtaining an IP address using DHCP");
    while(true);
  }

  // connect to web server on port 8080:
  if(client2.connect(HOST_NAME2, HTTP_PORT2)) {
    // if connected:
    Serial.println("Connected to server");
    // make a HTTP request:
    // send HTTP header
    client2.println(HTTP_METHOD2 + " " + PATH_NAME2 + " HTTP/1.1");
    client2.println("Host: " + String(HOST_NAME2));
    client2.println("Authorization: Basic TG9ydDpQYXNzd29yZA==");
    client2.println("Content-Type: application/json");
    client2.println("Content-Length: " + String(body.length()));
    client2.println("Connection: close");
    client2.println(); // end HTTP header
    client2.print(body);
    while(client2.connected()) {
      if(client2.available()){
        // read an incoming byte from the server and print it to serial monitor:
        char c = client2.read();
        Serial.print(c);
      }
    }

    // the server's disconnected, stop the client2:
    client2.stop();
    Serial.println();
    Serial.println("disconnected");
  } else {// if not connected:
    Serial.println("connection failed");
  }
}


