<script setup lang="ts">
import {useRoute} from "#vue-router";
import {MdPreview, MdCatalog} from 'md-editor-v3';
const route = useRoute();
const articleId = route.params.id as string;
// import 'md-editor-v3/lib/style.css';
import '~/assets/css/md-style.css'
import type {Article} from "~/types/article";
import WidthTest from "~/components/RedesignedComponents/WidthTest.vue";
import dayjs from 'dayjs'
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import ArticleHeader from "~/components/RedesignedComponents/ArticleHeader.vue";

const id = 'article-preview';

useScrollFadeIn({
  selector: '.animate-card-article-id',
  // y: 60,
  // duration: 0.6,
  // stagger: 0.15
  direction: 'up',
  x: 200,
  stagger: 0.1,
  duration: 0.6,
  start: 'top 90%',
  useScrollTrigger: false
})


const isLargeScreen = ref(false)

const checkScreen = () => {
  isLargeScreen.value = window.innerWidth >= 1024 // lg breakpoint
}


const p = `### 启动项目



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

const currentArticle = ref<Article>({
  id: 'fergvegvea',
  title: "基于Vue3+Golang的外卖系统的设计与实现",
  content: p,
  tags: [{id: 'tyhyrhbe', name: 'Vue3'}, {id: 'tyhyysegeag45', name: '服务端渲染'}],
  status: 'published',
  top: false,
  created_at: '2018-08-10T00:00:00.000Z',
})



// let mediaQuery: MediaQueryList | null = null

const { isDarkMode } = useDarkMode()

onMounted(() => {
  console.log(articleId)

  checkScreen()
  window.addEventListener('resize', checkScreen)


})

</script>

<template>
<!--  <div class="pt-[40px] px-4 md:px-8 flex flex-col md:flex-row gap-6">-->
<!--    &lt;!&ndash; 左侧目录 &ndash;&gt;-->
<!--    <aside-->
<!--        class="w-full md:w-1/4 flex-shrink-0 border-r border-gray-200 dark:border-gray-700 pr-4 md:pr-6 sticky top-[40px] h-[calc(100vh-40px)] overflow-auto"-->
<!--    >-->
<!--      <h2 class="text-lg font-semibold mt-10 mb-4">目录</h2>-->
<!--      <MdCatalog :editorId="id" scrollElement="html" />-->
<!--    </aside>-->

<!--    &lt;!&ndash; 右侧内容 &ndash;&gt;-->
<!--    <main-->
<!--        class="flex-1 w-full overflow-auto"-->
<!--        ref="scrollElement"-->
<!--    >-->
<!--      <MdPreview-->
<!--          :id="id"-->
<!--          :modelValue="p"-->
<!--          :theme="isDarkMode ? 'dark' : undefined"-->
<!--          :preview-theme="'github'"-->
<!--      />-->
<!--    </main>-->
<!--  </div>-->


  <!-- 页面级容器 -->
  <div class="pt-[40px] px-4 md:px-6 lg:px-8">

    <!-- 头部（独立，不参与 sticky 计算） -->
    <ArticleHeader
        v-if="isLargeScreen"
        :title="currentArticle.title"
        :tags="currentArticle.tags"
        :createdAt="dayjs(currentArticle.created_at).format('YYYY-MM-DD')"
    />

    <!-- 真正的内容区（和旧版本一样“干净”） -->
    <div class="flex flex-col lg:flex-row gap-6">

      <!-- 目录 -->
      <aside
          v-if="true"
          class="hidden lg:block w-1/4 flex-shrink-0
             border-r border-gray-200 dark:border-gray-700
             pr-6 sticky top-[40px]
             h-[calc(100vh-40px)] overflow-auto animate-card-article-id"
      >
        <p class="text-xl font-bold mt-6">目录</p>
        <MdCatalog :editorId="id" scrollElement="html" />
      </aside>

      <!-- 正文 -->
      <main class="flex-1 min-w-0 pb-[100px] animate-card-article-id">
        <div class="mx-auto max-w-[720px] lg:max-w-[900px]">
          <ArticleHeader
              v-if="!isLargeScreen"
              :title="currentArticle.title"
              :tags="currentArticle.tags"
              :createdAt="dayjs(currentArticle.created_at).format('YYYY-MM-DD')"
          />

          <Divider v-if="!isLargeScreen" class="mt-0 mb-0" />

          <MdPreview
              :id="id"
              :modelValue="p"
              :theme="isDarkMode ? 'dark' : undefined"
              :preview-theme="'github'"
          />
        </div>
      </main>

    </div>
  </div>


</template>

<style>

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
  border-radius: 6px;
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

}



</style>