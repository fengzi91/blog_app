const $ = require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
const sweetAlert = require("sweetalert");
$(() => {
    $(document).scroll(function(){
        if ($(this).scrollTop() >= 260) {
            $('#blog-tools').show();
        } else {
            $('#blog-tools').hide();
        }
    })
    /*
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
    }*/
});
document.addEventListener('DOMContentLoaded', (event) => {
    document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
    });
});
window.deletePost = function (id) {
    swal("确定要删除这篇文章吗?", {
        buttons: ["取消", "确认"],
    }).then((val) => {
        if (val) {
            $.get(id).then(() => {
                window.location.href = '/'
            })
        }
    });
}
window.deleteComment = function (id) {
    swal("确定要删除这条评论吗?", {
        buttons: ["取消", "确认"],
    }).then((val) => {
        if (val) {
            $.get(id).then(() => {
                window.location.reload()
            })
        }
    });
}
