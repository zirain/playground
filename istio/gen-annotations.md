
```bash
# tools
go build -o annotations_prep ./cmd/annotations_prep/main.go && chmod +x annotations_prep && cp annotations_prep /usr/local/bin
# api
annotations_prep --input annotation/annotations.yaml --output annotation/annotations.gen.go --html_output annotation/annotations.pb.html --collection_type annotation
cp annotation/annotations.pb.html ../istio.io/content/en/docs/reference/config/annotations/index.html

annotations_prep --input label/labels.yaml --output label/labels.gen.go --html_output label/labels.pb.html --collection_type label
cp label/labels.pb.html ../istio.io/content/en/docs/reference/config/labels/index.html

# istio.io
make serve
```