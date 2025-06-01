using BerryNameApi.DTO.Response;
using BerryNameApi.Entities;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.DTO.Request.Auth;
using blueberry_homework_dotnet.DTO.Response;
using blueberry_homework_dotnet.Repositories;
using blueberry_homework_dotnet.Utils;

namespace blueberry_homework_dotnet.UseCases
{
    public class AuthUseCase
    {
        private readonly AuthRepository _repository;

        public AuthUseCase(AuthRepository repository)
        {
            _repository = repository;
        }

        public Result<Unit> SignUp(SignUpRequest request)
        {
            if (_repository.FindByEmail(request.Email) != null)
            {
                return Result<Unit>.Fail(Constants.EmailAlreadyExist);
            }

            var currentTime = DateTime.UtcNow;
            var (hashedPassword, salt) = HashMaker.Hash(request.Password);

            var user = new UserEntity
            {
                Id = Guid.NewGuid(),
                Email = request.Email,
                PasswordHashed = hashedPassword,
                PasswordSalt = salt,
                Name = request.Name,
                Address = request.Address,
                Role = request.Role.ToLower() == "boss" ? Role.Boss : Role.Worker,
                CreatedAt = currentTime,
                UpdatedAt = currentTime
            };

            _repository.CreateUser(user);
            Console.WriteLine($"✅ Created user: {request.Email}, Role: {request.Role} id: {user.Id}");
            return Result<Unit>.Ok(Unit.Value);
        }

        public Result<AuthResponse> LogIn(LogInRequest request)
        {
            var user = _repository.FindByEmail(request.Email);

            if (user == null)
            {
                return Result<AuthResponse>.Fail(Constants.UserNotFound);
            }

            if (!HashMaker.Verify(request.Password, user.PasswordHashed, user.PasswordSalt))
            {
                return Result<AuthResponse>.Fail(Constants.IncorrectPassword);
            }

            // 토큰 발급
            string token = JwtTokenGenerator.GenerateToken(user.Id, user.Email, user.Name);

            AuthResponse authResponse = new AuthResponse
            {
                Id = user.Id,
                Email = user.Email,
                Name = user.Name,
                Address = user.Address,
                Role = user.getRole(),
                CreatedAt = user.CreatedAt,
                UpdatedAt = user.UpdatedAt,
                Token = token
            };

            return Result<AuthResponse>.Ok(authResponse);
        }

        public Result<Unit> ChangeUser(ChangeUserRequest request)
        {
            var user = _repository.FindById(request.Id);
            if (user == null) return Result<Unit>.Fail(Constants.UserNotFound);

            if (!string.IsNullOrEmpty(request.Password))
            {
                var (hashed, salt) = HashMaker.Hash(request.Password);
                user.PasswordHashed = hashed;
                user.PasswordSalt = salt;
            }

            if (!string.IsNullOrEmpty(request.Address))
            {
                user.Address = request.Address;
            }

            user.UpdatedAt = DateTime.UtcNow;
            _repository.UpdateUser(user);

            return Result<Unit>.Ok(Unit.Value);
        }

        public Result<AuthResponse> GetUser(Guid id)
        {
            var user = _repository.FindById(id);
            if (user == null)
            {
                return Result<AuthResponse>.Fail(Constants.UserNotFound);
            }

            var response = new AuthResponse
            {
                Id = user.Id,
                Email = user.Email,
                Name = user.Name,
                Address = user.Address,
                Role = user.Role.ToString(),
                CreatedAt = user.CreatedAt,
                UpdatedAt = user.UpdatedAt,
                Token = string.Empty
            };

            return Result<AuthResponse>.Ok(response);
        }
    }
}