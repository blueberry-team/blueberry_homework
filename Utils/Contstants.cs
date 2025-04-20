using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace BerryNameApi.Utils
{
    public static class Constants
    {
        // 공통 메시지
        public const string Success = "success";
        public const string Error = "error";

        // 유효성 검사 에러
        public const string NameRequired = "name is required";
        public const string NameLengthInvalid = "name must be between 1 and 50 characters";

        // 삭제 관련 에러
        public const string DeleteIndexRequired = "deleteIndex is required";
        public const string InvalidIndex = "Invalid index";

        // 이름 중복
        public const string DuplicateName = "name already exists";

        // 이름 삭제 실패
        public const string NameNotFound = "No user with that name";
    }
}