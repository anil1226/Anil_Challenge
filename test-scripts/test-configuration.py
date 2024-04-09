import requests
import socket

class ServerConfigurationTests:
    def __init__(self, server_address):
        self.server_address = server_address

    def test_ports(self):
        
        expected_ports = [80, 443]
        
        
        open_ports = []
        for port in expected_ports:
            s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            s.settimeout(5)
            try:
                s.connect((self.server_address, port))
                open_ports.append(port)
            except:
                pass
            s.close()
        
        
        assert set(expected_ports) == set(open_ports), f"Unexpected open ports: {set(open_ports) - set(expected_ports)}"

    def test_http_to_https(self):
       
        http_response = requests.get(f"http://{self.server_address}", verify=False)
        
        
        assert http_response.history, "No HTTP to HTTPS redirection"
        assert http_response.history[0].status_code == 301, "Unexpected HTTP status code"
        assert http_response.url.startswith("https://"), "Redirect is not to HTTPS"



server_address = "44.222.43.178"
tests = ServerConfigurationTests(server_address)
tests.test_ports()
tests.test_http_to_https()
print("All tests passed!")