import time,sys
import grovepi
import math



#Button ting
button = 3
grovepi.pinMode(button,"INPUT")


sensor = 4  # The Sensor goes on digital port 4.

blue = 0
white = 1   # The White colored sensor.

t = 0
h = 0

maxtemp = 25
mintemp = 18

maxhum = 80
minhum = 60


f = True

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
 
# example code


def tempstring(temp):
    
    result = ""
    
    if temp > 25:
        return "High temp"
    elif temp < 18:
        return "Low temp"
    else:
        return "Reg temp"

def humstring(hum):
    
    result = ""
    
    if hum > 80:
        return "High hum"
    elif hum < 60:
        return "Low hum"
    else:
        return "Reg hum"
    

def convert(c):
    
    x = (c * 1.8) + 32
    
    result = float("{:.2f}".format(x))    
    
    
    return result

while True:
    
    
    if grovepi.digitalRead(button) == 1:
        
        
        if f == True:
            f = False
        elif f == False:
            f = True
        
        print("Pressed")
        print(f)
        time.sleep(0.5)    
    
    try: 
        [temp,humidity] = grovepi.dht(sensor,blue)  
        if math.isnan(temp) == False and math.isnan(humidity) == False:
            #print("temp = %.02f C humidity =%.02f%%"%(temp, humidity))
            if temp > 0 and humidity > 0:
                
                
                fahrenheit = convert(temp)
                
                                
                
                hstring = str(humidity) + "% "
                hstring = hstring + humstring(humidity)
                
                if f:
                    
                    tstring = str(fahrenheit) + "F  " + tempstring(temp)    
                    setText(tstring + " " + hstring)                    
                   
                    
                else:   
                    tstring = str(temp) + "C  " + tempstring(temp)    
                    setText(tstring + " " + hstring)      
                
 
    except IOError:
        print ("Error")