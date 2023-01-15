$(document).ready(function () {
    $('#tj').click(function (){
        console.log("haha")
        let c=$('#content').val()
        if(c.trim()==""){
            alert("请输入留言内容")
            $('#content').focus()
            return
        }
        console.log($("form[name='f1']").serialize())
        $.ajax({
            type: "post",
            url: "/message/send",
            dataType: 'json',
            data:$("form[name='f1']").serialize(),
            success: function (res) {
                if(res.msg=="fail"){
                    alert("留言失败")
                }else{
                    alert("留言成功")
                    window.location.href="/"
                }
            },
        })
    })

})