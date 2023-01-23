from django.shortcuts import render

import requests
from django.shortcuts import redirect




#API-ting
url = 'http://10.130.54.111:8080/'
username = "Admin"
password = "Password"



def index (request):



    response = requests.get(url + "data",auth=(username, password))
    data = response.json()  


    device = request.GET.get("page") 

    
   

    return render(request, "hf5/index.html",{       

        "data" : data
        
      
    })


def warnings(request):
    

    return render(request, "hf5/warnings.html",{

        
        
    })