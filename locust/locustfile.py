from os import times
from locust import HttpUser, task, between
import json
default_headers = {'Content-Type': 'application/json', 'Cache-Control': "no-cache",
                   'User-Agent': 'locust', "Authorization": "Bearer abc8383xx"}
f = open('../3mb-growth_json.json',)
new_post = json.load(f)


class QuickstartUser(HttpUser):
    @task
    def get_ping(self):
        wait_time = between(1, 2)
        self.client.get("/ping", headers={'Content-Type': 'application/json'})

    @task
    def get_grow(self):
        wait_time = between(1, 2)
        for item_id in range(10):
            self.client.get("/api/v1/growth/brz/ngdp_r/2002",
                            name="/api/v1/growth/brz/ngdp_r/2002")
            times.sleep(1)
            
    @task
    def put_grow(self):
        wait_time = between(1, 2)
        for item_id in range(10):
             self.client.put("/api/v1/growth/brz/ngdp_r/2002",
                        json={'value': 345.55}, name="/api/v1/growth/brz/ngdp_r/2002")
             times.sleep(1)
