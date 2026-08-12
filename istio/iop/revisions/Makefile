ISTIOCTL_HOME ?= ~/Codes/istioctl

install-istio:
	istioctl install -f blue.yaml -y
	istioctl install -f green.yaml -y

install-istio-dev:
	istioctl-dev install -f blue.yaml -y --set tag=zirain-dev
	istioctl-dev install -f green.yaml -y --set tag=zirain-dev

install-eastwest:
	istioctl manifest generate -f eastwest.yaml | kubectl apply -f -
	kubectl apply -f expose-services.yaml -n istio-system

purge-eastwest:
	kubectl delete -f expose-services.yaml -n istio-system --ignore-not-found
	istioctl manifest generate -f eastwest.yaml | kubectl delete -f -

install-ns:
	kubectl create ns blue
	kubectl create ns green

purge-istio:
	istioctl uninstall --purge -y

purge-ns:
	kubectl delete ns blue --ignore-not-found
	kubectl delete ns green --ignore-not-found
	kubectl delete ns istio-system

.PHONY: cleanup
cleanup: purge-istio purge-ns