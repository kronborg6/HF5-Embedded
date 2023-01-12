from django.shortcuts import render

# Create your views here.



def index (request):

    

    return render(request, "hf5/index.html",{

        


        
      
    })


def hej(request,name):
    

    return render(request, "hf5/hej.html",{


        "name" : name

    })