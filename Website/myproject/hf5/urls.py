from django.urls import path
from . import views


urlpatterns = [
    path("", views.index, name="index"),
    path("hej/<str:name>", views.hej, name="hej")
]