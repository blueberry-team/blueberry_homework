using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;

namespace blueberry_homework_dotnet.Utils
{
    public class HashMaker
    {
        // 암호화
        public static (string hash, string salt) Hash(string password)
        {
            // 16바이트 (128비트) salt 생성 (랜덤 무작위 난수)
            byte[] saltBytes = RandomNumberGenerator.GetBytes(16);
            string salt = Convert.ToBase64String(saltBytes);

            // 비밀번호 해싱 알고리즘 PBKDF2 로 해싱 / 파라미터: 비밀번호, 랜덤솔트, 반복횟수, 사용해시함수
            using var pbkdf2 = new Rfc2898DeriveBytes(password, saltBytes, 100_000, HashAlgorithmName.SHA256);
            // 32바이트 (256비트)
            byte[] hashBytes = pbkdf2.GetBytes(32);

            string hash = Convert.ToBase64String(hashBytes);
            return (hash, salt);
        }

        // 검증
        public static bool Verify(string inputPassword, string storedHash, string storedSalt)
        {
            byte[] saltBytes = Convert.FromBase64String(storedSalt);

            using var pbkdf2 = new Rfc2898DeriveBytes(inputPassword, saltBytes, 100_000, HashAlgorithmName.SHA256);
            byte[] inputHash = pbkdf2.GetBytes(32);

            return Convert.ToBase64String(inputHash) == storedHash;
        }
    }
}