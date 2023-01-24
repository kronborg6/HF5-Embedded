from django.shortcuts import render

import requests, pickle
from django.shortcuts import redirect
from django.core.paginator import Paginator
from django.http import HttpResponse, HttpResponseRedirect
from django.urls import reverse
from django import forms
import json
import hashlib




#API-ting
url = 'http://10.130.54.111:8000/'
username = "Admin"
password = "Password"

print("Kronborg")




def index (request):

    
    token = request.session.get('token')
    headers = {"Authorization": "Bearer " + token }

    
    



    if token:
        response = requests.get(url + "data", headers=headers)
        data = response.json()    
    else:
        
        data = ""

    #Params
    device = request.GET.get("device")
    if device:
        device = int(device)    
    else:
        device = 0        

    session = requests.session()  # or an existing session

    with open('somefile', 'rb') as f:
        print(session.cookies.update(pickle.load(f)))




    return render(request, "hf5/index.html",{       

        "data" : data,
        "device" : device      
        
      
    })


def warnings(request):
    

    token = request.session.get('token')
    headers = {"Authorization": "Bearer " + token }


    if token:
        response = requests.get(url + "alarm", headers=headers)
        data = response.json()  
    else:
        data = ""        

    
    #Params
    device = request.GET.get("device")
    if device:
        device = int(device)    
    else:
        device = 0        


    return render(request, "hf5/warnings.html",{

        "data" : data,
        "device" : device
        
        
    })


class EditForm(forms.Form):    
    id = forms.CharField(label = "id")
    start_time = forms.CharField(label = "Start-time")    
    end_time = forms.CharField(label = "End-time")    
    



"""

    "start_time": 9,
"end_time": 18,
"min_temp": 5,
"max_temp": 100,
"min_hum": 10,
"max_hum": 20,
"max_noise_level": 400

"""
def startup(request):



    token = request.session.get('token')
    headers = {"Authorization": "Bearer " + token }

    response = requests.get(url + "startup", headers=headers)
    data = response.json()    

    

    if request.method == "POST":

        form = EditForm(request.POST)


        if form.is_valid():
            id = form.cleaned_data["id"]
            start_time = form.cleaned_data["start_time"] 
            end_time = form.cleaned_data["end_time"]   

            print(id + " " + start_time + " " + end_time)






        myobj = {
        "local": "Drivehus 24",
        "start_time": 70,
        "min_temp": 1000
        }      


        x = requests.put(url + "startup", data = myobj,auth=(username,password))
        print(myobj)
        print(x)

    return render(request, "hf5/startup.html",{

        "data" : data,
        "form" : EditForm
        

    })



def logout(request):


    request.session['token'] = ""
    print("lort")

    return HttpResponseRedirect(reverse("login"))  

def login(request):


    if request.method == "POST":
        # Attempt to sign user in
        username = request.POST["username"]
        password = request.POST["password"]   



        myobj = {'username': "kronborg", "password" : "Kronborg"}


        x = requests.post(url + "user/login", json = myobj)   
        json_response = x.json()

        token = json_response["token"]
        session = requests.session()     

        
       

        request.session['token'] = token 
        



        return HttpResponseRedirect(reverse("index"))    


        

    else:

        return render(request, "hf5/login.html")