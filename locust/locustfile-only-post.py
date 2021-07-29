from locust import HttpUser, task, between
import json
default_headers = {'Content-Type': 'application/json', 'Cache-Control': "no-cache",
                   'User-Agent': 'locust', "Authorization": "Bearer abc8383xx"}
f = open('../3mb-growth_json.json',)
new_post = json.load(f)
f.close()

class QuickstartUser(HttpUser):
    @task
    def post_growth(self):
        wait_time = between(1, 2)
        self.client.post("/api/v1/growth",
                        headers=default_headers, json=new_post, name="/api/v1/growth")
