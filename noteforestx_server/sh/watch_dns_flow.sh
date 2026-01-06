#!/usr/bin/env bash

DNS_IP="192.168.0.1"    # 你的 DNS
IFACE="any"

echo "[DNS FLOW WATCH]"
echo "Expect DNS -> $DNS_IP"
echo "Press Ctrl+C to stop"
echo

sudo tcpdump -i "$IFACE" port 53 -n | while read -r line; do
  if echo "$line" | grep -q "$DNS_IP"; then
    echo "✅ OK   $line"
  else
    echo "❌ BYPASS $line"
  fi
done