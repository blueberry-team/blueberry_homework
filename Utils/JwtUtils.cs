using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using Microsoft.IdentityModel.Tokens;

namespace blueberry_homework_dotnet.Utils
{
    public class JwtUtils
    {
        public static string GenerateToken(Guid userId, string email, string name)
        {
            var secretKey = Environment.GetEnvironmentVariable("JWT_SECRET_KEY")!;
            var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(secretKey));
            var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha512);

            var claims = new[]
            {
                // UID
                new Claim(JwtRegisteredClaimNames.Sub, userId.ToString()),
                // Email
                new Claim("email", email),
                // Name
                new Claim("name", name),
                // IAT
                new Claim(JwtRegisteredClaimNames.Iat, DateTimeOffset.UtcNow.ToUnixTimeSeconds().ToString(), ClaimValueTypes.Integer64)
            };

            var token = new JwtSecurityToken(
                claims: claims,
                // Expiration
                expires: DateTime.UtcNow.AddHours(5),
                signingCredentials: creds
            );

            return new JwtSecurityTokenHandler().WriteToken(token);
        }

        public static IDictionary<string, object>? Decode(string token)
        {
            var secret = Environment.GetEnvironmentVariable("JWT_SECRET_KEY") ?? throw new Exception("JWT secret not found");
            var handler = new JwtSecurityTokenHandler();

            var validations = new TokenValidationParameters
            {
                ValidateIssuer = false,
                ValidateAudience = false,
                ValidateLifetime = false, // 만료돼도 디코딩 가능
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(secret)),
                ValidAlgorithms = new[] { SecurityAlgorithms.HmacSha512 }
            };

            try
            {
                handler.ValidateToken(token, validations, out var validatedToken);
                return ((JwtSecurityToken)validatedToken).Payload;
            }
            catch
            {
                return null;
            }
        }
    }
}