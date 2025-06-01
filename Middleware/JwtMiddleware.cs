using System.IdentityModel.Tokens.Jwt;
using System.Text;
using BerryNameApi.Utils;
using Microsoft.Extensions.Primitives;
using Microsoft.IdentityModel.Tokens;

namespace blueberry_homework_dotnet.Middleware
{
    public class JwtMiddleware
    {
        private readonly RequestDelegate _next;
        private readonly string _secretKey;

        public JwtMiddleware(RequestDelegate next)
        {
            _next = next;
            _secretKey = Environment.GetEnvironmentVariable("JWT_SECRET_KEY")
                         ?? throw new Exception(Constants.KeyNotSet);
        }

        public async Task InvokeAsync(HttpContext context)
        {
            var path = context.Request.Path.ToString().ToLower();

            // 로그인, 회원가입은 제외
            if (path.Contains("/auth/sign-up") || path.Contains("/auth/log-in"))
            {
                await _next(context);
                return;
            }

            if (!context.Request.Headers.TryGetValue("Authorization", out StringValues authHeader))
            {
                context.Response.StatusCode = 401;
                await context.Response.WriteAsync(Constants.TokenMissingAuthorizationHeader);
                return;
            }

            var token = authHeader.ToString().Replace("Bearer ", "");

            var handler = new JwtSecurityTokenHandler();
            var validationParams = new TokenValidationParameters
            {
                ValidateIssuer = false,
                ValidateAudience = false,
                ValidateLifetime = true,
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_secretKey)),
                ClockSkew = TimeSpan.Zero,
                ValidAlgorithms = new[] { SecurityAlgorithms.HmacSha512 }
            };

            try
            {
                handler.ValidateToken(token, validationParams, out _);
                // 토큰 유효시 다음 미들웨어로
                await _next(context);
            }
            catch (Exception)
            {
                context.Response.StatusCode = 401;
                await context.Response.WriteAsync(Constants.TokenInvalidOrExpired);
            }
        }
    }

    public static class JwtMiddlewareExtensions
    {
        public static IApplicationBuilder UseJwtMiddleware(this IApplicationBuilder builder)
        {
            return builder.UseMiddleware<JwtMiddleware>();
        }
    }
}
