ISTIOCTL_HOME ?= ~/Codes/istioctl

install-istio:
	$(ISTIOCTL_HOME)/istio-1.25.0/bin/istioctl install -f blue.yaml -y
	istioctl install -f green.yaml -y

install-istio-dev:
	istioctl-dev install -f blue.yaml -y --set tag=zirain-dev
	istioctl-dev install -f green.yaml -y --set tag=zirain-dev

install-ns:
	kubectl create ns blue
	kubectl create ns green
	kubectl label ns blue istio-discovery=blue istio.io/rev=blue
	kubectl label ns green istio-discovery=green istio.io/rev=green

purge-istio:
	istioctl uninstall --purge -y

purge-ns:
	kubectl delete ns blue --ignore-not-found
	kubectl delete ns green --ignore-not-found
	kubectl delete ns istio-system

.PHONY: cleanup
cleanup: purge-istio purge-ns