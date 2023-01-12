

#include <Wire.h>
#include "rgb_lcd.h"

rgb_lcd lcd;

const int colorR = 255;
const int colorG = 255;
const int colorB = 255;


float maxt = 25;
float mint = 15;

float maxh = 80;
float minh = 60;

bool isF = true;

const int buttonPin = 6;

#include "DHT.h" 
#define DHTPIN A0 // what pin we're connected to
#define DHTTYPE DHT11 // DHT 11
 
DHT dht(DHTPIN, DHTTYPE);
 
void setup()
{
    Serial.begin(9600);


    pinMode(buttonPin,INPUT_PULLUP);


    lcd.begin(16, 2);
    lcd.setRGB(colorR, colorG, colorB);


    dht.begin();
}
 
void loop()
{  



    int digitalVal = digitalRead(buttonPin); // Take a reading

    if(HIGH == digitalVal)
    {
       Serial.println("ON");

       if (isF == true)
       {

         isF = false;

       }
       else if (isF == false)
       {

         isF = true;

       }
              
       delay(500);


    }
 


    float h = dht.readHumidity();
    float t = dht.readTemperature();
    float f = dht.readTemperature(true);
 

    if (isnan(t) || isnan(h))
    {
        Serial.println("Failed to read from DHT");
    }
    else
    {
        Serial.print("Humidity: ");
        Serial.print(h);
        Serial.print(" %\t");
        Serial.print("Temperature: ");
        Serial.print(t);
        Serial.println(" *C");
    }





  if (t > maxt)
  {

    lcd.setCursor(0,0);


    if (isF)
    {
      lcd.print(f);
      lcd.print("F");
    }
    else
    {
      lcd.print(t);
      lcd.print("C");      
    }



    lcd.print(" High temp");
    
  }

  else if (t < mint)
  {

    lcd.setCursor(0,0);

    if (isF)
    {
      lcd.print(f);
      lcd.print("F");
    }
    else
    {
      lcd.print(t);
      lcd.print("C");      
    }


    lcd.print(" Low temp");

  }

  else
  {

    lcd.setCursor(0, 0);

    if (isF)
    {
      lcd.print(f);
      lcd.print("F");
    }
    else
    {
      lcd.print(t);
      lcd.print("C");      
    }


    lcd.print(" Reg. temp");

  }


  if (h > maxh)
  {

    lcd.setCursor(0,1);
    lcd.print(h);
    lcd.print("%");
    lcd.print(" High hum");

  }
  else if (h < minh)
  {
    
    lcd.setCursor(0,1);
    lcd.print(h);
    lcd.print("%");
    lcd.print(" Low hum");
    
  }

  else
  {

    lcd.setCursor(0,1);
    lcd.print(h);
    lcd.print("%");
    lcd.print(" Reg hum");

  }

  


}