#include <ArduinoJson.h>
#include <ArduinoJson.hpp>


//Tidsting
#include <RTClib.h>
#include <Wire.h>
RTC_DS3231 rtc;
char t[32];

//Tidstagning
#include <MsTimer2.h>

#include <Wire.h>
#include "rgb_lcd.h"

rgb_lcd lcd;

const int colorR = 255;
const int colorG = 255;
const int colorB = 255;

int starttime = 8;
int endtime = 18;


float _h;
float _t;
float _f;

//Max temperature and humidity
float maxt = 25;
float mint = 15;

float maxh = 80;
float minh = 60;

//Max soundlevel
float maxsound = 300;

bool isF = false;
const int buttonPin = 6;

int h = 0;

#include "DHT.h"
#define DHTPIN A0      // what pin we're connected to
#define DHTTYPE DHT11  // DHT 11

DHT dht(DHTPIN, DHTTYPE);


//Sound sensor
const int pinAdc = A1;





void flash() {




  if (h > starttime - 1 && h < endtime + 1) {


    startget();
  }
}

void setup() {



  Serial.begin(9600);
  pinMode(buttonPin, INPUT_PULLUP);
  lcd.begin(16, 2);
  lcd.setRGB(colorR, colorG, colorB);

  //Tidsting
  Wire.begin();
  rtc.begin();
  rtc.adjust(DateTime(F(__DATE__), F(__TIME__)));
  dht.begin();


  startget();


  //900000ser
  MsTimer2::set(30000, flash);  // 500ms period
  MsTimer2::start();
}

void turnOffLCD() {
  lcd.setCursor(0, 0);
  lcd.print("                                  ");
  lcd.setCursor(0, 1);
  lcd.print("                                  ");
  lcd.setRGB(0, 0, 0);
}


void sendWarning() {

  //Serial.println("Sound is too loud!");
}


void loop() {

  //Sound sensor ting
  long sum = 0;
  for (int i = 0; i < 32; i++) {
    sum += analogRead(pinAdc);
  }

  sum >>= 5;


  if (sum > maxsound) {
    delay(10);
    sendWarning();
  }


  //Datetime ting

  DateTime now = rtc.now();
  sprintf(t, "%02d:%02d:%02d %02d/%02d/%02d", now.hour(), now.minute(), now.second(), now.day(), now.month(), now.year());
  h = now.hour();

  int digitalVal = digitalRead(buttonPin);  // Take a reading




  float h = dht.readHumidity();
  float t = dht.readTemperature();
  float f = dht.readTemperature(true);




  _t = t;
  _h = h;


  if (HIGH == digitalVal) {
    Serial.println("ON");
    //startpost(_t,_h);



    if (isF == true) {

      isF = false;

    } else if (isF == false) {

      isF = true;
    }

    delay(500);
  }




  if (isnan(t) || isnan(h)) {
    Serial.println("Failed to read from DHT");
  }


  if (now.hour() > starttime - 1 && now.hour() < endtime + 1) {

    if (t > maxt) {

      lcd.setCursor(0, 0);


      if (isF) {
        lcd.print(f);
        lcd.print("F");
      } else {
        lcd.print(t);
        lcd.print("C");
      }



      lcd.print(" High temp");

    }

    else if (t < mint) {

      lcd.setCursor(0, 0);

      if (isF) {
        lcd.print(f);
        lcd.print("F");
      } else {
        lcd.print(t);
        lcd.print("C");
      }


      lcd.print(" Low temp");

    }

    else {

      lcd.setCursor(0, 0);

      if (isF) {
        lcd.print(f);
        lcd.print("F");
      } else {
        lcd.print(t);
        lcd.print("C");
      }

      lcd.print(" Reg. temp");
    }

    if (h > maxh) {

      lcd.setCursor(0, 1);
      lcd.print(h);
      lcd.print("%");
      lcd.print(" High hum");

    } else if (h < minh) {

      lcd.setCursor(0, 1);
      lcd.print(h);
      lcd.print("%");
      lcd.print(" Low hum");

    }

    else {

      lcd.setCursor(0, 1);
      lcd.print(h);
      lcd.print("%");
      lcd.print(" Reg hum");
    }

  } else {

    turnOffLCD();
  }
}
