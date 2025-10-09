#!/bin/bash

POD_NS=${POD_NS:-default}
DEPLOYMENT=${DEPLOYMENT:-simple-http-waypoint}

istioctl pc all -n $POD_NS deploy/$DEPLOYMENT -o yaml > dumps/waypoint-all.yaml
istioctl pc c -n $POD_NS deploy/$DEPLOYMENT -o yaml > dumps/waypoint.cluster.yaml
istioctl pc l -n $POD_NS deploy/$DEPLOYMENT -o yaml > dumps/waypoint.listener.yaml
istioctl pc r -n $POD_NS deploy/$DEPLOYMENT -o yaml > dumps/waypoint.route.yaml


istioctl pc all -n $POD_NS deploy/$DEPLOYMENT -o json > dumps/waypoint-all.json
istioctl pc c -n $POD_NS deploy/$DEPLOYMENT -o json > dumps/waypoint.cluster.json
istioctl pc l -n $POD_NS deploy/$DEPLOYMENT -o json > dumps/waypoint.listener.json
istioctl pc r -n $POD_NS deploy/$DEPLOYMENT -o json > dumps/waypoint.route.json