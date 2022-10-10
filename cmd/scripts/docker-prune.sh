#!/bin/bash
docker system prune --force
docker system prune --volumes --force
docker image prune -a --force
