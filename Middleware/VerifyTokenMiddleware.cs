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

        public JwtMiddleware(RequestDelegate next, IConfiguration config)
        {
            _next = next;
            _secretKey = config["JWT_SECRET_KEY"] ?? throw new Exception(Constants.TokenNotSet);
        }

        public async Task Invoke(HttpContext context)
        {
            if (context.Request.Headers.TryGetValue("Authorization", out StringValues authHeader))
            {
                var token = authHeader.ToString().Replace("Bearer ", "");

                try
                {
                    var handler = new JwtSecurityTokenHandler();
                    var validationParams = new TokenValidationParameters
                    {
                        ValidateIssuer = false,
                        ValidateAudience = false,
                        ValidateIssuerSigningKey = true,
                        ValidateLifetime = true,
                        IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_secretKey)),
                        ClockSkew = TimeSpan.Zero,
                        ValidAlgorithms = new[] { SecurityAlgorithms.HmacSha512 }
                    };

                    var principal = handler.ValidateToken(token, validationParams, out SecurityToken validatedToken);
                    context.User = principal;

                    var jwtToken = validatedToken as JwtSecurityToken;
                    var exp = jwtToken?.Payload.Expiration;
                    var now = DateTimeOffset.UtcNow.ToUnixTimeSeconds();

                    if (exp.HasValue && exp.Value - now < 3600) // 1시간 이하 남았을 때
                    {
                        context.Response.Headers.Append(Constants.TokenRefreshRequired, "true");
                    }
                }
                catch
                {
                    context.Response.StatusCode = 401;
                    await context.Response.WriteAsJsonAsync(new { message = Constants.TokenUnauthorized });
                    return;
                }
            }
            else
            {
                context.Response.StatusCode = 401;
                await context.Response.WriteAsJsonAsync(new { message = Constants.TokenMissingAuthorizationHeader });
                return;
            }

            await _next(context);
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
