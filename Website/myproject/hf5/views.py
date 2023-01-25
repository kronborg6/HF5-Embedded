from django.shortcuts import render

import requests
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


def index (request):



   


    
    if request.session.get('token'):
        token = request.session.get('token')
        headers = {"Authorization": "Bearer " + token }    
    

    else:
        return HttpResponseRedirect(reverse("login")) 


        
    
   


    if token:
        response = requests.get(url + "data", headers=headers)
        data = response.json()    
    else:
        
        data = ""
        return HttpResponseRedirect(reverse("login")) 

    #Params
    device = request.GET.get("device")
    if device:
        device = int(device)    
    else:
        device = 0            

    _data = list()    

    for index, item in enumerate(data):

        if item["local_id"] == device:

            _data.append(item)        

        elif device == 0:

            _data.append(item)
    


    data_paginator = Paginator(_data, 9)       
    page_number = request.GET.get("page") 
    page = data_paginator.get_page(page_number)     


    return render(request, "hf5/index.html",{       

        "data" : data,
        "page" : page,
        "device" : device      
        
      
    })


def warnings(request):

    
    if request.session.get('token'):
        token = request.session.get('token')
        headers = {"Authorization": "Bearer " + token }    
    

    else:
        return HttpResponseRedirect(reverse("login")) 

    

    token = request.session.get('token')
    headers = {"Authorization": "Bearer " + token }


    if token:
        response = requests.get(url + "alarm", headers=headers)
        data = response.json()  
    else:
        data = ""   
        return HttpResponseRedirect(reverse("login"))      

    
    #Params
    device = request.GET.get("device")
    if device:
        device = int(device)    
    else:
        device = 0       


    _data = list()

    for index, item in enumerate(data):

        if item["local_id"] == device:

            _data.append(item)        

        elif device == 0:

            _data.append(item)




    data_paginator = Paginator(_data, 9)       
    page_number = request.GET.get("page") 
    page = data_paginator.get_page(page_number)   






    return render(request, "hf5/warnings.html",{

        "data" : data,
        "device" : device,
        "page" : page
        
        
    })


class EditForm(forms.Form):    
    id = forms.CharField(label = "id", required=True)
    start_time = forms.CharField(label = "Start-time",required=False)    
    end_time = forms.CharField(label = "End-time",required=False) 
    max_temp = forms.CharField(label = "Max-temp",required=False)
    min_temp = forms.CharField(label = "Min-temp",required=False)
    max_hum = forms.CharField(label = "Max-hum",required=False)
    min_hum = forms.CharField(label = "Min-hum",required=False)
    max_noise_level = forms.CharField(label = "Max-noise-level",required=False)

    


def startup(request):

    
    if request.session.get('token'):
        token = request.session.get('token')
        headers = {"Authorization": "Bearer " + token }    
    

    else:
        return HttpResponseRedirect(reverse("login")) 




    token = request.session.get('token')
   
    headers = {"Authorization": "Bearer " + token }

    data = ""    

    if token:
       

        response = requests.get(url + "startup", headers=headers)
        data = response.json()   


        
    

        if request.method == "POST":

            form = EditForm(request.POST)
            


            if form.is_valid():


                id = form.cleaned_data["id"]
                start_time = form.cleaned_data["start_time"] 
                end_time = form.cleaned_data["end_time"]   
                max_temp = form.cleaned_data["max_temp"]
                min_temp = form.cleaned_data["min_temp"]
                max_hum = form.cleaned_data["max_hum"]
                min_hum = form.cleaned_data["min_hum"]
                max_noise_level = form.cleaned_data["max_noise_level"]
                

            
            if start_time == "":
                start_time = data[int(id) -1]["start_time"]
            else:
                start_time = int(start_time)  

            if end_time == "":
                end_time = data[int(id) -1]["end_time"]
            else:
                end_time = int(end_time)       

            if max_temp == "":
                max_temp = data[int(id) -1]["max_temp"]
            else:
                max_temp = int(max_temp)         

            if min_temp == "":
                min_temp = data[int(id) -1]["min_temp"]
            else:
                min_temp = int(min_temp)    

            if max_hum == "":
                max_hum = data[int(id) -1]["max_hum"]
            else:
                max_hum = int(max_hum)  

            if min_hum == "":
                min_hum = data[int(id) -1]["min_hum"]
            else:
                min_hum = int(min_hum)  

            if max_noise_level == "":
                max_noise_level = data[int(id) -1]["max_noise_level"]
            else:
                max_noise_level = int(max_noise_level)                
               


            myobj = {            
            "start_time": start_time,
            "end_time": end_time,
            "max_temp" : max_temp,
            "min_temp" : min_temp,
            "max_hum" : max_hum,
            "min_hum" : min_hum,
            "max_noise_level" : max_noise_level            
            }      
           


            x = requests.put(url + "startup/" + str(id), headers=headers, json=myobj)

            
            return HttpResponseRedirect(reverse("startup")) 

    else:
        return HttpResponseRedirect(reverse("login")) 
           

        

    return render(request, "hf5/startup.html",{

        "data" : data,
        "form" : EditForm
        

    })



def logout(request):


    request.session['token'] = ""
   

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