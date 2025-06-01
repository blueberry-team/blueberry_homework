using BerryNameApi.DTO.Response;
using BerryNameApi.Entities;
using BerryNameApi.Repositories;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.UseCases;
using blueberry_homework_dotnet.Utils;

namespace BerryNameApi.UseCases
{
    public class NameUseCase
    {
        private readonly NameRepository _repository;

        public NameUseCase(NameRepository repository)
        {
            _repository = repository;
        }

        public Result<IEnumerable<UserResponse>> GetAll()
        {
            var users = _repository.GetAll();

            var response = users.Select(user => new UserResponse
            {
                Id = user.Id,
                Name = user.Name,
                CreatedAt = user.CreatedAt,
                UpdatedAt = user.UpdatedAt,
            });

            return Result<IEnumerable<UserResponse>>.Ok(response);
        }

        public Result<Unit> CreateName(string name)
        {
            // 이름 중복 검색
            if (_repository.FindByName(name) != null)
            {
                return Result<Unit>.Fail(Constants.DuplicateName);
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
            return Result<Unit>.Ok(Unit.Value);
        }

        public Result<Unit> ChangeName(Guid id, string newName)
        {
            var user = _repository.FindById(id);
            if (user == null)
            {
                return Result<Unit>.Fail(Constants.UserNotFound);
            }

            // 이전 이름과 동일 이름
            if (user.Name == newName)
            {
                return Result<Unit>.Fail(Constants.DuplicateName);
            }

            // 이름 중복
            if (_repository.FindByName(newName) != null)
            {
                return Result<Unit>.Fail(Constants.DuplicateName);
            }

            user.Name = newName;
            user.UpdatedAt = DateTime.UtcNow;

            // DB 입력
            if (!_repository.ChangeName(user))
            {
                return Result<Unit>.Fail(Constants.DatabaseError);
            }

            return Result<Unit>.Ok(Unit.Value);
        }

        public Result<Unit> DeleteByIndex(int index)
        {
            var deleted = _repository.DeleteByIndex(index);
            return deleted
                ? Result<Unit>.Ok(Unit.Value)
                : Result<Unit>.Fail($"{Constants.InvalidIndex}: {index}");
        }

        public Result<Unit> DeleteByName(string name)
        {
            var count = _repository.DeleteByName(name);
            return count > 0
                ? Result<Unit>.Ok(Unit.Value)
                : Result<Unit>.Fail(Constants.NameNotFound);
        }
    }
}