function show(res){
    if(res.msg==="success"){
        window.location.href="/"
    }else{
        alert("登录失败")
    }
}

$(document).ready(function () {
    $('#yy,#yzc').click(function (){
        window.location.href='/user/register'
    })
    $('#ydl').click(function (){
        window.location.href='/user/login'
    })
    $('#login').click(function (){
        let username=$('#username').val()
        let password=$('#password').val()
        //验证用户名是否为空
        if(username.trim()===''){
            alert('请输入用户名')
            $('#username').focus()
            return
        }
        //验证密码是否为空
        if(password.trim()===''){
            alert('请输入密码')
            $('#password').focus()
            return
        }
        //post到服务端
        let data=$('form[name="f1"]').serialize()
        console.log(data)
        $.ajax({
            type:"post",
            url:"/user/login",
            data:data,
            dataType:'json',
            success:show,
        })
    })
})