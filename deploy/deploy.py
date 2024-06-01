#!/usr/bin/env python3

import os
import subprocess

HOME = os.path.expanduser("~")


def run_command(command):
    subprocess.run(command, shell=True, check=True)


def rebuild():
    run_command("git pull")
    build = HOME + "/cmd/website.go build"
    run_command(f"go run {build}")


def restart_caddy():
    docker_compose = HOME + "/caddy/docker_compose.yaml"
    run_command(f"sudo docker compose -f {docker_compose} restart")


if __name__ == "__main__":
    rebuild()
    restart_caddy()
