#include <SPI.h>
#include <Ethernet.h>


byte mac[] = { 0xA8, 0x61, 0x0A, 0xAE, 0x98, 0x02 }; // MAC address of 

EthernetClient client;

int    HTTP_PORT   = 8080;
String HTTP_METHOD = "GET";
char   HOST_NAME[] = "10.130.54.111";
String PATH_NAME   = "/startup/2";


float _maxt;
float _mint;

float _maxh;
float _minh;

int _start_time;
int _end_time;

float _maxsound;

/*
void setvariables()
{
  maxt = _maxt;
  Serial.println("max temp");
  Serial.println(maxt);
  mint = _mint;
  maxh = _maxh;
  minh = _minh;
  starttime = _start_time;
  endtime = _end_time;

  


  

}
*/

void startget() {
  Serial.begin(9600);

  // initialize the Ethernet shield using DHCP:
  if (Ethernet.begin(mac) == 0) {
    Serial.println("Failed to obtaining an IP address using DHCP");
    while(true);
  }

  // connect to web server on port 80:
  if(client.connect(HOST_NAME, HTTP_PORT)) {
    // if connected:
    Serial.println("Connected to server");
    // make a HTTP request:
    // send HTTP header
    client.println(HTTP_METHOD + " " + PATH_NAME + " HTTP/1.1");
    client.println("Host: " + String(HOST_NAME));
    client.println("Authorization: Basic QWRtaW46UGFzc3dvcmQ=");
    client.println("Connection: close");
    client.println(); // end HTTP header

    while(client.connected()) {
      if(client.available()){    


        char c = client.read();        
        StaticJsonDocument<256> doc;
        
        deserializeJson(doc, client);      

        // Serial.println(doc)

        int id = doc[0]["id"];

        _maxt = doc[0]["max_temp"];        

        _mint = doc[0]["min_temp"];

        
        _maxh = doc[0]["max_hum"];
        _minh = doc[0]["min_hum"];

        _start_time = doc[0]["start_time"];
        _end_time = doc[0]["end_time"];

        _maxsound = doc[0]["max_noise_level"];

      


      }

    
      
    }

   

    if (_maxt != 0.00)
    {
      maxt = _maxt; 
    }

    if (_mint != 0.00)
    { 
      mint = _mint;
    }
  

    if (_maxh != 0.00)
    {
      maxh = _maxh; 
    }

    if (_minh != 0.00)
    { 
      minh = _minh;
    } 

    
    if(_start_time != 0)
    {
      starttime = _start_time;
    }

    if (_end_time !=0 )
    {
      endtime = _end_time;
    }

    if (_maxsound != 0.00)
    {
      maxsound = _maxsound;
    }
    

    client.stop();
    Serial.println();
    Serial.println("disconnected");
  } else {// if not connected:
    Serial.println("connection failed");
  }
}

