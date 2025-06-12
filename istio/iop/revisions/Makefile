.PHONY: setup cleanup

ISTIOCTL_HOME ?= ~/Codes/istioctl

install-istio:
	$(ISTIOCTL_HOME)/istio-1.25.2/bin/istioctl install -f blue.yaml -y
	istioctl install -f green.yaml -y

setup-ns:
	kubectl create ns blue
	kubectl create ns green
	kubectl label ns blue istio-discovery=enabled istio.io/rev=blue
	kubectl label ns green istio-discovery=enabled istio.io/rev=green

purge:
	istioctl uninstall --purge -y

cleanup-ns:
	kubectl delete ns blue
	kubectl delete ns green