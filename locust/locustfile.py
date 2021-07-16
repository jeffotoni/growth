from locust import HttpUser, task, between
import json
default_headers = {'Content-Type': 'application/json', 'Cache-Control': "no-cache",
                   'User-Agent': 'curl', "Authorization": "Bearer abc8383xx"}
f = open('../3mb-growth_json.json',)
new_post = json.load(f)

class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def post_growth(self):
        #f = open('../3mb-growth_json.json',)
        #new_post = json.load(f)
        self.client.post("/api/v1/growth",
                         headers=default_headers, json=new_post, name="/api/v1/growth")
        f.close()

    @task
    def get_ping(self):
        self.client.get("/ping", headers={'Content-Type': 'application/json'})

    #@task
    #def get_grow(self):
    #    self.client.get("/api/v1/growth/brz/ngdp_r/2002", headers={'Content-Type': 'application/json'},name="/api/v1/growth")
