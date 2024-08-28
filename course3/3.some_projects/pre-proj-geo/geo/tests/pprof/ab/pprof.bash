#!/bin/bash

# Запускаем ab в фоновом режиме
ab -n 1000 -c 10 -p data.json "http://localhost:8080/api/address/geocode?lat=47.6062&lng=-122.3321" &

# Запускаем curl в другом терминале
curl "http://localhost:8080/debug/profile/?seconds=5" > ./profile &

# Ждем завершения обеих команд
wait

echo "Обе команды завершены."
