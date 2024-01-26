# Github action for Cloudsmith

- `name` + `verion` must be unique for each upload.
- can't use `latest` as version name.

```yaml
      - name: "Upload Istioctl to Cloudsmith"
        uses: cloudsmith-io/action@fc70bf3268051bdf07305c56fd2ace32c3fc8348  # 0.5.4
        with:
          owner: "${{ secrets.CLOUDSMITH_REPO_OWNER }}"
          repo: "${{ secrets.CLOUDSMITH_REPO }}"
          api-key: ${{ secrets.CLOUDSMITH_API_KEY }}
          command: "push"
          format: "raw"
          file: "dist/istioctl.tar.gz"
          republish: "true"
          name: "istioctl"
          # This gives, for example:
          # - if push event is master: linux-amd64-latest, windows-amd64-latest.
          # - if push event is a tag with value 1.6.0: linux-amd64-1.6.0, windows-amd64-1.6.0.
          version: "${{ matrix.os }}-${{ matrix.arch }}-${{ needs.build.outputs.current-version }}"
```