
function checkEmail(email){
    let pattern=/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
    return pattern.test(email);
}

function show(res){
    if(res.msg==="success"){
        window.location.href="/"
    }else{
        alert("注册失败")
    }
}

$(document).ready(function () {
    $('#yzc').click(function (){
        window.location.href='/user/register'
    })
    $('#yy,#ydl').click(function (){
        window.location.href='/user/login'
    })
    $('#reg').click(function (){
        let username=$('#username').val()
        let password=$('#password').val()
        let password1=$('#password1').val()
        let email=$('#email').val()
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
        //验证确认密码是否为空
        if(password1.trim()===''){
            alert('请输入确认密码')
            $('#password1').focus()
            return
        }
        //验证密码与确认密码是否一致
        if(password!==password1){
            alert('密码不一致')
            $('#password1').focus()
            return
        }
        //验证邮箱是否为空
        if(email.trim()===''){
            alert('请输入邮箱')
            $('#email').focus()
            return
        }else{
            //验证邮箱是否合法
            if(checkEmail(email.trim())===false){
                alert("邮箱不合法")
                $('#email').focus()
                return
            }
        }
        //post到服务端
        let data=$('form[name="f1"]').serialize()
        console.log(data)
        $.ajax({
            type:"post",
            url:"/user/register",
            data:data,
            dataType:'json',
            success:show,
        })
    })
})