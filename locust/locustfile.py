from locust import HttpUser, task, between
import json
token_string = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEwODNiYTUzLWNmZTktNDI2OC1iMzA5LTA4MjM4OGU2N2M2MiIsImV4cCI6MTYxOTQ3NTA0MSwiaXNzIjoiZ293YWthbmRhIn0.Qqe0aU8uMDWlxK9cZwObetGEEA1tFztfq0ZYe6i3mGNyF-U4Kq4oYqAf1h6RzTDmKiVCKPNj2GLE6soe07gQZP4bQYbaj5woNl05g94NABcxsqXA6ujlij9GOUXHWMlDlBBjqAnx5WIu8kogox4rFkD_a8NMbmj1XGl942VxbOM"
default_headers = {'Content-Type': 'application/json', 'Cache-Control': "no-cache",
                   'User-Agent': 'curl', "Authorization": "Bearer " + token_string}


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def post_growth(self):
        f = open('../3mb-growth_json.json',)
        new_post = json.load(f)
        self.client.post("/api/v1/growth",
                         headers=default_headers, json=new_post, name="/api/v1/growth")
        f.close()

    @task
    def get_ping(self):
        self.client.get("/ping", headers={'Content-Type': 'application/json'})

    @task
    def get_grow(self):
        self.client.get("/api/v1/growth/brz/ngdp_r/2002", headers={'Content-Type': 'application/json'},name="/api/v1/growth")
