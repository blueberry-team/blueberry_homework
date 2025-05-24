#!/bin/bash

# API 테스트 스크립트
BASE_URL="http://localhost:8080"

echo "🧪 API 테스트 시작..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

echo "🔍 1. 서버 헬스체크"
curl -s -X GET "$BASE_URL/health" | jq .
echo -e "\n"

echo "👤 2. 사용자 생성 테스트"
echo "   2-1. 사용자 'Kim' 생성:"
curl -s -X POST "$BASE_URL/create-name" \
  -H "Content-Type: application/json" \
  -d '{"name": "Kim"}' | jq .

echo -e "\n   2-2. 사용자 'Park' 생성:"
curl -s -X POST "$BASE_URL/create-name" \
  -H "Content-Type: application/json" \
  -d '{"name": "Park"}' | jq .

echo -e "\n   2-3. 사용자 'Lee' 생성:"
curl -s -X POST "$BASE_URL/create-name" \
  -H "Content-Type: application/json" \
  -d '{"name": "Lee"}' | jq .

echo -e "\n📋 3. 사용자 목록 조회"
USERS_RESPONSE=$(curl -s -X GET "$BASE_URL/get-names")
echo "$USERS_RESPONSE" | jq .

# 첫 번째 사용자의 ID 추출 (이름 변경용)
USER_ID=$(echo "$USERS_RESPONSE" | jq -r '.data[0].id')
echo -e "\n   추출된 사용자 ID: $USER_ID"

echo -e "\n❌ 4. 중복 사용자 생성 시도"
curl -s -X POST "$BASE_URL/create-name" \
  -H "Content-Type: application/json" \
  -d '{"name": "Kim"}' | jq .

echo -e "\n✏️ 5. 사용자 이름 변경 테스트"
curl -s -X PUT "$BASE_URL/change-name" \
  -H "Content-Type: application/json" \
  -d "{\"id\": \"$USER_ID\", \"name\": \"KimChanged\"}" | jq .

echo -e "\n📋 6. 변경 후 사용자 목록 조회"
curl -s -X GET "$BASE_URL/get-names" | jq .

echo -e "\n🏢 7. 회사 생성 테스트"
echo "   7-1. KimChanged의 회사 'ABC Corp' 생성:"
curl -s -X POST "$BASE_URL/create-company" \
  -H "Content-Type: application/json" \
  -d '{"name": "KimChanged", "company_name": "ABC Corp"}' | jq .

echo -e "\n   7-2. Park의 회사 'XYZ Inc' 생성:"
curl -s -X POST "$BASE_URL/create-company" \
  -H "Content-Type: application/json" \
  -d '{"name": "Park", "company_name": "XYZ Inc"}' | jq .

echo -e "\n🏢 8. 회사 목록 조회"
curl -s -X GET "$BASE_URL/get-companies" | jq .

echo -e "\n❌ 9. 중복 회사 생성 시도"
curl -s -X POST "$BASE_URL/create-company" \
  -H "Content-Type: application/json" \
  -d '{"name": "KimChanged", "company_name": "Another Corp"}' | jq .

echo -e "\n🗑️ 10. 사용자 삭제 테스트"
echo "   10-1. 인덱스 2번 사용자 삭제:"
curl -s -X DELETE "$BASE_URL/delete-index?index=2" | jq .

echo -e "\n📋 11. 삭제 후 사용자 목록 조회"
curl -s -X GET "$BASE_URL/get-names" | jq .

echo -e "\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ API 테스트 완료!"
