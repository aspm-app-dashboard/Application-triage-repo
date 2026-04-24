import os
import subprocess
import sqlite3
import requests

# Hardcoded secrets (Gitleaks should catch these)
AWS_ACCESS_KEY = "AKIAIOSFODNN7EXAMPLE"
AWS_SECRET_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
DB_PASSWORD = "supersecretpassword123"
API_TOKEN = "ghp_1234567890abcdefghijklmnopqrst"

# Insecure use of eval (code injection risk)
def run_user_code(user_input):
    return eval(user_input)

# Command injection vulnerability
def ping_host(host):
    command = f"ping -c 1 {host}"
    return subprocess.getoutput(command)

# SQL injection vulnerability
def get_user(username):
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    query = f"SELECT * FROM users WHERE username = '{username}'"
    cursor.execute(query)
    return cursor.fetchall()

# Insecure HTTP request (no SSL verification)
def fetch_data(url):
    return requests.get(url, verify=False)

# Weak random number generation
def generate_token():
    import random
    return str(random.randint(100000, 999999))

# Hardcoded private key (major leak)
PRIVATE_KEY = """-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA7fakeexamplekeyfakeexamplekeyfakeexamplekey
-----END RSA PRIVATE KEY-----"""

# Unsafe file handling
def read_file(filename):
    with open(filename, "r") as f:
        return f.read()

# Debug mode enabled (sensitive info exposure)
DEBUG = True

if __name__ == "__main__":
    user_input = input("Enter command: ")
    print(run_user_code(user_input))
