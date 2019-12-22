require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
const hljs = require("highlight.js/lib/");
/*
hljs.registerLanguage('html', require('highlight.js/lib/languages/haml'));
hljs.registerLanguage('sql', require('highlight.js/lib/languages/sql'));
hljs.registerLanguage('php', require('highlight.js/lib/languages/php'));
hljs.registerLanguage('javascript', require('highlight.js/lib/languages/javascript'));
 */
$(() => {

});
document.addEventListener('DOMContentLoaded', (event) => {
    document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
    });
});
