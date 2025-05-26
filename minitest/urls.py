from django.urls import path
from . import views 

urlpatterns = [
    path('signUp/', views.SignUpAPIView.as_view(), name='signUp'),
    path('signIn/', views.SignInAPIView.as_view(), name='signIn'),
    path('myPage/', views.MyPageAPIView.as_view(), name='myPage'),
    # path('names/<int:idx>/', views.NameDeleteAPIView.as_view(), name='name-detail'),
    # path('names/<str:names>/', views.NameAPIView.as_view(), name='name-delete-by-name'),    
    path('company/', views.CompanyAPIView.as_view(), name='company'),
    path('company/<int:idx>/', views.CompanyDeleteAPIView.as_view(), name='company-detail'),
]