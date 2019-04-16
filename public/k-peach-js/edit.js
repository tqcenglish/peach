$(document).ready(function () {
    const el = document.querySelector('textarea');
    const stackedit = new Stackedit({url: "/stackedit"});
    // Open the iframe
    stackedit.openFile({
      content: {
        text: $.trim(el.value) // and the Markdown content.
      }
    });

    // Listen to StackEdit events and apply the changes to the textarea.
    stackedit.on('fileChange', (file) => {
      el.value = $.trim(file.content.text);
    });

    // 保存
    stackedit.on("close", () => {
      var settings = {
        "async": false,
        "url": window.location.href,
        "method": "POST",
        "headers": {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        "data": {
          "context": $.trim(el.value)
        }
      }

      $.ajax(settings).done( (res) => {
        window.location.href = res.path;
      });
    })
});