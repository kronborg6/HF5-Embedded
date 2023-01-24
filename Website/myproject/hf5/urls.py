from django.urls import path
from . import views


urlpatterns = [
    path("", views.index, name="index"),
    path("warnings", views.warnings, name="warnings"),
    path("login", views.login, name="login"),
    path("logout", views.logout, name="logout"),
    path("startup", views.startup, name="startup"),
]