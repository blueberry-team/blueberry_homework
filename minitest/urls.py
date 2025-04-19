from django.urls import path
from . import views 

urlpatterns = [
    path('names/', views.NameAPIView.as_view(), name='names'),
    path('names/<int:idx>/', views.NameDeleteAPIView.as_view(), name='name-detail'),
    path('names/<str:names>/', views.NameAPIView.as_view(), name='name-delete-by-name'),    
]