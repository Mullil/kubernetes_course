import time
import random
import string
import datetime

random_str = ''.join(random.choices(string.ascii_lowercase, k=10))

while True:
    current_time = datetime.datetime.now(datetime.timezone.utc)
    print(f"{current_time}: {random_str}", flush=True)
    time.sleep(5)
