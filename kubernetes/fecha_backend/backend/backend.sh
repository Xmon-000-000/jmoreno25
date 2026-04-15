#!/bin/bash
set -e

echo "Servidor HTTP escuchando en puerto 8080"

while true; do
  {
    printf "HTTP/1.1 200 OK\r\n"
    printf "Content-Type: text/plain; charset=utf-8\r\n"
    printf "\r\n"
    printf "Fecha y hora del sistema: %s\n" "$(date)"
  } | nc -lk -p 8080
done


