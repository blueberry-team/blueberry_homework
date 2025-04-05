from django.urls import path
from . import views 

urlpatterns = [
    path('names/', views.NameAPIView.as_view(), name='names'),
]