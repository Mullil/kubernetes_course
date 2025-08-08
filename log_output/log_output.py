import time
from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
import uvicorn
import requests
import datetime
import random
import string
import threading
import os

random_str = ''.join(random.choices(string.ascii_lowercase, k=10))
current_time = datetime.datetime.now(datetime.timezone.utc)
app = FastAPI()


@app.get("/", response_class=PlainTextResponse)
def index():
    message = os.getenv('MESSAGE')
    with open('/etc/config/information.txt', 'r') as f:
        content = f.read()
    pongs = requests.get("http://ping-pong-svc:2345/pings").text
    data = f"""file content: {content}\nenv variable: {message}\n
{current_time}: {random_str} \nPing / Pongs: {pongs}"""
    return PlainTextResponse(content=data)

def update_loop():
    global current_time
    while True:
        current_time = datetime.datetime.now(datetime.timezone.utc)
        print(f"{current_time}: {random_str}", flush=True)
        time.sleep(5)

if __name__=="__main__":
    threading.Thread(target=update_loop, daemon=True).start()
    uvicorn.run(app, host="0.0.0.0", port=3000)
