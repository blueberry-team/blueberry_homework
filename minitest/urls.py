from django.urls import path
from . import views 

urlpatterns = [
    path('signUp/', views.SignUpAPIView.as_view(), name='signUp'),
    path('signIn/', views.SignInAPIView.as_view(), name='signIn'),
    path('myPage/', views.MyPageAPIView.as_view(), name='myPage'),
    path('company/', views.CompanyAPIView.as_view(), name='company'),
    path('company/<int:idx>/', views.CompanyDeleteAPIView.as_view(), name='company-detail'),
    path('api/auth/refresh-token/', views.RefreshTokenView.as_view(), name='refresh_token'),
]