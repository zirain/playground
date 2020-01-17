当前遗留问题

logstash 到 Azure Blob 目前找到可用的插件是 [logstash-output-azure](https://rubygems.org/gems/logstash-output-azure/)

这个插件目前存在如下问题：

- 不支持 `storage_blob_host` 配置项，目前默认只能在Global Azure上使用，修改本版在[GitHub](https://github.com/zirain/Logstash-output-to-Azure-Blob)