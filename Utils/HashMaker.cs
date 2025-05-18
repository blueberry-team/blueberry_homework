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
        public static string Hash(string password)
        {
            using var sha = SHA256.Create();
            var bytes = sha.ComputeHash(Encoding.UTF8.GetBytes(password));
            return Convert.ToBase64String(bytes);
        }

        // 검증
        public static bool Verify(string plain, string hash)
        {
            return Hash(plain) == hash;
        }
    }
}