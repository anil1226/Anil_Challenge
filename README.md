This repository contains the code for SED Challenge. Hosted Here: https://44.222.43.178/

# Infrastructure

1. **Web Server Deployment**: 
--> Provisioned AWS EC2 instance with Terraform
--> Configured NGINX and Deployed Static Content using Ansible
--> Created seperate Github Actions pipelines for provisioning and configuring/deploying the app

2. **Security and HTTPS Configuration**: 
--> The web server is configured to use HTTPS with a self-signed certificate. 
--> The Ansible playbook handles the generation and configuration of the self-signed certificate, as well as the redirection of HTTP traffic to HTTPS.

3. **Automated Testing**: 
--> test script is developed using python to validate the correctness of the server configuration. 
--> Created a Github Actions pipeline to automate testing

# Coding

1. **Validating Credit Card Numbers**: 

--> Solved the problem in golang
--> Used regexp package to check validity of credit card numbers

