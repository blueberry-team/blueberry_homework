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
            if (_repository.FindByName(newName) != null)
            {
                return Result.Fail(Constants.DuplicateName);
            }

            user.Name = newName;
            user.UpdatedAt = DateTime.UtcNow;

            // DB 입력
            if (!_repository.ChangeName(user))
            {
                return Result.Fail(Constants.DatabaseError);
            }

            return Result.Ok();
        }

        public Result DeleteByIndex(int index)
        {
            var deleted = _repository.DeleteByIndex(index);
            return deleted
                ? Result.Ok()
                : Result.Fail($"{Constants.InvalidIndex}: {index}");
        }

        public Result DeleteByName(string name)
        {
            var count = _repository.DeleteByName(name);
            return count > 0
                ? Result.Ok()
                : Result.Fail(Constants.NameNotFound);
        }
    }
}