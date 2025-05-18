using BerryNameApi.Entities;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.DTO.Request.Auth;
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

        public Result SignUp(SignUpRequest request)
        {
            if (_repository.FindByEmail(request.Email) != null)
            {
                return Result.Fail(Constants.EmailAlreadyExist);
            }

            var currentTime = DateTime.UtcNow;

            var user = new UserEntity
            {
                Id = Guid.NewGuid(),
                Email = request.Email,
                PasswordHashed = HashMaker.Hash(request.Password),
                Name = request.Name,
                Address = request.Address,
                Role = request.Role.ToLower() == "boss" ? Role.Boss : Role.Worker,
                CreatedAt = currentTime,
                UpdatedAt = currentTime
            };

            _repository.CreateUser(user);
            Console.WriteLine($"✅ Created user: {request.Email}, Role: {request.Role} id: {user.Id}");
            return Result.Ok();
        }

        public Result LogIn(LogInRequest request)
        {
            var user = _repository.FindByEmail(request.Email);

            if (user == null)
            {
                return Result.Fail(Constants.UserNotFound);
            }

            if (!HashMaker.Verify(request.Password, user.PasswordHashed))
            {
                return Result.Fail(Constants.IncorrectPassword);
            }

            return Result.Ok();
        }

        public Result ChangeUser(ChangeUserRequest request)
        {
            var user = _repository.FindById(request.Id);
            if (user == null) return Result.Fail(Constants.UserNotFound);

            if (!string.IsNullOrEmpty(request.Password))
            {
                user.PasswordHashed = HashMaker.Hash(request.Password);
            }
            if (!string.IsNullOrEmpty(request.Address))
            {
                user.Address = request.Address;
            }

            user.UpdatedAt = DateTime.UtcNow;
            _repository.UpdateUser(user);

            return Result.Ok();
        }

        public UserEntity? GetUser(Guid id)
        {
            return _repository.FindById(id);
        }
    }
}