#include <SPI.h>
#include <Ethernet.h>

byte mac2[] = { 0xA8, 0x61, 0x0A, 0xAE, 0x98, 0x02 }; // mac2 address of 

EthernetClient client2;

int    HTTP_PORT2   = 8080;
// String HTTP_METHOD2 = "POST"; 
char   HOST_NAME2[] = "10.130.54.111";
// String PATH_NAME2   = "/data";
String c;
char test;
char g;

String api(String method, String path, String body) {
  Serial.begin(9600);


  Serial.println(method);
  Serial.println(path);
  byte mac2[] = { 0xA8, 0x61, 0x0A, 0xAE, 0x98, 0x02 }; // mac2 address of 
  int    HTTP_PORT2   = 8080;
  char   HOST_NAME2[] = "10.130.54.111";

  if (Ethernet.begin(mac2) == 0) {
    Serial.println("Failed to obtaining an IP address using DHCP");
    // while(true);
  }
  if(client2.connect(HOST_NAME2, HTTP_PORT2)) {
    // if connected:
      Serial.println("Connected to server");
    // make a HTTP request:
    // send HTTP header
      client2.println(method + " " + path + " HTTP/1.1");
      client2.println("Host: " + String(HOST_NAME2));
      client2.println("Authorization: Basic QWRtaW46UGFzc3dvcmQ=");
      client2.println("Content-Type: application/json");
      client2.println("Content-Length: " + String(body.length()));
      client2.println("Connection: close");
      client2.println(); // end HTTP header
      client2.print(body);

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
    client2.println(method + " " + path + " HTTP/1.1");
    client2.println("Host: " + String(HOST_NAME2));
    client2.println("Authorization: Basic QWRtaW46UGFzc3dvcmQ=");
    client2.println("Content-Type: application/json");
    client2.println("Content-Length: " + String(body.length()));
    client2.println("Connection: close");
    client2.println(); // end HTTP header
    client2.print(body);
    while(client2.connected()) {
      if(client2.available()){
        // read an incoming byte from the server and print it to serial monitor:
        // Serial.println("Min mor er lort #SkydHende");        
        g = client2.read();
        test += g;
        // Serial.println(g);
        // test += c;  
        // return c;
        
      }
    }
    Serial.println("her stå se hvad har han lavet af lort??");
    Serial.println(test);
    // return c;

    // the server's disconnected, stop the client2:
    client2.stop();
    Serial.println();
    Serial.println("disconnected");
    return c;
  } else {// if not connected:
    Serial.println("connection failed");
  }
  }

  
}



