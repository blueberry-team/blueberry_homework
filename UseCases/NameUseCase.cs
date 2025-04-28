using BerryNameApi.DTO.Response;
using BerryNameApi.Entities;
using BerryNameApi.Repositories;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.UseCases;

namespace BerryNameApi.UseCases
{
    public class NameUseCase
    {
        private readonly NameRepository _repository;

        public NameUseCase(NameRepository repository)
        {
            _repository = repository;
        }

        public IEnumerable<UserResponse> GetAll() => _repository.GetAll().Select(user => new UserResponse
        {
            Id = user.Id,
            Name = user.Name,
            CreatedAt = user.CreatedAt,
            UpdatedAt = user.UpdatedAt,
        });

        public Result CreateName(string name)
        {
            // 이름 중복 검색
            if (_repository.FindByName(name) != null)
            {
                return Result.Fail(Constants.DuplicateName);
            }

            var UtcNow = DateTime.UtcNow;

            var user = new UserEntity
            {
                Id = Guid.NewGuid(),
                Name = name,
                CreatedAt = UtcNow,
                UpdatedAt = UtcNow,
            };
            _repository.CreateName(user);
            return Result.Ok();
        }

        public Result ChangeName(Guid id, string newName)
        {
            var user = _repository.FindById(id);
            if (user == null)
            {
                return Result.Fail(Constants.UserNotFound);
            }

            // 이전 이름과 동일 이름
            if (user.Name == newName)
            {
                return Result.Fail(Constants.DuplicateName);
            }

            // 이름 중복
            var otherUser = _repository.FindByName(newName);
            if (otherUser != null)
            {
                return Result.Fail(Constants.DuplicateName);
            }

            user.Name = newName;
            user.UpdatedAt = DateTime.UtcNow;

            return Result.Ok();
        }

        public Result DeleteByIndex(int index)
        {
            if (!_repository.DeleteByIndex(index))
            {
                return Result.Fail($"{Constants.InvalidIndex}: {index}");
            }

            return Result.Ok();
        }

        public Result DeleteByName(string name)
        {
            if (!_repository.DeleteByName(name))
            {
                return Result.Fail(Constants.NameNotFound);
            }

            return Result.Ok();
        }
    }
}