import time,sys
import grovepi
import math
from datetime import datetime
import threading
import requests

#API-ting

url = 'http://10.130.54.111:8080/'
username = "Admin"
password = "Password"

#Soundsensorting
sound_sensor = 0
grovepi.pinMode(sound_sensor,"INPUT")




#Data sender hver 15 min.


def startup():
    
    threading.Timer(10,startup).start()
    response = requests.get(url + "startup/1",auth=(username, password))
    data = response.json()
    
    
    global max_temp
    global min_temp
    
    global max_hum
    global min_hum    
    
    global threshold_value
    
    
    #Tidsperiode
    global starttime
    global endtime
        
    
    dict = data[0]       
    
    #TEMP    
    _max_temp = dict["max_temp"]
    _min_temp = dict["min_temp"]   
    
    max_temp = _max_temp
    min_temp = _min_temp
    
    
    #HUM
    _max_hum = dict["max_hum"]
    _min_hum = dict["min_hum"]   
    
    max_hum = _max_hum
    min_hum = _min_hum
    
    
    #START&ENDTIME
    
    starttime = dict["start_time"]
    endtime = dict["end_time"]
    
    #SOUND       
    _max_noise_level = dict["max_noise_level"]    
    threshold_value = _max_noise_level       
    
    
startup()
        
    
def resetter():
    #1800
    threading.Timer(15,resetter).start()
    global wasSent
    wasSent = False      

resetter()   



global _temp
global _hum

_temp = 0.0
_hum = 0.0


def senddata():  
#900 er 15min
    threading.Timer(10.0, senddata).start()  
      
    myobj = {"local_id": 1,"hum" : _hum,"temp" : _temp}
    response = requests.post(url + "data",json=myobj,auth=(username, password)) 
    print(response.status_code)
  
    

senddata()





#Button ting
button = 3
grovepi.pinMode(button,"INPUT")


sensor = 4  # The Sensor goes on digital port 4.

blue = 0
white = 1   # The White colored sensor.







f = False

string = ""
 

 
if sys.platform == 'uwp':
    import winrt_smbus as smbus
    bus = smbus.SMBus(1)
else:
    import smbus
    import RPi.GPIO as GPIO
    rev = GPIO.RPI_REVISION
    if rev == 2 or rev == 3:
        bus = smbus.SMBus(1)
    else:
        bus = smbus.SMBus(0)
 
# this device has two I2C addresses
DISPLAY_RGB_ADDR = 0x62
DISPLAY_TEXT_ADDR = 0x3e
 
# set backlight to (R,G,B) (values from 0..255 for each)
def setRGB(r,g,b):
    bus.write_byte_data(DISPLAY_RGB_ADDR,0,0)
    bus.write_byte_data(DISPLAY_RGB_ADDR,1,0)
    bus.write_byte_data(DISPLAY_RGB_ADDR,0x08,0xaa)
    bus.write_byte_data(DISPLAY_RGB_ADDR,4,r)
    bus.write_byte_data(DISPLAY_RGB_ADDR,3,g)
    bus.write_byte_data(DISPLAY_RGB_ADDR,2,b)
 
# send command to display (no need for external use)    
def textCommand(cmd):
    bus.write_byte_data(DISPLAY_TEXT_ADDR,0x80,cmd)
 
# set display text \n for second line(or auto wrap)     
def setText(text):
    textCommand(0x01) # clear display
    time.sleep(.05)
    textCommand(0x08 | 0x04) # display on, no cursor
    textCommand(0x28) # 2 lines
    time.sleep(.05)
    count = 0
    row = 0
    for c in text:
        if c == '\n' or count == 16:
            count = 0
            row += 1
            if row == 2:
                break
            textCommand(0xc0)
            if c == '\n':
                continue
        count += 1
        bus.write_byte_data(DISPLAY_TEXT_ADDR,0x40,ord(c))
 
#Update the display without erasing the display
def setText_norefresh(text):
    textCommand(0x02) # return home
    time.sleep(.05)
    textCommand(0x08 | 0x04) # display on, no cursor
    textCommand(0x28) # 2 lines
    time.sleep(.05)
    count = 0
    row = 0
    while len(text) < 32: #clears the rest of the screen
        text += ' '
    for c in text:
        if c == '\n' or count == 16:
            count = 0
            row += 1
            if row == 2:
                break
            textCommand(0xc0)
            if c == '\n':
                continue
        count += 1
        bus.write_byte_data(DISPLAY_TEXT_ADDR,0x40,ord(c)) 



def post(alarm_type_id,type_id,value):    
    
    
    global wasSent
    
   
    if wasSent == False: 

        
        myobj = {"local_id": 1,
             "alarm_type_id" : alarm_type_id,
             "type_id" : type_id,
             "value" : value
             }        

        response = requests.post(url + "alarm",json=myobj,auth=(username, password))        


        print(response.status_code)
        wasSent = True
               


def tempstring(temp):
    
    result = ""
    
    
       
    if temp > max_temp:
        
        post(2,2,temp)    
        return "High temp"
    
    elif temp < min_temp:
        post(1,2, temp)
        return "Low temp"
    else:        
        return "Reg temp"

def humstring(hum):
    
    result = ""
    
    if hum > max_hum:
        post(2,3,hum)
        return "High hum"
    elif hum < min_hum:
        post(1,3,hum)
        return "Low hum"
    else:
        return "Reg hum"
    

def convert(c):
    
    x = (c * 1.8) + 32    
    result = float("{:.2f}".format(x))        
    
    return result

setRGB(255,255,255)

while True:
    
    
    
    
    try:
        
        
          
        now = datetime.now()
        currenthour = now.strftime("%H")        
        currenthour = now.strftime("%H")
        currenthour = int(currenthour)
        
        
        
        # Read the sound level
        sensor_value = grovepi.analogRead(sound_sensor)
        # If loud, illuminate LED, otherwise dim
        
        
        
        
        
        if sensor_value > threshold_value:
            
            if currenthour < endtime and currenthour >= starttime:
            
                                   
                #SEND ADVARSEL, INDE FOR START OG SLUTTID
                post(1,1, sensor_value)
                print("Sound too loud, sending warning")   
                
                
                
            else:
                
                
                #SEND ALARM, UDE FOR START OG SLUTTID
                post(2,1, sensor_value)
                print("Sound too loud, sending alarm")                
                
           
 
    except IOError:
        print ("Error")
        
        
  
    
    if grovepi.digitalRead(button) == 1:
        
        
        if f == True:
            f = False
        elif f == False:
            f = True
        
        print("Pressed")
        print(f)
        time.sleep(0.5)    
    


    try: 

        #Sætter temp og humidity fra hardwaren, kan den ikke registrere hardwaren springer programmet til except.
        [temp,humidity] = grovepi.dht(sensor,blue)
        
        global _temp
        global _hum
        
       
        
        if math.isnan(temp) == False and math.isnan(humidity) == False:
        
            
            #Tjekker om temperaturen og luftfugtigheden er over 0
            if temp > 0 and humidity > 0:
                
                _temp = temp
                _hum = humidity
                
                now = datetime.now()
                currenthour = now.strftime("%H")
                currenthour = int(currenthour)
                
                #Tjekker om nuværende tidspunkt er inden for tidsperioden
                if currenthour < endtime and currenthour >= starttime:
                    
                    setRGB(255,255,255)
                    
                    #Sætter fahrenheit variablen vha. convert funktionen.
                    fahrenheit = convert(temp)    

                    #Sætter hsttring til at være humidity + %                       
                    hstring = str(humidity) + "% "
                    #Bagefter bliver humstring tilføjet vha. af humstring funktionen. Humstring funktionen tilføjer advarsler til stringen som f.eks high temp osv.
                    hstring = hstring + humstring(humidity)
                    
                    #Hvis fahrenheit er True udskriver skærmen den passende string.
                    if f:                        
                        tstring = str(fahrenheit) + "F  " + tempstring(temp)                        
                        setText_norefresh(tstring + " " + hstring)      
                    else:   
                        tstring = str(temp) + "C  " + tempstring(temp) 
                        setText_norefresh(tstring + " " + hstring)
                else:
                    
                    #Er skærmen ude for tidsperioden slukkes skærmen
                    setRGB(0,0,0)
                    setText("")
    except IOError:
        print ("Error")