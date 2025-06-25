#!/bin/bash

istioctl pc all -n ambient deploy/simple-http-waypoint -o yaml > dumps/waypoint-1.26.yaml
istioctl pc c -n ambient deploy/simple-http-waypoint -o yaml > dumps/waypoint-1.26.cluster.yaml
istioctl pc l -n ambient deploy/simple-http-waypoint -o yaml > dumps/waypoint-1.26.listener.yaml
istioctl pc r -n ambient deploy/simple-http-waypoint -o yaml > dumps/waypoint-1.26.route.yaml