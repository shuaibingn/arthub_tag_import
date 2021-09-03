## ArtHub Tag Import 

### 使用方法

1. 配置文件

| 参数 | 说明 |
| ---- | ---- |
| domain | arthub api接口 |
| token | arhub管理后台生成的token |
| depot | arthub资源库 |
| path | 要上传excel的文件路径 |
| name | 要上传excel的文件名称 |

2. 命令参数

```shell
arthub_tag_import.exe --domain "domain" -- token "token" --depot "depot" --file.path "filepath" --file.name "filename"
```

当命令行存在时优先会使用命令行解析参数, 当命令行不存在是默认使用`config.yaml`文件来解析参数
