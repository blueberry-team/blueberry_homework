namespace blueberry_homework_dotnet.UseCases
{
    public class Result<T>
    {
        public bool Success { get; }
        public string? ErrorMessage { get; }
        public T? Data { get; }

        private Result(bool success, string? errorMessage, T? data)
        {
            Success = success;
            ErrorMessage = errorMessage;
            Data = data;
        }

        public static Result<T> Ok(T data) => new Result<T>(true, null, data);
        public static Result<T> Fail(string errorMessage) => new Result<T>(false, errorMessage, default);
    }
}
