namespace BerryNameApi.Utils
{
    public static class Constants
    {
        // 공통 메시지
        public const string Success = "success";
        public const string Error = "error";

        // DB
        public const string DatabaseError = "Database Error";

        // Name
        public const string NameRequired = "name is required";
        public const string NameLengthInvalid = "name must be between 1 and 50 characters";
        public const string DuplicateName = "name already exists";
        public const string NameNotFound = "No user with that name";
        public const string DeleteIndexRequired = "deleteIndex is required";
        public const string InvalidIndex = "Invalid index";

        // Company
        public const string UserAlreadyCompany = "User Already Has Company";

        // Auth
        public const string UserNotFound = "User not found";
        public const string EmailAlreadyExist = "Email already exist";
        public const string IncorrectPassword = "Incorrect password";
    }
}