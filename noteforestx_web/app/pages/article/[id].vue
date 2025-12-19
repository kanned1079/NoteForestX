<script setup lang="ts">
import {useRoute} from "#vue-router";
import {MdPreview, MdCatalog} from 'md-editor-v3';
const route = useRoute();
const articleId = route.params.id as string;
import 'md-editor-v3/lib/style.css';
const id = 'article-preview';


const p = `
## 基于Vue3+Golang的外卖系统的设计与实现

### 启动项目

#### 开发软件以及环境下载

根据你的项目用到的工具来下载 **不需要全部下载** 如果不确定自己能不能完成可以联系我们

- **Java 21 LTS**

注意下载21版本 [下载链接](https://www.oracle.com/java/technologies/downloads/#jdk21-windows)
打开直接安装即可 不需要手动配置环境

![java install](https://ikanned.com:2444/d/dell_r730xd_storage_pool2/public/javadown.png)

- **Golang 1.24**

注意下载24版本 [下载链接](https://go.dev/dl/)
打开直接安装即可 不需要手动配置环境

![java install](https://ikanned.com:2444/d/dell_r730xd_storage_pool2/public/golangdown.png)

- **JetBrains IltelliJ IDEA**

目前IDEA的免费版和Ultimate已经合并 直接下载即可 [下载链接](https://www.jetbrains.com.cn/idea/download)

![idea install](https://ikanned.com:2444/d/dell_r730xd_storage_pool2/public/ideadown.png)

- **Nodejs**

用于安装配置前端依赖 注意使用\`v22.21\`  [下载链接](https://nodejs.org/zh-cn/download)

![nodejs install](https://ikanned.com:2444/d/dell_r730xd_storage_pool2/public/nodejsdown.png)

- **MySQL**

MySQL数据库  [下载链接](https://dev.mysql.com/downloads/installer/)

![mysql install](https://ikanned.com:2444/d/dell_r730xd_storage_pool2/public/mysqldown.png)

#### 组织文件夹

我们给到的将会由两个部分组成 \`前端\` \`后端\` 你需要新建一个 **英文** 文件夹来存放 例如(myapp)：

\`\`\`shell
myapp:
│
├─ xxxapp_web (前端)
├─ xxxapp_server (后端)
└─ 其他文件
\`\`\`

在安装好环境后使用 \`JetBrains IltelliJ IDEA\` 打开即可



#### 安装依赖

##### 前端

开启一个命令行进入 \`xxxapp_web\` 运行下面的命令安装依赖

\`\`\`shell
npm install
\`\`\`

##### 后端

开启一个命令行进入 \`xxxapp_server\` 运行下面的命令安装依赖

\`\`\`shell
go mod tidy #如果你使用Golang 或者你可以告诉我们你用的什么设备我们编译好你直接运行即可
\`\`\`


`
const isDarkModeEnabled = ref<boolean>(false)

const scrollElement = document.documentElement;

onMounted(() => {
  console.log(articleId)
})

</script>

<template>
  <div class="pt-[40px] px-4 md:px-8 flex flex-col md:flex-row gap-6">
    <!-- 左侧目录 -->
    <aside
        class="w-full md:w-1/4 flex-shrink-0 border-r border-gray-200 dark:border-gray-700 pr-4 md:pr-6 sticky top-[40px] h-[calc(100vh-40px)] overflow-auto"
    >
      <h2 class="text-lg font-semibold mb-4">目录</h2>
      <MdCatalog :editorId="id" :scrollElement="scrollElement" />
    </aside>

    <!-- 右侧内容 -->
    <main
        class="flex-1 w-full overflow-auto"
        ref="scrollElement"
    >
      <MdPreview
          :id="id"
          :modelValue="p"
          :theme="isDarkModeEnabled ? 'dark' : undefined"
          :preview-theme="'github'"
      />
    </main>
  </div>
</template>

<style scoped>

aside::-webkit-scrollbar {
  width: 6px;
}
aside::-webkit-scrollbar-thumb {
  background-color: rgba(100,100,100,0.3);
  border-radius: 3px;
}

.md-editor {
  z-index: 1 !important;
  --md-border-color: #7c7c7c !important;
  border-radius: 4px;
  //border: #919191 solid 1px;
  --md-border-active-color: #000 !important;
  --md-bk-color: rgba(0, 0, 0, 0.0) !important;
}

.md-editor-dark {
  --md-bk-color: rgba(0, 0, 0, 0.0) !important;
  --md-border-color: #989898 !important;
}

.operate-part-root {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  margin: 20px 0;
  .operate-l {

  }
}
</style>