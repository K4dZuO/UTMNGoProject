import requests
import random
import time
import json

URL = "http://localhost:8081/reviews"   # твой reviews сервис
REQUESTS_PER_SECOND = 10

def send_review():
    payload = {
        "product_id": random.randint(1, 100000),
        "rate": random.randint(1, 5)
    }

    try:
        r = requests.post(URL, json=payload)
        print("Status:", r.status_code, "Response:", r.text)
    except Exception as e:
        print("Error:", e)

if __name__ == "__main__":
    print("Sending 10 reviews per second...")
    delay = 1 / REQUESTS_PER_SECOND

    while True:
        send_review()
        time.sleep(delay)
