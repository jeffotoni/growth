from locust import HttpUser, task, between
token_string = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEwODNiYTUzLWNmZTktNDI2OC1iMzA5LTA4MjM4OGU2N2M2MiIsImV4cCI6MTYxOTQ3NTA0MSwiaXNzIjoiZ293YWthbmRhIn0.Qqe0aU8uMDWlxK9cZwObetGEEA1tFztfq0ZYe6i3mGNyF-U4Kq4oYqAf1h6RzTDmKiVCKPNj2GLE6soe07gQZP4bQYbaj5woNl05g94NABcxsqXA6ujlij9GOUXHWMlDlBBjqAnx5WIu8kogox4rFkD_a8NMbmj1XGl942VxbOM"
default_headers = {'Content-Type':'application/json','Cache-Control': "no-cache",'User-Agent': 'curl',"Authorization":"Bearer " + token_string}
class QuickstartUser(HttpUser):
    wait_time = between(1, 2)
    @task
    def post_buy(self):
        self.client.post("/api/v1/order/buy",headers=default_headers,json={"user_uuid":"a035c945-e052-4d7a-ba77-419a73ec6580","msguuid":"356bd99a-894d-47de-bf65-842bc8660639","share_code":"VIBR5","share_name":"VIBRANIUM PN N2","broker":"Wakanda.Broker","operation":"daytrade","duration":"hoje","datevenc":"2021-04-26","qtd":100,"price":22.23,"timestamp":"2021-04-01T08:17:24-03:00"},name="POST")

    @task
    def get_healthz(self):
   		self.client.get("/api/v1/order/buy/healthz",headers={'Content-Type': 'application/json'})