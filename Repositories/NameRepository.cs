using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BerryNameApi.Entities;

namespace BerryNameApi.Repositories
{
    public class NameRepository
    {
        private readonly List<UserEntity> _store = new();

        public IEnumerable<UserEntity> GetAll() => _store;

        // 생성
        public void CreateName(UserEntity user) => _store.Add(user);

        // 인덱스로 삭제

        public bool DeleteByIndex(int index)
        {
            if (index < 0 || index >= _store.Count) return false;
            _store.RemoveAt(index);
            return true;
        }

        // 이름으로 삭제

        public int DeleteByName(string name)
        {
            var countBefore = _store.Count;
            _store.RemoveAll(x => x.Name == name);
            return countBefore - _store.Count;
        }
    }
}