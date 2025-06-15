import time
from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
import uvicorn

app = FastAPI()


@app.get("/", response_class=PlainTextResponse)
def index():
    with open ("./files/logs.txt", "r") as file:
        logs = file.read()
    return PlainTextResponse(content=logs)

if __name__=="__main__":
    uvicorn.run(app, host="0.0.0.0", port=3000)
