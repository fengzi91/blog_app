const $ = require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
const hljs = require("highlight.js/lib/");
const sweetAlert = require("sweetalert");
const Editor = require('tui-editor');
/*
hljs.registerLanguage('html', require('highlight.js/lib/languages/haml'));
hljs.registerLanguage('sql', require('highlight.js/lib/languages/sql'));
hljs.registerLanguage('php', require('highlight.js/lib/languages/php'));
hljs.registerLanguage('javascript', require('highlight.js/lib/languages/javascript'));
 */
$(() => {
    $(document).scroll(function(){
        if ($(this).scrollTop() >= 260) {
            $('#blog-tools').show();
        } else {
            $('#blog-tools').hide();
        }
    })
    if ($('#editSection').length > 0) {
        const editor = new Editor({
            el: document.querySelector('#editSection'),
            initialEditType: 'markdown',
            previewStyle: 'vertical',
            height: '300px',
            hooks: {
                addImageBlobHook: (fileOrBlob, callback, source) => {
                    uploadImage(fileOrBlob).then(res => {
                        const customDescription = `${res.url}_foo`;
                        console.log(res);
                        callback(res.url, customDescription);
                    });
                }
            }
        });
        $('#form').on('submit', (e) => {
            var content = editor.getMarkdown();
            $('#content').val(content);
            var canSubmit = false;
            if ($('#form input[name="Title"]').val() === '') {
                swal({
                    title: "请填写标题",
                    icon: "error"
                });
            } else if($('#content').val() === '') {
                swal({
                    title: "请输入正文",
                    icon: "error"
                });
            } else if($('#category').val() == 0) {
                swal({
                    title: "请选择分类",
                    icon: "error"
                });
            } else {
                canSubmit = true;
            }
            canSubmit ? '' : e.preventDefault();
        })
    }
});
function uploadImage(fileOrBlob) {
    const result = {}

    if (Object.prototype.toString.call(fileOrBlob) === '[object File]') {
        const formData = new FormData();
        formData.append('file', fileOrBlob);
        return new Promise((resolve, reject) => {
            $.ajax({
                type: "POST",
                url: url,
                data: formData,
                contentType: false,
                processData: false,
                dataType: 'json',
            }).then(res => {
                resolve(res);
            }).catch(err => {
                reject(error)
            })
        });
    }
    return false;
}

document.addEventListener('DOMContentLoaded', (event) => {
    document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
    });
});
