namespace BerryNameApi.DTO.Response
{
    public class ApiSuccessResponse<T>
    {
        public string Message { get; set; } = "success";
        private T? _data;

        // 반환시 data 입력 안하면 [] 빈 배열 반환
        public T? Data
        {
            get => _data ?? (typeof(T).IsGenericType && typeof(T).GetGenericTypeDefinition() == typeof(IEnumerable<>)
                ? (T)Activator.CreateInstance(typeof(List<>).MakeGenericType(typeof(T).GetGenericArguments()[0]))!
                : default);
            set => _data = value;
        }
    }
}