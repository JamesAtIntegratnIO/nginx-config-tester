from locust import between, task
from locust.contrib.fasthttp import FastHttpUser

class HealthTest(FastHttpUser):
    wait_time = between(5, 15)

    
    @task
    def health(self):
        self.client.get("/health")
    @task
    def health_with_header(self):
        h = {
            'my-dumb-token': 8675309,
        }

        self.client.get("/health2", headers=h)