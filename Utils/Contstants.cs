namespace BerryNameApi.Utils
{
    public static class Constants
    {
        // 공통 메시지
        public const string Success = "success";
        public const string Error = "error";

        // DB
        public const string DatabaseError = "Database Error";

        // Token
        public const string KeyNotSet = "JWT_SECRET_KEY not set";
        public const string TokenRefreshRequired = "X-Token-Refresh-Required";
        public const string TokenUnauthorized = "Unauthorized";
        public const string TokenMissingAuthorizationHeader = "Missing Authorization Header";
        public const string InvalidToken = "Invalid Token";
        public const string TokenInvalidOrExpired = "Invalid or expired token.";

        // Name
        public const string NameRequired = "name is required";
        public const string NameLengthInvalid = "name must be between 1 and 50 characters";
        public const string DuplicateName = "name already exists";
        public const string NameNotFound = "No user with that name";
        public const string DeleteIndexRequired = "deleteIndex is required";
        public const string InvalidIndex = "Invalid index";

        // Company
        public const string UserAlreadyCompany = "User Already Has Company";
        public const string BossCreateCompany = "Only boss can create a company";
        public const string CompanyNotFound = "Company not found";

        // Auth
        public const string UserNotFound = "User not found";
        public const string EmailAlreadyExist = "Email already exist";
        public const string IncorrectPassword = "Incorrect password";
        public const string InvalidUID = "Invalid UserID";
    }
}