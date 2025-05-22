import os
from cassandra.cluster import Cluster, Session
from cassandra.auth import PlainTextAuthProvider
from cassandra.query import dict_factory
from dotenv import load_dotenv

# 환경 변수 로드
load_dotenv()


class ScyllaDB:
    _instance = None
    cluster = None
    session = None
    keyspace = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(ScyllaDB, cls).__new__(cls)
            cls._connect()
        return cls._instance

    @classmethod
    def _connect(cls):
        """ScyllaDB 클러스터에 연결하고 세션을 생성합니다."""
        try:
            # 환경 변수에서 연결 정보 가져오기
            host = os.getenv("SCYLLA_HOST", "scylla")
            port = int(os.getenv("SCYLLA_PORT", "9042"))
            cls.keyspace = os.getenv("SCYLLA_KEYSPACE", "mykeyspace")

            # 세션 생성 (db랑 연결하는 세션)
            cls.cluster = Cluster(
                contact_points=[
                    host,
                ],
                port=port,
                # auth_provider=PlainTextAuthProvider(
                #     username="scylla", password="your-awesome-password"
                # ),
            )
            cls.session = cls.cluster.connect()
            cls.session.row_factory = dict_factory

            # 키스페이스 생성 (없는 경우)
            # 키스페이스는 데이터베이스의 논리적 구분을 위한 단위
            cls._create_keyspace_if_not_exists()

            # 키스페이스 사용
            # 키스페이스를 설정하면 해당 키스페이스에 속한 테이블에 접근할 수 있음
            cls.session.set_keyspace(cls.keyspace)

            print(f"ScyllaDB에 성공적으로 연결되었습니다. 키스페이스: {cls.keyspace}")
        except Exception as e:
            print(f"ScyllaDB 연결 중 오류 발생: {e}")
            raise

    @classmethod
    def _create_keyspace_if_not_exists(cls):
        """키스페이스가 없을 경우 생성합니다."""
        query = f"""
        CREATE KEYSPACE IF NOT EXISTS {cls.keyspace}
        WITH REPLICATION = {{
            'class' : 'SimpleStrategy',
            'replication_factor' : 1
        }}
        """
        cls.session.execute(query)

    @classmethod
    def get_session(cls) -> Session:
        """ScyllaDB 세션을 반환합니다."""
        if cls.session is None:
            cls._connect()
        return cls.session

    @classmethod
    def close(cls):
        """연결을 종료합니다."""
        if cls.cluster:
            cls.cluster.shutdown()
            cls.session = None
            cls.cluster = None
            print("ScyllaDB 연결이 종료되었습니다.")


# 데이터베이스 초기화 함수
def init_db():
    """데이터베이스를 초기화하고 필요한 테이블을 생성합니다."""
    db = ScyllaDB()
    session = db.get_session()

    # 데이터베이스 초기화
    # 테이블 삭제, 테스트용
    delete_all(session)

    # 테이블 생성
    create_user_table(session)
    create_user_index_table(session)

    create_company_table(session)
    create_company_index_table(session)

    return db


def delete_all(session: Session):
    session.execute(
        """
    DROP TABLE IF EXISTS users
    """
    )

    session.execute(
        """
    DROP TABLE IF EXISTS companies
    """
    )


def create_user_table(session: Session):
    """필요한 테이블을 생성합니다."""
    # 사용자 테이블 생성
    session.execute(
        """
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        username TEXT,
        email TEXT,
        password TEXT,
        address TEXT,
        role TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
    """
    )

    print("유저 테이블이 성공적으로 생성되었습니다.")


def create_company_table(session: Session):
    """필요한 테이블을 생성합니다."""
    # 회사 테이블 생성
    session.execute(
        """
    CREATE TABLE IF NOT EXISTS companies (
        id UUID PRIMARY KEY,
        user_id TEXT,
        company_name TEXT,
        company_address TEXT,
        total_staff INT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
    """
    )

    print("컴퍼니 테이블이 성공적으로 생성되었습니다.")


def create_user_index_table(session: Session):
    """필요한 인덱스 테이블을 생성합니다."""
    # 인덱스 테이블 생성
    session.execute(
        """
    CREATE INDEX IF NOT EXISTS username_index ON users (username)
    """
    )

    print("유저 인덱스 테이블이 성공적으로 생성되었습니다.")


def create_company_index_table(session: Session):
    """필요한 인덱스 테이블을 생성합니다."""
    # 인덱스 테이블 생성
    session.execute(
        """
    CREATE INDEX IF NOT EXISTS name_index ON companies (company_name)
    """
    )

    print("컴퍼니 인덱스 테이블이 성공적으로 생성되었습니다.")


# 금지
# 이유 : 스칼라디비에서 키 없이 완전 탐색을 하면 찾아지지도 않고 성능도 구림
# let query = "SELECT id FROM company WHERE name = ? ALLOW FILTERING";


# 싱글톤 인스턴스 가져오기
def get_db():
    """데이터베이스 인스턴스를 반환합니다."""
    return ScyllaDB()
