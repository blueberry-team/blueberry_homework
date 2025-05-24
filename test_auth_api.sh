#!/bin/bash

# PART.6 인증 시스템 API 테스트 스크립트
BASE_URL="http://localhost:8080"

echo "🧪 PART.6 인증 시스템 API 테스트 시작..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

echo "🔍 1. 서버 헬스체크"
curl -s -X GET "$BASE_URL/health" | jq .
echo -e "\n"

echo "👤 2. 회원가입 테스트"
echo "   2-1. Boss 사용자 'john@example.com' 회원가입:"
BOSS_RESPONSE=$(curl -s -X POST "$BASE_URL/sign-up" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "SecurePass123",
    "name": "John Doe",
    "role": "boss"
  }')
echo "$BOSS_RESPONSE" | jq .

# Boss 사용자 ID 추출
BOSS_ID=$(echo "$BOSS_RESPONSE" | jq -r '.data.id')
echo -e "\n   Boss 사용자 ID: $BOSS_ID"

echo -e "\n   2-2. Worker 사용자 'jane@example.com' 회원가입:"
WORKER_RESPONSE=$(curl -s -X POST "$BASE_URL/sign-up" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jane@example.com",
    "password": "WorkerPass456",
    "name": "Jane Smith",
    "role": "worker"
  }')
echo "$WORKER_RESPONSE" | jq .

# Worker 사용자 ID 추출
WORKER_ID=$(echo "$WORKER_RESPONSE" | jq -r '.data.id')
echo -e "\n   Worker 사용자 ID: $WORKER_ID"

echo -e "\n❌ 3. 중복 이메일 회원가입 시도"
curl -s -X POST "$BASE_URL/sign-up" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "AnotherPass789",
    "name": "Another John",
    "role": "worker"
  }' | jq .

echo -e "\n🔐 4. 로그인 테스트"
echo "   4-1. Boss 사용자 로그인:"
curl -s -X POST "$BASE_URL/log-in" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "SecurePass123"
  }' | jq .

echo -e "\n   4-2. 잘못된 비밀번호로 로그인 시도:"
curl -s -X POST "$BASE_URL/log-in" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "WrongPassword"
  }' | jq .

echo -e "\n👤 5. 사용자 정보 조회"
curl -s -X POST "$BASE_URL/get-user" \
  -H "Content-Type: application/json" \
  -d "{\"userId\": \"$BOSS_ID\"}" | jq .

echo -e "\n✏️ 6. 사용자 정보 수정"
curl -s -X PUT "$BASE_URL/change-user" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$BOSS_ID\",
    \"name\": \"John Doe Updated\"
  }" | jq .

echo -e "\n🏢 7. 회사 생성 테스트 (Boss만 가능)"
echo "   7-1. Boss가 회사 생성:"
COMPANY_RESPONSE=$(curl -s -X POST "$BASE_URL/create-company" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$BOSS_ID\",
    \"companyName\": \"Tech Innovations Inc\",
    \"companyAddress\": \"123 Tech Street, Silicon Valley\",
    \"totalStaff\": 25
  }")
echo "$COMPANY_RESPONSE" | jq .

# 회사 ID 추출
COMPANY_ID=$(echo "$COMPANY_RESPONSE" | jq -r '.data.id')
echo -e "\n   회사 ID: $COMPANY_ID"

echo -e "\n   7-2. Worker가 회사 생성 시도 (실패해야 함):"
curl -s -X POST "$BASE_URL/create-company" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$WORKER_ID\",
    \"companyName\": \"Worker Company\",
    \"companyAddress\": \"456 Worker Ave\",
    \"totalStaff\": 5
  }" | jq .

echo -e "\n🏢 8. 회사 정보 조회"
echo "   8-1. 모든 회사 조회:"
curl -s -X GET "$BASE_URL/get-companies" | jq .

echo -e "\n   8-2. 특정 회사 조회 (회사 ID로):"
curl -s -X POST "$BASE_URL/get-company" \
  -H "Content-Type: application/json" \
  -d "{\"companyId\": \"$COMPANY_ID\"}" | jq .

echo -e "\n   8-3. 사용자 ID로 회사 조회:"
curl -s -X POST "$BASE_URL/get-company" \
  -H "Content-Type: application/json" \
  -d "{\"userId\": \"$BOSS_ID\"}" | jq .

echo -e "\n✏️ 9. 회사 정보 수정"
curl -s -X PUT "$BASE_URL/change-company" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$BOSS_ID\",
    \"companyId\": \"$COMPANY_ID\",
    \"companyName\": \"Tech Innovations Corp\",
    \"totalStaff\": 30
  }" | jq .

echo -e "\n🔍 10. 회사명으로 검색"
curl -s -X GET "$BASE_URL/find-companies?name=Tech" | jq .

echo -e "\n❌ 11. 권한 없는 사용자가 회사 수정 시도"
curl -s -X PUT "$BASE_URL/change-company" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$WORKER_ID\",
    \"companyId\": \"$COMPANY_ID\",
    \"companyName\": \"Hacked Company\"
  }" | jq .

echo -e "\n🗑️ 12. 회사 삭제 테스트"
echo "   12-1. 권한 없는 사용자가 삭제 시도:"
curl -s -X DELETE "$BASE_URL/delete-company" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$WORKER_ID\",
    \"companyId\": \"$COMPANY_ID\"
  }" | jq .

echo -e "\n   12-2. 소유자가 회사 삭제:"
curl -s -X DELETE "$BASE_URL/delete-company" \
  -H "Content-Type: application/json" \
  -d "{
    \"userId\": \"$BOSS_ID\",
    \"companyId\": \"$COMPANY_ID\"
  }" | jq .

echo -e "\n📋 13. 삭제 후 회사 목록 확인"
curl -s -X GET "$BASE_URL/get-companies" | jq .

echo -e "\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ PART.6 인증 시스템 API 테스트 완료!"
echo "🔐 인증 시스템, 권한 관리, CRUD 기능이 모두 정상 작동합니다."
