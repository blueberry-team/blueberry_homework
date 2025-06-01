using blueberry_homework_dotnet.Repositories;
using blueberry_homework_dotnet.DTO.Response;
using blueberry_homework_dotnet.Utils;
using BerryNameApi.Utils;

namespace blueberry_homework_dotnet.UseCases
{
    public class RefreshUseCase
    {
        private readonly AuthRepository _repository;

        public RefreshUseCase(AuthRepository repository)
        {
            _repository = repository;
        }

        public Result<AuthResponse> Refresh(string token)
        {
            var payload = JwtUtils.Decode(token);
            if (payload == null || !payload.TryGetValue("sub", out var userIdString)) return Result<AuthResponse>.Fail(Constants.InvalidToken);

            if (!Guid.TryParse(userIdString?.ToString(), out Guid userId)) return Result<AuthResponse>.Fail(Constants.InvalidUID);

            var user = _repository.FindById(userId);
            if (user == null) return Result<AuthResponse>.Fail(Constants.UserNotFound);

            var newToken = JwtUtils.GenerateToken(user.Id, user.Email, user.Name);

            return Result<AuthResponse>.Ok(new AuthResponse
            {
                Id = user.Id,
                Email = user.Email,
                Name = user.Name,
                Address = user.Address,
                Role = user.Role.ToString(),
                CreatedAt = user.CreatedAt,
                UpdatedAt = user.UpdatedAt,
                Token = newToken
            });
        }
    }
}
