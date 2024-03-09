

tools/bin/crd-ref-docs \
        --source-path=api/v1alpha1 \
        --config=tools/crd-ref-docs/config.yaml \
        --templates-dir=tools/crd-ref-docs/templates \
        --output-path=site/content/en/latest/api/extension_types.md \
        --max-depth 10 \
        --renderer=markdown


go run main.go --source-path=test/api \
        --config=test/config.yaml \
        --max-depth 10 \
        --renderer=markdown \
        --output-path=test/expected.md \
        --templates-dir=./templates/markdown

go run main.go --source-path=test/api/ \
        --config=test/config.yaml \
        --max-depth 10 \
        --renderer=asciidoctor \
        --output-path=test/expected.asciidoc \
        --templates-dir=./templates/asciidoctor