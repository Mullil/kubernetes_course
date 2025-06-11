import time
import random
import string
import datetime
from fastapi import FastAPI
from fastapi.responses import HTMLResponse
import uvicorn

app = FastAPI()

current_time = datetime.datetime.now(datetime.timezone.utc)
random_str = ''.join(random.choices(string.ascii_lowercase, k=10))


@app.get("/")
def index():
    current_time = datetime.datetime.now(datetime.timezone.utc)
    html = f'<html> <meta http-equiv="refresh" content="5"> <p>Time: {current_time}</p> <p>Random string: {random_str}</p> </html>'
    return HTMLResponse(content=html)

def update_loop():
    global current_time
    while True:
        current_time = datetime.datetime.now(datetime.timezone.utc)
        print(f"{current_time}: {random_str}", flush=True)
        time.sleep(5)

if __name__=="__main__":
    uvicorn.run(app, host="0.0.0.0", port=3000)
