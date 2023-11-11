<template>
    <div style="border: 1px solid #ccc">
        <Toolbar style="border-bottom: 1px solid #ccc" :editor="editor" :defaultConfig="toolbarConfig" :mode="mode" />
        <Editor style="height: 300px; overflow-y: hidden;" @customPaste="customPaste" v-model="content"
            :defaultConfig="editorConfig" :mode="mode" @onChange="onChange" @onCreated="onCreated" />
    </div>
</template>
<script >
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { onBeforeUnmount, ref, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

export default {
    components: { Editor, Toolbar },
    props: ['html'],
    data() {
        let _this = this;
        return {
            editor: null,
            content:"",
            toolbarConfig: {},
            editorConfig: {
                placeholder: '请输入内容...',
                // 所有的菜单配置，都要在 MENU_CONF 属性下
                MENU_CONF: {
                    uploadImage: {
                        server: _this.$theia.env.join("/file/upload"),
                        metaWithUrl: true, // join params to url
                        headers: { 'Content-Type': 'multipart/form-data' },
                        customUpload(file, insertFn) {
                            _this.$emit('uploadingChange', true);
                            let formData = new FormData();
                            formData.append('file', file);
                            formData.append('dir', 'supplier');
                            return _this.$theia.http.upload(_this.$theia.env.join("/file/upload"), formData, {
                                'Content-Type': 'multipart/form-data',
                            })
                                .then(async (res) => {
                                    console.log("res:", res)
                                    insertFn(res.path);
                                    _this.$nextTick(() => {
                                        setTimeout(() => {
                                            _this.editor?.focus(true);
                                            _this.$emit('uploadingChange', false);
                                        }, 500);
                                    });
                                    return res.path;
                                })
                                .catch((err) => {
                                    _this.$message.error('上传失败');
                                    _this.$nextTick(() => {
                                        setTimeout(() => {
                                            _this.editor?.focus(true);
                                            _this.$emit('uploadingChange', false);
                                        }, 500);
                                    });
                                });
                        },
                    },
                },
            },
        }
    },
    watch: {
        html(){
        this.content = this.html
    },
  },
    methods: {
        onChange(editor) {
            this.$emit('onChange', editor.getHtml());
        },
        onCreated(editor) {
            this.editor = Object.seal(editor); // 【注意】一定要用 Object.seal() 否则会报错
        },

        insertTextHandler() {
            const editor = this.editor;
            if (editor == null) return;
            editor.insertText(' hello ');
        },
        printEditorHtml() {
            const editor = this.editor;
            if (editor == null) return;
            console.log(editor.getHtml());
        },
        disableHandler() {
            const editor = this.editor;
            if (editor == null) return;
            editor.disable();
        },
        // onCreated(editor) {
        //     this.editor = editor // 一定要用 Object.seal() ，否则会报错
        // },
        customPaste(editor, event) {

            // 获取粘贴的html部分（？？没错粘贴word时候，一部分内容就是html），该部分包含了图片img标签
            let html = event.clipboardData.getData('text/html');

            // 获取rtf数据（从word、wps复制粘贴时有），复制粘贴过程中图片的数据就保存在rtf中
            const rtf = event.clipboardData.getData('text/rtf');

            if (html && rtf) { // 该条件分支即表示要自定义word粘贴

                // 列表缩进会超出边框，直接过滤掉
                html = html.replace(/text\-indent:\-(.*?)pt/gi, '')

                // 从html内容中查找粘贴内容中是否有图片元素，并返回img标签的属性src值的集合
                const imgSrcs = this.findAllImgSrcsFromHtml(html);

                // 如果有
                if (imgSrcs && Array.isArray(imgSrcs) && imgSrcs.length) {

                    // 从rtf内容中查找图片数据
                    const rtfImageData = this.extractImageDataFromRtf(rtf);

                    // 如果找到
                    if (rtfImageData.length) {

                        // TODO：此处可以将图片上传到自己的服务器上
                        console.log("rtfImageData:", rtfImageData)

                        // 执行替换：将html内容中的img标签的src替换成ref中的图片数据，如果上面上传了则为图片路径
                        html = this.replaceImagesFileSourceWithInlineRepresentation(html, imgSrcs, rtfImageData)
                        editor.dangerouslyInsertHtml(html);

                    }
                }

                // 阻止默认的粘贴行为
                event.preventDefault();
                return false;
            } else {
                return true;
            }

        },
        findAllImgSrcsFromHtml(htmlData) {

            let imgReg = /<img.*?(?:>|\/>)/gi; //匹配图片中的img标签
            let srcReg = /src=[\'\"]?([^\'\"]*)[\'\"]?/i; // 匹配图片中的src

            let arr = htmlData.match(imgReg); //筛选出所有的img
            if (!arr || (Array.isArray(arr) && !arr.length)) {
                return false;
            }


            let srcArr = [];
            for (let i = 0; i < arr.length; i++) {
                let src = arr[i].match(srcReg);
                // 获取图片地址
                srcArr.push(src[1]);
            }

            return srcArr;
        },

        /**
        * 从rtf内容中匹配返回图片数据的集合
        * @param rtfData
        * @return Array
        */
        extractImageDataFromRtf(rtfData) {
            if (!rtfData) {
                return [];
            }

            const regexPictureHeader = /{\\pict[\s\S]+?({\\\*\\blipuid\s?[\da-fA-F]+)[\s}]*/
            const regexPicture = new RegExp('(?:(' + regexPictureHeader.source + '))([\\da-fA-F\\s]+)\\}', 'g');
            const images = rtfData.match(regexPicture);
            const result = [];

            if (images) {
                for (const image of images) {
                    let imageType = false;

                    if (image.includes('\\pngblip')) {
                        imageType = 'image/png';
                    } else if (image.includes('\\jpegblip')) {
                        imageType = 'image/jpeg';
                    }

                    if (imageType) {
                        result.push({
                            hex: image.replace(regexPictureHeader, '').replace(/[^\da-fA-F]/g, ''),
                            type: imageType
                        });
                    }
                }
            }

            return result;
        },

        /**
        * 将html内容中img标签的属性值替换
        * @param htmlData html内容
        * @param imageSrcs html中img的属性src的值的集合
        * @param imagesHexSources rtf中图片数据的集合，与html内容中的img标签对应
        * @param isBase64Data 是否是Base64的图片数据
        * @return String
        */
        replaceImagesFileSourceWithInlineRepresentation(htmlData, imageSrcs, imagesHexSources, isBase64Data =
            true) {
            if (imageSrcs.length === imagesHexSources.length) {
                for (let i = 0; i < imageSrcs.length; i++) {
                    const newSrc = isBase64Data ?
                        `data:${imagesHexSources[i].type};base64,${_convertHexToBase64(imagesHexSources[i].hex)}` :
                        imagesHexSources[i];

                    htmlData = htmlData.replace(imageSrcs[i], newSrc);
                }
            }

            return htmlData;
        },

        /**
        * 十六进制转base64
        */
        _convertHexToBase64(hexString) {
            return btoa(hexString.match(/\w{2}/g).map(char => {
                return String.fromCharCode(parseInt(char, 16));
            }).join(''));
        },
    },
    mounted() {
        this.editor = shallowRef()
    },
    beforeDestroy() {
        const editor = this.editor
        if (editor == null) return
        editor.destroy() // 组件销毁时，及时销毁编辑器
    }

}
</script>