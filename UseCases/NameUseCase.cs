using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BerryNameApi.DTO.Response;
using BerryNameApi.Entities;
using BerryNameApi.Repositories;

namespace BerryNameApi.UseCases
{
    public class NameUseCase
    {
        private readonly NameRepository _repository;

        public NameUseCase(NameRepository repository)
        {
            _repository = repository;
        }

        public IEnumerable<UserResponse> GetAll() => _repository.GetAll().Select(u => new UserResponse
        {
            Name = u.Name,
            CreatedAt = u.CreatedAt
        });

        public void CreateName(string name)
        {
            var entity = new UserEntity
            {
                Name = name,
                CreatedAt = DateTime.UtcNow
            };
            _repository.CreateName(entity);
        }

        public bool DeleteByIndex(int index) => _repository.DeleteByIndex(index);

        public int DeleteByName(string name) => _repository.DeleteByName(name);
    }
}