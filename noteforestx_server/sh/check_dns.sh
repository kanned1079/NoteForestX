#!/usr/bin/env bash

# ===== 可配置项 =====
DNS_SERVER="192.168.0.1"     # 你的 DNS 服务器（如 AdGuard）
TEST_DOMAINS=(
  "www.google.com"
  "www.youtube.com"
  "www.baidu.com"
  "www.qq.com"
)
TIMEOUT=2
# ===================

echo "== DNS Check Script =="
echo "Using DNS server: $DNS_SERVER"
echo

for domain in "${TEST_DOMAINS[@]}"; do
  echo "Query: $domain"

  result=$(dig @"$DNS_SERVER" "$domain" +time=$TIMEOUT +tries=1 +short)

  if [[ -z "$result" ]]; then
    echo "  ❌ FAIL: no response"
  else
    echo "  ✅ OK:"
    echo "$result" | sed 's/^/     -> /'
  fi

  echo
done