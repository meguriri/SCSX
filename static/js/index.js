
$(document).ready(function () {
    $('#exit').click(function (){
        $.ajax({
            type:"get",
            url:"/home/exit",
            dataType:'json',
            success:function (res){
                if(res.msg==="success"){
                    window.location.href="/user/login"
                }
            },
        })
    })
    $('#ydl,#gwcdl').click(function (){
        window.location.href='/user/login'
    })
    $.ajax({
        type:"get",
        url:"/home/userShow",
        dataType:'json',
        success:function (res){
            $('#memberInfo').empty()
            if(res.msg==="success"){
                $('#memberInfo').append(res.username)
            }
        },
    })
    $.ajax({
        type:"get",
        url:"/home/allCategory",
        dataType:'json',
        success:function (res){
            let obj=eval(res.list)
            console.log("list:",res.list)
            console.log("obj:",obj)
            let str="<ul>"
            for(let i=0;i<obj.length;i++){
                str+="<li>\n" +
                "                    \t<div class=\"fj\">\n" +
                "                        \t<span class=\"n_img\"><span></span><img src=\"../static/resources/images/nav"+(i+1)+".png\" /></span>\n" +
                "                            <span class=\"fl\"><a style='color:white;' onclick='selByCid("+obj[i].id+")'>"+obj[i].name+"</a></span>\n" +
                "                        </div>\n" +
                "                        \n" +
                "                    </li>"
            }
            str+="</ul>"
            $('#categoryContent').empty()
            $('#categoryContent').html(str)
        },
    })
})

function selByCid(cid){
    window.location.href="/product?cid="+cid
}